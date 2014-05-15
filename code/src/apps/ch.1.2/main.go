//Código de ejemplo para el capítulo 1.2 del "Build aplicación Web con Golang"
//Propósito: Ejecute el archivo para comprobar si el espacio de trabajo está correctamente configurado.
// Para ejecutar, vaya al directorio actual en una consola y teclear `go run main.go
// Si no se muestra el texto "Hello World", la configuración no esta bien.
package main

import (
	"fmt"
	"mymath" //este paquete debe estar ya creado
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
