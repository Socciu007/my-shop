package initalize

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql   MysqlConfig `mapstructure:"mysql"`
	MongoDB MongoConfig `mapstructure:"mongodb"`
	Security SercurityConfig `mapstructure:"security"`
	Server SercurityConfig `mapstructure:"server"`
}

type MysqlConfig struct {
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type SercurityConfig struct {
	AccessKey string `mapstructure:"accesskey"`
	RefreshKey string `mapstructure:"refreshkey"`
	Port int `mapstructure:"port"`
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

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return &config
}