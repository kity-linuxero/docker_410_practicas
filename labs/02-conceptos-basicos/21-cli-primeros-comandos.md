# Laboratorio 2.1 - Primeros comandos en Docker

### Objetivos:

- Practicar los primeros comandos en Docker
- Interactuar con la Docker CLI.

## 1. Iniciando un contenedor Ubuntu

Vamos a iniciar un contenedor Ubuntu de la siguiente manera

```bash
docker run -d --name duerme_30 ubuntu sleep 30
```

Analicemos:
- `docker run`: Ejecutará un contenedor de una imágen especificada, en este caso `ubuntu`.
- `-d` ó `--detach`: Ejecuta un contenedor en background. Imprime el container ID.
- `--name`: Setea un nombre personalizado al contenedor.
- `sleep 30`: Es un comando que queremos enviar al contenedor. `sleep` hace que quede en espera una cantidad determinada de segundos, como hemos enviado el parámetro `30` quedará ejecutando en espera 30 segundos.

> [!TIP]  
> Para mas info consulte documetación de Docker <a href="https://docs.docker.com/reference/cli/docker/container/run/" target="_blank">`docker run`</a>

### Listando contenedores en ejecución

Luego de ejecutar el comando, antes de los 30 segundos ejecutamos `docker ps`:

```bash
docker ps
CONTAINER ID   IMAGE     COMMAND      CREATED         STATUS         PORTS     NAMES
56bdfd527c78   ubuntu    "sleep 30"   4 seconds ago   Up 3 seconds             duerme_30
```

El comando nos devuelve información de los contenedores en ejecución:
- **`CONTAINER ID`**: Es un identificador único (hash) asignado a cada contenedor, utilizado para diferenciarlo de otros contenedores en el sistema.
- **`IMAGE`**: Muestra el nombre de la imagen desde la cual se ha creado y está corriendo el contenedor.
- **`COMMAND`**: Indica el comando que se ejecutó dentro del contenedor al iniciarlo. En este caso, `sleep 30`.
- **`CREATED`**: Indica el tiempo transcurrido desde que el contenedor fue creado.
- **`STATUS`**: Refleja el estado actual del contenedor, como si está en ejecución, detenido, o en otro estado.
- **`PORTS`**: Lista los puertos que el contenedor está escuchando y aquellos en el host que están redirigidos al contenedor. Este concepto se explicará más a fondo al trabajar con contenedores que interactúan mediante la red.
- **`NAMES`**: Muestra el nombre asignado al contenedor. Si no se especifica un nombre al crear el contenedor, Docker le asigna uno de forma aleatoria, lo que permite una identificación alternativa además del ID.

Si luego de los 30 segundos, volvemos a ejecutar `docker ps`:

```bash
docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
# Ninguna información
```

No mostrará nada, porque el contenedor ya no está en ejecución.

Para ver información de TODOS los contenedores, `docker ps -a`:

```bash
docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED         STATUS                     PORTS    NAMES
56bdfd527c78   ubuntu   "sleep 30"  2 minutes ago   Exited (0) 2 minutes ago            duerme_30
# Observar info status
```

> [!IMPORTANT]  
> Una vez que el contenedor termine su ejecución NO se eliminará, sino que estará en estado de `Exited`.  

> [!TIP]  
> Para mas info consulte documetación de Docker <a href="https://docs.docker.com/reference/cli/docker/container/ps/" target="_blank">`docker ps`</a>

## 2. Ejecutando comandos en el contenedor

Como podemos deducir, el último de los parámetros enviados en el comando `docker run` es la órden que ejecutará el contenedor.
Así, podemos ejecutar cualquiera que se nos ocurra, por ejemplo: 

  ```bash
  docker run ubuntu ls -l
  ```
> [!NOTE]  
> El comando `ls` de Linux listará los archivos y directorios.


```bash
$ docker run ubuntu ls -l
total 48
lrwxrwxrwx   1 root root    7 Apr 22 13:08 bin -> usr/bin
drwxr-xr-x   2 root root 4096 Apr 22 13:08 boot
drwxr-xr-x   5 root root  340 Aug 12 02:23 dev
drwxr-xr-x   1 root root 4096 Aug 12 02:23 etc
drwxr-xr-x   3 root root 4096 Jun  5 02:06 home
lrwxrwxrwx   1 root root    7 Apr 22 13:08 lib -> usr/lib
lrwxrwxrwx   1 root root    9 Apr 22 13:08 lib64 -> usr/lib64
drwxr-xr-x   2 root root 4096 Jun  5 02:02 media
drwxr-xr-x   2 root root 4096 Jun  5 02:02 mnt
drwxr-xr-x   2 root root 4096 Jun  5 02:02 opt
dr-xr-xr-x 370 root root    0 Aug 12 02:23 proc
drwx------   2 root root 4096 Jun  5 02:05 root
drwxr-xr-x   4 root root 4096 Jun  5 02:06 run
lrwxrwxrwx   1 root root    8 Apr 22 13:08 sbin -> usr/sbin
drwxr-xr-x   2 root root 4096 Jun  5 02:02 srv
dr-xr-xr-x  13 root root    0 Aug 12 02:23 sys
drwxrwxrwt   2 root root 4096 Jun  5 02:05 tmp
drwxr-xr-x  12 root root 4096 Jun  5 02:02 usr
drwxr-xr-x  11 root root 4096 Jun  5 02:05 var
```
Esto creará un contenedor nuevo partiendo de la imágen `ubuntu`, ejecutará el comando `ls -l` y luego sale del contenedor. El contenedor habrá finalizado su ejecución.


## 3. Eliminando contenedores

Vamos a eliminar el contenedor `duerme_30` para eso debemos ejecutar `docker rm <CONTAINER ID> o <CONTAINER_NAME>`. Luego ejecutamos `docker ps -a` para verificar que efectivamente se haya eliminado.

```bash
docker rm duerme_30
56bdfd527c78

docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED             STATUS                          PORTS    NAMES
f89187d1d5b4   ubuntu   "ls -l"     About a minute ago  Exited (0) About a minute ago            vibrant_jones
```
Observe que:

- El contenedor con ID `56bdfd527c78` y name `duerme_30` se habrá eliminado.
- Aún está el contenedor creado para hacer el listado de directorio del punto 2.


Eliminemos el contenedor del listado, ejecutando `docker rm <CONTAINER_ID>`en mi caso será el ID `f89187d1d5b4`.

```bash
docker rm f89187d1d5b4
f89187d1d5b4

docker ps -a
CONTAINER ID   IMAGE    COMMAND     CREATED         STATUS                     PORTS    NAMES
# No debería haber ningún contenedor.
```



## 4. Eliminar automáticamente un contenedor luego de su ejecución

Luego de varias pruebas de ejecutar contenedores con el comando `docker run` notará que, si consulta con el comando `docker ps -a` se habrán creado un gran número de contenedores que, quizás ya no sean necesarios. Puede agregar el parámetro `--rm` para que se eliminen automáticamente luego de la ejecución.

Ejecutemos:

```bash
docker run --rm --name lab21-rm ubuntu ls -l
```

Output:

```bash
$ docker run --rm --name lab21-rm ubuntu ls -l
total 48
lrwxrwxrwx   1 root root    7 Apr 22 13:08 bin -> usr/bin
drwxr-xr-x   2 root root 4096 Apr 22 13:08 boot
drwxr-xr-x   5 root root  340 Oct  7 02:31 dev
drwxr-xr-x   1 root root 4096 Oct  7 02:31 etc
drwxr-xr-x   3 root root 4096 Jun  5 02:06 home
lrwxrwxrwx   1 root root    7 Apr 22 13:08 lib -> usr/lib
lrwxrwxrwx   1 root root    9 Apr 22 13:08 lib64 -> usr/lib64
drwxr-xr-x   2 root root 4096 Jun  5 02:02 media
drwxr-xr-x   2 root root 4096 Jun  5 02:02 mnt
drwxr-xr-x   2 root root 4096 Jun  5 02:02 opt
dr-xr-xr-x 383 root root    0 Oct  7 02:31 proc
drwx------   2 root root 4096 Jun  5 02:05 root
drwxr-xr-x   4 root root 4096 Jun  5 02:06 run
lrwxrwxrwx   1 root root    8 Apr 22 13:08 sbin -> usr/sbin
drwxr-xr-x   2 root root 4096 Jun  5 02:02 srv
dr-xr-xr-x  13 root root    0 Oct  7 02:31 sys
drwxrwxrwt   2 root root 4096 Jun  5 02:05 tmp
drwxr-xr-x  12 root root 4096 Jun  5 02:02 usr
drwxr-xr-x  11 root root 4096 Jun  5 02:05 var
```

Si buscamos el contenedor de la siguiente manera:

```bash
docker ps -a
```

Podrá verificar que el contenedor llamado `lab21-rm` no aparece en el listado. Esto es porque luego de la ejecución del mismo, se ha eliminado.

## Referencias

- <a href="https://docs.docker.com/reference/cli/docker/container" target="_blank">docker container commands</a>
- <a href="https://docs.docker.com/reference/cli/docker/container/run/" target="_blank">docker run</a>
- <a href="https://docs.docker.com/reference/cli/docker/container/ls/" target="_blank">docker ps</a>
- <a href="https://docs.docker.com/reference/cli/docker/container/rm/" target="_blank">docker rm</a>


---------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="../../img/logos.footer.gray.webp">
  </a>
</p>