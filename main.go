package main

import (
	"github.com/didntpot/bliss/bliss"
	"log/slog"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelInfo)
}

func main() {
	conf := bliss.Conf{}
	b := bliss.New(conf)
	if err := b.Start(); err != nil {
		panic(err)
	}
}
