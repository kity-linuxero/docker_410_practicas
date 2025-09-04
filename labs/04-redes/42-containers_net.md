# Laboratorio 4.2 - Conexión de Múltiples Contenedores en una Red Bridge

## Objetivo
Crear y conectar múltiples contenedores en una red bridge para que puedan comunicarse entre sí, utilizando nombres de contenedor y comprobando la redirección de puertos.


## 1. Crear una red bridge:

- Crea una red bridge llamada `red_app`
- Para eso haremos uso del comando `docker network create`

    ```bash
    docker network create red_app
    ```

## 2. Crear dos contenedores conectados a la misma red:

- Crea dos contenedores basados en alpine y nginx, nombrados app1 y app2 respectivamente, ambos conectados a la red `red_app`:

    ```bash
    docker run -d --rm -it --name app1 --network red_app alpine
    docker run -d --rm --name app2 --network red_app nginx
    ```
- Verificar que los contenedores estén corriendo:
    ```bash
    docker ps
    ```

## 3. Comprobar la conectividad entre contenedores:

- Accede al contenedor web1 e intenta hacer ping al contenedor web2 utilizando su nombre.
- Para acceder a un contenedor que está corriendo utilizaremos el comando `docker exec`. <a href="https://docs.docker.com/reference/cli/docker/container/exec/" target="_blank">Doc Ref</a>
    ```
    docker exec -it app1 ping -c 3 app2
    ```
- ¿El resultado del ping fue exitoso?
- ¿Cuál es la IP interna de app2?

## 4. Inspeccionar la red

- Inspecciona red_app para ver cómo están configurados los contenedores en la red:

    ```bash
    docker network inspect red_app
    ```
- ¿Podés obtener las ips internas de los contenedores?


## 5. Accedemos por web al contenedor nginx

- Accedamos al contenedor nginx vía web.
- Haz clic <a href="http://localhost" target="_blank">aquí</a> para acceder al nginx vía web o bien accede desde tu navegador a `http://localhost`.
- ¿Qué ha pasado?... ¿Por qué no se puede acceder al contenedor?

>Pista: No hemos publicado los puertos al host.

- ¿Será posible acceder desde la red `red_app` al puerto 80 del nginx?
- Nos conectamos en una sesión `sh` interactiva al contenedor `app1` con el siguiente comando:
    ```bash
    docker exec -it app1 sh
    ```

- Instalemos el paquete `curl`

    ```bash
    apk add curl
    ```
- Accedamos por web al contenedor nginx:
    ```bash
    curl app2
    ```
- ¿Qué ha pasado?. ¿Fué exitosa la conexión?
- Salimos del contenedor:
    ```sh
    exit
    ```



## 6. Eliminar los contenedores y la red:

- Detén y elimina ambos contenedores, y luego elimina la red:

    ```bash
    docker stop app1 app2
    docker network rm red_app
    ```

--------

<p align="center">
  <img src="../../../img/logos.footer.gray.webp">
</p>