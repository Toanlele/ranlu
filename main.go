package main

import (
	"bytes"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func main() {
	// Read the mp3 file into memory
	fileBytes, err := os.ReadFile("src/1234.mp3")
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	samplingRate := 54400

	numOfChannels := 12 //播放速度

	audioBitDepth := 2
	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
