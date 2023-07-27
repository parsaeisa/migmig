package config

var C *Config

type Config struct {
	OriginDatabase      DatabaseConnection `yaml:"origin_database"`
	DestinationDatabase DatabaseConnection `yaml:"destination_database"`
	Bulk                bool               `yaml:"bulk"`
}

type DatabaseConnection struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func Init(filename string) *Config {
	return nil
}
