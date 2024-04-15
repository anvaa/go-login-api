package main

import (
	"appconf"
	"filefunc"
	"embedfiles"
	"initializers"
	"appsec"
	
	"log"
	"os"
	"routers"
)

var WD string
var https bool

func init() {
	WD = getWD()

	configFile := WD + "/.app" // check for .app file
	if !filefunc.IsExists(configFile) {
		log.Println("No .app file found. Creating one...")
		appconf.WriteDefaultConfig(WD)
	}
	appconf.ReadConfig() // read the .app file

	// Create the http.FileSystem using the embedded files
	embedfiles.GetWebFS()
	embedfiles.EmbedFilesToDisk()
	
	err := embedfiles.EmbedFiles()
	if err != nil {
		log.Fatal(err)
	}

	dataFolder := appconf.GetVal("root_folder") + "/data"
	if !filefunc.IsExists(dataFolder) {
		log.Println("No data folder found. Creating one...")
		filefunc.CreateFolder(dataFolder)
		filefunc.CreateFile(appconf.GetVal("db_path"))
	}

	dbpath := appconf.GetVal("db_path")
	initializers.ConnectToDB(dbpath)
	initializers.SyncDB()

	certFile := appconf.GetVal("root_folder") + "/app.crt"
	keyFile := appconf.GetVal("root_folder") + "/app.key"
	if !filefunc.IsExists(certFile) || !filefunc.IsExists(keyFile) {
		log.Println("No RSA files found. Creating key pair ...")

		err := appsec.GenerateTLS(keyFile, certFile, "2048")
		if err != nil {
			log.Fatal(err)
		}
	}

	https = appconf.GetVal("https") == "true"
}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("WorkDir: %s", wd)
	return wd
}

func main() {
	
	// set new JWT_SECRET environment variable if in release mode
	err := os.Setenv("JWT_SECRET", appsec.GetSecret())
	if err != nil {
		log.Fatal(err)
	}

	r := routers.SetupRouter()

	if https {
		certFile := appconf.GetVal("root_folder") + "/app.crt"
		keyFile := appconf.GetVal("root_folder") + "/app.key"

		log.Println("Starting HTTPS server on port " + appconf.GetVal("port"))
		r.RunTLS(":"+appconf.GetVal("port"), certFile, keyFile)
	} else {
		log.Println("Starting HTTP server on port " + appconf.GetVal("port"))
		r.Run(":" + appconf.GetVal("port"))
	}
}

