package appconf

import (
	"log"
	"github.com/spf13/viper"
)

var AppConf = viper.New()

func init() {
	AppConf.SetConfigName(".app")
	AppConf.AddConfigPath(".")
	AppConf.SetConfigType("env")
}

func WriteDefaultConfig(appRoot string) {
	AppConf.SetDefault("https", "false")
	AppConf.SetDefault("port", "8090")
	AppConf.SetDefault("gin_mode", "debug")
	AppConf.SetDefault("root_folder", appRoot)
	AppConf.SetDefault("db_path", appRoot+"/data/app.db")
	AppConf.SetDefault("jwt_secret", "DEBUG_MODE_TEST_KEY_v41_J_ZQc022qQUGCzXq_Iu")
	AppConf.SetDefault("access_time", "3600*24*7")

	err := AppConf.WriteConfigAs(".app")
	if err != nil {
		log.Fatal("Error writing .app file")
	}
}

func ReadConfig() {
	err := AppConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file")
	}
}

func GetVal(key string) string {
	return AppConf.GetString(key)
}

func SetVal(key string, val string) {
	AppConf.Set(key, val)
	err := AppConf.WriteConfigAs(".app")
	if err != nil {
		log.Fatal("Error writing .app file")
	}
}