package env

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	JWT_ALG string `env:"JWT_ALG"`
	JWT_KEY string `env:"JWT_KEY"`
	JWT_ISS string `env:"JWT_ISS"`
	JWT_AUD string `env:"JWT_AUD"`
}

var CONFIG *config

func init() {
	CONFIG = &config{}
	if err := env.Parse(CONFIG); err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}
}
