# Laboratorio 5.3 - Montando shares de red (SMB/CIFS) en red interna

## Objetivos

  - Crear una red bridge personalizada para comunicación aislada entre contenedores.
  - Desplegar un servidor Samba (SMB) sin exponer puertos al host para evitar conflictos con Windows.
  - Configurar un volumen de Docker que conecte al share a través de la red interna.
  - Solucionar problemas de permisos comunes en montajes de red.

## Parte 1: Configuración de la Red y el Servidor

### 1\. Crear la red interna:

  - Crea una red de tipo bridge llamada `red_lab`:

    ```powershell
    docker network create red_lab
    ```

### 2\. Levantar el servidor Samba:

  - Ejecuta el servidor conectado a `red_lab`. 

    ```powershell
    docker run -d --name servidor_samba `
      --network red_lab `
      dperson/samba `
      -u "cristian;password123" `
      -s "publico;/data;0;0;0;cristian"
    ```

### 3\. Ajustar permisos internos:

  - Para evitar errores de "Permission denied" desde el cliente, forzamos que la carpeta interna del servidor sea accesible:

    ```powershell
    docker exec -u root servidor_samba chmod -R 777 /data
    ```

### 4\. Preparar datos de prueba:

  - Creamos un archivo directamente dentro del servidor:

    ```powershell
    docker exec servidor_samba bash -c "echo 'Datos compartidos en red interna' > /data/archivo_red.txt"
    ```

## Parte 2: Creación del Volumen de Red (CIFS)

### 1\. Obtener la IP del servidor:

  - Necesitamos la IP que Docker le asignó al servidor dentro de `red_lab`:

    ```powershell
    docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' servidor_samba
    ```

    *(Anota la IP resultante, por ejemplo: `172.18.0.2`)*.

### 2\. Crear el volumen en PowerShell:

  - **IMPORTANTE:** Reemplaza `<IP_INTERNA>` con el dato del paso anterior. Usamos la opción `noperm` para que el cliente no bloquee la escritura localmente:

    ```powershell
    docker volume create --driver local `
      --opt type=cifs `
      --opt device=//<IP_INTERNA>/publico `
      --opt o="username=cristian,password=password123,file_mode=0777,dir_mode=0777,noperm,rw" `
      volumen_cifs
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

## Parte 4: Limpieza

  - Eliminamos los recursos creados para mantener el entorno limpio:

    ```powershell
    docker stop servidor_samba
    docker rm servidor_samba
    docker volume rm volumen_cifs
    docker network rm red_lab
    ```


-----

<p align="center"\>
<img src="../../img/logos.footer.gray.webp"\>
</p\>

-----