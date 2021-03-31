# crypto-CBC
## Comandos Necesarios

### Encriptar
~~~
openssl enc -aes-256-cbc -nosalt -e -a -A -in image.jpg -K '3132333435363738393031323334353637383930313233343536373839303132' -iv '31323334353637383930313233343536' -out image.enc
~~~

### Desencriptar
~~~
openssl enc -aes-256-cbc -nosalt -d -in image.enc -K '3132333435363738393031323334353637383930313233343536373839303132' -iv '31323334353637383930313233343536' -out nuevo.jpg
~~~


