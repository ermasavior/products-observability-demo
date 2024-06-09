package configs

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppPort    string `env:"APP_PORT"`
	AppEnv     string `env:"APP_ENV"`
	DbName     string `env:"DB_NAME"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbHost     string `env:"DB_HOST"`
	DbPort     int    `env:"DB_PORT"`
}

type configLoader struct {
}

type ConfigLoader interface {
	LoadConfig() *Config
}
