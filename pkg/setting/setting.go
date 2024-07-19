package setting

type Config struct {
	Mysql    MysqlConfig     `mapstructure:"mysql"`
	MongoDB  MongoConfig     `mapstructure:"mongodb"`
	Security SercurityConfig `mapstructure:"security"`
	Server   SercurityConfig `mapstructure:"server"`
}

type MysqlConfig struct {
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type SercurityConfig struct {
	AccessKey  string `mapstructure:"accesskey"`
	RefreshKey string `mapstructure:"refreshkey"`
	Port       int    `mapstructure:"port"`
}

type MongoConfig struct {
	URI string `mapstructure:"uri"`
}