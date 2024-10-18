# Laboratorio 2.3: Gestión de imágenes y versiones en Docker

## Objetivos:
- Entender por qué las imágenes de Docker son inmutables.
- Correr un contenedor de prueba y hacer modificaciones.
- Crear nuevas versiones de una imagen usando el comando [`commit`](#referencias).
- Guardar una imagen en un archivo utilizando el comando [`save`](#referencias).
- Cargar una imagen desde un archivo usando el comando [`load`](#referencias).

### Inmutabilidad de las Imágenes Docker

**Concepto:**
Las imágenes Docker son inmutables, lo que significa que una vez creadas, no se pueden cambiar. Esta característica garantiza que los entornos sean consistentes y reproducibles. Cualquier cambio que quieras hacer se debe realizar creando una nueva versión de la imagen.


## Gestión de versiones de imágenes

### 1. Correr un contenedor Linux de prueba

1. **Ejecuta el siguiente comando para correr un contenedor basado en la imagen de Ubuntu o AlmaLinux:**

    ```bash
    docker run -it --name contenedor-prueba ubuntu
    ```

    o bien

    ```bash
    docker run -it --name contenedor-prueba-alma almalinux bash
    ```



2. **Una vez dentro del contenedor, instala el paquete `curl`:**


#### En el caso de Ubuntu será:
    
```bash
apt update            # Actualiza la lista de paquetes
apt install -y curl   # Instala curl
```

#### En el caso de AlmaLinux será:

```bash
dnf install curl --skip-broken        # Instala Curl
```


3. **Comprobamos que `curl` esté correctamente instalado**

    ```bash
    curl ifconfig.co  # ifconfig.co es para obtener nuestra ip pública
    181.220.190.3 
    # Si una dirección ip significa que curl está instalado y fue posible obtener info de nuestra ip
    ```

4. **Salir del contenedor escribiendo:**

    ```bash
    exit
    ```

### 2. Crear una nueva versión de la imágen

1. **Guarda el contenedor como una nueva imagen con la etiqueta `v1.0`:**

    ```bash
    docker commit contenedor-prueba mi-imagen:v1.0
    ```

    Con el comando commit realizaremos una nueva versión de la imágen.


2. **Verifica que la nueva imagen se haya creado:**

    ```bash
    docker images

    REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
    mi-imagen    v1.0      9337aa68f959   5 seconds ago   133MB
    ubuntu       latest    35a88802559d   8 weeks ago     78.1MB
    ```



### 3. Hacer otro cambio y crear una nueva versión de la imágen

1. **Volvemos a correr un contenedor usando la imagen `mi-imagen:v1.0`:**

    ```bash
    docker run -it --name contenedor-prueba-v1.0 mi-imagen:v1.0
    ```

2. **Dentro del contenedor, instalamos otro paquete, por ejemplo, `nano`. En Ubuntu ejecutamos:**

    ```bash
    apt update
    apt install -y nano
    ```

    ó en AlmaLinux:

   ```bash
   dnf install nano
   ```

4. **Salir del contenedor:**

    ```bash
    exit
    ```

5. **Guardamos el contenedor como una nueva imagen con la etiqueta `v1.1`:**

    ```bash
    docker commit contenedor-prueba-v1.0 mi-imagen:v1.1
    ```

6. **Verifica que la nueva imagen se haya creado:**

    ```bash
    docker images

    REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
    mi-imagen    v1.1      b9ed48533c28   10 seconds ago   134MB
    mi-imagen    v1.0      9337aa68f959   45 minutes ago   133MB
    ubuntu       latest    35a88802559d   8 weeks ago      78.1MB
    
    ```


En estos pasos, aprendimos a correr un contenedor, hacer modificaciones y crear nuevas versiones de una imagen Docker.Este proceso demuestra cómo las imágenes Docker son inmutables y cómo se pueden gestionar versiones de imágenes mediante el comando commit.


## Guardado y cargando una imágen

Podemos guardar nuestra imágen personalizada a una registry o un archivo. En este lab, veremos el segundo caso.
La gestión local (en archivos) de imágenes nos serán útiles si no queremos subir nuestra imágen a un repositorio público o bien, podemos transportarla de un equipo a otro sin necesidad de tener internet para *pullear* las imágenes base.

### 4. Guardar imágen a un archivo

1. **Guarda la imagen `mi-imagen:v1.1` en un archivo:**

    ```bash
    docker save -o mi-imagen_v1.1.tar mi-imagen:v1.1
    ```

    Otra forma de hacerlo es:

    ```bash
    docker save mi-imagen:v1.1 > mi-imagen_v1.1.tar
    ```

    Ambas nos darán el mismo resultado.

2. **Verificamos que el archivo se ha creado:**

  - Explore la carpeta local y verifique que se haya creado el archivo
  - Si está en una terminal Bash, escriba el siguiente comando: 

    ```bash
    ls -sh mi-imagen_v1.1.tar 
    131M mi-imagen_v1.1.tar
    ```

### 5: Cargar la imagen desde un archivo

1. **Elimina la imagen `mi-imagen:v1.1` para simular que no está en el sistema:**

    ```bash
    docker rmi mi-imagen:v1.1
    ```

2. **Verifica que la imagen ya no está disponible:**

    ```bash
    docker images

    REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
    mi-imagen    v1.0      9337aa68f959   58 minutes ago   133MB
    ubuntu       latest    35a88802559d   8 weeks ago      78.1MB
    ```

3. **Carga la imagen desde el archivo tar:**

    ```bash
    docker load -i mi-imagen_v1.1.tar
    ```

    Otra forma de hacerlo es:

    ```bash
    docker load < mi-imagen_v1.1.tar
    ```

    *Resultado*
    ```bash
    docker load < mi-imagen_v1.1.tar
    92acd04bfc2a: Loading layer [==================================================>]  1.389MB/1.389MB
    Loaded image: mi-imagen:v1.1
    ```

4. **Verifica que la imagen se haya cargado correctamente:**

    ```bash
    docker images
    REPOSITORY   TAG       IMAGE ID       CREATED             SIZE
    mi-imagen    v1.1      b9ed48533c28   14 minutes ago      134MB
    mi-imagen    v1.0      9337aa68f959   About an hour ago   133MB
    ubuntu       latest    35a88802559d   8 weeks ago         78.1MB
    ```

## Conclusión

En este laboratorio, has aprendido sobre la inmutabilidad de las imágenes Docker y cómo crear y gestionar versiones de imágenes utilizando los comandos `commit`, `save` y `load`. Este conocimiento te permitirá mantener entornos reproducibles y versiones controladas de tus aplicaciones en contenedores Docker.




## Referencias

- <a href="https://docs.docker.com/reference/cli/docker/container/commit/" target="_blank">docker commit</a>
- <a href="https://docs.docker.com/reference/cli/docker/image/save/" target="_blank">docker image save</a>
- <a href="https://docs.docker.com/reference/cli/docker/image/load/" target="_blank">docker image load</a>

--------------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="../../img/logos.footer.gray.webp">
  </a>
</p>
