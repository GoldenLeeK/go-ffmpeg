package format

type Formatable interface {
	GetVideoCode() string
	GetAudioCode() string
	SetVideoCode(string)
	SetAudioCode(string)
	SetKiloBitrate(int)
	GetKiloBitrate() int
	SetInitParams([]string)
	GetInitParams() []string
	SetAddParams([]string)
	GetAddParmas() []string
}
