package mongoconnector

type Config struct {
	DB       string `yaml:"db"`
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
