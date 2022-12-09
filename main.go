package main

import (
	"fmt"

	"github.com/GoldenLeeK/go-ffmpeg/driver"
)

func main() {
	config := &Config{}
	ffmpeg := config.NewDrvier(&driver.FFmpeg{})
	

}
