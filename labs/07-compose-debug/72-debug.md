# Laboratorio 7.2 - Depuración de contenedores
## Objetivo
Utilizar diferentes herramientas de depuración para entender el estado mirando logs del contenedor con `docker logs` y el funcionamiento de un contenedor en ejecución con `docker exec`.


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

- Vamos a usar `docker exec` para acceder a la shell del contenedor y realizar cambios en caliente sin detener el contenedor. Lo que se dice hacer cambios _en caliente_:

    ```bash
    docker exec -it nginx sh
    ```

- Una vez en la shell cambiaremos el `index.html` de la siguiente manera:

    ```bash
    echo "Hola Curso" > /usr/share/nginx/html/index.html
    exit # Salimos de la sesión interactiva
    ```

  > Esto cambiará la página que estamos viendo en el navegador.

- Refresquemos la página con `F5`.

## 4. Limpieza

- Detener el contenedor:

    ```bash
    docker stop nginx
    ```

## Resumen de Lab

En este lab hemos visto los logs del contenedor web con `docker logs` y ejecutado comandos dentro del contenedor con `docker exec` para realizar cambios en caliente y poder navegar sobre un contenedor en ejecución.



--------

<p align="center">
  <img src="../../img/logos.footer.gray.webp">
</p>