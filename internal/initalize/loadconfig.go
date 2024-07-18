package initalize

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql   MysqlConfig `mapstructure:"mysql"`
	MongoDB MongoConfig `mapstructure:"mongodb"`
}

type MysqlConfig struct {
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type MongoConfig struct {
	URI string `mapstructure:"uri"`
}

var config Config

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("internal/config/") // path to config file
	v.SetConfigName("local")            // name of file
	v.SetConfigType("yaml")             // type of the config file

	// read config
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	// read server configuration
	fmt.Println("Server port:: ", v.GetInt("server.port"))

	// configure structure
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %v", err))
	}
}