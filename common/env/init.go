package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port                 int    `env:"PORT,unset" envDefault:"5000"`
	WebURL               string `env:"WEB_URL,unset"`
	DatabaseURL          string `env:"DB_URL,unset"`
	AccessTokenKey       string `env:"ACCESS_TOKEN_KEY,unset"`
	GinMode              string `env:"GIN_MODE,unset" envDefault:"debug"`
	EmailUsername        string `env:"EMAIL_USERNAME,unset"`
	EmailPassword        string `env:"EMAIL_PASSWORD,unset"`
	EmailHost            string `env:"EMAIL_HOST,unset"`
	EmailPort            int    `env:"EMAIL_PORT,unset" envDefault:"587"`
	MidtransServerKey    string `env:"MIDTRANS_SERVER_KEY,unset"`
	MidtransEnvirontment string `env:"MIDTRANS_ENVIRONMENT,unset"`
	MidtransPaymentURL   string `env:"MIDTRANS_PAYMENT_URL,unset"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
