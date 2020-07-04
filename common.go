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
