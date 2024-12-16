package bliss

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// Bliss ...
type Bliss struct {
	log  *slog.Logger
	conf Conf

	c chan struct{}
}

// New ...
func New(conf Conf) *Bliss {
	return &Bliss{
		conf: conf,
		c:    make(chan struct{}),
	}
}

// Start ...
func (b *Bliss) Start() error {
	return b.CloseOnProgramEnd()
}

// CloseOnProgramEnd ...
func (b *Bliss) CloseOnProgramEnd() error {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	b.log.Info("received shutdown signal")
	close(b.c)
	return nil
}
