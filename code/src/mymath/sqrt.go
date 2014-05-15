// Código de ejemplo para el capítulo 1.2 del "Build aplicación Web con Golang"
// Objetivo: Muestra cómo crear un paquete sencillo llamado `MyMath`
// Este paquete puede  ser importado de otro archivo para que funcione.
package mymath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
