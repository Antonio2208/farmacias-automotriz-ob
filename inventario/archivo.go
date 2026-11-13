package inventario

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"farmacias-automotriz-ob/productos"
)

const archivoExistencias = "existencias.json"

// CargarExistencias recupera el saldo de ambos repuestos al abrir el programa.
func CargarExistencias() error {
	datos, err := os.ReadFile(archivoExistencias)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	// Un puntero vacío distingue un dato ausente o nulo de un saldo válido de cero.
	var cantidades map[string]*int
	if err := json.Unmarshal(datos, &cantidades); err != nil {
		return fmt.Errorf("%s no contiene un inventario válido: %w", archivoExistencias, err)
	}
	carro := cantidades[productos.CodigoBujiaCarro]
	moto := cantidades[productos.CodigoBujiaMoto]
	if len(cantidades) != 2 || carro == nil || moto == nil || *carro < 0 || *moto < 0 {
		return fmt.Errorf("%s debe contener un saldo entero no negativo por cada repuesto de prueba", archivoExistencias)
	}
	unidadesCarro, unidadesMoto = *carro, *moto
	return nil
}

func guardarExistencias(codigo string, nuevoSaldo int) error {
	cantidades := map[string]int{
		productos.CodigoBujiaCarro: unidadesCarro,
		productos.CodigoBujiaMoto:  unidadesMoto,
	}
	cantidades[codigo] = nuevoSaldo
	datos, err := json.MarshalIndent(cantidades, "", "  ")
	if err != nil {
		return err
	}

	// Escribe primero un temporal y reemplaza el archivo solo cuando está completo.
	temporal, err := os.CreateTemp(".", ".existencias-*.tmp")
	if err != nil {
		return err
	}
	defer os.Remove(temporal.Name())
	defer temporal.Close()

	if _, err := temporal.Write(append(datos, '\n')); err != nil {
		return err
	}
	if err := temporal.Sync(); err != nil {
		return err
	}
	if err := temporal.Close(); err != nil {
		return err
	}
	return os.Rename(temporal.Name(), archivoExistencias)
}
