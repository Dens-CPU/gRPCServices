package orderconfig

type Config struct {
	Server   Server   `mapstructure:"server"`
	Postgres Postgres `mapstructure:"postgres"`
}
type Server struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	Network string `mapstructure:"network"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Sslmode  string `mapstructure:"sslmode"`
}
