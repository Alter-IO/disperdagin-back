package logger

import (
	"io"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Global logger instance
var (
	logger *slog.Logger
	once   sync.Once
)

// InitializeLogger sets up structured logging with slog
func InitializeLogger() *slog.Logger {
	once.Do(func() {
		// Set up Lumberjack for log rotation
		fileLumberjack := &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    100,
			MaxBackups: 10,
			Compress:   true,
		}

		// Use multi-handler (write to both console and file)
		logger = slog.New(slog.NewJSONHandler(io.MultiWriter(fileLumberjack, os.Stdout), &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		// Handle SIGHUP to rotate logs
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP)

		go func() {
			for range c {
				fileLumberjack.Rotate()
			}
		}()
	})

	return logger
}

// Get returns the global logger instance
func Get() *slog.Logger {
	return InitializeLogger()
}
