# Laboratorio 3.1 - Hola Docker

## Objetivos:
- Aprender a crear un Dockerfile básico.
- Construir una imagen Docker, y ejecutar un contenedor a partir de esa imagen.
- Familiarización con la construcción de imágenes y la ejecución de contenedores.

### Archivos a usar en esta actividad:

- `Dockerfile`: Archivo Dockerfile con las instrucciones para armar la imaǵen.

## 1. Crear el archivo `Dockerfile`:

- Crea una carpeta llamada `mi_app`.
- Dentro de esa carpeta, crea un archivo llamado `Dockerfile`.
- Escribe un Dockerfile que utilice la imagen base oficial de Alpine y que ejecute un script simple que muestre un mensaje en la terminal.
- Copie y pegue el siguiente contenido en un archivo Dockerfile

  ```Dockerfile
  # Usar la imagen base oficial de Alpine Linux
  FROM alpine:latest

  # Añadir un mensaje de bienvenida
  RUN echo "¡Hola, Docker!" > /mensaje.txt

  # Comando que se ejecutará cuando se inicie el contenedor
  CMD ["cat", "/mensaje.txt"]
  ```


> [!TIP]  
> Si utiliza el editor de textos de Windows, es probable que le agregue la extensión `.txt`. Utilice algún editor de texto como VSCode o Notepad++ para tener mas control sobre los nombre de archivos.

## 2. Construir la Imágen Docker

- Abre una terminal en la carpeta `mi_app`.
- Construye la imagen usando el siguiente comando:

  ```bash
  docker build -t mi_imagen:1.0 .
  ```
- Esto creará una imagen llamada `mi_imagen` con la etiqueta `1.0`.
- Verificar que la imágen se haya creado con el comando:

  ```bash
  docker images 
  ```


## 3. Ejecutar un Contenedor desde la Imagen

- Ejecuta un contenedor usando la imagen creada:

  ```bash
  docker run --rm mi_imagen:1.0
  ```
- Deberías ver el mensaje `¡Hola, Docker!` en la terminal.


> [!TIP]  
> El parámetro `--rm` enviado al comando `docker run` indicará que el contenedor debe eliminarse una vez que haya terminado su ejecución.

## 4. Seteando variables de entorno

- Vamos a modificar un poco el `Dockerfile` para que tenga un mensaje predeterminado y podamos personalizandolo en el `docker run`.

  ```dockerfile
  # Usar la imagen base oficial de Alpine Linux
  FROM alpine:latest

  # Añadir un mensaje de bienvenida. Esta vez mediante una variable de entorno.
  ENV MENSAJE="Hola Docker"

  # Comando que se ejecutará cuando se inicie el contenedor
  CMD ["/bin/sh", "-c", "echo $MENSAJE"]
  ```

- Contruimos la imágen y corremos el contenedor

  ```bash
  docker build -t mi-app:v2.0 . # Contruimos imágen
  docker run --rm mi-app:v2.0 # Corremos contenedor
  ```
  Debería mostrar el mismo mensaje, `Hola Docker`.

- Cambiaremos el comando de ejecución para mostrar otro mensaje:

  ```bash
  docker run -e MENSAJE="Mensaje personalizado" mi-app:v2.0
  ```

- Deberías ver tu mensaje personalizado. 👍



---------------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="../../../img/logos.footer.gray.webp">
  </a>
</p>
