package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

// NewZapLogger создает новый экземпляр логгера типа *zap.Logger.
// Принимает аргумент level - уровень логирования в виде строки.
// Возвращает указатель на *zap.Logger.
//
// Аргументы:
// - level: строка, уровень логирования
//
// Пример использования:
// logger := NewZapLogger("info")
func NewZapLogger(level string, toFile bool) *zap.Logger {

	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeDuration = zapcore.MillisDurationEncoder
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	out := os.Stdout

	if toFile {
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}

		out = file
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(out),
		lvl,
	)

	return zap.New(core).WithOptions(zap.AddCaller())
}
