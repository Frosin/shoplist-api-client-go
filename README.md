Generate api scheme for server:
`oapi-codegen --package api --generate types,server,spec ./docs/swagger.yaml > ./api/swagger.gen.go`

Generate api scheme for client:
`oapi-codegen --package client --generate types,client ./docs/swagger.yaml > ./client/swagger.gen.go` 

Generate sqlc go files: 
`sqlc generate`

Compile for ARM v7:
Install gcc arm compiler:
`sudo apt update`
`sudo apt install gcc-arm-linux-gnueabihf`
`export CC=arm-linux-gnueabihf-gcc`
`GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build`

To use: 
`$ ./shoplist serve -p 8080`