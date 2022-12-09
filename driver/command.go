package driver

import (
	"github.com/GoldenLeeK/go-ffmpeg/format"
	"github.com/GoldenLeeK/go-ffmpeg/media"
)

type Command struct {
}

func (c *Command) Build(format format.Formatable, input media.Mediaable, output media.Mediaable) (command string) {
	initParams := format.GetInitParams()

	for _, value := range initParams {
		command += value + " "
	}

	command += " -i " + input.Filepath()

	code := format.GetVideoCode()

	if code != "" {
		command += " -vcodec " + code
	}

	code = format.GetAudioCode()

	if code != "" {
		command += " -acodec " + code
	}

	addParams := format.GetAddParmas()

	for _, value := range addParams {
		command += value + " "
	}

	command += " " + output.Filepath()

	return command

}
