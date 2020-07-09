package utilpkg

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Keep a Go process alive by listening to process interrupt calls.
func KeepAlive() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Log an error with given context and the error itself
func Report(context string, err error) {
	log.Printf("%s.\n%s", context, err.Error())
}

// RemoveFromSlice takes <removeString string> <list []string>
// RemoveFromSlice removes a specific string form a slice.
// Returns []string
func RemoveFromSlice(removeString string, list []string) []string {
	indexItem := 0

	for index, item := range list {
		if item == removeString {
			indexItem = index
		}
	}

	return append(list[:indexItem], list[indexItem+1:]...)
}
