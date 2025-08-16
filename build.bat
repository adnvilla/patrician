@echo off

if "%1"=="" goto default
if "%1"=="build" goto build
if "%1"=="install" goto install
if "%1"=="run" goto run
if "%1"=="test" goto test
if "%1"=="test-verbose" goto test-verbose
if "%1"=="test-coverage" goto test-coverage
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
go test ./...
goto end

:test-verbose
go test -v ./...
goto end

:test-coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
goto end

:end
