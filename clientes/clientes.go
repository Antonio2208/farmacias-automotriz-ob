package clientes

// RegistrarCliente recibirá un identificador interno y el nombre del cliente.
func RegistrarCliente(idCliente, nombre string) (bool, string) {
	return false, "Registro de clientes previsto para una etapa posterior."
}

// ConsultarCliente buscará el nombre asociado con un identificador interno.
func ConsultarCliente(idCliente string) (nombre string, encontrado bool, mensaje string) {
	return "", false, "Consulta de clientes prevista para una etapa posterior."
}
