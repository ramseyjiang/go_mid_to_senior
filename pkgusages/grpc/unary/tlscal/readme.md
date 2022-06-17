If you are already using Mac, it’s probably installed. 
You can run the following command to check which version you are using
% openssl version
LibreSSL 2.8.3

Step 1: Generate CA private key and self-signed certificate
In the cert folder,
% openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=test/CN=test.com/emailAddress=ramseyjiang@gmail.com"

Generating a 4096 bit RSA private key
..........++
................................................................................................................................................++
writing new private key to 'ca-key.pem'
-----


-x509 flag provides self-signed certificate instead of a certificate request
X509 is format of public key certificates. 
The public key certificates also known as a digital certificates or identity verification using to prove ownership of public key.

-newkey rsa:4096 
flog means provides both RSA key with 4096-bit and its certificate request at the same time
-nodes 
flag means to not encrypt the private key
-days
flag means certificate valid date.
-keyout
flag means write the created private key to ca-key.pem file
-out 
flag means write the certificate to ca-cert.pem file

Notice that: We are adding -nodes flag to develop and test the certificate without asking passphrase key.

-subj command explanation

/C=TR is for Country
/ST=ASIA is for State or province
/L=ISTANBUL is for Locality name or city
/O=DEV is for Organization
/OU=TUTORIAL is for Organization Unit
/CN=*.tutorial.dev is for Common Name or domain name
/emailAddress=mert@tutorial.com is for email address

Notice that: in the subj quotes, you cannot add any space in.

When the run commands above , you will see the files that named ca-cert.pem and ca-key.pem has been created.


Step2: Generate Web Server’s Private Key and CSR (Certificate Signing Request)

% openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=test/CN=ramsey.org/emailAddress=ramseyjiang@gmail.com"

Generating a 4096 bit RSA private key
......................................................................................................................++
....++
writing new private key to 'server-key.pem'
-----


-x509 flag is deleted because we don't want to self-sign certificate as like as CA certificate.
-days flag is deleted because we are creating CSR instead of certificate
-keyout the name of the output key
-out the name of the certificate request

When the run commands above , you will see the files that named server-key.pem and server-cert.pem has been created. 
At this point, It’s not certificate just Certificate Request.


Step3: Sign the Web Server Certificate Request (CSR)
In the cert folder, create a server-ext.conf file first.

% openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.conf
Signature ok

subject=/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=test/CN=ramsey.org/emailAddress=ramseyjiang@gmail.com
Getting CA Private Key


-req 
It means it is going to pass in certificate request
-in
It is the name of the request file that is server-req.pem
-CA
it is pass Certificate File of CA:ca-cert.pem
-CAKey 
It is pass private key of CA: ca-key.pem
-CAcreateserial
It is CA must ensure that each certificate it signs with a unique serial number
-out 
It is output to certificate file: server-cert.pem

-extfile 
It is an option tell to openssl we have extra options such as alternative name,


After run the above command, server certificate will be created with server-cert.pem file name.


Step 4: Generate client’s private key and CSR(Certificate Signing Request) 

% openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=CLIENT/CN=*.client.com/emailAddress=client@gmail.com"

Generating a 4096 bit RSA private key
.....................++
............................++
writing new private key to 'client-key.pem'
-----


Step5: Sign the Client Certificate Request (CSR)

Create a client-ext.conf file in the cert folder.

% openssl x509 -req -in client-req.pem -days 180 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.conf

Signature ok
subject=/C=NZ/ST=PACIFIC/L=AUCKLAND/O=DEV/OU=CLIENT/CN=*.client.com/emailAddress=client@gmail.com
Getting CA Private Key


Step6: verify certificate

% openssl verify -CAfile ca-cert.pem server-cert.pem
server-cert.pem: OK


For this demonstration, we are going use same Certificate Authority (CA) for both server and client certificate process. 



Makefiles are an incredibly useful automation tool that you can use to run and build not just your Go applications, 
but for most programming languages.



There are 3 types of gRPC connections:

1. In this connection, all data transfered between client and server is not encrypted. So please don’t use it in production!
2. The second type is connection secured by server-side TLS. 
In this case, all the data is encrypted, but only the server needs to provide its TLS certificate to the client. 
You can use this type of connection if the server does not care which client is calling its API.
3. The third and strongest type is connection secured by mutual TLS. 
We use it when the server also needs to verify who’s calling its services. So in this case, both client and server must provide their TLS certificates to the other.



Execute following command in the unary/cal folder
% protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/*