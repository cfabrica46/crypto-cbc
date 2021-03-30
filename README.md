# crypto-CBC
## Comandos Necesarios

### Encriptar
~~~
openssl enc -aes-256-cbc -nosalt -e -a -A -in image.jpg -K '12345678901234567890123456789012' -iv '1234567890123456' -out image.enc
~~~

### Desencriptar
~~~
openssl enc -aes-256-cbc -nosalt -d -a -A -in image.enc -K '12345678901234567890123456789012' -iv '1234567890123456' -out decrypt.jpg
~~~

### Desencriptar (Alternativa con hexadecimal)
~~~
openssl enc -aes-256-cbc -nosalt -d -a -A -in image.enc -K '3132333435363738393031323334353637383930313233343536373839303132' -iv '31323334353637383930313233343536' -out decrypt.jpg
~~~

### Desencriptar (Alternativa con base64)
~~~
openssl enc -aes-256-cbc -nosalt -d -a -A -in image.enc -K 'MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=' -iv 'MTIzNDU2Nzg5MDEyMzQ1Ng==' -out decrypt.jpg
~~~
