// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"trade-shop/pkg/restapi/operations"
	"trade-shop/pkg/restapi/operations/login"
	"trade-shop/pkg/restapi/operations/register"
	"trade-shop/pkg/restapi/operations/sales"
	"trade-shop/pkg/restapi/operations/user"
)

//go:generate swagger generate server --target ../../pkg --name TradeShop --spec ../../tmp/swagger.yaml --exclude-main

func configureFlags(api *operations.TradeShopAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TradeShopAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.SalesBuyHandler = sales.BuyHandlerFunc(func(params sales.BuyParams) middleware.Responder {
		return middleware.NotImplemented("operation sales.Buy has not yet been implemented")
	})
	api.UserCancelHandler = user.CancelHandlerFunc(func(params user.CancelParams) middleware.Responder {
		return middleware.NotImplemented("operation user.Cancel has not yet been implemented")
	})
	api.UserInventoryHandler = user.InventoryHandlerFunc(func(params user.InventoryParams) middleware.Responder {
		return middleware.NotImplemented("operation user.Inventory has not yet been implemented")
	})
	api.LoginLoginHandler = login.LoginHandlerFunc(func(params login.LoginParams) middleware.Responder {
		return middleware.NotImplemented("operation login.Login has not yet been implemented")
	})
	api.UserLogoutHandler = user.LogoutHandlerFunc(func(params user.LogoutParams) middleware.Responder {
		return middleware.NotImplemented("operation user.Logout has not yet been implemented")
	})
	api.RegisterRegisterHandler = register.RegisterHandlerFunc(func(params register.RegisterParams) middleware.Responder {
		return middleware.NotImplemented("operation register.Register has not yet been implemented")
	})
	api.SalesSaleHandler = sales.SaleHandlerFunc(func(params sales.SaleParams) middleware.Responder {
		return middleware.NotImplemented("operation sales.Sale has not yet been implemented")
	})
	api.SalesSalesListHandler = sales.SalesListHandlerFunc(func(params sales.SalesListParams) middleware.Responder {
		return middleware.NotImplemented("operation sales.SalesList has not yet been implemented")
	})
	api.UserUserSalesListHandler = user.UserSalesListHandlerFunc(func(params user.UserSalesListParams) middleware.Responder {
		return middleware.NotImplemented("operation user.UserSalesList has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
