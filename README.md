# Farmacias Automotriz OB

Farmacias Automotriz OB es un proyecto de gestión para una tienda de repuestos e insumos de autos y motos en Quito. Busca que el administrador pueda consultar en un solo lugar el catálogo, las existencias y los pedidos que hoy se apuntarían en cuadernos o en WhatsApp. La tienda es una propuesta de negocio. La instalación de repuestos se plantea como servicio futuro del taller.

## Aprendizaje Autónomo 1

El proyecto organiza diez funciones de negocio en cuatro paquetes. El menú de consola ejecuta tres operaciones con datos de prueba. Los otros siete casos aparecen como previstos en el menú y tienen sus firmas en el código.

| Paquete | Operaciones de demostración | Funciones declaradas para después |
|---|---|---|
| `productos` | Consultar producto con su precio y existencias | Registrar y modificar producto |
| `inventario` | Consultar y ajustar existencias con guardado automático | |
| `clientes` | | Registrar y consultar cliente |
| `pedidos` | | Crear, consultar y cancelar pedido |

La demostración empieza con ocho bujías para carro a USD 6.50 y cinco para moto a USD 4.50. Escribe `bujía carro` o `bujía moto` para elegir el producto. Se aceptan también sin tilde y en mayúsculas. Los ajustes se guardan automáticamente en `existencias.json`. Al iniciar desde la misma carpeta se recuperan las últimas cantidades guardadas. La etapa usa funciones y tipos básicos de Go, sin estructuras propias ni paquetes de terceros. El diseño contempla un administrador. Los pagos, la facturación, el envío, la base de datos y la atención del taller quedan fuera de esta etapa.

## Cómo ejecutar

Abre una terminal en la carpeta que contiene `go.mod`. Usa Go 1.22 o posterior.

```sh
go build ./...
go run .
```

1. Elige `2` y escribe `bujía carro`. Verás el nombre, el precio unitario y las unidades disponibles. En una copia nueva empieza con ocho.
2. Elige `2` con `bujía moto` para ver su precio y su saldo. Al inicio son cinco unidades.
3. Elige `5`, escribe `bujía carro` e ingresa `-2`. Si partiste de ocho, quedarán seis unidades. Sal con `0`, ejecuta otra vez `go run .` desde la misma carpeta y consulta con la opción `2` o `4`. Seguirán siendo seis.
4. Elige `8` para ver cómo se anuncia una operación prevista. Elige `0` para salir.

## Organización

```text
farmacias-automotriz-ob/
├── go.mod
├── README.md
├── .gitignore
├── main.go
├── consola.go
├── productos/productos.go
├── inventario/
│   ├── inventario.go
│   └── archivo.go
├── clientes/clientes.go
└── pedidos/pedidos.go
```

`go.mod` identifica el módulo de Go `farmacias-automotriz-ob`. `main.go` inicia el programa y `consola.go` gestiona el menú, la entrada y los mensajes. Ambos archivos pertenecen a `package main`. La consola pide el nombre, `productos` lo traduce a un código interno e `inventario` usa ese código para consultar o ajustar el saldo del repuesto correcto. `inventario/archivo.go` lee los saldos al iniciar y los guarda al aceptar un ajuste. El archivo de datos se crea en la carpeta desde la que ejecutas el programa. `.gitignore` evita que Git incluya ese archivo local y los temporales de guardado.

## Datos guardados

El archivo guarda las cantidades de los dos repuestos. Para respaldarlas, copia `existencias.json` con el programa cerrado. Para recuperarlas en otra copia, coloca ese archivo junto a `go.mod` antes de ejecutar `go run .`. Usa una sesión del programa a la vez para trabajar con ese inventario.

## Evolución prevista

Una versión posterior podrá listar varias marcas por tipo de repuesto y registrar clientes y pedidos. Cada pedido tendrá un código de cliente y una referencia al repuesto. El diseño prevé reservar existencias al registrar el pedido y devolverlas una sola vez si se cancela antes de entregar el producto. Para una futura instalación, una orden de trabajo vincularía pedido, vehículo y mano de obra. El repuesto reservado se entregaría al taller sin volver a descontar las mismas unidades.

## Material del curso

- [Guía del Aprendizaje Autónomo 1](https://github.com/carlosandresat/uide-poo-go-online/blob/main/entregas/aa1.md).
- [Semana 1 · Menú con `for` y `switch`](https://github.com/carlosandresat/uide-poo-go-online/blob/main/semana-01/ejemplos/menu/main.go).
- [Semana 2 · Ejercicio de funciones y paquetes](https://github.com/carlosandresat/uide-poo-go-online/tree/main/semana-02/ejemplos/ejercicio-paquetes).
