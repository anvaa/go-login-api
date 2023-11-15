package main

import (
	"log"
	"initializers"
	"routers"
	"filefunc"
	"os"
)

var WD string

func init() {
	WD = getWD()
	
	if !filefunc.IsExists(WD + "/.env") {
		log.Println("No .env file found. Creating one...")
		initializers.WriteEnv(WD)
	}
	
	initializers.LoadEnv(WD)
	initializers.ConnectToDB(os.Getenv("DB_PATH"))
	initializers.SyncDB()
}

func main() {

	log.Println("Port: " + os.Getenv("PORT"))

	r := routers.SetupRouter(WD)

	r.Run()
}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Root: " + wd)
	return wd
}