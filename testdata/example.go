package testdata

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Log messages
	// Отрицательное тестирование
	log.Println("This is a log")
	log.Println("Bad message!")
	log.Fatalf("Русский лог с заглавной")
	log.Println("пример лога!")
	// Поизитвное тестирование
	log.Println("server started")
	log.Println("connection failed")
	log.Println("something went wrong")

	// Slog messages
	// Отрицательное тестирование
	slog.Debug("api_key=sk_live_abcd1234")
	slog.Info("user password: password123")
	slog.Debug("token: 1233qwee")
	// Поизитвное тестирование
	slog.Info("user authenticated successfully")
	slog.Error("api request completed")
	slog.Info("token validated")

	// Zap messages
	// Отрицательное тестирование
	logger.Debug("server started! 🚀")
	logger.Info("connection failed!!!")
	logger.Debug("warning: something went wrong...")
	// Поизитвное тестирование
	logger.Error("failed to connect to database")
	logger.Debug("failed to connect to database")
	logger.Info("starting server on port 8080")
}
