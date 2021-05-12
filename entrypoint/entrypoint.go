package entrypoint

import (
	"seawise.com/capture/Reader"
	"seawise.com/capture/log"
	"seawise.com/capture/scheduler"
)

func Execute(){
	reader, err := Reader.Create(3, 2)
	if err != nil {
		log.Fatal("Failed to create reader: ", err)
	}

	interval := scheduler.SetInterval(reader.ReadChannels, 1000, false)

	for {
		if reader.Counter == 60 {
			interval <- true
			return
		}
	}
}
