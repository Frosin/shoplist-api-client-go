gen:
	oapi-codegen --package api --generate types,server,spec ./docs/swagger.yaml > ./api/swagger.gen.go
	oapi-codegen --package client --generate types,client ./docs/swagger.yaml > ./client/swagger.gen.go
	sqlc generate
build:
	export CC=
	CGO_ENABLED=1 GOOS=linux go build ./cmd/shoplist
barm:
	export CC=arm-linux-gnueabihf-gcc
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build -o shoplist_arm ./cmd/shoplist
run:
	./shoplist serve -p 8081
br:
	make build && make run
