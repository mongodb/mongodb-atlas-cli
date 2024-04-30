import copy
import functools
import multiprocessing
import signal
import socket
import ssl
import sys
import six
from kmip.services.server import KmipServer, engine, monitor

# Monkey patch the KMIP server and engine classes with the following fixes:
# - When the KMIP engine processes an encrypt request with empty cryptographic parameters, it will use a set of default cryptographic parameters to process the request.
# - The default socket timeout on the server is removed -- servers shouldn't just kill a connection after 10 seconds of inactivity.


def patch_server():
    from kmip.core import enums
    from kmip.core import exceptions
    from kmip.core.messages import payloads
    from kmip.core import policy as operation_policy
    from kmip.core.attributes import CryptographicParameters

    def _process_encrypt_patched(self, payload):
        self._logger.info("Processing operation: Encrypt")

        unique_identifier = self._id_placeholder
        if payload.unique_identifier:
            unique_identifier = payload.unique_identifier

        # The KMIP spec does not indicate that the Encrypt operation should
        # have it's own operation policy entry. Rather, the cryptographic
        # usage mask should be used to determine if the object can be used
        # to encrypt data (see below).
        managed_object = self._get_object_with_access_controls(
            unique_identifier, enums.Operation.GET
        )

        cryptographic_parameters = payload.cryptographic_parameters

        if cryptographic_parameters is None:
            # Monkey patched here -- rather than exception, we set to default params.
            default_crypto_params = CryptographicParameters(
                block_cipher_mode=enums.BlockCipherMode.CBC,
                padding_method=enums.PaddingMethod.PKCS5,
                cryptographic_algorithm=enums.CryptographicAlgorithm.AES,
            )
            cryptographic_parameters = default_crypto_params

        if managed_object._object_type != enums.ObjectType.SYMMETRIC_KEY:
            raise exceptions.PermissionDenied(
                "The requested encryption key is not a symmetric key. "
                "Only symmetric encryption is currently supported."
            )

        if managed_object.state != enums.State.ACTIVE:
            raise exceptions.PermissionDenied(
                "The encryption key must be in the Active state to be used "
                "for encryption."
            )

        masks = managed_object.cryptographic_usage_masks
        if enums.CryptographicUsageMask.ENCRYPT not in masks:
            raise exceptions.PermissionDenied(
                "The Encrypt bit must be set in the encryption key's "
                "cryptographic usage mask."
            )

        result = self._cryptography_engine.encrypt(
            cryptographic_parameters.cryptographic_algorithm,
            managed_object.value,
            payload.data,
            cipher_mode=cryptographic_parameters.block_cipher_mode,
            padding_method=cryptographic_parameters.padding_method,
            iv_nonce=payload.iv_counter_nonce,
            auth_additional_data=payload.auth_additional_data,
            auth_tag_length=cryptographic_parameters.tag_length,
        )

        response_payload = payloads.EncryptResponsePayload(
            unique_identifier,
            result.get("cipher_text"),
            result.get("iv_nonce"),
            result.get("auth_tag"),
        )
        return response_payload

    _process_encrypt_patched = engine.KmipEngine._kmip_version_supported("1.2")(
        _process_encrypt_patched
    )

    def _process_decrypt_patched(self, payload):
        self._logger.info("Processing operation: Decrypt")

        unique_identifier = self._id_placeholder
        if payload.unique_identifier:
            unique_identifier = payload.unique_identifier

        # The KMIP spec does not indicate that the Decrypt operation should
        # have it's own operation policy entry. Rather, the cryptographic
        # usage mask should be used to determine if the object can be used
        # to decrypt data (see below).
        managed_object = self._get_object_with_access_controls(
            unique_identifier, enums.Operation.GET
        )

        cryptographic_parameters = payload.cryptographic_parameters
        if cryptographic_parameters is None:
            # Monkey patched here -- rather than exception, we set to default params.
            default_crypto_params = CryptographicParameters(
                block_cipher_mode=enums.BlockCipherMode.CBC,
                padding_method=enums.PaddingMethod.PKCS5,
                cryptographic_algorithm=enums.CryptographicAlgorithm.AES,
            )
            cryptographic_parameters = default_crypto_params

        if managed_object._object_type != enums.ObjectType.SYMMETRIC_KEY:
            raise exceptions.PermissionDenied(
                "The requested decryption key is not a symmetric key. "
                "Only symmetric decryption is currently supported."
            )

        if managed_object.state != enums.State.ACTIVE:
            raise exceptions.PermissionDenied(
                "The decryption key must be in the Active state to be used "
                "for decryption."
            )

        masks = managed_object.cryptographic_usage_masks
        if enums.CryptographicUsageMask.DECRYPT not in masks:
            raise exceptions.PermissionDenied(
                "The Decrypt bit must be set in the decryption key's "
                "cryptographic usage mask."
            )

        result = self._cryptography_engine.decrypt(
            cryptographic_parameters.cryptographic_algorithm,
            managed_object.value,
            payload.data,
            cipher_mode=cryptographic_parameters.block_cipher_mode,
            padding_method=cryptographic_parameters.padding_method,
            iv_nonce=payload.iv_counter_nonce,
            auth_additional_data=payload.auth_additional_data,
            auth_tag=payload.auth_tag,
        )

        response_payload = payloads.DecryptResponsePayload(unique_identifier, result)
        return response_payload

    _process_decrypt_patched = engine.KmipEngine._kmip_version_supported("1.2")(
        _process_decrypt_patched
    )

    def no_placeholder(fn):
        @functools.wraps(fn)
        def inner(self, *args, **kwargs):
            self._logger.info("Clearing out the current _id_placeholder value")
            self._id_placeholder = None
            return fn(self, *args, **kwargs)

        return inner

    def start_patched(self):
        """
        Prepare the server to start serving connections.

        Configure the server socket handler and establish a TLS wrapping
        socket from which all client connections descend. Bind this TLS
        socket to the specified network address for the server.

        Raises:
            NetworkingError: Raised if the TLS socket cannot be bound to the
                network address.
        """
        self.manager = multiprocessing.Manager()
        self.policies = self.manager.dict()
        policies = copy.deepcopy(operation_policy.policies)
        for policy_name, policy_set in six.iteritems(policies):
            self.policies[policy_name] = policy_set

        self.policy_monitor = monitor.PolicyDirectoryMonitor(
            self.config.settings.get("policy_path"), self.policies, self.live_policies
        )

        def interrupt_handler(trigger, frame):
            self.policy_monitor.stop()

        signal.signal(signal.SIGINT, interrupt_handler)
        signal.signal(signal.SIGTERM, interrupt_handler)

        self.policy_monitor.start()

        self._engine = engine.KmipEngine(
            policies=self.policies,
            database_path=self.config.settings.get("database_path"),
        )

        self._logger.info("Starting server socket handler.")

        # Create a TCP stream socket and configure it for immediate reuse.
        # Monkey patched here - remove default socket timeout.
        # socket.setdefaulttimeout(10)
        self._socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self._socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

        self._logger.debug(
            "Configured cipher suites: {0}".format(
                len(self.config.settings.get("tls_cipher_suites"))
            )
        )
        for cipher in self.config.settings.get("tls_cipher_suites"):
            self._logger.debug(cipher)
        auth_suite_ciphers = self.auth_suite.ciphers.split(":")
        self._logger.debug(
            "Authentication suite ciphers to use: {0}".format(len(auth_suite_ciphers))
        )
        for cipher in auth_suite_ciphers:
            self._logger.debug(cipher)

        self._socket = ssl.wrap_socket(
            self._socket,
            keyfile=self.config.settings.get("key_path"),
            certfile=self.config.settings.get("certificate_path"),
            server_side=True,
            cert_reqs=ssl.CERT_REQUIRED,
            ssl_version=self.auth_suite.protocol,
            ca_certs=self.config.settings.get("ca_path"),
            do_handshake_on_connect=False,
            suppress_ragged_eofs=True,
            ciphers=self.auth_suite.ciphers,
        )

        try:
            self._socket.bind(
                (
                    self.config.settings.get("hostname"),
                    int(self.config.settings.get("port")),
                )
            )
        except Exception as e:
            self._logger.exception(e)
            raise exceptions.NetworkingError(
                "Server failed to bind socket handler to {0}:{1}".format(
                    self.config.settings.get("hostname"),
                    self.config.settings.get("port"),
                )
            )
        else:
            self._logger.info(
                "Server successfully bound socket handler to {0}:{1}".format(
                    self.config.settings.get("hostname"),
                    self.config.settings.get("port"),
                )
            )
            self._is_serving = True

    engine.KmipEngine._process_encrypt = _process_encrypt_patched
    engine.KmipEngine._process_decrypt = _process_decrypt_patched
    engine.KmipEngine._process_get = no_placeholder(engine.KmipEngine._process_get)
    engine.KmipEngine._process_get_attributes = no_placeholder(
        engine.KmipEngine._process_get_attributes
    )
    KmipServer.start = start_patched


patch_server()


def main():
    if len(sys.argv) != 2:
        print("missing <pykmipConfigFolder>: start.py <pykmipConfigFolder>")
        sys.exit()
    config_folder = sys.argv[1]
    server = KmipServer(
        hostname="localhost",
        port=8081,
        key_path=f"{config_folder}/tls-localhost.pem",
        certificate_path=f"{config_folder}/tls-localhost.pem",
        ca_path=f"{config_folder}/tls-rootCA.pem",
        config_path=None,
        policy_path="",
        log_path=f"{config_folder}/log",
        logging_level="DEBUG",
        database_path=f"{config_folder}/kmip.db",
        auth_suite="TLS1.2",
        enable_tls_client_auth=False,
    )

    server.start()
    server.serve()


if __name__ == "__main__":
    main()
