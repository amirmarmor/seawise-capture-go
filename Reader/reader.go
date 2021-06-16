package Reader

import (
	"fmt"
	"seawise.com/capture/channel"
	"seawise.com/capture/log"
)

type Reader struct {
	Counter  int
	channels  []*channel.Channel
}

func Create(channels int, offset int) (*Reader, error) {
	reader := &Reader{
		0,
		make([]*channel.Channel, 0, channels),
	}

	for ch := offset; ch <= channels*2; ch += 2 {
		camera, err := channel.Create(ch)
		if err != nil {
			return nil, fmt.Errorf("failed to create reader: %v", err)
		}

		reader.channels = append(reader.channels, camera)
	}

	return reader, nil
}

func (r *Reader) ReadChannels() error {
	for _, ch := range r.channels {
		err := ch.Read(r.Counter)
		if err != nil {
			return fmt.Errorf("failed to read channels: %v", err)
		}
		log.V5(fmt.Sprintf("taking picture %v", r.Counter))
		r.Counter++
	}
	return nil
}

