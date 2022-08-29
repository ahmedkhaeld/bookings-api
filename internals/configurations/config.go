package configurations

import (
	"github.com/spf13/viper"
	"log"
)

// Env stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Env struct {
	PgxDriver     string `mapstructure:"PGX_DRIVER"`
	PgxSource     string `mapstructure:"PGX_SOURCE"`
	POSTGRES      string `mapstructure:"POSTGRES"`
	MongoURI      string `mapstructure:"MONGO_URI"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Port          string `mapstructure:"PORT"`
	SMTPHost      string `mapstructure:"SMTP_HOST"`
	SMTPPort      string `mapstructure:"SMTP_PORT"`
	STMPUsername  string `mapstructure:"SMTP_USERNAME"`
	STMPPassword  string `mapstructure:"SMTP_PASSWORD"`
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	ENV           string `mapstructure:"ENV"`
}

// LoadConfig reads configuration from environment file or variables
func LoadConfig(path string) (config Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
