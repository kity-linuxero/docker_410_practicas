# Guía de Laboratorios curso Fundamentos de Docker

![Docker](https://img.shields.io/badge/Docker-Container-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Version](https://img.shields.io/badge/Version-1.0-orange)




## Índice
1. [Introducción a Docker](#introducción-a-docker)
2. [Conceptos básicos - Imágenes, registry y contenedores](#conceptos-básicos)
3. [Dockerfile y personalizar imágenes](#dockerfile-y-conteinerizando-nuestras-primeras-apps)
4. [Redes y persistencia de datos en Docker](#redes-y-persistencia-de-datos-en-docker)


## Introducción a Docker

En estos laboratorios aprenderemos a instalar Docker y daremos los primeros pasos con una terminal de Docker y la terminal de bash sobre un GNU/Linux Ubuntu.

### Laboratorio 1

1. [Instalación de Docker](./labs/01-introduccion/instalacion.md)
2. [Primeros comandos](./labs/01-introduccion/primeros-comandos.md)


## Conceptos básicos

En estos laboratorios entenderemos en la práctica la diferencia entre imágen y contenedores. Versionaremos imágenes para para crear nuestras imágenes y exportaremos nuestras imágenes a archivos para luego poder importarlas.

### Laboratorio 2

1. [Imagenes, registry y contenedores](./labs/02-conceptos-basicos/21-images-registry-container.md)
2. [Imágenes y versiones](./labs/02-conceptos-basicos/22-images-tags.md)


## Dockerfile y Conteinerizando nuestras primeras apps

En estos laboratorios vamos a crear nuestras imagenes de forma automática usando `Dockerfiles` para conteinerizar aplicaciones. Vamos a hostear una página estática html y una aplicación Python.

### Laboratorio 3

1. [Hola Docker](./labs/03-dockerfiles/31-holamundo/README.md)
2. [Una vaca saluda](./labs/03-dockerfiles/31-cow/README.md)
3. [La mejor página](./labs/03-dockerfiles/32-best-page/README.md)

## Redes y persistencia de datos en Docker

En estos laboratorios vamos crear nuestras primeras redes y volúmenes para persistir datos. Comprobaremos la conectividad entre contenedores y migración de datos.

### Laboratorio 4

1. [Introducción a redes en Docker](./labs/04-redes_volumes/redes/41-introduccion.md)
2. [Conexión de Múltiples Contenedores en una Red Bridge](./labs/04-redes_volumes/redes/42-containers_net.md)
3. [Trabajando datos persistentes en Docker](./labs/04-redes_volumes/volumes/43-volumenes_docker.md)
3. [Backup y migración de volúmenes](./labs/04-redes_volumes/volumes/44-volumenes_bkp.md)

---------

Centro de Formación CFL 410 - Omar Nuñez

IDEP | Instituto de Estudios sobre Estado y Participación
