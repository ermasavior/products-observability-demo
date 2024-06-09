package main

import (
	"context"
	orders "products-observability/internal/modules/orders/handler"
	"products-observability/pkg/configs"
	"products-observability/pkg/db"
	httpPkg "products-observability/pkg/http"
	"products-observability/pkg/logger"
)

func main() {
	cfg := configs.NewConfigLoader().Load()
	pgDb := db.NewPostgresDB(db.PostgresDsn{
		Host:     cfg.DbHost,
		User:     cfg.DbUsername,
		Password: cfg.DbPassword,
		Port:     cfg.DbPort,
		Db:       cfg.DbName,
	})
	logger.InitLogger(cfg.AppName, cfg.AppEnv)
	server := httpPkg.NewHTTPServer(cfg.AppName)

	orders.RegisterController(orders.Domain{
		Router:   server,
		NewRelic: nil,
		DB:       pgDb,
	})

	logger.Info(context.Background(), "Server running on Port "+cfg.AppPort)
	if err := server.Run(":" + cfg.AppPort); err != nil {
		logger.Fatal(context.Background(), err.Error())
	}
}
