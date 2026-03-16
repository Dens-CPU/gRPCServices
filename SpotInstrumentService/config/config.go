package spotconfig

type Config struct {
	Redis  Redis  `mapstructure:"redis"`
	Server Server `mapstructure:"server"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Server struct {
	Port    string `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	Network string `mapstructure:"network"`
}
