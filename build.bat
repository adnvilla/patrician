@echo off

if "%1"=="" goto default
if "%1"=="build" goto build
if "%1"=="install" goto install
if "%1"=="run" goto run
if "%1"=="test" goto test
if "%1"=="test-verbose" goto test-verbose
if "%1"=="test-coverage" goto test-coverage
if "%1"=="test-cover" goto test-cover
if "%1"=="test-integration" goto test-integration
if "%1"=="test-domain" goto test-domain
if "%1"=="test-use-cases" goto test-use-cases
if "%1"=="test-handlers" goto test-handlers
if "%1"=="test-repositories" goto test-repositories
goto end

:default
:build
call :commit_hash
call :install
goto end

:install
set GOBIN=%CD%\bin
go install -v
goto end

:commit_hash
ruby generate_commit_hash_file.rb
goto end

:run
call :build
.\bin\patrician.exe
goto end

:test
go test -v ./...
goto end

:test-verbose
go test -v ./...
goto end

:test-coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
goto end

:test-cover
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
goto end

:test-integration
go test -v ./integration/...
goto end

:test-domain
go test -v ./src/domain/...
goto end

:test-use-cases
go test -v ./src/application/use_cases/...
goto end

:test-handlers
go test -v ./src/interfaces/handlers/...
goto end

:test-repositories
go test -v ./src/infrastructure/data/postgresql/...
goto end

:end
