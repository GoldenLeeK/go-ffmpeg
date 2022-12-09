package media

type Video struct {
	Path string
}

func (v *Video) Filepath() string {
	return v.Path
}
