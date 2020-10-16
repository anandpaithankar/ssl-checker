## Certificate Checker

This utility connects to any SSL server on the given port and displays various SSL handshake parameters like negotiated SSL(TLS) version, Ciphers suites, Server certificate chain etc.

```
$ make all
$ ./bin/ssl_checker amazon.com 443

Connecting to amazon.com ...

TLS version : TLS 1.2
Cipher Suite: TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
OCSP Stapling Support: true

Server certificate chain:
[0](ca=false): CN=*.peg.a2z.com,O=Amazon.com\, Inc.,L=Seattle,ST=Washington,C=US [Expires in: 130 days]
[1](ca=true): CN=DigiCert Global CA G2,O=DigiCert Inc,C=US [Expires in: 2846 days]
[2](ca=true): CN=DigiCert Global Root G2,OU=www.digicert.com,O=DigiCert Inc,C=US [Expires in: 751 days]

```
