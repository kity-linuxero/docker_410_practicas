# Laboratorio 1.2 - Primeros contenedores y comandos Linux

### Objetivos:

- Correr nuestro primer contenedor `hello world`.
- Entender que sucede cuando corremos un contenedor.
- Correr nuestro primer contenedor Linux.
- Ejecución de primeros comandos sobre un contenedor Linux.

## Hola mundo. Nuestro primer contenedor

Como en todo lenguaje de programación, lo que primero hacemos es el clásico `hello world` y en Docker haremos lo mismo.

En Windows, abrimos una ventana de PowerShell y escribimos lo siguiente:


```powershell
docker run hello-world
```

Observamos la salida de la terminal:

```bash
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
c1ec31eb5944: Pull complete 
Digest: sha256:1408fec50309afee38f3535383f5b09419e6dc0925bc69891e79d84cc4cdcec6
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/

```

### Entendiendo la salida de la terminal

Si la salida de la terminal es como la que se indica arriba, la instalación de Docker fue correcta. `hello-world` es una imágen que se usa para corroborar que Docker esté funcionando. En el texto de salida podemos observar lo siguiente:

```bash
...
To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

    To try something more ambitious, you can run an Ubuntu container with:
    
    $ docker run -it ubuntu bash
...
```

Lo que indicá que:
1. El cliente Docker se contactó con el daemon de Docker.
2. El daemon descargó (pulled) la imágen `hello-world` desde <a href="https://hub.docker.com/" target="_blank">Docker Hub</a>
3. El daemon de Docker creó un nuevo contenedor a partir de esa imagen que ejecuta el ejecutable que produce la salida que estábamos leyendo.
4. El daemon de Docker transmitió esa salida al cliente de Docker, que la envió a la terminal.



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

Ya tenemos un contenedor Linux corriendo. Lo podemos usar como si de una VM se tratara. 

Ejecutamos el siguiente comando para verificar conectividad y si tenemos red:

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

Como nos encontramos en una terminal Linux, es este caso es `sh` pero podría ser `bash`, podemos ejecutar comandos en el intérprete del contenedor. A continuación tiene una lista de comandos básicos para probar dentro de una terminal Linux. Pruebe ejecutar alguno de los siguientes comandos para familiarizarse con la terminal:

- `pwd`: Muestra el directorio de trabajo actual.
- `ls`: Lista los archivos y directorios en el directorio actual.
- `whoami`: Muestra el nombre del usuario actual.
- `uname -r`: Muestra versión del kernel del sistema operativo.
- `cat /etc/os-release`: Muestra información sobre la distribución de Linux.
- `df -h`: Muestra el uso del espacio en disco en formato legible.
- `free -h`: Muestra la cantidad de memoria libre y utilizada en el sistema.
- `top`: Muestra una vista dinámica en tiempo real de los procesos en ejecución. Para salir, presione `Control+C`
- `apk update`:  (Alpine Linux) Actualiza la lista de paquetes disponibles.
- `apk add [paquete]`:  (Alpine Linux) Instala un paquete específico (reemplaza [paquete] con el nombre del paquete que deseas instalar).
- `exit`: Sale de la terminal. En este caso se terminará el contenedor.

Los comandos les será util cuando tenga que depurar un contenedor.

### 4. Salir del contenedor

Para salir del contenedor y detener su ejecución:

```sh
exit
```

## Resumen

En este laboratorio práctico se abordaron los siguientes puntos clave:
- **Primer contenedor `hello-world`:** Ejecución de nuestro primer contenedor de prueba para verificar la correcta instalación y funcionamiento de Docker, comprendiendo los pasos del cliente y del daemon (descarga de imagen, creación y ejecución).
- **Contenedor interactivo Alpine Linux:** Inicio de una sesión interactiva en una distribución liviana utilizando `docker run -it alpine sh`.
- **Gestión de paquetes:** Utilización de `apk update` y `apk add curl` para instalar herramientas adicionales dentro del contenedor.
- **Comandos básicos de GNU/Linux:** Ejecución de comandos del sistema (`pwd`, `ls`, `whoami`, `uname`, `cat`, `df`, `free`, `top`) útiles para explorar y diagnosticar el entorno del contenedor.
- **Ciclo de vida:** Salida y detención del contenedor de manera segura a través de la instrucción `exit`.

---------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="../../img/logos.footer.gray.webp">
  </a>
</p>
