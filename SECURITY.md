# Security Policy

## Reporting a Vulnerability

Any security concerns or vulnerabilities discovered in one of MongoDB's products or hosted services 
can be responsibly disclosed by utilizing one of the methods described in our [create a vulnerability report](https://docs.mongodb.com/manual/tutorial/create-a-vulnerability-report/) docs page.

While we greatly appreciate community reports regarding security issues, at this time MongoDB does not provide compensation for vulnerability reports.

## Credential Storage

The MongoDB Atlas CLI uses a hybrid approach to store sensitive credentials, prioritizing security while maintaining compatibility across different environments and systems.

### Secure Storage (Preferred)

When available, the CLI uses your operating system's native keyring services to securely store sensitive credentials. This provides the highest level of security by leveraging OS-level encryption and access controls.

**Credentials stored securely include:**
- API Keys (public and private)
- Access Tokens
- Refresh Tokens  
- Client ID and Client Secret

**Supported platforms:**
- **macOS**: Uses Keychain Services for secure credential storage
- **Windows**: Integrates with Windows Credential Manager
- **Linux**: Works with desktop [Secret Service](https://specifications.freedesktop.org/secret-service-spec/latest/) dbus interface

### Insecure Storage (Fallback)

If the operating system's keyring services are unavailable, the CLI automatically falls back to storing credentials in a local configuration file (`config.toml`). This ensures the CLI remains functional even in restricted environments.

**Important security considerations for fallback storage:**
- Credentials are stored in plain text in the configuration file

### Checking Your Storage Method

You can verify whether your CLI is using secure storage by running any `atlas` command. You'll see the message `Warning: Secure storage is not available, falling back to insecure storage`, when secure storage is not available.

### Technical Implementation

The CLI uses the [go-keyring](https://github.com/zalando/go-keyring) library to interface with OS keyring services. Each profile's credentials are stored under a service name prefixed with `atlascli_` (e.g., `atlascli_default` for the default profile).

For detailed information about how your operating system manages keyring encryption and security:

- **macOS Keychain**: [Apple Keychain Services Documentation](https://developer.apple.com/documentation/security/keychain_services)
- **Windows Credential Manager**: [Windows Credential Manager Documentation](https://docs.microsoft.com/en-us/windows/security/identity-protection/credential-guard/)
- **Linux GNOME Keyring**: [GNOME Keyring Documentation](https://wiki.gnome.org/Projects/GnomeKeyring)
