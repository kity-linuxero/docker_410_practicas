# Laboratorio 4.3 - Trabajando con Volúmenes en Docker

## Objetivo
Familiarizarse con la creación de volúmenes, montarlos en contenedores, verificar la persistencia de datos, y manipular volúmenes.



## 1. Crear un volumen:

- Crea un volumen llamado `mi_volumen`:

    ```bash
    docker volume create mi_volumen
    ```

## 2. Montar el volumen en un contenedor:

- Ejecuta un contenedor basado en alpine y monta mi_volumen en el directorio /datos dentro del contenedor:

    ```bash
    docker run -it --name mi_contenedor -v mi_volumen:/datos alpine
    ```

## 3. Guardar datos en el volumen:

- Dentro del contenedor, creamos un archivo de texto en el directorio `/datos`:

    ```
    echo "Hola desde Docker!" > /datos/mensaje.txt
    ```

- Salimos del contenedor:

    ```
    exit
    ```


## 4. Comprobar la persistencia de datos:

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


## 5. Inspeccionar el volumen:

- Inspecciona el volumen mi_volumen para ver su configuración y ubicación en el host:

    ```bash
    docker volume inspect mi_volumen
    ```
- ¿Qué dato relevante podés ver?

## 6. Eliminar el volumen:

- Elimina el contenedor nuevo_contenedor y luego elimina el volumen mi_volumen:

    ```bash
    docker rm nuevo_contenedor
    docker volume rm mi_volumen
    ```



--------