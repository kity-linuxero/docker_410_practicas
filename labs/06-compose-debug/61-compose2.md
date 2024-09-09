# Laboratorio 6.1 - Docker Compose | Parte II

## Objetivo
Configurar un archivo `compose.yml` que utilice variables de entorno desde un archivo `.env`, interpolación de variables y una política de despliegue avanzada con `restart` y `resources`.


## 1. Crear el archivo `.env`:

- Crear el archivo .env en la raíz de tu proyecto para definir las variables de entorno con el siguiente contenido.

    ```env
    # Apps env
    APP_NAME=miapp
    APP_PORT=8080

    # DB env
    DB_USER=user
    DB_PASS=q3BFUyMZvL8CbkzYgxGut4Es
    ```

## 2. Crear archivo `compose.yml`:

- Crear un archivo `compose.yml` en la carpeta raíz del proyecto con el siguiente contenido

    ```compose
    services:
        web:
            image: nginx:alpine
            container_name: ${APP_NAME}_web
            ports:
            - "${APP_PORT}:80"
            env_file: ".env"
            deploy:
                restart_policy:
                    condition: on-failure
                    max_attempts: 3
                resources:
                    limits:
                      cpus: "0.5"
                      memory: 128M
                    reservations:
                      cpus: "0.25"
                      memory: 64M
        db:
            image: postgres
            env_file: ".env"
            environment:
                POSTGRES_USER: ${DB_USER}
                POSTGRES_PASSWORD: ${DB_PASS}
            depends_on:
                - web
    ```
### Explicación de los conceptos:

- **Interpolation:**
    - `${APP_NAME}`: se usa para asignar un nombre dinámico al contenedor.
    - `${APP_PORT}`: se usa para utilizar un puerto del que escucha el host. Se cambia desde el `.env`.
    - `${DB_USER}`: Es utilizado para asignar el usuario de la base de datos.
    - `${DB_PASS}`: Es la contraseña que se establecerá en la base de datos.
- **Restart Policy:** Define que el contenedor se reinicie en caso de fallo, con un máximo de tres intentos.
- **Resources:** Limita el uso del CPU y memoria del contenedor y reserva recursos mínimos.

## 3. Levantar los servicios:

- Desde la terminal ejecutamos:

    ```bash
    docker-compose up -d
    ```


## 4. Verificar la interpolación y configuración:

```bash
docker ps
docker inspect web
```

- Observar:

```
"PortBindings": {
    "80/tcp": [
        {
            "HostIp": "",
            "HostPort": "8080"
        }
    ]
},
"RestartPolicy": {
    "Name": "on-failure",
    "MaximumRetryCount": 3
},

"Memory": 134217728, # 128MiB
"MemoryReservation": 67108864, #64MiB
"NanoCpus": 500000000, # 0.5

```


## 5. Verificar el límite de recursos

- Verificar las estadísticas con `docker stats`

```
CONTAINER ID   NAME        CPU %     MEM USAGE / LIMIT     MEM %     NET I/O       BLOCK I/O       PIDS
f97f61992451   lab6-db-1   0.46%     40.92MiB / 15.51GiB   0.26%     16.7kB / 0B   19MB / 53.5MB   6
64e8079e2ead   miapp_web   0.00%     4.309MiB / 128MiB     3.37%     17.1kB / 0B   0B / 12.3kB     5

```

- Observar límite en memoria en `miapp_web`.

## 6. Modificar el archivo `.env`

Modificar la clave `APP_NAME` y setearla en `miapp2`.

## 7. Levantar los servicios

- Ejecutar nuevamente:

```bash
docker compose up -d

[+] Running 3/3
 ✔ Container miapp_web  Recreated                                  0.3s 
 ✔ Container lab6-db-1   Started                                   0.7s 
 ✔ Container miapp2_web  Started                                   0.2s
```

- Verificar que haya cambiado el nombre del contenedor:

```bash
docker compose ps

NAME        IMAGE         COMMAND           SERVICE  CREATED          STATUS          PORTS
lab6-db-1   postgres      "docker-entr…"  db       49 seconds ago   Up 47 seconds   5432/tcp
miapp2_web  nginx:alpine  "/docker-ent…"  web      49 seconds ago   Up 48 seconds   0.0.0.0:8080->80/tcp

```


## 8. Detener los servicios

- Detenga los servicios con:

    ```bash
    docker-compose down
    ```


--------