# Imágen base
FROM ubuntu:latest

# Ejecuto actualización de repositorios e instalo el servidor web nginx
RUN apt-get update && apt-get install -y nginx

# Copio el archivo index para que sea visible desde el nginx del contenedor
COPY index.html /var/www/html/

# Comando a ejecutar
CMD ["nginx", "-g", "daemon off;"]

# Expongo el puerto 80
EXPOSE 80
