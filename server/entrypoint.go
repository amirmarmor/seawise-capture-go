package server

import (
	"seawise.com/capture/Reader"
	"seawise.com/capture/core"
	"seawise.com/capture/log"
	"seawise.com/capture/scheduler"
)

type Entrypoint struct {}

func (p *Entrypoint) Execute() {
	log.InitFlags()
	core.InitFlags()
	log.Info("starting")

	reader, err := Reader.Create(core.Config.Channels, core.Config.Offset)
	if err != nil {
		log.Fatal("Failed to create reader: ", err)
	}

	interval := scheduler.SetInterval(reader.ReadChannels, core.Config.Rate, false)

	for {
		if reader.Counter >= core.Config.MaxPics {
			interval <- true
			return
		}
	}
}

