package main

import "MerchantBank/app"

func init() {
	app.LoadEnvVariable()
	app.ConnnectToDB()
	app.SycDatabase()

}

func main() {

}
