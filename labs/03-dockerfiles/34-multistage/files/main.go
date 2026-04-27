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
