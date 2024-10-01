# Laboratorio 1.1 - Instalación de Docker

### Objetivos:

- Correcta instalación de Docker en nuestra PC de prueba.

## 1. Instalación Docker Desktop

Docker Desktop se describe como _una aplicación de instalación con un solo clic para su entorno Mac, Linux* o Windows que le permite crear, compartir y ejecutar aplicaciones y microservicios en contenedores_. Nos permitirá una instalación sencilla en entornos Windows.

* [No compatible con Ubuntu 24.04](https://docs.docker.com/desktop/install/ubuntu/#prerequisites). Para tal caso se recomienda la [siguiente instalación](https://docs.docker.com/engine/install/ubuntu/).

### Instalación en Windows

Asegurese de [cumplir con los requisitos del sistema](https://docs.docker.com/desktop/install/windows-install/#system-requirements).
- WSL version 1.1.3.0 o superior.
- Windows 11 64-bit: Home or Pro version 21H2 or higher, or Enterprise or Education version 21H2 or higher.
- Windows 10 64-bit:
    - We recommend Home or Pro 22H2 (build 19045) or higher, or Enterprise or Education 22H2 (build 19045) or higher.
    - Minimum required is Home or Pro 21H2 (build 19044) or higher, or Enterprise or Education 21H2 (build 19044) or higher.
- Procesador de 64 bits con [SLAT](https://en.wikipedia.org/wiki/Second_Level_Address_Translation)
- 4 GB RAM
- Virtualización por HW habilitada [Ver virtualización](https://docs.docker.com/desktop/troubleshoot/topics/#virtualization)


[Ver mas](https://docs.docker.com/desktop/install/windows-install/#system-requirements).


#### Descargue Docker Desktop <a href="https://hub.docker.com/" target="_blank">Docker Hub</a>


### Una vez descargado el ejecutable lo ejecutamos.
![](https://docker.idepba.com.ar/img/clase1/Docker_install.png)

### Aceptamos los términos

![](https://docker.idepba.com.ar/img/clase1/Docker_install2.png)


### Elegimos una opción

![](https://docker.idepba.com.ar/img/clase1/Docker_install3.png)

### Docker Instalado

![](https://docker.idepba.com.ar/img/clase1/Docker_install4.png)

### Elegimos una opción o bien "Skip survey"

![](https://docker.idepba.com.ar/img/clase1/Docker_install5.png)

### Iniciando Docker

Puede tardar en iniciar por primera vez el Docker Engine. En background se ejecutan tareas que utilizando <a href="https://learn.microsoft.com/es-es/windows/wsl/about" target="_blank">WSL</a>, descarga una imágen de Linux para poder correr Docker.

![](https://docker.idepba.com.ar/img/clase1/Docker_install6.png)

### Docker iniciado

![](https://docker.idepba.com.ar/img/clase1/Docker_install7.png)

:bulb: Observe que tiene disponible dos tutoriales para explorar Docker Desktop: How do I run a container? y Run Docker Hub Images. Puede verlos si desea.

### Abrimos una ventana de PowerShell:

Escribimos:

```powershell
docker --version
```

![](https://docker.idepba.com.ar/img/clase1/Docker_install8.png)

## 2. Hello World

Como en todo lenguaje de programación, lo que primero hacemos es el clásico `hello world` en Docker haremos lo mismo.

En una ventana de PowerShell ejecutaremos:

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

Y para probar algo mas ambicioso, podemos ejecutar un container Ubuntu. Lo haremos en el [Lab 1.2](./primeros-comandos.md)


## Conclusión

En este laboratorio instalamos Docker Desktop que nos permitió correr comandos con la CLI de Docker. Los mismos pasos podrían haberse realizado desde la interface gráfica, pero en este curso nos enfocaremos en la CLI de Docker.
Además, ejecutamos un contenedor `hello-world` que se utiliza para verificar el correcto funcionamiento del entorno Docker en nuestro entorno.

---------

<img src="../../img/footer.svg" alt="Descripción de la imagen" style="width:100%; filter: grayscale(100%); opacity: 40%">