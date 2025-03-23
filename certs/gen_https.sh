#!/bin/bash

# Check if domain name is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <domain>"
    exit 1
fi

DOMAIN=$1
# 证书过期日期
DAYS=3650
COUNTRY="CN"
STATE="GZ"
CERT_DIR="./certs"
CONFIG_FILE="$CERT_DIR/openssl.cnf"

# Create directory for certificates if it doesn't exist
mkdir -p $CERT_DIR

# 生成 private key
openssl genrsa -out $CERT_DIR/${DOMAIN}.key 2048

# Create OpenSSL configuration file 通过该配置文件生成对应的csr
cat > $CONFIG_FILE <<EOL
[ req ]
distinguished_name = req_distinguished_name
prompt = no
default_bits = 2048
req_extensions = cert_ext

[ req_distinguished_name ]
countryName         = $COUNTRY
stateOrProvinceName = $STATE
localityName        = $STATE
organizationName    = $STATE
commonName          = $DOMAIN

[ cert_ext ]
basicConstraints = CA:TRUE # CA配置 
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName = @alt_names
# SAN配置
[ alt_names ]
DNS.1 = $DOMAIN
DNS.2 = *.$DOMAIN

EOL

# Generate CSR with extensions 生成csr
openssl req -new -key $CERT_DIR/${DOMAIN}.key -out $CERT_DIR/${DOMAIN}.csr -config $CONFIG_FILE

# Generate self-signed certificate with both extensions 生成crt
openssl x509 -req -days $DAYS \
    -in $CERT_DIR/${DOMAIN}.csr \
    -signkey $CERT_DIR/${DOMAIN}.key \
    -out $CERT_DIR/${DOMAIN}.crt \
    -extensions cert_ext \
    -extfile $CONFIG_FILE

echo "Certificate and key have been generated in the $CERT_DIR directory."