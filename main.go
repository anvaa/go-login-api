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
	
	envFile := WD + "/.env"
	if !filefunc.IsExists(envFile) {
		log.Println("No .env file found. Creating one...")
		initializers.WriteEnv(WD)
	}

	dataFolder := WD + "/data"
	if !filefunc.IsExists(dataFolder) {
		log.Println("No data folder found. Creating one...")
		filefunc.CreateFolder(dataFolder)
	}

	shareFolder := WD + "/share"
	if !filefunc.IsExists(shareFolder) {
		log.Println("No data folder found. Creating one...")
		filefunc.CreateFolder(shareFolder)
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