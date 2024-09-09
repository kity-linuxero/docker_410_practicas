# Laboratorio 6.2 - Verificación de recursos
## Objetivo
Utilizar `compose` para definir recursos limitados para un servicio. Poner a prueba los recursos limitados.


## 1. Copiamos el siguiente archivo compose:

- Creamos un archivo `compose.yml` con el siguiente contenido. Definiendo políticas de deploy con límites y reservas y definimos la red "red_interna".

```compose
services:
    web:
        image: docker/welcome-to-docker
        ports:
        - "8080:80"

        networks:
            - red_interna

        deploy:
            restart_policy:
                condition: on-failure
                max_attempts: 3
            resources:
                limits:
                  cpus: "0.02"
                  memory: 32M
                reservations:
                  cpus: "0.01"
                  memory: 16M

networks:
    red_interna:
```

## 2. Levantamos el servicio `web`:

- Desde la terminal ejecutamos:

    ```bash
    docker-compose -p curso up -d
    ```
    - El `-p` adicional es para indicar el nombre del proyecto. Por default el nombre del proyecto será la carpeta raíz.
- Verifiquemos el contenedor esté arriba:
    ```bash
    docker compose ps
    ```
- Abrimos un navegador en la URL http://localhost:8080
- Deberíamos ver una página de felicitaciones por correr nuestro primer contenedor en Docker.


## 3. Realizar prueba de stress

- Según el archivo `compose.yml`, hemos reservado `0.01` de CPU y puesto límite en `0.02`. Eso indica que máximo, el contenedor debe utilizar un `2%` del CPU.

- Abrimos una terminal nueva y escribimos `docker stats` para ir observando el uso de recursos.

- En otra terminal ejecutamos el siguiente comando, pero antes [vamos a analizar qué hace](#analizando-el-comando):

    ```bash
    docker run --rm -it --name linux --network curso_red_interna alpine sh -c "apk add curl && while true; do curl http://web; done"
    ```

### Analizando el comando

- `docker run` : Ejecuta un contenedor con imágen `alpine`
- `--rm`: El contenedor se elimina luego de terminar su ejecución.
- `--it`: Una terminal interactiva en modo TTY para que nos muestre la salida por pantalla y podamos detener el comando.
- `--name`: Es el nombre que tendrá el contenedor, en este caso `linux`.
- `--network`: Indica que se va a conectar a la misma red en el que se encuentra el contenedor `web`.
- `sh -c`: Es la shell por defecto de Alpine. El `-c` adicional es para indicar que le vamos a mandar parámetros.
- `apk add curl`: Es la forma que tiene Alpine para indicar la instalación de paquetes, en este caso instalamos `curl`.
- `while true; do curl http://web; done`: Es un pequeño script que está en loop infinito consultando la url http://web.
    - Observar que, al estar conectado a la misma red que el contenedor que sirve la página, pueden verse sin necesidad de conocer la IP.
    - La idea del script en loop infinito es consumir recursos del contenedor para verificar los límites.

- La terminal con `docker stats` debería dar un resultado similar al siguiente, donde el CPU de `curso-web` no debería pasar el 2% o si lo hace es solo por unos momentos que tendrá un pico de consumo para luego volver a bajar.

```bash
CONTAINER ID   NAME          CPU %     MEM USAGE / LIMIT     MEM %     NET I/O          BLOCK I/O  
bce229c6ff9a   curso-web-1   1.94%     4.395MiB / 32MiB      13.73%    715kB / 1.98MB   0B / 12.3kB
052d08a66a5c   linux         40.90%    7.656MiB / 15.51GiB   0.05%     5.43MB / 328kB   0B / 8.07MB
```

## 4. Detener prueba de stress y aumentar el límite de CPU

- En la terminal donde ejecutó el Linux Alpine, presione `Control+C`.
    - Debería haberse detenido el script.

- Modifique el archivo `compose.yml` para que el límite del CPU del servicio `web` esté en `1.0`.

- Vuelva a levantar los servicios. El contenedor del servicio `web` debería ser reconstruído: 

    ```bash
    docker-compose -p curso up -d
    ```

- Volvamos a ejecutar la prueba de stress con:

    ```bash
    docker run --rm -it --name linux --network curso_red_interna alpine sh -c "apk add curl && while true; do curl http://web; done"
    ```
- En la terminal donde está corriendo `docker stats` el porcentaje utilizado por `curso-web` debería ser mayor. 

```bash
CONTAINER ID   NAME          CPU %     MEM USAGE / LIMIT     MEM %     NET I/O          BLOCK I/O  
fd2e2c172856   curso-web-1   3.11%     4.371MiB / 32MiB      13.66%    331kB / 749kB    0B / 12.3kB
038b020396ed   linux         70.56%    7.516MiB / 15.51GiB   0.05%     5.47MB / 346kB   0B / 8.07MB
```

## 5. Terminar lab

- En la terminal donde se ejecuta Alpine Linux terminar el proceso con `Control+C`.
- En la terminal donde se ejecuta `docker stats` terminal el proceso con `Control+C`.
- Ejecutar el comando `docker compose down` para detener el contenedor `web` y eliminar las redes creadas.

--------