# Actividad 2 - La mejor página del mundo

>Al finalizar esta actividad, usted habrá conteinerizado una página web estática.

### Archivos a usar en esta actividad:
- `index.html`: El código de la página que vamos a servir.
- `Dockerfile`: Archivo Dockerfile con las instrucciones para armar la imaǵen.


## 1. Analizar el archivo `Dockerfile`:

```dockerfile
# Imágen base
FROM ubuntu:latest

# Ejecuto actualización de repositorios e instalo el servidor web nginx
RUN apt-get update && apt-get install -y nginx

# Copio el archivo index para que sea visible desde el nginx del contenedor
COPY index.html /var/www/html/

# Comando a ejecutar
CMD ["nginx", "-g", "daemon off;"]

# Expongo el puerto 80
EXPOSE 80
```

Según podemos ver, arrancamos de una imágen base Ubuntu. Sin nada instalado. Así que, se procede a actualizar la base de repositorios con `apt-get update` y luego instalamos nginx con `apt-get install nginx`.

En la capa 3, se copia el archivo `index.html` a la ruta `/var/www/html` que es la ruta por defecto donde nginx almacena los archivos.

## 2. Crear la imaǵen

Ejecutar:

```bash
docker build . -t best-page
```
Eso generará la imágen Docker, llamada `best-page`.

## 3. Verificar que la imágen fue creada

Verificar que la imágen fue creada con el comando `docker images`

```bash
docker images

REPOSITORY                    TAG       IMAGE ID       CREATED          SIZE
best-page                     latest    b77f72c4c183   4 minutes ago    125MB 
```

:ok_hand: ¡Perfecto! La imágen ha sido creada.

## 4. Correr el contenedor:

```bash
docker run -d -p 80:80 best-page
```

En este caso, se le manda parámetros a `docker run`:
- `-d`: Indica que el contenedor se ejecutará como *daemon*.
- `-p`: Indica que se publicarán puertos hacia afuera del contenedor. En este caso, `80:80` mapea el puerto 80 de nuestra PC al puerto 80 del contenedor.

## 5. Verificar que el contenedor esté corriendo

Ejecutar el comando `docker ps`

```bash
docker ps

CONTAINER ID   IMAGE       COMMAND                  CREATED         STATUS         PORTS                               NAMES
ea914d8b91b3   best-page   "nginx -g 'daemon of…"   5 seconds ago   Up 5 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp   funny_elgamal
```

## 6. Comprobar el funcionamiento

Ingrese a un navegador y escriba en la dirección [localhost](http://localhost)

![](./screenshot.png)

---------------

![](../../img/footer.svg)