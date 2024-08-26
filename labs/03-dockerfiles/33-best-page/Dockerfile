# Usar la imagen base de Rocky Linux
FROM almalinux:latest

# Instalar Nginx
RUN yum update -y && \
    yum install -y epel-release && \
    yum install -y nginx && \
    yum clean all

# Copiar el archivo index.html al directorio por defecto de Nginx
COPY index.html /usr/share/nginx/html/index.html

# Exponer el puerto 80
EXPOSE 80

# Comando para iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
