# Laboratorio 5.2 - Backup y migración de volúmenes

## Objetivo
- Usar volúmenes para hacer backups, restaurar datos y migrar datos entre contenedores.
- Utilizar el flag `--volumes-from` para crear contenedores que monten volúmenes.

> **Nota:** Con el flag [`--volumes-from`](#referencias) podemos crear contenedores que monten los mismos volúmenes de otros contenedores.

> [!IMPORTANT] 
> **tar:** En este lab utilizaremos `tar` para empaquetar el contenido del volumen `dbdata` en un archivo `backup.tar`. Para mas info, consulte la [manual de tar](https://www.commandlinux.com/man-page/man1/tar.1.html). Este comando está fuera del alcance de este curso. Solo se utiliza para fines ilustrativos.

## 1. Preparación:

- Crear un directorio de backup:

    ```bash
    mkdir backup
    cd backup
    ```

- Crear un nuevo contenedor llamado `dbstore` con un volumen llamado `/dbdata`.:

    ```bash
    docker run -v /dbdata --name dbstore ubuntu /bin/bash
    ```
    En este caso utilizamos un volumen [anónimo](https://docs.docker.com/engine/storage/volumes#named-and-anonymous-volumes), ya que no le seteamos el nombre. El nombre se genera automáticamente por Docker. De todas maneras, en este caso, se monta en la ruta `/dbdata` del contenedor.
    Podemos ver el nombre del volumen que Docker le ha asignado con el siguiente comando:

    ```bash
    docker volume ls
    ```

    Probablemente el nombre asignado por Docker sea el hash del volumen. Algo como `e877e92467ac3ec0108...`.

- Crear un archivo de texto que será como nuestra info a resguardar:

    ```bash
    echo "Datos de la base de datos" > info.txt
    ```

- Copiamos el archivo `info.txt` al volumen `dbdata`:

    ```bash
    docker cp info.txt dbstore:/dbdata
    ```

- Deberíamos ver el siguiente mensaje que indica que el archivo fue correctamente copiado al volumen:

    ```bash
    Successfully copied 2.05kB to dbstore:/dbdata
    ```



## 2. Realizar el backup:

En el siguiente comando, haremos lo siguiente:

- Lanzar un contenedor temporal y montar el volumen del contenedor `dbstore`.
- Montar un directorio del host local como `/backup`.
- Pasar un comando que empaquete el contenido del volumen `dbdata` en un archivo `backup.tar` dentro del directorio `/backup`.

    ```bash
    docker run --rm --volumes-from dbstore -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar /dbdata
    ```

    Cuando se completa el comando y el contenedor se detiene, crea una copia de seguridad del volumen `dbdata` en el archivo `backup.tar` dentro del directorio `backup`. 

- Verifique en su explorador de archivos o por consola que el archivo `backup.tar` se ha creado en el directorio `backup`.


## 3. Restaurar volumen desde el backup:

Con el backup recién creado, podemos restaurarlo en el mismo contenedor o en otro contenedor que creemos en otro lugar.

- Para este lab, crearemos un nuevo contenedor llamado `dbstore2`.

    ```bash
    docker run -v /dbdata --name dbstore2 -dit ubuntu /bin/bash
    ```

- Luego, descomprimimos el archivo de backup en el volumen de datos del nuevo contenedor:

    ```bash
    docker run --rm --volumes-from dbstore2 -v $(pwd):/backup ubuntu bash -c "cd /dbdata && tar xvf /backup/backup.tar --strip 1"
    ```

## 4. Verificar la restauración:

- Verificamos si el archivo restaurado se encuentra en el contenedor `dbstore2`:

    ```bash
    docker exec -it dbstore2 ls /dbdata
    ```
- Deberíamos ver listado el archivo `info.txt` restaurado.


## 5. Eliminar contenedores y volúmenes

- Eliminamos los contenedores `dbstore` y `dbstore2`:

    ```bash
    docker rm -f dbstore dbstore2
    ```
- Eliminar volúmenes. Como se han creado volúmenes anónimos, vamos a eliminarlos con el comando `docker volume prune`:

    ```bash
    docker volume prune
    ```
- Eliminar directorio `backup`:

    ```bash
    cd ..
    rm -rf backup
    ```

## Referencias

- [Docker Docs: Backup, Restore or Migrate Data Volumes](https://docs.docker.com/engine/storage/volumes#back-up-restore-or-migrate-data-volumes)
- [Docker Docs: Volumes from](https://docs.docker.com/reference/cli/docker/container/run/#volumes-from)
- [Docker Docs: Named and anonymous volumes](https://docs.docker.com/engine/storage/volumes#named-and-anonymous-volumes)
- [Docker Docs: Volumes](https://docs.docker.com/engine/storage/volumes/)

--------

<p align="center">
  <img src="../../img/logos.footer.gray.webp">
</p>
