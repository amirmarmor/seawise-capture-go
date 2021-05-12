package videofile

import (
	"fmt"
	"gocv.io/x/gocv"
	"os"
	"seawise.com/capture/log"
	"strconv"
	"strings"
	"time"
)

type Reader struct {
	Counter  int
	file     string
	stream   *gocv.VideoCapture
	img      gocv.Mat
	window   *gocv.Window
	savePath string
}

func Create(file string) *Reader {
	path := fmt.Sprintf("videos/%s", file)
	img := gocv.NewMat()
	window := gocv.NewWindow("Hello")

	savePath, err := createSavePath(file)
	if err != nil {
		log.Error("failed to create path", err)
	}

	stream, err := gocv.VideoCaptureFile(path)
	if err != nil {
		log.Error("Failed to open file", err)
	}

	return &Reader{
		0,
		file,
		stream,
		img,
		window,
		savePath,
	}
}

func createSavePath(file string) (string, error) {
	_, err := os.Stat("images")

	if os.IsNotExist(err) {
		log.V5("images directory doesnt exist. creating it now!")
		err := os.Mkdir("images", 0777)
		if err != nil {
			log.Error("couldnt create images directory", err)
			return "", err
		}
	}

	vidFileNameParts := strings.Split(file, ".")
	source := vidFileNameParts[0]
	path := fmt.Sprintf("images/%s", source)
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

func (r *Reader) Test() error {
	r.Counter++
	fmt.Printf("%s, %s, %s", time.Now(), "running", r.Counter)
	return nil
}

func (r *Reader) ReadVideoFile() error {
	saveFile := fmt.Sprintf("%s/%s.jpg", r.savePath, strconv.Itoa(r.Counter))

	if ok := r.stream.Read(&r.img); !ok {
		err := fmt.Errorf("Device closed %s\n", r.file)
		return err
	}
	if r.img.Empty() {
		return nil
	}
	gocv.IMWrite(saveFile, r.img)
	r.Counter++
	return nil
}
