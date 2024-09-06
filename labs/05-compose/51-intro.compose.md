# Laboratorio 5.1 - Deployar un contenedor web sencillo con Docker Compose

## Objetivo
Utilizar Docker Compose para ejecutar un solo contenedor que sirva una página web estática con Nginx.


## 1. Crear el archivo `docker-compose.yml`:

- En un nuevo directorio, crea un archivo llamado docker-compose.yml que defina el contenedor de Nginx.

    ```bash
    services:
      web:
        image: nginx:alpine
        ports:
          - "8080:80"
        volumes:
          - ./html:/usr/share/nginx/html
    ```

Explicación `docker-compose.yml`:
- `services`: Acá declaramos el servicio web, que usará la imagen de nginx basada en Alpine Linux.
- `ports`: Mapea el puerto 8080 del host al puerto 80 del contenedor (donde Nginx escucha).
- `volumes`: Monta el directorio local ./html en el contenedor para que sirva los archivos web desde ese directorio.

## 2. Crear contenido HTML

- Crear un archivo HTML simple en un directorio html/ dentro de tu directorio de trabajo:

    ```html
    <!DOCTYPE html>
    <html lang="es">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>¡Hola Docker Compose!</title>
    </head>
    <body>
        <h1>¡Hola desde Docker Compose con Nginx!</h1>
    </body>
    </html>

    ```
Este archivo será servido por Nginx desde el contenedor.

## 3. Ejecutar Docker Compose

- Desde la terminal navegamos al directorio donde se encuentra el archivo `docker-compose.yml` y ejecutamos:

    ```bash
    docker-compose up -d
    ```

Este comando descargará la imagen de Nginx y levantará el contenedor con la configuración indicada. Además, por el parámetro `-d` se ejecutará en _daemon_.

## 4. Verificar que Funciona

- Abrimos nuestro navegador web en la dirección http://localhost:8080. Deberías ver el mensaje "¡Hola desde Docker Compose con Nginx!".



## 5. Detener el Contenedor

- Para detener el contenedor y liberar los recursos:
    ```bash
    docker-compose down
    ```

Este comando detendrá el contenedor y limpiará la red creada.


--------