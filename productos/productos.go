package productos

import "strings"

const (
	CodigoBujiaCarro = "BUJ-001"
	CodigoBujiaMoto  = "BUJ-002"
)

// ResolverCodigo convierte el nombre escrito en el código interno del repuesto.
// Devuelve una cadena vacía si no reconoce el producto.
func ResolverCodigo(consulta string) string {
	texto := strings.Join(strings.Fields(strings.ToLower(consulta)), " ")
	texto = strings.ReplaceAll(texto, "í", "i")

	switch texto {
	case "bujia carro", "bujia de carro", "buj-001":
		return CodigoBujiaCarro
	case "bujia moto", "bujia de moto", "buj-002":
		return CodigoBujiaMoto
	default:
		return ""
	}
}

// RegistrarProducto recibirá los datos de un producto nuevo y validará su código.
func RegistrarProducto(codigo, nombre, categoria string, precioCentavos int) (bool, string) {
	return false, "Registro de productos previsto para una etapa posterior."
}

// ConsultarProducto busca un repuesto de demostración por su código interno.
func ConsultarProducto(codigo string) (nombre, categoria string, precioCentavos int, encontrado bool) {
	switch codigo {
	case CodigoBujiaCarro:
		return "Bujía para carro", "Bujías", 650, true
	case CodigoBujiaMoto:
		return "Bujía para moto", "Bujías", 450, true
	default:
		return "", "", 0, false
	}
}

// ModificarProducto actualizará los datos comerciales del código indicado.
func ModificarProducto(codigo, nombre, categoria string, precioCentavos int) (bool, string) {
	return false, "Modificación de productos prevista para una etapa posterior."
}
