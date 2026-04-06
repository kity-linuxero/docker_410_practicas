# Laboratorio 5.2 - Backup y migración de volúmenes

## Objetivo

  - Gestionar persistencia mediante volumes.
  - Utilizar la sintaxis `--mount` para mayor claridad y control.
  - Realizar backups comprimidos, simular la pérdida de datos y restaurarlos en un nuevo volumen.


> [\!IMPORTANT]
> **tar & busybox:** Utilizaremos la imagen `busybox` para ejecutar `tar` y empaquetar el contenido del volumen. El flag `-C` es fundamental para que el backup no incluya rutas absolutas del contenedor.

## 1. Preparación:

  - Crear un directorio de backup en tu host:

    ```bash
    mkdir backup
    cd backup
    ```

  - Crear un **Named Volume** llamado `v_datos_app`:

    ```bash
    docker volume create v_datos_app
    ```

  - Crear un contenedor llamado `dbstore` que monte este volumen usando la sintaxis `--mount`:

    ```bash
    docker run -d --name dbstore --mount type=volume,src=v_datos_app,dst=/data ubuntu tail -f /dev/null
    ```

  - Crear un archivo de texto con "información crítica" directamente en el volumen:

    ```bash
    docker exec dbstore sh -c "echo 'Backup realizado el $(date)' > /data/info.txt"
    ```

  - Verificar que el archivo existe:

    ```bash
    docker exec dbstore cat /data/info.txt
    ```

## 2. Realizar el backup (Empaquetado):

En este paso, lanzamos un contenedor temporal que "conecta" nuestro volumen con una carpeta de nuestro host para extraer la información:

  - **src=v\_datos\_app**: El volumen origen.

  - **$(pwd):/backup**: Nuestra carpeta actual en el host.

    ```bash
    docker run --rm --mount type=volume,src=v_datos_app,dst=/source_data -v ${pwd}:/backup_dir busybox tar czf /backup_dir/backup_2026.tar.gz -C /source_data .
    ```

  - Verifique que el archivo `backup_2026.tar.gz` se ha creado en su directorio actual.

## 3\. Simulación de desastre (Borrado):

Para demostrar la utilidad del backup, eliminaremos el contenedor y el volumen original:

  - Detener y eliminar el contenedor:

    ```bash
    docker rm -f dbstore
    ```

  - Eliminar el volumen (Pérdida total de datos):

    ```bash
    docker volume rm v_datos_app
    ```

## 4\. Restaurar volumen desde el backup:

Ahora restauraremos los datos en un volumen completamente nuevo, simulando una migración a otro host o recuperación de desastre.

  - Crear el nuevo volumen de destino:

    ```bash
    docker volume create v_datos_recuperados
    ```

  - Descomprimir el backup dentro del nuevo volumen:

    ```bash
    docker run --rm --mount type=volume,src=v_datos_recuperados,dst=/target_data -v ${pwd}:/backup_dir busybox sh -c "cd /target_data && tar xzf /backup_dir/backup_2026.tar.gz"
    ```

## 5. Verificar la restauración:

  - Verificamos si el archivo restaurado se encuentra en el nuevo volumen usando un contenedor efímero:

    ```bash
    docker run --rm --mount type=volume,src=v_datos_recuperados,dst=/data alpine cat /data/info.txt
    ```

  - Deberías ver el contenido original que creaste en el paso 1.

## 6\. Limpieza final

  - Eliminar el volumen de prueba y el archivo de backup:

    ```bash
    docker volume rm v_datos_recuperados
    cd ..
    rm -r backup #Borrar la carpeta backup
    ```


-----

<p align="center"\>
<img src="../../img/logos.footer.gray.webp"\>
</p\>

-----

