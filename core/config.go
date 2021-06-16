package core

import (
	"encoding/json"
	"github.com/namsral/flag"
	"seawise.com/capture/log"
)

type Configuration = struct {
	Channels int
	Offset   int
	Rate     int
	MaxPics  int
}

var Config Configuration

func InitFlags() {
	flag.IntVar(&Config.Channels, "channels", 3, "The number of channels to use")
	flag.IntVar(&Config.Offset, "offset", 0, "The number of the first channel - if no webcam = 0")
	flag.IntVar(&Config.Rate, "rate", 10000, "The rate of pictures in ms")
	flag.IntVar(&Config.MaxPics, "maxpics", 6000, "maximal number of pics to take")

	log.AddNotify(postParse)
}

func postParse() {
	marshal, err := json.Marshal(Config)
	if err != nil {
		log.Fatal("marshal config failed: %v", err)
	}

	log.V5("configuration loaded: %v", string(marshal))
}
