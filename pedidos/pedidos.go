package pedidos

// CrearPedido recibirá una solicitud de compra de un repuesto y su cantidad.
func CrearPedido(idPedido, idCliente, codigoProducto string, cantidad int) (bool, string) {
	return false, "Creación de pedidos prevista para una etapa posterior."
}

// ConsultarPedido buscará el estado y el total del pedido.
func ConsultarPedido(idPedido string) (estado string, totalCentavos int, encontrado bool, mensaje string) {
	return "", 0, false, "Consulta de pedidos prevista para una etapa posterior."
}

// CancelarPedido marcará como cancelado un pedido que admita cancelación.
func CancelarPedido(idPedido string) (bool, string) {
	return false, "Cancelación de pedidos prevista para una etapa posterior."
}
