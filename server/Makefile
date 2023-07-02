server:
	go run cmd/main.go

build: 
	go build -o bin/server cmd/main.go

swagger:
	go get -u github.com/swaggo/swag/cmd/swag
	${GOPATH}/bin/swag init -g cmd/main.go