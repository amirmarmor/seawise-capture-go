package scheduler

import (
	"time"
)

type Worker func() error

func SetInterval(work Worker, milliseconds int, async bool) chan bool{
	interval := time.Duration(milliseconds) * time.Millisecond
	ticker := time.NewTicker(interval)
	clear := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					go work()
				} else {
					work()
				}
			case <- clear:
				ticker.Stop()
				return
			}
		}
	}()

	return clear
}