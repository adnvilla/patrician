# Patrician Codebase Instructions

## Project Overview
Patrician is a Go HTTP API simulating a trading system inspired by the Patrician game series. The API exposes endpoints for querying cities, commodities, and distances, as well as updating market data for each city.

## Architecture

### Domain-Driven Design
The codebase follows a clean architecture with Domain-Driven Design principles:
- `domain/`: Core business entities and rules (City, Commodity, MarketHall)
- `application/use_cases/`: Application business logic implementing the use case pattern
- `interfaces/handlers/`: HTTP API controllers using Gin framework
- `infrastructure/`: External dependencies like database connections
- `pkg/`: Shared utilities and interfaces

### Key Patterns
1. **Use Case Pattern**: The `UseCase<TRequest, TResponse>` interface in `pkg/use_case/use_case.go` is the core abstraction for business logic. All use cases implement this interface.

2. **Repository Pattern**: Data access is abstracted through repositories like `CityRepository` and `CommoditiesRepository` in `infrastructure/data/postgresql/`.

3. **OpenAPI-driven Development**: The API is defined in `docs/openapi.yaml` and validated at runtime.

## Developer Workflows

### Building and Running
- `make build` or `build.bat build` - Compiles the binary to `./bin/patrician`
- `make run` or `build.bat run` - Builds and runs the server on port 8080
- `go run main.go` - Alternative way to run the server

### Testing
- `make test` or `build.bat test` - Run all tests
- `make test-verbose` or `build.bat test-verbose` - Run tests with verbose output
- `make test-coverage` or `build.bat test-coverage` - Generate coverage report

### Database
- Set `POSTGRES_DSN` environment variable for database connection
- Default connection falls back to local PostgreSQL: `host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable`

## Key Concepts

### City Trading Network
Cities have:
- Market Halls containing commodities
- Distance relationships with other cities
- Supply/demand metrics for commodities

### Stock Calculation
Stock is calculated as `Production - Consumption` (see `domain/commodity.go`).

Supply between cities considers distance:
```go
// In domain/city.go
func (c *City) GetSupplyCommodityFromCity(namecommodity, namecity string) int16 {
    // ... distance calculation logic
    return int16((float32(stock) * distanceInDays) / (Week / 30))
}
```

## Integration Points

1. **OpenAPI Validation**: API requests are validated against `docs/openapi.yaml` using middleware
   ```go
   router.Use(middleware.OapiRequestValidator(swagger))
   ```

2. **CORS**: Default CORS configuration is enabled
   ```go
   router.Use(cors.Default())
   ```

3. **PostgreSQL**: The `adnvilla/db-go` package handles database connectivity

## Common Tasks

- Adding a new endpoint: Define it in `docs/openapi.yaml`, create a handler in `interfaces/handlers/`, and implement the use case in `application/use_cases/`
- Adding a new domain entity: Create the entity in `domain/` and implement repositories in `infrastructure/data/postgresql/`

## Example Patterns

### Creating a new Use Case:
```go
type MyNewUseCase struct {}

func NewMyNewUseCase() use_case.UseCase[MyInput, MyOutput] {
    return new(MyNewUseCase)
}

func (u *MyNewUseCase) Handle(ctx context.Context, in MyInput) (MyOutput, error) {
    // implementation
}
```