# Laboratorio 5.2 - Aplicaciones multicontenedor

## Objetivo
- Comprobar la sencillez de docker compose en aplicaciones multicontenedor en contraste de ejecutar con docker run.


### Explicación 

Ejecutar aplicaciones en un solo contenedor es sencillo cuando se trata de tareas simples, como un script de Python o una aplicación Node.js que sirva un sitio estático. Sin embargo, a medida que las aplicaciones crecen, manejar varios contenedores de forma individual se vuelve más complicado. Por ejemplo, si el script necesita conectarse a una base de datos o requiere autenticación, el contenedor se vuelve más complejo y grande.

Una de las mejores prácticas para contenedores es que cada uno haga una sola tarea y la haga bien. Aunque existen excepciones, es importante evitar que un contenedor maneje múltiples procesos.

Cuando tenemos aplicaciones que dependen entre sí, manejar múltiples contenedores con comandos docker run se vuelve propenso a errores y complicado de escalar, especialmente si necesitamos mantener el orden de arranque y manejar configuraciones de red.

Ahí es donde entra Docker Compose. Este permite definir toda tu aplicación multicontenedor en un archivo YAML (compose.yml), especificando configuraciones, variables de entorno, volúmenes y redes de manera centralizada. Algunas de sus ventajas incluyen:

- No necesitas ejecutar múltiples comandos docker run, simplificando la gestión de contenedores.
- Puedes ejecutar contenedores en un orden específico y manejar las conexiones de red.
- Facilita la escalabilidad de servicios individuales.
- Implementar volúmenes persistentes es sencillo.
- Permite configurar variables de entorno de forma centralizada.

Docker Compose hace que gestionar aplicaciones complejas sea modular, escalable y consistente.

En esta guía práctica, primero verás cómo construir y ejecutar una aplicación web de contadores basada en Node.js, un proxy inverso Nginx y una base de datos Redis utilizando los comandos docker run. También verás cómo puedes simplificar todo el proceso de despliegue usando Docker Compose.

# Parte I - Sin Docker Compose

## 1. Obtener la aplicación


- Para este ejercicio, el desarrollador nos provee de su repositorio Git donde tiene alojada la aplicación. Si tenemos git instalado ejecutamos:

    ```bash
    git clone https://github.com/dockersamples/nginx-node-redis
    ```

- Otra alternativa, podemos obtener el archivo <a href="https://github.com/dockersamples/nginx-node-redis/archive/refs/heads/main.zip" download>ZIP</a> del repositorio.

- Una vez descargada y descomprimida, nos ubicamos en el directorio de la aplicación `nginx-node-redis`.

- La estructura de directorios será la siguiente:

```
└── nginx-node-redis
    ├── nginx
    │   ├── Dockerfile
    │   └── nginx.conf
    ├── web
    │   ├── Dockerfile
    │   ├── package.json
    │   ├── package-lock.json
    │   └── server.js
    └── compose.yml

```

## 2. Buildear imágenes

- Navegar a la carpeta `nginx` para buildear la imágen:

    ```bash
    docker build -t nginx .
    ```

- Navegar a la carpeta `web` para buildear la imágen:

    ```bash
    docker build -t web .
    ```

## 3. Crear redes

- Antes de correr una aplicación multi-contenedor necesitamos crear redes para que los contenedores se comuniquen.

    ```bash
    docker network create sample-app
    ```



## 4. Correr los contenedores

#### Base de datos

- Iniciaremos el contenedor Redis ejecutando el siguiente comando, que lo adjuntará a la red creada previamente y creará un alias de red (dns lockup)

    ```bash
    docker run -d  --name redis --network sample-app --network-alias redis redis
    ```

#### Servidor web 1 

- Iniciemos el servidor web primario con el siguiente comando:

    ```bash
    docker run -d --name web1 -h web1 --network sample-app --network-alias web1 web
    ```

#### Servidor web 2

- Iniciemos el servidor web secundario con el siguiente comando:

    ```bash
    docker run -d --name web2 -h web2 --network sample-app --network-alias web2 web
    ```

#### Servidor NGINX (Proxy reverso)

- Iniciaremos el contenedor con el proxy reverso Nginx con el siguiente comando:

    ```bash
    docker run -d --name nginx --network sample-app  -p 80:80 nginx
    ```

Este contenedor actúa como proxy reverso, enrutando el tráfico a otro servidores. En este caso dirigirá tráfico a los contenedores de backend `web1`y `web2`.

## 5. Verificar que los contenedores estén corriendo

- Verificamos que los contenedores estén corriendo con el siguiente comando:

```bash
docker ps
```

Deberíamos ver algo similar a esto:

```bash
CONTAINER ID   IMAGE     COMMAND                  CREATED              STATUS              PORTS                NAMES
2cf7c484c144   nginx     "/docker-entrypoint...."   9 seconds ago        Up 8 seconds        0.0.0.0:80->80/tcp   nginx
7a070c9ffeaa   web       "docker-entrypoint.s..."   19 seconds ago       Up 18 seconds                            web2
6dc6d4e60aaf   web       "docker-entrypoint.s..."   34 seconds ago       Up 33 seconds                            web1
008e0ecf4f36   redis     "docker-entrypoint.s..."   About a minute ago   Up About a minute   6379/tcp             redis
```

- Si verificamos en Docker Desktop Dashboard deberíamos ver algo similar a la siguiente imágen:

![](https://docs.docker.com/get-started/docker-concepts/running-containers/images/multi-container-apps.webp)

#### Verificar correcto funcionamiento:

Si todo funciona bien, deberíamos poder acceder desde nuestro navegador a http://localhost donde podremos visualizar la cantidad de visitas y qué servidor atendió nuestra solicitud.

```bash
web2: Number of visits is: 9
web1: Number of visits is: 10
web2: Number of visits is: 11
web1: Number of visits is: 12
```

## 6. Detenemos los contenedores

- Detener y borrar todos los contenedores con el siguiente comando:

    ```bash
    docker rm -f nginx web1 web2 redis
    ```
- Eliminamos la red:

    ```bash
    docker network rm sample-app
    ```


# Parte II - Utilizando Docker Compose

En este paso, simplificaremos todo el despliegue de una aplicación multi-contenedor utilizando Docker Compose. 


## 1. Archivo Compose

- En la carpeta raíz del proyecto hay un archivo llamado `compose.yml` con el siguiente contenido:

    ```dokcer-compose
    services:
    redis:
        image: redis
        ports:
        - '6379:6379'
    web1:
        restart: on-failure
        build: ./web
        hostname: web1
        ports:
        - '81:5000'
    web2:
        restart: on-failure
        build: ./web
        hostname: web2
        ports:
        - '82:5000'
    nginx:
        build: ./nginx
        ports:
        - '80:80'
        depends_on:
        - web1
        - web2
    ```

- Analice el archivo compose.

## 2. Corriendo aplicación con docker compose:

Ejecute el siguiente comando:

```bash
docker compose up -d --build
```

Verá la salida del comando como `compose` empieza a construir las imágenes, levantar los contenedores y creando las redes necesarias.

```bash
Running 5/5
✔ Network nginx-nodejs-redis_default    Created                                                0.0s
✔ Container nginx-nodejs-redis-web1-1   Started                                                0.1s
✔ Container nginx-nodejs-redis-redis-1  Started                                                0.1s
✔ Container nginx-nodejs-redis-web2-1   Started                                                0.1s
✔ Container nginx-nodejs-redis-nginx-1  Started
```

### 3. Verificando el correcto funcionamiento

Acceda desde su navegador a http://localhost.

Debería ver la misma página que antes.

### 4. Eliminando los contenedores

- Desde la terminal ejecute el siguiente comando:

    ```bash
    docker compose down
    ```

- Debería ver en la salida de pantalla como se eliminan los contenedores y las redes creadas.

    ```bash
    [+] Running 5/5
    ✔ Container nginx-node-redis-redis-1  Removed                                                                                    0.3s 
    ✔ Container nginx-node-redis-nginx-1  Removed                                                                                    0.3s 
    ✔ Container nginx-node-redis-web2-1   Removed                                                                                    0.8s 
    ✔ Container nginx-node-redis-web1-1   Removed                                                                                    0.8s 
    ✔ Network nginx-node-redis_default    Removed  
    ```


### Referencias:

- [Docker Docs: Multi-container applications](https://docs.docker.com/get-started/docker-concepts/running-containers/multi-container-applications/)


--------