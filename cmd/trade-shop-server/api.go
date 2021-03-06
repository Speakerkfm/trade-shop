package main

import (
	"net/http"
	"trade-shop/pkg/handlers/exchange"
	"trade-shop/pkg/handlers/login"
	"trade-shop/pkg/handlers/register"
	"trade-shop/pkg/handlers/sales"
	"trade-shop/pkg/handlers/user"
	"trade-shop/pkg/restapi/operations"
	exchangeApi "trade-shop/pkg/restapi/operations/exchange"
	loginApi "trade-shop/pkg/restapi/operations/login"
	registerApi "trade-shop/pkg/restapi/operations/register"
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
	userService := service.NewUserService(st)
	exchangeService := service.NewExchangeService()

	salesContext := sales.NewContext(st, saleService, authService)
	api.SalesSaleHandler = salesApi.SaleHandlerFunc(salesContext.SaleItems)
	api.SalesSalesListHandler = salesApi.SalesListHandlerFunc(salesContext.GetSalesList)
	api.SalesBuyHandler = salesApi.BuyHandlerFunc(salesContext.BuyLot)

	loginContext := login.NewContext(st, rst, authService)
	api.LoginLoginHandler = loginApi.LoginHandlerFunc(loginContext.AuthByEmailAndPassword)

	userContext := user.NewContext(st, rst, authService, saleService, invService)
	api.UserInventoryHandler = userApi.InventoryHandlerFunc(userContext.GetInventoryList)
	api.UserLogoutHandler = userApi.LogoutHandlerFunc(userContext.LogoutUser)
	api.UserUserSalesListHandler = userApi.UserSalesListHandlerFunc(userContext.GetUserSalesList)
	api.UserCancelHandler = userApi.CancelHandlerFunc(userContext.SaleCancel)

	registerContext := register.NewContext(userService)
	api.RegisterRegisterHandler = registerApi.RegisterHandlerFunc(registerContext.Register)

	exchangeContext := exchange.NewContext(exchangeService, authService)
	api.ExchangeExchangeRatesHandler = exchangeApi.ExchangeRatesHandlerFunc(exchangeContext.GetExchangeRates)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddleware))
}

func setupMiddleware(handler http.Handler) http.Handler {
	return handler
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
