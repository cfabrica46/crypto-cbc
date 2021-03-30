# crypto-CBC
## Comandos Necesarios

### Encriptar
~~~
openssl enc -aes-256-cbc -nosalt -e -a -A -in image.jpg -K '12345678901234567890123456789012' -iv '1234567890123456' -out image.jpg.enc
~~~

### Desencriptar
~~~
openssl enc -aes-256-cbc -nosalt -d -a -A -in image.jpg.enc -K '12345678901234567890123456789012' -iv '1234567890123456' -out decrypt.jpg
~~~
