# Laboratorio 1.2 - Primeros contenedores y comandos Linux

### Objetivos:

- Correr nuestro primer contenedor Linux.
- Ejecución de primeros comandos sobre un contenedor Linux.

## Contenedores Linux

### 1. Ejecutar nuestro primer contenedor Linux

Vamos a ejecutar un contenedor de <a href="https://www.alpinelinux.org/" target="_blank">Alpine Linux</a>, una distribución súper liviana que está creciendo en popularidad. Para eso ejecutemos el siguiente comando:

```bash
docker run -it alpine sh
```

> [!TIP]
> El comando `docker run` se utiliza para correr una imágen de Docker en un contenedor. El parámetro `-it` inicia una sesión interactiva. Se verá en mayor detalle mas adelante en el curso.

Luego que la imágen se haya descargado, tendremos un promt de terminal de sh de la siguiente manera:

```
/#
```

Ya tenemos un contenedor Linux corriendo. Lo podemos usar como si de una VM se tratara. Solo para probar, vamos a instalar algunos paquetes.

Ejecutamos el siguiente comando para verificar si TCP/IP está OK:

```bash
ping -c 3 localhost
```

```
# ping -c 3 localhost
PING localhost (::1) 56 data bytes
64 bytes from localhost (::1): icmp_seq=1 ttl=64 time=0.028 ms
64 bytes from localhost (::1): icmp_seq=2 ttl=64 time=0.025 ms
64 bytes from localhost (::1): icmp_seq=3 ttl=64 time=0.022 ms
```


### 2. Instalación de paquetes

Instalaremos paquetes dentro del contenedor para ampliar sus capacidades, en este caso instalaremos el paquete `curl`.

> [!NOTE]
> **cURL**, que significa 'Client for URLs', es una herramienta de línea de comando utilizada para transferir datos con URLs.


```sh
# Actualizamos la base de datos de los repositorios
apk update

# Instalamos el paquete curl
apk add curl
```
Una vez que haya terminado la secuencia de instalación 
Probamos utilizar `curl` de la siguiente manera para saber nuestra ip pública.

```sh
curl ifconfig.co
```

### 3. Comandos de GNU/Linux 

Como nos encontramos en una terminal de Linux, es este caso es `sh` pero podría ser `bash`, podemos ejecutar comandos en el intérprete del contenedor. A continuación tiene una lista de comandos básicos para probar dentro de una terminal Linux. Pruebe ejecutar alguno de los siguientes comandos para familiarizarse con la terminal:

- `pwd`: Muestra el directorio de trabajo actual.
- `ls`: Lista los archivos y directorios en el directorio actual.
- `whoami`: Muestra el nombre del usuario actual.
- `uname -a`: Muestra información del sistema operativo.
- `cat /etc/os-release`: Muestra información sobre la distribución de Linux.
- `df -h`: Muestra el uso del espacio en disco en formato legible.
- `free -h`: Muestra la cantidad de memoria libre y utilizada en el sistema.
- `top`: Muestra una vista dinámica en tiempo real de los procesos en ejecución. Para salir, presione `Control+C`
- `apk update`:  (Alpine Linux) Actualiza la lista de paquetes disponibles.
- `apk add [paquete]`:  (Alpine Linux) Instala un paquete específico (reemplaza [paquete] con el nombre del paquete que deseas instalar).

- `apt update`: (Debian, Ubuntu) Actualiza la lista de paquetes disponibles (Debian, Ubuntu).
- `app install -y [paquete]`: (Debian, Ubuntu) Instala un paquete específico (reemplaza [paquete] con el nombre del paquete que deseas instalar).
- `exit`: Sale de la terminal. En este caso se terminará el contenedor.

Los comandos les será util cuando tenga que depurar un contenedor.

### 4. Salir del contenedor

Para salir del contenedor y detener su ejecución:

```sh
exit
```


---------

<p align="center">
  <img src="https://docker.idepba.com.ar/img/logos/logos.footer.gray.webp">
</p>
