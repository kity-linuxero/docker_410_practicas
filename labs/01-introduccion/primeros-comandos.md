# Laboratorio 1.2 - Primeros comandos

### Objetivos:

- Ejecución de nuestro primer contenedor.
- Ejecutar algunos comandos básicos de GNU/Linux y Docker

## Contenedor Ubuntu

Vamos a ejecutar un contenedor de la distribución de GNU/Linux Ubuntu.

```bash
docker run -it ubuntu bash
```

Luego que la imágen se haya descargado, tendremos un promt de terminal de bash de la siguiente manera:

```
root@d9080b400c84:/#
```

Ya tenemos un contenedor Ubuntu corriendo. Lo podemos usar como si de una VM se tratara. Solo para probar, vamos a instalar algunos paquetes.

```bash
apt update
apt install -y iputils-ping
```

Una vez que haya terminado la secuencia de instalación ejecutamos:

```bash
ping -c 3 localhost
```

```
# ping -c 3 localhost
PING localhost (::1) 56 data bytes
64 bytes from localhost (::1): icmp_seq=1 ttl=64 time=0.028 ms
64 bytes from localhost (::1): icmp_seq=2 ttl=64 time=0.025 ms
64 bytes from localhost (::1): icmp_seq=3 ttl=64 time=0.022 ms
```

### Comandos de GNU/Linux 

Como nos encontramos en una terminal `bash`, así que podemos ejecutar comandos en el intérprete a nuestro sistema operativo. A continuación tiene una lista de comandos básicos para probar dentro de una terminal de Ubuntu.

- `pwd`: Muestra el directorio de trabajo actual.
- `ls`: Lista los archivos y directorios en el directorio actual.
- `whoami`: Muestra el nombre del usuario actual.
- `uname -a`: Muestra información del sistema operativo.
- `cat /etc/os-release`: Muestra información sobre la distribución de Linux.
- `df -h`: Muestra el uso del espacio en disco en formato legible.
- `free -h`: Muestra la cantidad de memoria libre y utilizada en el sistema.
- `top`: Muestra una vista dinámica en tiempo real de los procesos en ejecución. Para salir, presione `Control+C`
- `apt update`: Actualiza la lista de paquetes disponibles.
- `apt install -y [paquete]`: Instala un paquete específico (reemplaza [paquete] con el nombre del paquete que deseas instalar).
- `exit`: Sale de la terminal. En este caso se terminará el contenedor.

Los comandos les será util cuando tenga que depurar un contenedor "en caliente".

#### Salir del contenedor

Ejecutamos:

```bash
exit
```

## Conclusión

En este laboratorio instalamos Docker Desktop que nos permitió correr comandos con la CLI de Docker. Los mismos pasos podrían haberse realizado desde la interface gráfica, pero en este curso nos enfocaremos en la CLI de Docker.
Además, ejecutamos un contenedor de Ubuntu, donde podemos instalar paquetes de la misma manera que lo haríamos en un GNU/Linux Ubuntu nativo.

---------
