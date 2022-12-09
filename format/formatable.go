package format

type Formatable interface {
	getFormatCode() string
	setKiloBitrate()
	getKiloBitrate() int
	setInitParams([]string)
	getInitParams() []string
	setAddParams([]string)
	getAddParmas() []string
}
