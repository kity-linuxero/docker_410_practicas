# Laboratorio 3.4 - Multi-stage builds: Optimizando imágenes

## Objetivos:

* Comprender el concepto de **Multi-stage build**.
* Reducir el tamaño de las imágenes finales.
* Separar el entorno de desarrollo y compilación (herramientas no necesarias para el entorno de producción) del entorno de ejecución.

> [!TIP]
> Esta técnica es la que se utiliza en para que las imágenes sean más seguras y livianas.


En este lab tenemos una aplicación escrita en **Go**. Para compilarla y generar el ejecutable, necesitamos el compilador de Go, pero para que la aplicación funcione en el servidor, solo necesitamos el archivo binario resultante.


## 1. El escenario: El entorno de desarrollo

* Crea una carpeta llamada `multi_stage`.
* Dentro de esa carpeta crea un archivo llamado `main.go` con el siguiente contenido, que será el código fuente de la aplicación:

```go
package main
import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola! Soy un binario optimizado con Multi-stage")
	})
	fmt.Println("Servidor corriendo en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}

```

## 2. Analizar el archivo `Dockerfile`:

Crea el archivo `Dockerfile` y analiza cómo dividimos el proceso en dos etapas usando la palabra clave `AS`:

```dockerfile
# ETAPA 1: Compilación (La llamamos "builder")
FROM golang:1.22-alpine AS builder

# Seteamos directorio de trabajo
WORKDIR /app

# Copiamos el código fuente
COPY main.go .

# Compilamos el binario
RUN go build -o mi-app-binario main.go


# ETAPA 2: Ejecución (Imagen final - Sin el entorno de desarrollo)
FROM alpine:latest

# Seteamos directorio de trabajo
WORKDIR /root/

# Copiamos SOLO el binario desde la etapa "builder"
COPY --from=builder /app/mi-app-binario .

# Exponemos el puerto y ejecutamos
EXPOSE 8080
CMD ["./mi-app-binario"]

```

> [!IMPORTANT]
> Fíjate que en la segunda etapa (`FROM alpine`), no instalamos Go. Solo traemos el archivo compilado de la etapa anterior mediante el flag `--from=builder`.

## 3. Construir la Imágen Docker

* Abre una terminal en la carpeta y construye la imagen:

```bash
docker build -t app-optimizada:1.0 .
```

## 4. Verificar y comparar tamaños

Este es el punto clave del laboratorio. Vamos a comparar el peso de nuestra imagen optimizada vs. la imagen que necesitamos para compilar.

* Ejecuta el siguiente comando para ver las imágenes:

```bash
docker images
```


### Análisis del resultado:

Solamente verás la imágen final generada. Ya que la imágen intermedia quedó en la `build cache`.


```bash
IMAGE                     ID             DISK USAGE   CONTENT SIZE
app-optimizada:1.0        cc5a47326e4f         24MB          7.9MB
```


## 5. Correr el contenedor:

```bash
docker run -d --name mi-app-multi -p 80:8080 app-optimizada:1.0
```

* Accede desde tu navegador a: [http://localhost](http://localhost)
* Deberías ver el mensaje del programa en Go.

**¿Cuánto debería haber ocupado la imágen si no hubiesemos usado multi-stage?**

## 6. Imágen sin usar multi-stage

- Crear un nuevo archivo llamado `Dockerfile.dev` con el siguiente contenido:

```Dockerfile
# Usamos la imagen de Go (que contiene todo el kit de compilación)
FROM golang:1.22-alpine

# Seteamos directorio de trabajo
WORKDIR /app

# Copiamos el código fuente
COPY main.go .

# Compilamos el binario
RUN go build -o mi-app-binario main.go

# Exponemos el puerto
EXPOSE 8080

# Ejecutamos el binario resultante
CMD ["./mi-app-binario"]
```

- Construir la imágen:

```bash
docker build -t app-no-multi:1.0 -f Dockerfile.dev .
```

- Verificar que la imágen fue creada:

```bash
docker images
```

```bash
IMAGE                     ID             DISK USAGE   CONTENT SIZE   EXTRA
app-no-multi:1.0          8ef6bdbd6d96        445MB         93.3MB        
app-optimizada:1.0        cc5a47326e4f         24MB          7.9MB    U   
```

Observar la diferencia de tamaño.

> [!TIP]
> `DISK USAGE`: Tamaño de la imagen en el disco.
> `CONTENT SIZE`: Tamaño de la imagen en la registry.


## 7. Limpiar imágenes contenedores e imágenes generadas

- Eliminar contenedor
    ```bash
    docker rm mi-app-multi
    ```
- Eliminar imágenes:
    ```bash
    docker rmi app-optimizada:1.0 app-no-multi:1.0
    ```
    


## Conclusión

En este lab hemos aplicado **Multi-stage builds**. Esta técnica es fundamental porque:

1. **Seguridad:** La imagen final no tiene el código fuente ni herramientas de compilación.
2. **Eficiencia:** Las imágenes son mucho más pequeñas, lo que acelera el despliegue y ahorra espacio en disco.

---------------

<p align="center">
  <a href="https://centro410laplata.edu.ar/">
    <img src="../../../img/logos.footer.gray.webp">
  </a>
</p>
