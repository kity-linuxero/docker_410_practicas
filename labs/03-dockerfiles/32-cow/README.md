# Laboratorio 3.2 - Una vaca te saluda

## Objetivos:
- Analizar y entender básicamente archivos `Dockerfile`
- Crear una imágen usando `Dockerfile` que sirva para correr nuestra aplicación con todas sus dependencias y scripts para instalarlas.
- Ejecutar contenedores de nuestra imágen
- Conteinerizar una aplicación Python.

### Archivos a usar en esta actividad:

- <a href="/cowsay-hello.py" download>`cowsay-hello.py`</a>: El código fuente escrito en Python que vamos a correr en un contenedor.
- <a href="/Dockerfile" download>`Dockerfile`</a>: Archivo Dockerfile con las instrucciones para armar la imaǵen.

- <a href="/requirements.txt" download>`requirements.txt`</a>: Archivo generalmente usado en programas Python para listar las dependencias.

Puede descargar todos estos archivos desde un solo link <a href="./32-cow.zip" download>acá</a>.

## 1. Analizar el archivo `Dockerfile`:
```Dockerfile
# Imagen base
FROM python:3.9.19-alpine3.20

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de la aplicación
COPY . /app

# Instalar las dependencias
RUN pip install --no-cache-dir -r requirements.txt

# Comando para ejecutar la aplicación
CMD ["python", "cowsay-hello.py"]
```
Como puede verse, se trata de un programa escrito en Python. Las dependencias están en el archivo `requirements.txt`. La capa de instalación de dependencias, instalará las dependencias para este programa.


## 2. Crear la imaǵen

Ejecutar:

```bash
docker build . -t vaquita
```
La imágen `vaquita` ha sido creada:
- Se ha descargado la imágen `python:3.9.19-alpine3.20` de la registry.
- Se han ejecutado los comandos indicados en el Dockerfile.


## 3. Verificar que la imágen fue creada

Ejecutar el comando `docker images`

```bash
docker images

REPOSITORY                    TAG       IMAGE ID       CREATED          SIZE
vaquita                       latest    4bfe8b137eec   10 minutes ago   54.4MB 
```

:ok_hand: ¡Perfecto! La imágen ha sido creada.

## 4. Correr el contenedor:

```bash
docker run vaquita
```
### Resultado

```bash
  ______________________                                                                                                      
| ¡Hola curso de Docker! |
  ======================
                      \
                       \
                         ^__^
                         (oo)\_______
                         (__)\       )\/\
                             ||----w |
                             ||     ||

```

En este Lab hemos _dockerizado_ una app desarrollada en Python. Si bien la app es muy simple, el mecanismo para dockerizar aplicaciones en Python es el mismo.

---------------

![](../../img/footer.svg)