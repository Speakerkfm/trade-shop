package main

import (
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"gopkg.in/boj/redistore.v1"
	"net/http"
	"trade-shop/pkg/flags"
	"trade-shop/pkg/handlers/inventory"
	"trade-shop/pkg/handlers/login"
	"trade-shop/pkg/handlers/sales"
	"trade-shop/pkg/restapi/operations"
	inventoryApi "trade-shop/pkg/restapi/operations/inventory"
	loginApi "trade-shop/pkg/restapi/operations/login"
	salesApi "trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/service"
	"trade-shop/pkg/store"
)

func configureAPI(api *operations.TradeShopAPI, db *gorm.DB, rst *redistore.RediStore, conf *flags.Config) http.Handler {
	st := store.NewStore(db)
	salesService := service.NewSalesListService(st)
	authService := service.NewAuthService(rst)

	salesContext := sales.NewContext(st, salesService, authService)
	loginContext := login.NewContext(st, rst, authService)
	inventoryContext := inventory.NewContext(st, authService)

	api.SalesSalesListHandler = salesApi.SalesListHandlerFunc(salesContext.GetSalesList)
	api.LoginLoginHandler = loginApi.LoginHandlerFunc(loginContext.AuthByEmailAndPassword)
	api.InventoryInventoryHandler = inventoryApi.InventoryHandlerFunc(inventoryContext.GetInventoryList)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddleware))
}

func setupMiddleware(handler http.Handler) http.Handler {
	return handler
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func sessionListener(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		if err := sessions.Save(r, w); err != nil {
			panic(err)
		}
	})
}
