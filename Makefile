.PHONY:*
run:
	export `cat .env` && https_proxy=http://localhost:8080 go run cmd/trade-shop-server/*.go  --scheme=http --port=8000

server: flatten
	rm -f pkg/restapi/configure_trade_shop.go
	swagger generate server -f tmp/swagger.yaml -t pkg --exclude-main

doc: flatten
	swagger serve -p 8095 -F swagger tmp/swagger.yaml

flatten:
	mkdir -p tmp
	swagger -o=tmp/flatten_log.txt flatten ./swagger/swagger.yaml > tmp/swagger.yaml && swagger validate tmp/swagger.yaml

iface:
	ifacemaker -f pkg/store -s Store -i StoreInterface -p store -o pkg/store/store_interface.go
	ifacemaker -f pkg/service/inventory.go -s Inventory -i Inventory -p serviceiface -o pkg/service/serviceiface/inventory.go
	ifacemaker -f pkg/service/mailer.go -s Mailer -i Mailer -p serviceiface -o pkg/service/serviceiface/mailer.go
	ifacemaker -f pkg/service/sale.go -s Sale -i Sale -p serviceiface -o pkg/service/serviceiface/sale.go
	ifacemaker -f pkg/service/auth.go -s AuthService -i AuthService -p serviceiface -o pkg/service/serviceiface/auth.go

mock: iface
	mockery -dir pkg/store --all -output pkg/mocks -case underscore
	mockery -dir pkg/service/serviceiface --all -output pkg/mocks -case underscore
