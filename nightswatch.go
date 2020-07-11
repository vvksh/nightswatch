package nightswatch

import (
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/vvksh/amigo"
)

var watchers []Watcher

// Register a watcher
func Register(watcher Watcher) {
	watchers = append(watchers, watcher)
}

// Start watching
func Start() {
	for _, watcher := range watchers {
		go watch(watcher)
	}
	select {}
}

func watch(watcher Watcher) {
	log.Printf("watching now %s", reflect.TypeOf(watcher).String())
	for {
		updates := watcher.Check()
		handleUpdates(updates, watcher.SlackChannel(), watcher)
		r := rand.Intn(100)
		time.Sleep(watcher.Interval() + time.Duration(r)*time.Millisecond)
	}
}

func handleUpdates(updates []string, channel string, watcher Watcher) {
	if len(updates) == 0 {
		log.Printf("\nNo new updates from %s \n", reflect.TypeOf(watcher).String())
		return
	}
	for _, update := range updates {
		err := amigo.SendSlackNotification(update, channel)
		if err != nil {
			log.Printf(err.Error())
		}
	}

}
