# Laboratorio 5.4 - Benchmark de Almacenamiento en Docker

## Objetivo

  - Comparar el rendimiento de lectura/escritura y latencia entre **Bind Mounts**, **Named Volumes** y **tmpfs**.
  - Utilizar la herramienta `dd` para pruebas de rendimiento secuencial.
  - Compartir el resultado en el foro de la clase [foro clase en Campus](https://campus.idepba.com.ar/mod/forum/discuss.php?d=21)

## 1\. Preparación del Entorno

Primero, creamos los recursos necesarios en el host:

```bash
mkdir ~/bench_test
docker volume create bench_vol
```

## 2\. Prueba 1: Bind Mount (Directorio del Host)

Los Bind Mounts dependen directamente del sistema de archivos y el rendimiento del disco de tu host.

```bash
docker run --rm --mount type=bind,src=$PWD/bench_test,dst=/data alpine sh -c "echo '--- BIND MOUNT ---' && dd if=/dev/zero of=/data/test.bin bs=1M count=1000 conv=fsync"
```

#### Resultados de ejemplo*

```bash
--- BIND MOUNT ---
1000+0 records in
1000+0 records out
1048576000 bytes (1000.0MB) copied, 0.566914 seconds, 1.7GB/s
```

Velocidad: 1.7GB/s

## 3\. Prueba 2: Named Volume (Gestionado por Docker)

En Linux, los volúmenes suelen tener un rendimiento casi idéntico al host, pero son gestionados íntegramente por el motor de Docker.

```bash
docker run --rm --mount type=volume,src=bench_vol,dst=/data alpine sh -c "echo '--- NAMED VOLUME ---' && dd if=/dev/zero of=/data/test.bin bs=1M count=1000 conv=fsync"
```

#### Resultados de ejemplo*

```bash
--- NAMED VOLUME ---
1000+0 records in
1000+0 records out
1048576000 bytes (1000.0MB) copied, 0.660572 seconds, 1.5GB/s

```

Velocidad: 1.5GB/s

## 4\. Prueba 3: tmpfs (Memoria RAM)

Esta es la opción más rápida. Los datos se escriben directamente en la **memoria RAM**. Al no haber un disco físico, no usamos `conv=fsync`.

```bash
docker run --rm --mount type=tmpfs,dst=/data alpine sh -c "echo '--- TMPFS (RAM) ---' && dd if=/dev/zero of=/data/test.bin bs=1M count=1000"
```

#### Resultados de ejemplo*

```bash
--- TMPFS (RAM) ---
1000+0 records in
1000+0 records out
1048576000 bytes (1000.0MB) copied, 0.396952 seconds, 2.5GB/s
```

Velocidad: 2.5GB/s


*Los resultados pueden variar según el hardware utilizado.


-----

## 5\. Análisis de Resultados

Resultados realizado en una Laptop con disco SSD corriendo GNU/Linux.

| Tipo de Montaje | Velocidad | Caso de Uso Ideal |
| :--- | :--- | :--- |
| **Bind Mount** | 1.7GB/s | Código fuente en desarrollo (Hot-reload). |
| **Named Volume** | 1.5GB/s | Bases de datos y persistencia en producción. |
| **tmpfs** | 2.5GB/s | Secretos, caché temporal o sesiones de alta velocidad. |


## 6\. Compartir resultados

Inicie sesión en el [campus](https://campus.idepba.com.ar/mod/forum/discuss.php?d=21) y comparta sus resultados en el hilo de la clase correspondiente.

## 7\. Limpieza

```bash
docker volume rm bench_vol
rm -r ~/bench_test
```

## Resumen de Lab

En este laboratorio vimos las diferencias de performance entre bind mounts, named volumes y tmpfs, utilizando la herramienta dd para medir el rendimiento de los volúmenes.


-----

<p align="center"\>
<img src="../../img/logos.footer.gray.webp"\>
</p\>

-----
