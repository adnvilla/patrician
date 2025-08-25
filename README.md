# Patrician
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fadnvilla%2Fpatrician.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fadnvilla%2Fpatrician?ref=badge_shield)

Simulador y API HTTP inspirada en el comercio de la saga *Patrician*. El servidor expone endpoints para consultar ciudades, mercancías y distancias, así como actualizar el mercado de cada ciudad.

## Requisitos

- [Go](https://go.dev/) 1.24 o superior
- `make` (opcional, para utilizar el Makefile)

## Instalación y ejecución

```bash
git clone https://github.com/adnvilla/patrician.git
cd patrician
make build    # compila el binario en ./bin/patrician
make run      # inicia el servidor en el puerto 8080
```

Alternativamente puede ejecutarse directamente:

```bash
go run main.go
```

## Endpoints principales

- `GET /cities`
- `GET /commodities`
- `GET /distances`
- `GET /city/{name}/commodities`
- `POST /city/{name}/commodity`
- `POST /city/{name}/commodities`
- `GET /city/{name}/stock`
- `GET /city/{name}/supply/{city}`

La especificación completa de la API se encuentra en [`docs/openapi.yaml`](docs/openapi.yaml).

## Pruebas

```bash
go test ./...
```

## Licencia

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fadnvilla%2Fpatrician.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fadnvilla%2Fpatrician?ref=badge_large)
