Generate api scheme:
`oapi-codegen --package api --generate types,server,spec ./docs/swagger.yaml > ./api/swagger.gen.go` 
Generate sqlc go files: 
`sqlc generate`
To use: 
`$ ./shoplist-api-client-go serve -p 8080`