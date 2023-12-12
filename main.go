package main

import (
	"MerchantBank/app"
	"MerchantBank/router"
)

func init() {
	app.LoadEnvVariable()
	app.ConnnectToDB()
	app.SycDatabase()

}

func main() {
	router.RouterApi()
}
