package main

import (
	"net/http"
	"trade-shop/pkg/handlers/login"
	"trade-shop/pkg/handlers/sales"
	"trade-shop/pkg/handlers/user"
	"trade-shop/pkg/restapi/operations"
	loginApi "trade-shop/pkg/restapi/operations/login"
	salesApi "trade-shop/pkg/restapi/operations/sales"
	userApi "trade-shop/pkg/restapi/operations/user"
	"trade-shop/pkg/service"
	"trade-shop/pkg/store"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	redistore "gopkg.in/boj/redistore.v1"
)

func configureAPI(api *operations.TradeShopAPI, db *gorm.DB, redisClient *redis.Client, rst *redistore.RediStore, amqpClient *service.Queue) http.Handler {
	st := store.NewStore(db, redisClient)
	mailer := service.NewMailer(amqpClient)

	saleService := service.NewSale(st, mailer)
	authService := service.NewAuthService(rst)
	invService := service.NewInventory(st)

	salesContext := sales.NewContext(st, saleService, authService)
	api.SalesSaleHandler = salesApi.SaleHandlerFunc(salesContext.SaleItems)
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
