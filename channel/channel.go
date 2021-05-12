package channel

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"seawise.com/capture/log"
)

type Channel struct {
	name    int
	vc      *gocv.VideoCapture
	mat     gocv.Mat
	window  *gocv.Window
	path    string
}

func Create(channel int) (*Channel, error) {
	vc, err := gocv.OpenVideoCapture(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to open channel %v: %v", channel, err)
	}

	savePath, err := createSavePath(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to create path", err)
	}

	mat := gocv.NewMat()
	window := gocv.NewWindow(fmt.Sprintf("channel %v", channel))

	return &Channel{
		channel,
		vc,
		mat,
		window,
		savePath,
	}, nil
}

func (c *Channel) Read(fname int) error {
	saveFile := fmt.Sprintf("%s/%v.jpg", c.path, fname)

	if ok := c.vc.Read(&c.mat); !ok {
		err := fmt.Errorf("channel closed %v\n", c.name)
		return err
	}
	if c.mat.Empty() {
		return nil
	}
	gocv.IMWrite(saveFile, c.mat)
	return nil
}

func createSavePath(channel int) (string, error) {
	_, err := os.Stat("images")

	if os.IsNotExist(err) {
		log.V5("images directory doesnt exist. creating it now!")
		err := os.Mkdir("images", 0777)
		if err != nil {
			log.Error("couldnt create images directory", err)
			return "", err
		}
	}

	path := fmt.Sprintf("images/%v", channel)
	_, err = os.Stat(path)

	if !os.IsNotExist(err) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Error("couldnt remove folder", path)
		}
	}

	log.V5("creating file direcotry!")
	err = os.Mkdir(path, 0777)
	if err != nil {
		log.Error("couldnt create images directory", err)
		return "", err
	}
	return path, nil
}
