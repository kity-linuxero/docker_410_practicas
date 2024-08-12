# Laboratorio 1.2 - Primeros comandos

### Objetivos:

- Ejecución de nuestro primer contenedor GNU/Linux.
- Ejecutar algunos comandos básicos de GNU/Linux y Docker

## 1. Contenedor Ubuntu

Vamos a ejecutar un contenedor de la distribución de GNU/Linux Ubuntu. Para eso ejecutemos el siguiente comando:

```bash
docker run -it ubuntu bash
```

Luego que la imágen se haya descargado, tendremos un promt de terminal de bash de la siguiente manera:

```
root@d9080b400c84:/#
```

Ya tenemos un contenedor Ubuntu corriendo. Lo podemos usar como si de una VM se tratara. Solo para probar, vamos a instalar algunos paquetes.

```bash
apt update
apt install -y iputils-ping
```

Una vez que haya terminado la secuencia de instalación ejecutamos:

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

### Comandos de GNU/Linux 

Como nos encontramos en una terminal `bash`, podemos ejecutar comandos en el intérprete del contenedor. A continuación tiene una lista de comandos básicos para probar dentro de una terminal de Ubuntu.

- `pwd`: Muestra el directorio de trabajo actual.
- `ls`: Lista los archivos y directorios en el directorio actual.
- `whoami`: Muestra el nombre del usuario actual.
- `uname -a`: Muestra información del sistema operativo.
- `cat /etc/os-release`: Muestra información sobre la distribución de Linux.
- `df -h`: Muestra el uso del espacio en disco en formato legible.
- `free -h`: Muestra la cantidad de memoria libre y utilizada en el sistema.
- `top`: Muestra una vista dinámica en tiempo real de los procesos en ejecución. Para salir, presione `Control+C`
- `apt update`: Actualiza la lista de paquetes disponibles.
- `apt install -y [paquete]`: Instala un paquete específico (reemplaza [paquete] con el nombre del paquete que deseas instalar).
- `exit`: Sale de la terminal. En este caso se terminará el contenedor.

Los comandos les será util cuando tenga que depurar un contenedor "en caliente".

#### Salir del contenedor

Ejecutamos:

```bash
exit
```

## 2. Primeros comandos de Docker

A continuación vamos a ejercitar en comandos básicos para ir practicando interacción con la Docker CLI.

### Iniciando otro contenedor Ubuntu

Vamos a iniciar un contenedor Ubuntu de la siguiente manera

```bash
docker run -d ubuntu sleep 10
```

Analicemos:
- `docker run`: Ejecutará un contenedor de una imágen especificada, en este caso `ubuntu`.
- `-d` ó `--detach`: Ejecuta un contenedor en background. Imprime el container ID.
- `sleep 10`: Es un comando que queremos enviar al contenedor. `sleep` hace que quede en espera una cantidad determinada de segundos, como hemos enviado el parámetro `10` quedará ejecutando en espera 10 segundos.

Luego de ejecutar el comando, antes de los 10 segundos ejecutamos `docker ps`:

```bash
docker ps
CONTAINER ID   IMAGE     COMMAND      CREATED         STATUS         PORTS     NAMES
56bdfd527c78   ubuntu    "sleep 10"   4 seconds ago   Up 3 seconds             intelligent_wilson
```

El comando nos devuelve información de los contenedores en ejecución:
- CONTAINER ID
- IMAGE
- COMMAND
- CREATED
- STATUS
- PORTS
- NAMES

Si luego de los 10 segundos, volvemos a ejecutar `docker ps`:

```bash
docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
# Ninguna información
```

No mostrará nada, porque el contenedor ya no está en ejecución.

Para ver información de TODOS los contenedores, `docker ps -a`:

```bash
docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED         STATUS                     PORTS    NAMES
56bdfd527c78   ubuntu   "sleep 10"  2 minutes ago   Exited (0) 2 minutes ago            intelligent_wilson
# Observar info status
```

### Ejecutando comandos en el contenedor

Como podemos deducir, el último de los parámetros enviados en el comando `docker run` es la órden que ejecutará el contenedor.
Así, podemos ejecutar cualquiera que se nos ocurra, por ejemplo `docker run ubuntu ls -l`

```bash
$ docker run ubuntu ls -l
total 48
lrwxrwxrwx   1 root root    7 Apr 22 13:08 bin -> usr/bin
drwxr-xr-x   2 root root 4096 Apr 22 13:08 boot
drwxr-xr-x   5 root root  340 Aug 12 02:23 dev
drwxr-xr-x   1 root root 4096 Aug 12 02:23 etc
drwxr-xr-x   3 root root 4096 Jun  5 02:06 home
lrwxrwxrwx   1 root root    7 Apr 22 13:08 lib -> usr/lib
lrwxrwxrwx   1 root root    9 Apr 22 13:08 lib64 -> usr/lib64
drwxr-xr-x   2 root root 4096 Jun  5 02:02 media
drwxr-xr-x   2 root root 4096 Jun  5 02:02 mnt
drwxr-xr-x   2 root root 4096 Jun  5 02:02 opt
dr-xr-xr-x 370 root root    0 Aug 12 02:23 proc
drwx------   2 root root 4096 Jun  5 02:05 root
drwxr-xr-x   4 root root 4096 Jun  5 02:06 run
lrwxrwxrwx   1 root root    8 Apr 22 13:08 sbin -> usr/sbin
drwxr-xr-x   2 root root 4096 Jun  5 02:02 srv
dr-xr-xr-x  13 root root    0 Aug 12 02:23 sys
drwxrwxrwt   2 root root 4096 Jun  5 02:05 tmp
drwxr-xr-x  12 root root 4096 Jun  5 02:02 usr
drwxr-xr-x  11 root root 4096 Jun  5 02:05 var
```
Esto creará un contenedor nuevo partiendo de la imágen `ubuntu`, ejecutará el comando `ls -l` y luego sale del contenedor. El contenedor habrá finalizado su ejecución.

### Eliminamos el contenedor

Escribamos, `docker rm <CONTAINER ID>` y luego `docker ps -a`:

```bash
docker rm 56bdfd527c78
56bdfd527c78

docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED             STATUS                          PORTS    NAMES
f89187d1d5b4   ubuntu   "ls -l"     About a minute ago  Exited (0) About a minute ago            vibrant_jones
# El contenedor con ID 56bdfd527c78 se habrá eliminado
```

Eliminemos el último contenedor: 

```bash
docker rm f89187d1d5b4
f89187d1d5b4

docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED         STATUS                     PORTS    NAMES
# No debería haber ningún contenedor.
```

---------
