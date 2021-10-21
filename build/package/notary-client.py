#!/usr/bin/env python
#
# Generate signature and checksum files for local archive files for push by MCI
#
# usage: notary-client.py [-h] [--dry-run] [--notary-url NOTARY_URL]
#                         [--key-name KEY_NAME] [--filter FILTER]
#                         [--archive-file-ext ARCHIVE_FILE_EXT]
#                         [--package-file-suffix PACKAGE_FILE_SUFFIX]
#                         [--auth-token AUTH_TOKEN]
#                         [--auth-token-file AUTH_TOKEN_FILE]
#                         [--outputs md5,sha1,sha256,sig] [--skip-missing]
#                         --comment COMMENT
#                         filename [filename ...]

import argparse
import requests
import re
import os.path

from poster.encode import multipart_encode
from poster.streaminghttp import register_openers
import urllib2
import json

# These rest of the imports are for authentication support
# they'll silently fail if auth is unavailable, but the client
# will fail if auth is requested but not supported.
try:
    from Crypto.Protocol.KDF import PBKDF2
    from Crypto.Hash import HMAC
    from Crypto.Hash import SHA
    auth_token_support = True
except ImportError:
    auth_token_support = False
from datetime import datetime

package_files = (".msi", ".rpm")
possible_outputs = ("md5", "sha1", "sha256", "sig")
def parse_commalist(instr):
    return set(re.split(r'\s*,\s*', instr))

# parse command line
#
parser = argparse.ArgumentParser(description='Signs packages and files with the notary service')
parser.add_argument('--dry-run', action='store_true', required=False,
                    help='Print changes that would be made without actually making them', default = False);
parser.add_argument('--notary-url', required=False,
                    help='URL base for notary service', default = 'http://localhost:5000');
parser.add_argument('--key-name', required=False,
                    help='Key parameter to notary service', default = 'test');
parser.add_argument('--filter', required=False,
                    help='Only sign files matching case-insensitive substring filter', default = None);
parser.add_argument('--archive-file-ext', required=False,
                    help='File extension to use for non-package files (.tgz,.zip,etc)', default="sig")
parser.add_argument('--package-file-suffix', required=False, default='-signed',
                    help='Suffix to add to signed package files (i.e. package.rpm -> package-signed.rpm)')
parser.add_argument('--auth-token', required=False,
                    help='Pre-shared authentication key for using signing key (mutually exclusive with --auth-token-file)')
parser.add_argument('--auth-token-file', required=False, type=argparse.FileType("r"),
                    help='Path to file containing pre-shared authentication key for using signing key (mutually exclusive with --auth-token)')
parser.add_argument('--outputs', required=False, type=parse_commalist,
                    help='Comma-separated list of checksums to produce',
                    default=",".join(possible_outputs), metavar=",".join(possible_outputs))
parser.add_argument('--skip-missing', required=False, action='store_true',
                    help='Skip missing files, instead of exiting with an error', default=False)
parser.add_argument('--comment', required=True,
                    help='Comment to add to signing request (will be stored in auditing logs on server)')
parser.add_argument('filename', nargs='+', help="Path(s) to files to sign")
args = parser.parse_args()

notary_urlbase = args.notary_url
notary_url = notary_urlbase + '/api/sign'
notary_payload = { 'key': args.key_name, 'comment': args.comment}

def write_checksum(filename, which, value):
    # outputs has the list of checksums/sigs to produce
    # by default it contains all the possible checksums/sigs
    # skip signatures, since those are handled elsewhere
    if which not in args.outputs:
        return
    outfilename = "{0}.{1}".format(filename, which)
    if args.dry_run:
        print outfilename, value
    else:
        with open(outfilename, "w") as ckfp:
            ckfp.write("{0}  {1}\n".format(value, filename))

def sign_file(fp):
    if args.filter and args.filter.lower() not in filename.lower():
        return

    response_json = {}
    signature = ""

    postdata = notary_payload
    postdata['file'] = fp

    filename = fp.name
    print "signing {0}".format(filename)

    try:
        datagen, headers = multipart_encode(postdata)
        request = urllib2.Request(notary_url, datagen, headers)
        response_json = json.loads(urllib2.urlopen(request).read())
    except Exception as e:
        print("Error contacting notary service for {0}: {1}".format(filename, e))
        return False

    if "permalink" not in response_json:
        print("Signing service didn't return a permalink for {0}: {1}".format(
            filename, response_json["message"]));
        return False

    # Package files get a special suffix to indicate that they are signed
    if filename.endswith(package_files) and len(args.package_file_suffix) > 0:
        (fileroot, fileext) = os.path.splitext(filename)
        filename = fileroot + args.package_file_suffix + fileext

    # If we're writing for tarballs, we want the checksum for the original file (input)
    # otherwise, we want the checksum for the signed package (output)
    whichchecksum = "outchecksum"
    if not filename.endswith(package_files):
        whichchecksum = "inchecksum"

    for cksumname in possible_outputs:
        if cksumname == 'sig':
            continue
        write_checksum(filename, cksumname, response_json[whichchecksum][cksumname])

    # If the user didn't specify a signature, then we're done here.
    if 'sig' not in args.outputs:
        return True

    # Downloading the signature depends on whether it's a signed package or not
    try:
        # Tarballs get the signature content downloaded inline, and written out as text
        if not filename.endswith(package_files):
            signature = requests.get(notary_urlbase + response_json["permalink"]).text
            if args.dry_run == True:
                print filename + "." + args.archive_file_ext
                print signature
            else:
                with open(filename + "." + args.archive_file_ext, "w") as sigfp:
                    sigfp.write(signature)
        # Signed packages get streamed to the original filename
        else:
            download_url = notary_urlbase + response_json["permalink"]
            if args.dry_run == True:
                print filename
                print "Would download {0}".format(download_url)
            else:

                sigrep = requests.get(download_url, stream=True)
                with open(filename, "wb") as sigfp:
                    for chunk in sigrep.iter_content(65536):
                        sigfp.write(chunk)

    except Exception as e:
        print("Error contacting notary service for {0}: {1}".format(filename, e))
        return False

    return True

def auth_key_for_str(auth_key):
    auth_key = PBKDF2(auth_key, auth_key[::-1])
    signed_data = datetime.now().isoformat()
    auth_key = HMAC.new(auth_key, signed_data, SHA).hexdigest() + signed_data
    return auth_key

if __name__ == "__main__":
    if not auth_token_support and (args.auth_token or args.auth_token_file):
        print "An auth token was specified, but pycrypto is not installed to perform authentication."
        exit(1)

    auth_key = None
    if args.auth_token:
        auth_key = auth_key_for_str(args.auth_token)
    elif args.auth_token_file:
        auth_key = auth_key_for_str(args.auth_token_file.readline().strip())
        args.auth_token_file.close()
    if auth_key:
        notary_payload["auth_token"] = auth_key

    register_openers()

    for filename in args.filename:
        try:
            fp = open(filename, "rb")
        except IOError as e:
            if e.errno == 2 and args.skip_missing:
                print "Could not find {0} to sign - skipping.".format(filename)
                continue
            print "Error occurred while opening {0}: {1}".format(filename, e)
            exit(1)
        try:
            if not sign_file(fp):
                exit(1)
        except Exception as e:
            print "Error occured while signing {0}: {1}".format(filename, e)
            exit(1)
    exit(0)
