# Laboratorio 4.3 - Trabajando datos persistentes en Docker

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
- Vuelve a ejecutar un nuevo contenedor montando el mismo volumen y comprueba si el archivo aún está presente:

    ```bash
    docker run -it --name nuevo_contenedor -v mi_volumen:/datos alpine
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

- Crea un directorio en tu máquina llamado `mi_binds_mounts`:

    ```bash
    mkdir mi_binds_mounts
    ```

- Nos ubicamos en el directorio recien creado:

    ```bash
    cd mi_binds_mounts
    ```

### 2. Crear un archivo en el directorio del host

- Dentro de `mi_binds_mounts`, creamos un archivo de texto:

    ```bash
    echo "Hola desde el host" > archivo_host.txt
    ```

- Deberíamos ver el archivo recien creado en el directorio con el siguiente comando:

    ```bash
    ls
    ```
### 3. Montar el directorio del host en un contenedor

- Inicia un contenedor de Alpine montando `mi_binds_mounts` en `/data` dentro del contenedor:

    ```bash
    docker run -it --rm -v .:/data alpine
    ```

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

- El bind mount se elimina automáticamente al detener y eliminar el contenedor.

--------