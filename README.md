# Guía de Laboratorios curso Fundamentos de Docker



[![Docker](https://badgen.net/badge/icon/docker?icon=docker&label)](https://docker.com/)
[![Supported by](https://img.shields.io/badge/Supported%20by-CFL410-green.svg)](https://centro410laplata.edu.ar/)
[![Supported by](https://img.shields.io/badge/Supported%20by-IDEP-green.svg)](https://idepba.com.ar/)
[![Powered](https://img.shields.io/badge/Powered%20by-ATE-green.svg)](https://atepba.org.ar/)
![GitHub commits](https://badgen.net/github/commits/kity-linuxero/docker_410)
![Version](https://img.shields.io/badge/Version-1.1-orange)


## 🐳 Índice de temas
1. [Introducción a Docker](#1-introducción-a-docker)
2. [Conceptos básicos - Imágenes, registry y contenedores](#2-conceptos-básicos)
3. [Dockerfile y personalizar imágenes](#3-dockerfile-y-containerizando-nuestras-primeras-apps)
4. [Redes y persistencia de datos en Docker](#4-redes-y-persistencia-de-datos-en-docker)
5. [Docker compose](#5-docker-compose)
6. [Docker compose Parte II y depuración de contenedores](#6-docker-compose-parte-ii-y-depuración)
7. [Awesome Compose](#7-awesome-compose)

## 1. Introducción a Docker

En estos laboratorios aprenderemos a instalar Docker y daremos los primeros pasos con una terminal sobre un contenedor Linux.


### Prácticas Lab 1

1. [Instalación de Docker](./labs/01-introduccion/instalacion.md)
2. [Primeros comandos](./labs/01-introduccion/primeros-comandos.md)


## 2. Conceptos básicos

En estos laboratorios entenderemos en la práctica la diferencia entre imágen y contenedores. Versionaremos imágenes para para crear nuestras imágenes y exportaremos nuestras imágenes a archivos para luego poder importarlas.

### Prácticas Lab 2

1. [Primeros comandos en Docker](./labs/02-conceptos-basicos/21-cli-primeros-comandos.md)
2. [Imagenes, registry y contenedores](./labs/02-conceptos-basicos/22-images-registry-container.md)
3. [Imágenes y versiones](./labs/02-conceptos-basicos/23-images-tags.md)
4. [Guardando imágenes en DockerHub](./labs/02-conceptos-basicos/24-images-push.md)


## 3. Dockerfile y Containerizando nuestras primeras apps

En estos laboratorios vamos a crear nuestras imagenes de forma automática usando `Dockerfiles` para conteinerizar aplicaciones. Vamos a hostear una página estática html y una aplicación Python.

### Prácticas Lab 3

1. [Hola Docker](./labs/03-dockerfiles/31-holamundo/README.md)
2. [Una vaca saluda](./labs/03-dockerfiles/31-cow/README.md)
3. [La mejor página](./labs/03-dockerfiles/32-best-page/README.md)

## 4. Redes y persistencia de datos en Docker

En estos laboratorios vamos crear nuestras primeras redes y volúmenes para persistir datos. Comprobaremos la conectividad entre contenedores y migración de datos.

#### Prácticas Lab 4: Redes
1. [Introducción a redes en Docker](./labs/04-redes_volumes/redes/41-introduccion.md)
2. [Conexión de Múltiples Contenedores en una Red Bridge](./labs/04-redes_volumes/redes/42-containers_net.md)

#### Prácticas Lab4: Volúmenes
3. [Trabajando datos persistentes en Docker](./labs/04-redes_volumes/volumes/43-volumenes_docker.md)
4. [Backup y migración de volúmenes](./labs/04-redes_volumes/volumes/44-volumenes_bkp.md)


## 5. Docker compose

En estos laboratorios comprobará las ventajas y sencillez de ejecutar contenedores con Docker Compose.

### Prácticas Lab 5

1. [Deployar un contenedor web sencillo con Docker Compose](./labs/05-compose/51-intro.compose.md)
2. [Aplicaciones multicontenedor](./labs/05-compose/52-multicontainer.md)

## 6. Docker compose parte II y depuración

En estos laboratorios veremos algunas atributos mas avanzados de Docker Compose. También veremos nociones básicas de depuración de contenedores.

> [!CAUTION]
> **PROXIMAMENTE**


## 7. Awesome Compose

En este laboratorio exploraremos composes compartidos por la comunidad listos para deployar y usar.

> [!CAUTION]
> **PROXIMAMENTE**


---------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="img/logos.footer.gray.webp">
  </a>
</p>

