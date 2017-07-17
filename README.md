# Install

```bash
go get -d github.com/jedvardsson/go-ws-demo
cd $GOPATH/src/github.com/jedvardsson/go-ws-demo
dep ensure
```

# Run Tests
```
go test
```

# Build and Run
```bash
go build
./go-ws-demo
```

# API
```
# List products
curl -s http://localhost:8080/products | json_pp

# Add product
curl -s -X POST http://localhost:8080/product -T - <<<'{"name":"prod-x","price":11.22}' | json_pp

# Get product
curl -s http://localhost:8080/product/2 | json_pp

# Update product
curl -s -X PUT http://localhost:8080/product/3 -T - <<<'{"name":"prod-x","price":9.99}'|json_pp

# Delete product
curl -s -X DELETE http://localhost:8080/product/2 | json_pp
```
