package server

type Config struct {
	SessionSecret string `env:"SESSION_SECRET,required"`
}
