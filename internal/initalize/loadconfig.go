package initalize

import (
	"fmt"
	"my_shop/global"
	"my_shop/pkg/setting"

	"github.com/spf13/viper"
)

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
  var config setting.Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %v", err))
	}

	// Assign the loaded configuration to the global.Config variable
  global.Config = config
}
