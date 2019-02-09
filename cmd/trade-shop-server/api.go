package main

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"trade-shop/pkg/flags"
	"trade-shop/pkg/handlers/sales"
	"trade-shop/pkg/restapi/operations"
	salesApi "trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/service"
	"trade-shop/pkg/store"
)

func configureAPI(api *operations.TradeShopAPI, db *gorm.DB, conf *flags.Config) http.Handler {
	st := store.NewStore(db)
	salesService := service.NewSalesList(st)

	salesContext := sales.NewContext(st, salesService)

	api.SalesSalesListHandler = salesApi.SalesListHandlerFunc(salesContext.GetSalesList)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return corsListener(handler)
}

func corsListener(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("123", "123")

		h.ServeHTTP(w, r)
	})
}
