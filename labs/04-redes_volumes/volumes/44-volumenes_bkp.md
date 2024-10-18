# Laboratorio 4.4 - Backup y migración de volúmenes

## Objetivo
Aprender a hacer backups de volúmenes, migrar datos a otro volumen, y manipular archivos en volúmenes montados en contenedores.



## 1. Crear un volumen y un contenedor:

- Crea un volumen llamado data_vol y un contenedor que monte este volumen en `/app`:

    ```bash
    docker volume create data_vol
    docker run -it --name app_container -v data_vol:/app alpine
    ```

## 2. Añadir datos al volumen:

- Dentro del contenedor, crea un archivo de texto en `/app`:

    ```bash
    echo "Datos importantes" > /app/datos.txt
    ```
- Sal del contenedor:
    ```bash
    exit
    ```

## 3. Crear un backup del volumen:

- Crearemos un directorio para trabajo:
    ```bash
    mkdir workdir
    ```
- Nos ubicamos dentro del directorio creado:
    ```bash
    cd workdir
    ```

- Realizar un backup del volumen data_vol en un archivo `tar`:

    ```bash
    docker run --rm -v data_vol:/data -v .:/backup alpine tar czf /backup/backup_data_vol.tar.gz -C /data .
    ```

Si todo salió bien, deberías tener en el directorio donde estás ubicado un archivo llamado `backup_data_vol.tar.gz`. Que no es mas que un archivo comprimido del contenido completo del volumen.

### Desglosando el comando

- `docker run`: Es el comando que inicia un nuevo contenedor.
- `--rm`: Esta opción le indica a Docker que elimine automáticamente el contenedor cuando termine de ejecutarse.
- `-v data_vol:/data`: Son los parámetros que indican como montar el volumen:
    - `data_vol`: es el volumen de Docker que ya hemos creado previamente y contiene los datos que queremos respaldar.
    - `/data`: es el directorio dentro del contenedor donde se montará el volumen. Esto significa que todo el contenido del volumen data_vol estará accesible dentro del contenedor bajo el directorio /data.
- `-v .:/backup`: Realiza un montaje. PERO en esta vez se trata de un <a href="https://docs.docker.com/engine/storage/bind-mounts/" target="_blank">`bind mount`</a> que permite montar un directorio del host directamente al contenedor.
    - `.`: Indica que queremos montar el directorio actual.
    - `/backup`: es el directorio dentro del contenedor donde se montará el directorio del host. De esa forma, cualquier archivo que se cree en `/backup` dentro del contenedor aparecerá en el directorio actual de tu máquina local.
- `alpine`: Es el nombre de la imágen que se usará para crear el contenedor. En este caso la distribución liviana de Linux Alpine.
- `tar czf /backup/backup_data_vol.tar.gz -C /data .`:
    - `tar`: es una herramienta de Unix que se utiliza para crear archivos del tupo tarball, que son básicamente archivos empaquetados en uno solo.
    - `czf` son opciones de `tar`:
        - `c`: Crea un nuevo archivo tar.
        - `z`: Comprime el archivo tar usando gzip.
        - `f`: Indica el nombre de archivo tar a crear.
    - `/backup/backup_data_vol.tar.gz`: es la ruta donde se guardará el archivo tar comprimido dentro del contenedor, que se corresponde con el directorio actual del host (por el bind mount).
    - `-C /data`: le indica a `tar` que cambie al directorio `/data` dentro del contenedor antes de iniciar la creación del archivo tar.
    - `.`: le dice a `tar` que archive todo el contenido del directorio actual (`/data`)

#### Resumen
- Este comando inicia un contenedor temporal basado en Alpine.
- Monta el volumen data_vol en el directorio /data dentro del contenedor y monta el directorio actual de tu host en /backup.
- Luego, utiliza tar para crear un archivo comprimido (.tar.gz) de todo el contenido de /data (es decir, todo lo que hay en el volumen data_vol) y lo guarda en /backup, que corresponde al directorio actual de tu host.

## 4. Restaurar el backup en un nuevo volumen:

- Crea un nuevo volumen llamado restored_vol y restaura los datos del backup en él:

    ```bash
    docker volume create restored_vol
    ```
- Ejecutamos un contenedor temporal para guardar los datos en el volumen

    ```bash
    docker run --rm -v restored_vol:/data -v .:/backup alpine sh -c "tar xzf /backup/backup_data_vol.tar.gz -C /data"
    ```

## 5. Montar el nuevo volumen en un contenedor y verificar los datos:

- Ejecuta un contenedor que monte restored_vol en /app y verifica si los datos han sido restaurados correctamente:

    ```bash
    docker run -it --name restored_container -v restored_vol:/app alpine
    cat /app/datos.txt
    ```
- Deberías ver el contenido `Datos importantes`.



## 6. Copiar un archivo al contenedor:

- En tu host, crea un archivo llamado nuevos_datos.txt con el contenido "Nuevos datos para Docker".

- Copia este archivo al contenedor restored_container en el directorio `/app`:

    ```bash
    docker cp nuevos_datos.txt restored_container:/app/nuevos_datos.txt
    ```

## 7. Verificar el archivo copiado:

- Dentro del contenedor, verifica que el archivo ha sido copiado correctamente:

    ```bash
    cat /app/nuevos_datos.txt
    ```

## 8. Eliminar volúmenes:

```bash
docker rm app_container restored_container
docker volume rm data_vol restored_vol
```
--------

<p align="center">
  <img src="../../../img/logos.footer.gray.webp">
</p>