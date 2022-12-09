package main

import (
	"fmt"

	"github.com/GoldenLeeK/go-ffmpeg/driver"
	"github.com/GoldenLeeK/go-ffmpeg/format"
	"github.com/GoldenLeeK/go-ffmpeg/media"
)

func main() {
	config := &driver.Config{}
	ffmpeg := config.NewDrvier(&driver.FFmpeg{})
	fmt.Println(ffmpeg)

	x264 := &format.X264{}
	x264.SetVideoCode("libx264")
	x264.SetInitParams([]string{" -y"})

	command := &driver.Command{}
	test := command.Build(x264, &media.Video{Path: "dsdsa"}, &media.Video{Path: "12321321"})

	fmt.Println(test)

}
