BUILDPATH=$(CURDIR)
API_NAME=etl-app

build: 
	@echo "Creando Binario ..."
	@go build  -o $(BUILDPATH)/build/${API_NAME} main.go
	@echo "Binario generado en build/${API_NAME}"

test: 
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

.PHONY: test build
