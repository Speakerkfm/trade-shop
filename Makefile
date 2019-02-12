.PHONY: *
run:
	export `cat .env` && https_proxy=http://localhost:8080 go run cmd/trade-shop-server/*.go  --scheme=http --port=8000

swaggergen: flatten
	rm -f pkg/restapi/configure_trade_shop.go
	swagger generate server -f tmp/swagger.yaml -t pkg --exclude-main --skip-operations

flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten ./swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

swaggerdoc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml