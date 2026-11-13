package main

import (
	"fmt"
	"os"

	"farmacias-automotriz-ob/inventario"
)

func main() {
	if err := inventario.CargarExistencias(); err != nil {
		fmt.Fprintln(os.Stderr, "No se pudo cargar el inventario:", err)
		os.Exit(1)
	}
	iniciarConsola()
}
