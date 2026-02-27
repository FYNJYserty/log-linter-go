package testdata

import (
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func test() {
	log.Println("This is a log")
	log.Println("Bad message!")

	slog.Info("This is an error message slog")

	slog.Debug("api_key=sk_live_abcd1234")

	zap.L().Info("This is standard library logging")
}
