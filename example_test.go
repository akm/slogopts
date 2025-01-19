package slogopts_test

import (
	"log/slog"
	"os"

	"github.com/akm/slogopts"
)

func ExampleNew() {
	logger := slogopts.New(os.Stdout)
	logger.Info("hello")
}

func ExampleNew_withLevel() {
	logger := slogopts.New(os.Stdout, slogopts.Level(slog.LevelInfo))
	logger.Info("hello")
}

func ExampleNew_withJSON() {
	logger := slogopts.New(os.Stdout, slogopts.JSON())
	logger.Info("hello")
}

func ExampleNew_withReplace() {
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
