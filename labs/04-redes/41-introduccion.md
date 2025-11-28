# Laboratorio 4.1 - Introducción a redes en Docker

## Objetivo
Familiarizarse con la creación de redes bridge, la redirección de puertos, y la conexión de contenedores a estas redes.



## 1. Crear una red bridge:

- Crea una red bridge llamada `mi_red_bridge`
- Para eso haremos uso del comando `docker network create`

    ```bash
    docker network create mi_red_bridge
    ```

## 2. Listar redes

- Lista todas las redes disponibles para asegurarte de que mi_red_bridge ha sido creada:

    ```bash
    docker network ls
    ```

## 3. Crear un contenedor web

- Ejecutaremos un contenedor basado en la imágen `docker/welcome-to-docker`.
- Redireccionaremos los puertos con el `8080` del host hacia el `80` del contenedor.

    ```
    docker run -d --rm --name mi_contenedor -p 8080:80 --network mi_red_bridge docker/welcome-to-docker
    ```

- Haz clic <a href="http://localhost:8080" target="_blank">aquí</a> verificar el correcto funcionamiento o bien accede desde tu navegador a `http://localhost:8080`.


## 4. Inspeccionar la red

- Inspecciona mi_red_bridge para ver sus detalles y comprender mejor su configuración:

    ```bash
    docker network inspect mi_red_bridge
    ```
- ¿Podés obtener la subred y el gateway con la información brindada?

- Inspeccionemos el contenedor:
    ```bash
    docker inspect mi_contenedor
    ```
- ¿Podés obtener la IP interna del contenedor `mi_contenedor` con la información brindada?
    - 🗒️Pista: Buscar en la clave `Networks`.


## 5. Desconectar el contenedor de la red

- Desconecta el contenedor mi_contenedor de la red mi_red_bridge:

    ```bash
    docker network disconnect mi_red_bridge mi_contenedor
    ```

- ¿Es posible acceder mediante red al contenedor?
- Haz clic <a href="http://localhost:8080" target="_blank">aquí</a> para comprobarlo o bien accede desde tu navegador a `http://localhost:8080`.


## 6. Reconectar el contenedor a la red:

- Vuelve a conectar el contenedor a la red:

    ```bash
    docker network connect mi_red_bridge mi_contenedor
    ```

## 7. Eliminar la red:

- Detén el contenedor y elimina la red:

    ```bash
    docker stop mi_contenedor
    docker network rm mi_red_bridge
    ```

--------

<p align="center">
  <img src="../../img/logos.footer.gray.webp">
</p>