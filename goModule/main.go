package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	file, err := os.Open("example.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stream, format, err := mp3.Decode(file)

	if err != nil {
		fmt.Println("error in decoding file,  ", err)
		return
	}
	defer stream.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.PlayAndWait(stream)

}
