# Laboratorio 6.3 - Depuración de contenedores
## Objetivo
Utilizar diferentes herramientas de depuración para entender el estado y el funcionamiento de un contenedor en ejecución.


## 1. Ejecutar contenedor web

- Ejecutamos un contenedor nginx

    ```bash
    docker run -d --rm --name nginx -p 8080:80 nginx:alpine
    ```

## 2. Ver los logs del contenedor:

- Los logs son útiles para ver las salidas estándar del contenedor y posibles errores.

    ```bash
    docker logs nginx
    ```
- Veamos los logs en tiempo real del contenedor

    ```bash
    docker logs --follow nginx
    ```
- En un browser web abrimos la siguiente URL: <a href="http://localhost:8080" target="_blank">http://localhost:8080</a>

- Observemos como cambia el log.
- Actualicemos la página en el navegador con la tecla `F5` y veamos como muestra log.


## 3. Ejecutar comandos dentro del contenedor

- Vamos a usar `docker exec` para acceder a la shell del contenedor y realizar cambios en caliente sin detener el contenedor:

    ```bash
    docker exec -it nginx sh
    ```

- Una vez en la shell cambiaremos el `index.html` de la siguiente manera:

    ```bash
    echo "Hola Curso" > /usr/share/nginx/html/index.html
    exit # Salimos de la sesión interactiva
    ```

- Refresquemos la página con `F5`.

--------

<img src="../../img/footer.svg" alt="Descripción de la imagen" style="width:100%; filter: grayscale(100%); opacity: 40%">