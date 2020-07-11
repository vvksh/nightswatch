package nightswatch

import "time"

type Watcher interface {
	Check() []string
	Interval() time.Duration
	SlackChannel() string
}
