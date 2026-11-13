package inventario

import (
	"fmt"
	"math"

	"farmacias-automotriz-ob/productos"
)

// Cantidades de prueba que se usan cuando todavía no hay un archivo guardado.
var (
	unidadesCarro = 8
	unidadesMoto  = 5
)

// saldoDe devuelve la dirección del saldo del repuesto solicitado.
func saldoDe(codigo string) *int {
	switch codigo {
	case productos.CodigoBujiaCarro:
		return &unidadesCarro
	case productos.CodigoBujiaMoto:
		return &unidadesMoto
	default:
		return nil
	}
}

// ConsultarExistencias devuelve las unidades disponibles del código solicitado.
func ConsultarExistencias(codigo string) (unidades int, encontrado bool) {
	saldo := saldoDe(codigo)
	if saldo == nil {
		return 0, false
	}
	return *saldo, true
}

// AjustarExistencias valida y guarda una entrada o salida.
func AjustarExistencias(codigo string, variacion int) (unidades int, ajustado bool, mensaje string) {
	saldo := saldoDe(codigo)
	if saldo == nil {
		return 0, false, "Producto sin inventario en la demostración."
	}
	actuales := *saldo
	if variacion == 0 {
		return actuales, false, "Escribe un ajuste distinto de cero."
	}
	if variacion < -actuales {
		return actuales, false, "Existencias insuficientes. El ajuste dejaría saldo negativo."
	}
	if variacion > math.MaxInt-actuales {
		return actuales, false, "El ajuste supera el máximo que admite el programa."
	}

	nuevoSaldo := actuales + variacion
	if err := guardarExistencias(codigo, nuevoSaldo); err != nil {
		return actuales, false, fmt.Sprintf("No se guardó el ajuste: %v", err)
	}
	*saldo = nuevoSaldo
	return nuevoSaldo, true, "Ajuste aplicado y guardado."
}
