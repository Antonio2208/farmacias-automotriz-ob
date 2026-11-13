package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"farmacias-automotriz-ob/inventario"
	"farmacias-automotriz-ob/productos"
)

func iniciarConsola() {
	entrada := bufio.NewScanner(os.Stdin)
	ejecutarMenu(entrada)

	if err := entrada.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "No se pudo leer la entrada:", err)
	}
}

func ejecutarMenu(entrada *bufio.Scanner) {
	fmt.Println("Farmacias Automotriz OB | Repuestos para autos y motos")
	fmt.Println("Productos de prueba: bujía carro y bujía moto.")

	for {
		mostrarMenu()
		opcion, leido := leerLinea(entrada, "Opción: ")
		if !leido {
			return
		}

		switch opcion {
		case "0":
			fmt.Println("Hasta luego.")
			return
		case "2", "4", "5":
			consulta, leido := leerLinea(entrada, "Producto: ")
			if !leido {
				return
			}
			codigo := productos.ResolverCodigo(consulta)
			if codigo == "" {
				fmt.Println("Producto no encontrado. Prueba con bujía carro o bujía moto.")
				continue
			}
			switch opcion {
			case "2":
				mostrarProducto(codigo)
			case "4":
				mostrarExistencias(codigo)
			case "5":
				if !ajustarInventario(entrada, codigo) {
					return
				}
			}
		case "1", "3", "6", "7", "8", "9", "10":
			fmt.Println("Función declarada. Su implementación corresponde a una etapa posterior.")
		default:
			fmt.Println("Opción no válida. Elige un número del 0 al 10.")
		}
	}
}

func mostrarMenu() {
	fmt.Print(`
PRODUCTOS
 1. Registrar producto (prevista)
 2. Consultar producto
 3. Modificar producto (prevista)
INVENTARIO
 4. Consultar existencias
 5. Ajustar existencias
CLIENTES
 6. Registrar cliente (prevista)
 7. Consultar cliente (prevista)
PEDIDOS
 8. Crear pedido (prevista)
 9. Consultar pedido (prevista)
10. Cancelar pedido (prevista)
 0. Salir
`)
}

func leerLinea(entrada *bufio.Scanner, mensaje string) (string, bool) {
	fmt.Print(mensaje)
	if !entrada.Scan() {
		return "", false
	}
	return strings.TrimSpace(entrada.Text()), true
}

func mostrarProducto(codigo string) {
	nombre, _, precioCentavos, encontrado := productos.ConsultarProducto(codigo)
	if !encontrado {
		fmt.Println("Producto no encontrado en la demostración.")
		return
	}

	fmt.Println("Producto:", nombre)
	fmt.Printf("Precio unitario: USD %d.%02d\n", precioCentavos/100, precioCentavos%100)
	mostrarExistencias(codigo)
}

func mostrarExistencias(codigo string) {
	unidades, encontrado := inventario.ConsultarExistencias(codigo)
	if !encontrado {
		fmt.Println("Existencias no disponibles.")
		return
	}
	fmt.Printf("Existencias: %d unidades\n", unidades)
}

// ajustarInventario devuelve false cuando se termina la entrada de la consola.
func ajustarInventario(entrada *bufio.Scanner, codigo string) bool {
	linea, leido := leerLinea(entrada, "Variación (ej. +3 para entrada o -2 para salida): ")
	if !leido {
		return false
	}

	variacion, err := strconv.Atoi(linea)
	if err != nil {
		fmt.Println("Escribe un número entero para ajustar el inventario.")
		return true
	}

	unidades, ajustado, mensaje := inventario.AjustarExistencias(codigo, variacion)
	fmt.Println(mensaje)
	if ajustado {
		fmt.Println("Existencias después del ajuste:", unidades)
	}
	return true
}
