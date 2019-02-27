package main

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/boj/redistore.v1"
	"net/http"
	"trade-shop/pkg/flags"
	"trade-shop/pkg/handlers/login"
	"trade-shop/pkg/handlers/sales"
	"trade-shop/pkg/handlers/user"
	"trade-shop/pkg/restapi/operations"
	loginApi "trade-shop/pkg/restapi/operations/login"
	saleApi "trade-shop/pkg/restapi/operations/sale"
	salesApi "trade-shop/pkg/restapi/operations/sales"
	userApi "trade-shop/pkg/restapi/operations/user"
	"trade-shop/pkg/service"
	"trade-shop/pkg/store"
)

func configureAPI(api *operations.TradeShopAPI, db *gorm.DB, rst *redistore.RediStore, conf *flags.Config) http.Handler {
	st := store.NewStore(db)

	saleService := service.NewSale(st)
	authService := service.NewAuthService(rst)
	invService := service.NewInventory(st)

	salesContext := sales.NewContext(st, saleService, authService)
	api.SaleSaleHandler = saleApi.SaleHandlerFunc(salesContext.SaleItems)
	api.SalesSalesListHandler = salesApi.SalesListHandlerFunc(salesContext.GetSalesList)
	api.SalesBuyHandler = salesApi.BuyHandlerFunc(salesContext.BuyLot)

	loginContext := login.NewContext(st, rst, authService)
	api.LoginLoginHandler = loginApi.LoginHandlerFunc(loginContext.AuthByEmailAndPassword)

	userContext := user.NewContext(st, rst, authService, invService)
	api.UserInventoryHandler = userApi.InventoryHandlerFunc(userContext.GetInventoryList)
	api.UserLogoutHandler = userApi.LogoutHandlerFunc(userContext.LogoutUser)

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
	})
}
