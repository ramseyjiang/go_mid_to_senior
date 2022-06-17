# delete pem file
cd cert
rm *.pem 

# 1. Create CA private key and self-signed certificate
# adding -nodes to not encrypt the private key
openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=test/CN=*.test.com/emailAddress=test.com@gmail.com"

# 2. Create Web Server private key and CSR
# adding -nodes to not encrypt the private key
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=test/CN=*.ramsey.org/emailAddress=ramsey@gmail.com"

# 3. Use CA's private key to sign the Web Server Certificate Request (CSR) and get back its certificate
openssl x509 -req -in server-req.pem -days 180 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.conf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=CLIENT/CN=*.client.com/emailAddress=client@gmail.com"

# 5. Sign the Client Certificate Request (CSR)
openssl x509 -req -in client-req.pem -days 180 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.conf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text

# Verify certificate
echo "Verifying certificate"
openssl verify -CAfile ca-cert.pem server-cert.pem