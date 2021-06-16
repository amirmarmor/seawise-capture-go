package entrypoint

import (
	"seawise.com/capture/Reader"
	"seawise.com/capture/log"
	"seawise.com/capture/scheduler"
)

func Execute(){
	rate := 10000 //milliseconds
	maxPics := 6000
	reader, err := Reader.Create(3, 2)
	if err != nil {
		log.Fatal("Failed to create reader: ", err)
	}

	interval := scheduler.SetInterval(reader.ReadChannels, rate, false)

	for {
		if reader.Counter >= maxPics {
			interval <- true
			return
		}
	}
}
