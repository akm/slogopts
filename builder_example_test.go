package slogopts_test

import (
	"os"

	"github.com/akm/slogopts"
)

func ExampleJSON() {
	logger := slogopts.New(os.Stdout, slogopts.JSON())
	logger.Info("hello")
}

func ExampleText() {
	logger := slogopts.New(os.Stdout, slogopts.Text())
	logger.Info("hello")
}
