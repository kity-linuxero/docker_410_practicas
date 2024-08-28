# Laboratorio 4.4 - Backup y Migración de Volúmenes

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

- Realiza un backup del volumen data_vol en un archivo tar:

    ```
    docker run --rm -v data_vol:/data -v $(pwd):/backup alpine tar czf /backup/backup_data_vol.tar.gz -C /data .
    ```

## 4. Restaurar el backup en un nuevo volumen:

- Crea un nuevo volumen llamado restored_vol y restaura los datos del backup en él:

    ```bash
    docker volume create restored_vol
    docker run --rm -v restored_vol:/data -v $(pwd):/backup alpine sh -c "tar xzf /backup/backup_data_vol.tar.gz -C /data"
    ```

## 5. Montar el nuevo volumen en un contenedor y verificar los datos:

- Ejecuta un contenedor que monte restored_vol en /app y verifica si los datos han sido restaurados correctamente:

    ```bash
    docker run -it --name restored_container -v restored_vol:/app alpine
cat /app/datos.txt
    ```
- Deberías ver el contenido "Datos importantes".

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