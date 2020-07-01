package utilpkg

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func KeepAlive() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func Report(context string, err error) {
	log.Printf("%s.\n%s", context, err.Error())
}
