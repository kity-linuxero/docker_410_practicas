# Laboratorio 1.1 - Instalación de Docker

### Objetivos:
- Correcta instalación de Docker en nuestra PC de prueba.



## 1. Instalación Docker Desktop

Docker Desktop se describe como _una aplicación de instalación con un solo clic para su entorno Mac, Linux* o Windows que le permite crear, compartir y ejecutar aplicaciones y microservicios en contenedores_. Nos permitirá una instalación sencilla en entornos Windows.

* [No compatible con Ubuntu 24.04](https://docs.docker.com/desktop/install/ubuntu/#prerequisites). Para tal caso se recomienda la [siguiente instalación](https://docs.docker.com/engine/install/ubuntu/).

### Instalación en Windows

Asegurese de [cumplir con los requisitos del sistema](https://docs.docker.com/desktop/install/windows-install/#system-requirements).
- WSL version 1.1.3.0 o superior.
- Windows 11 64-bit: Home or Pro version 21H2 or higher, or Enterprise or Education version 21H2 or higher.
- Windows 10 64-bit:
    - We recommend Home or Pro 22H2 (build 19045) or higher, or Enterprise or Education 22H2 (build 19045) or higher.
    - Minimum required is Home or Pro 21H2 (build 19044) or higher, or Enterprise or Education 21H2 (build 19044) or higher.
- Procesador de 64 bits con [SLAT](https://en.wikipedia.org/wiki/Second_Level_Address_Translation)
- 4 GB RAM
- Virtualización por HW habilitada [Ver virtualización](https://docs.docker.com/desktop/troubleshoot/topics/#virtualization)


[Ver mas](https://docs.docker.com/desktop/install/windows-install/#system-requirements).


#### Descargue Docker Desktop <a href="https://hub.docker.com/" target="_blank">Docker Hub</a>


### Una vez descargado el ejecutable lo ejecutamos.
![](https://docker.idepba.com.ar/img/clase1/Docker_install.png)

### Aceptamos los términos

![](https://docker.idepba.com.ar/img/clase1/Docker_install2.png)


### Elegimos una opción

![](https://docker.idepba.com.ar/img/clase1/Docker_install3.png)

### Docker Instalado

![](https://docker.idepba.com.ar/img/clase1/Docker_install4.png)

### Elegimos una opción o bien "Skip survey"

![](https://docker.idepba.com.ar/img/clase1/Docker_install5.png)

### Iniciando Docker
![](https://docker.idepba.com.ar/img/clase1/Docker_install6.png)

### Docker iniciado
![](https://docker.idepba.com.ar/img/clase1/Docker_install7.png)

### Abrimos una ventana de PowerShell:

Escribimos:

```powershell
docker --version
```

![](https://docker.idepba.com.ar/img/clase1/Docker_install8.png)



