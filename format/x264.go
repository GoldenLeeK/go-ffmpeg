package format

type X264 struct {
	audioCode   string
	videoCode   string
	kiloBitrate int
	initParams  []string
	addParams   []string
}

func (x264 *X264) SetVideoCode(code string) {
	x264.videoCode = code
}
func (x264 *X264) GetVideoCode() string {
	return x264.videoCode
}
func (x264 *X264) SetAudioCode(code string) {
	x264.audioCode = code
}
func (x264 *X264) GetAudioCode() string {
	return x264.audioCode
}
func (x264 *X264) SetKiloBitrate(kiloBitrate int) {
	x264.kiloBitrate = kiloBitrate
}
func (x264 *X264) GetKiloBitrate() int {
	return x264.kiloBitrate
}
func (x264 *X264) SetInitParams(params []string) {
	for _, value := range params {
		x264.initParams = append(x264.initParams, value)
	}
}

func (x264 *X264) GetInitParams() []string {
	return x264.initParams
}

func (x264 *X264) SetAddParams(params []string) {
	for _, value := range params {
		x264.initParams = append(x264.addParams, value)
	}
}

func (x264 *X264) GetAddParmas() []string {
	return x264.addParams
}
