package slogopts_test

import (
	"log/slog"
	"os"

	"github.com/akm/slogopts"
)

func ExampleLevel() {
	logger := slogopts.New(os.Stdout, slogopts.Level(slog.LevelInfo))
	logger.Info("hello")
}

func ExampleAddSource() {
	logger := slogopts.New(os.Stdout, slogopts.AddSource(true))
	logger.Info("hello")
}

func ExampleReplaceAttr() {
	logger := slogopts.New(os.Stdout, slogopts.ReplaceAttr(
		func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{Key: "t", Value: a.Value}
			}
			return a
		},
	))
	logger.Info("hello")
}
