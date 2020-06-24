package grin

type GrinConfig struct {
	URL      string `envconfig:"GRIN_API_URL" required:"true"`
	Password string `envconfig:"GRIN_PASSWORD" required:"true"`
}
