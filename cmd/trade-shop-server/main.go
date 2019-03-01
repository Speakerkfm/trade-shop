package main

import (
	"github.com/go-openapi/swag"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/boj/redistore.v1"
	"log"
	"os"
	"strconv"
	"time"
	pkg_flags "trade-shop/pkg/flags"
	"trade-shop/pkg/restapi"
	"trade-shop/pkg/restapi/operations"
	"trade-shop/pkg/service"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTradeShopAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	//config
	conf := &pkg_flags.Config{}
	configureFlags(api, server, swaggerSpec, conf)

	//mysql
	mysqlConf := mysql.NewConfig()
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = conf.DatabaseHost + ":" + conf.DatabasePort
	mysqlConf.User = conf.DatabaseUser
	mysqlConf.Passwd = conf.DatabasePassword
	mysqlConf.DBName = conf.DatabaseName
	mysqlConf.MultiStatements = true
	mysqlConf.ParseTime = true
	mysqlConf.Loc = time.Local

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	defer db.Close()
	db.SingularTable(true)

	//rabit
	amqpClient, err := service.NewQueue(conf)
	defer amqpClient.Connection.Close()

	//redistore
	rstoreSize, _ := strconv.Atoi(conf.RedisStoreSize)
	rstore, err := redistore.NewRediStore(
		rstoreSize,
		"tcp",
		conf.RedisStoreHost+":"+conf.RedisStorePort,
		conf.RedisStorePassword,
		[]byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer rstore.Close()

	//set handlers
	handler := configureAPI(api, db, rstore, amqpClient, conf)
	server.SetHandler(handler)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func configureFlags(api *operations.TradeShopAPI, server *restapi.Server, swaggerSpec *loads.Document, conf *pkg_flags.Config) {
	params := swag.CommandLineOptionsGroup{
		LongDescription: "Additional configuration parameters",
		Options:         conf,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{params}

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Trade Shop"
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			panic(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
}
