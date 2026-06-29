# Laboratorio 5.1 - Trabajando datos persistentes en Docker

## Objetivos
- Aprender a crear y gestionar volúmenes en Docker.
- Utilizar bind mounts para montar directorios específicos del host en contenedores y verificar la persistencia de datos entre sesiones de contenedores.


## Parte 1: Creación de volumenes y persistencia de datos

### 1. Crear un volumen:

- Crea un volumen llamado `mi_volumen`:

    ```bash
    docker volume create mi_volumen
    ```

### 2. Montar el volumen en un contenedor:

- Ejecuta un contenedor basado en alpine y monta mi_volumen en el directorio /datos dentro del contenedor:

    ```bash
    docker run -it --name mi_contenedor -v mi_volumen:/datos alpine
    ```

### 3. Guardar datos en el volumen:

- Dentro del contenedor, creamos un archivo de texto en el directorio `/datos`:

    ```
    echo "Hola desde Docker!" > /datos/mensaje.txt
    ```

- Salimos del contenedor:

    ```
    exit
    ```


### 4. Comprobar la persistencia de datos:

- Elimina el contenedor mi_contenedor

    ```bash
    docker rm mi_contenedor
    ```
- Vuelve a ejecutar un nuevo contenedor basado en otra imágen, esta vez en vez de `alpine` usa `ubuntu`, montando el mismo volumen y comprueba si el archivo aún está presente:

    ```bash
    docker run -it --name nuevo_contenedor -v mi_volumen:/datos ubuntu
    cat /datos/mensaje.txt
    ```
- Deberías ver el mensaje "Hola desde Docker!" en la salida.

- Salimos del contenedor:

    ```bash
    exit
    ```

### 5. Inspeccionar el volumen:

- Inspecciona el volumen mi_volumen para ver su configuración y ubicación en el host:

    ```bash
    docker volume inspect mi_volumen
    ```
- ¿Qué dato relevante podés ver?

### 6. Eliminamos el contenedor:

- Elimina el contenedor nuevo_contenedor:

    ```bash
    docker rm nuevo_contenedor
    ```

## Parte 2: Uso de Bind Mounts para montar directorios del host

### 1. Crea un directorio en el host

- Crea un directorio o carpeta en tu máquina llamada `mi_binds_mounts`. Lo podés hacer con el explorador de archivos o con el comando:

    ```bash
    mkdir mi_binds_mounts
    ```

- Nos ubicamos en el directorio recien creado:

    ```bash
    cd mi_binds_mounts
    ```

### 2. Crear un archivo en el directorio del host

- Dentro de `mi_binds_mounts`, creamos un archivo de texto. Lo podés hacer con el bloc de notas o con el siguiente comando:

    ```bash
    echo "Hola desde el host" > archivo_host.txt
    ```

- Deberíamos ver el archivo recien creado en el directorio con el siguiente comando:

    ```bash
    ls
    ```
### 3. Montar el directorio del host en un contenedor

- Inicia un contenedor de Alpine montando `mi_binds_mounts` en `/data` dentro del contenedor:

    **Si estás usando Windows, hay que especificar la ruta completa del directorio. Por ejemplo:**

    ```bash
    docker run -it --rm -v "${PWD}:/data" alpine
    ```

    **Si estás usando Linux o MacOS, el comando es:**

    ```bash
    docker run -it --rm -v .:/data alpine
    ```

    En los comandos anteriores, `${PWD}` o `.` representa el directorio actual en tu máquina local.

> Ahora tendremos el directorio de nuestra máquina local montado en el contenedor en la carpeta `/data`.

### 4. Verificar el contenido del bind mount

- Dentro del contenedor, lista los archivos en `/data` para verificar que el archivo `archivo_host.txt` está disponible:

    ```bash
    ls -l /data
    ```

- Mostramos el contenido del archivo

    ```bash
    cat /data/archivo_host.txt
    ```

### 5. Crear un archivo dentro del contenedor

- Crea un archivo nuevo dentro del directorio montado desde el contenedor:

    ```bash
    echo "Archivo creado desde el contenedor" > /data/archivo_contenedor.txt
    ```

- Salimos del contenedor:

    ```bash
    exit
    ```

### 6. Verificar la persistencia en el host

- En el directorio `mi_bind_mount` de nuestra máquina local, verificamos que el archivo `archivo_contenedor.txt` fue creado en el host:

    ```bash
    cat archivo_contenedor.txt
    ```

## Parte 3: Eliminación de Volúmenes y Bind Mounts

### 1. Eliminar un volumen

- Eliminamos el volumen `mi_volumen`:

    ```bash
    docker volume rm mi_volumen
    ```

### 2. Eliminar un bind mount:

- El bind mount se elimina automáticamente al detener y eliminar el contenedor, ya que no es un recurso gestionado por Docker. Sin embargo, podemos eliminar el directorio `mi_binds_mounts` como eliminaríamos una carpeta normal en tu sistema operativo.

## Referencias

- [bind mounts](https://docs.docker.com/engine/storage/bind-mounts/)
- [volumes](https://docs.docker.com/engine/storage/volumes/)


--------

<p align="center">
  <img src="../../img/logos.footer.gray.webp">
</p>