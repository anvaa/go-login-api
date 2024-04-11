package main

import (
	
	"filefunc"
	"embedfiles"
	"initializers"
	"log"
	"os"
	"routers"
)

var WD string

func init() {
	WD = getWD()
	dbpath := WD + "/data/users.db"

	// Create the http.FileSystem using the embedded files
	embedfiles.GetWebFS()
	embedfiles.EmbedFilesToDisk()
	
	err := embedfiles.EmbedFiles()
	if err != nil {
		log.Fatal(err)
	}
	
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

	// shareFolder := WD + "/embededfiles/share"
	// if !filefunc.IsExists(shareFolder) {
	// 	log.Println("No share folder found. Creating one...")
	// 	filefunc.CreateFolder(shareFolder)
	// }

	
	
	initializers.LoadEnv(WD)
	initializers.ConnectToDB(dbpath)
	initializers.SyncDB()
}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Root: " + wd)
	return wd
}

func main() {
	// generate new 64-bit secret key for JWT on startup
	err := os.Setenv("JWT_SECRET", initializers.GetSecret())
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(os.Environ())

	r := routers.SetupRouter()

	r.Run()
}

