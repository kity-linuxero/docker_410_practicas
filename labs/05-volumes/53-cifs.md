# Laboratorio 5.3 - Montando shares de red (SMB/CIFS) en red interna

## Objetivos

  - Desplegar un servidor Samba (SMB) para compartir archivos. Simulando que tenemos un share de Windows.
  - Configurar un volumen de Docker que conecte al share de Windows a través de la red.

>Para simplificar la implementación, usaremos una imagen de Samba que ya viene preconfigurada y trabajaremos en una red bridge interna de Docker.

## ¿Qué son CIFS y Samba?

- **CIFS (Common Internet File System)**: Es un protocolo de red propietario originalmente desarrollado por Microsoft, basado en SMB (Server Message Block). Se utiliza principalmente en entornos Windows para compartir archivos, impresoras y recursos a través de una red de forma transparente para el usuario.
- **Samba**: Es un software libre y de código abierto que implementa el protocolo SMB/CIFS en sistemas Unix y Linux. Samba permite que un servidor Linux actúe como un servidor de archivos e impresión compatible con redes Windows, facilitando la interoperabilidad entre ambos sistemas operativos.

## Parte 1: Configuración de la Red y el Servidor

### 1\. Crear la red interna:

  - Crea una red de tipo bridge llamada `red_lab`:

    ```powershell
    docker network create red_lab
    ```

### 2\. Levantar el servidor Samba:

  - Ejecuta el servidor conectado a `red_lab`. 

    ```powershell
    docker run -d --name servidor_samba --network red_lab dperson/samba -u "user;password123" -s "publico;/data;0;0;0;user"
    ```

    > Con esto simulamos un share de Windows.

### 3\. Ajustar permisos internos:

  - Para evitar errores de `Permission denied` desde el cliente, forzamos que la carpeta interna del servidor sea accesible:

    ```powershell
    docker exec -u root servidor_samba chmod -R 777 /data
    ```

> [!NOTE]
> No se recomienda usar los permisos `777` en Linux, ya que el mismo otorga permisos de lectura, escritura y ejecución a todos los usuarios. Pero en este caso, lo haremos a modo de prueba.
 
### 4\. Preparar datos de prueba:

  - Creamos un archivo directamente dentro del servidor:

    ```powershell
    docker exec servidor_samba bash -c "echo 'Datos compartidos en red interna' > /data/archivo_red.txt"
    ```

    > Con esto simulamos que se crea un archivo en una carpeta compartida de un servidor Windows.

## Parte 2: Creación del Volumen de Red (CIFS)

### 1\. Obtener la IP del servidor:

  - Necesitamos la IP que Docker le asignó al servidor dentro de `red_lab`:

    ```powershell
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' servidor_samba
    ```

    *(Anota la IP resultante, por ejemplo: `172.18.0.2`)*.

### 2\. Crear el volumen en PowerShell:

  > [!IMPORTANT]  
  > Reemplaza `<IP_INTERNA>` con el dato del paso anterior. Usamos la opción `noperm` para que el cliente no bloquee la escritura localmente:

  ```powershell
  docker volume create --driver local --opt type=cifs --opt device=//<IP_INTERNA>/publico --opt o="username=user,password=password123,file_mode=0777,dir_mode=0777,noperm,rw" volumen_cifs
  ```

Si sale lo siguiente:
```powershell
docker: Error response from daemon: error while mounting volume '/var/lib/docker/volumes/volumen_cifs/_data': error resolving passed in network volume address: lookup <IP_INTERNA>: no such host
```
Significa que no reemplazaste `<IP_INTERNA>` con la IP del servidor.

## Parte 3: Verificación de Persistencia y Acceso

### 1\. Montar el volumen en un cliente:

  - Iniciamos un contenedor Alpine en la misma red para validar el acceso:

    ```powershell
    docker run -it --rm --network red_lab -v volumen_cifs:/mnt/red alpine
    ```

### 2\. Comprobar lectura y escritura:

  - Dentro del contenedor cliente, ejecuta:

    ```bash
    # Leer el archivo original
    cat /mnt/red/archivo_red.txt

    # Crear un nuevo archivo (acá probamos la escritura)
    echo "Actualización desde el cliente" > /mnt/red/nota_cliente.txt

    # Salir
    exit
    ```

### 3\. Validar actualización en el servidor:

  - Verificamos que el servidor "ve" el archivo creado por el cliente:

    ```powershell
    docker exec servidor_samba ls /data
    ```
  - Deberíamos ver los siguientes archivos:

    ```powershell
    archivo_red.txt
    nota_cliente.txt
    ```

## Parte 4: Limpieza

  - Eliminamos los recursos creados para mantener el entorno limpio:

    ```powershell
    docker stop servidor_samba
    docker rm servidor_samba
    docker volume rm volumen_cifs
    docker network rm red_lab
    ```


## Resumen de Lab

En este lab vimos que es posible montar un share de red de Windows en un contenedor Docker. En el lab se usaron usuarios preestablecidos, pero se puede configurar en una red con usuarios de un dominio.

-----

<p align="center"\>
<img src="../../img/logos.footer.gray.webp"\>
</p\>

-----
