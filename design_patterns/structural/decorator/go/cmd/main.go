package main

import (
	logger "decorator.example/internal/logger"
)

func main() {
	basic := &logger.ConsoleLogger{}
	basic.Log("Basic log")

	decorated := &logger.LevelDecorator{
		LoggerDecorator: logger.LoggerDecorator{Logger: &logger.TimestampDecorator{
			LoggerDecorator: logger.LoggerDecorator{Logger: basic},
		}},
		Level: "INFO",
	}

	decorated.Log("Decorated log")
}
