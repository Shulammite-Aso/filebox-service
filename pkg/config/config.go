package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	DBUrl          string `mapstructure:"DB_URL"`
	EmailSvcUrl    string `mapstructure:"EMAIL_SVC_URL"`
	CloudName      string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudApiKey    string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudApiSecret string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudFolder    string `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
