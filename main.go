package main

import "MerchantBank/repository"

func init() {
	repository.LoadEnvVariable()
	repository.ConnnectToDB()
	repository.SycDatabase()

}

func main() {

}
