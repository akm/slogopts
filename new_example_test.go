package slogopts_test

import (
	"log/slog"
	"os"

	"github.com/akm/slogopts"
)

func ExampleNew_withoutOptions() {
	logger := slogopts.New(os.Stdout)
	logger.Info("hello")
}

func ExampleNew_withMultipleOptions() {
	logger := slogopts.New(os.Stdout,
		slogopts.JSON(),
		slogopts.Level(slog.LevelDebug),
		slogopts.AddSource(true),
	)
	logger.Info("hello")
}

func ExampleNewHandler_withoutOptions() {
	logger := slog.New(slogopts.NewHandler(os.Stdout))
	logger.Info("hello")
}

func ExampleNewHandler_withMultipleOptions() {
	logger := slog.New(slogopts.NewHandler(os.Stdout,
		slogopts.JSON(),
		slogopts.Level(slog.LevelDebug),
		slogopts.AddSource(true),
	))
	logger.Info("hello")
}
