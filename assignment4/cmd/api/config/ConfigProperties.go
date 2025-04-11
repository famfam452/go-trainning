package config
type Properties struct {
	Postgresql Postgresql `yaml:"postgresql"`
	Redis      Redis      `yaml:"redis"`
}

type Postgresql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Sslmode string `yaml:"sslmode"`
}
type Redis struct {
	Address string `yaml:"address"`
	Password string `yaml:"password"`
	Database int `yaml:"database"`
}