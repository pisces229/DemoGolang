package singleton

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppConfiguration = initConfiguration()

type Configuration struct {
	Mode     string `yaml:"mode"`
	LogPath  string `yaml:"logPath"`
	TempPath string `yaml:"tempPath"`
	Mkcert   struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"mkcert"`
	Connection struct {
		Pool struct {
			MaxLifetime  int `yaml:"maxLifetime"`
			MaxOpenConns int `yaml:"maxOpenConns"`
			MaxIdleConns int `yaml:"maxIdleConns"`
		} `yaml:"pool"`
		Sqlserver struct {
			Server   string `yaml:"server"`
			Port     int    `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"sqlserver"`
		Postgres struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"postgres"`
	} `yaml:"connection"`
}

func initConfiguration() (result *Configuration) {
	fmt.Println("initConfiguration")
	// viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetDefault("key", "value")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper.ReadInConfig Fail:" + err.Error())
	} else {
		// fmt.Printf("%+v\n", viper.AllSettings())
		err = viper.Unmarshal(&result)
		if err != nil {
			panic("viper.Unmarshal Fail:" + err.Error())
		}
		// fmt.Printf("%+v\n", result)
	}
	return
}
