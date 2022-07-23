[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=8024865&assignment_repo_type=AssignmentRepo)
# empty
Empty repo for github classroom starter

## Creating a new project

 - Rename `cmd/golang-starter` folder
 - Rename `APP_NAME` in `Makefile`
 - Rename `go.mod` with proper username and repository names
 - Update imports accordingly

## Usage

- `make build`
- `make tests`
- `make run`
- `make clean`

## Endpoint
- GET /product
- GET /product/:id | id = int
- POST /product | body {"name" : "es teh", "price" : 500}
- PUT /product/:id | body {"name" : "es teh segar", "price" : 700}
- DELETE /product/:id | id = int

## RUN with echo fw
go run ./cmd/echo/main.go 

## RUN with echo fw and GORM
go run ./cmd/gorm/main.go 