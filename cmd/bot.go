package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/PalindromeCorp/TODOBOT/src/server"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {

	var logger *log.Logger
	logger = log.New()
	logger.Formatter = new(log.TextFormatter)
	logger.Formatter.(*log.TextFormatter).DisableColors = true
	logger.Formatter.(*log.TextFormatter).FullTimestamp = true
	logger.Level = log.TraceLevel
	logger.WithField("ts", logger.WithTime(time.Now()))

	srv := server.New(logger.WithField("component", "bot"))
	serverBot := srv.Serve()
	// Wait here until CTRL-C or other term signal is received.
	logger.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	errs := make(chan error, 2)

	// Cleanly close down the Discord session.
	// Wait here until CTRL-C or other term signal is received.
	go func() {
		errs <- serverBot.Close()
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	logger.Log(log.InfoLevel, "terminated ", <-errs)

}
