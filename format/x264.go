package format

type X264 struct {
	formatCode  string
	kiloBitrate int
	initParams  []string
	addParams   []string
}

func (x264 *X264) init() {
	x264.setFormatCode("libx264")
	x264.initParams = append(x264.initParams, "-y -i")
}
func (x264 *X264) setFormatCode(code string) {
	x264.formatCode = code
}
func (x264 *X264) getFormatCode() string {
	return x264.formatCode
}
func (x264 *X264) setKiloBitrate(kiloBitrate int) {
	x264.kiloBitrate = kiloBitrate
}
func (x264 *X264) getKiloBitrate(kiloBitrate int) int {
	return x264.kiloBitrate
}
func (x264 *X264) setInitParams(params []string) {
	for _, value := range params {
		x264.initParams = append(x264.initParams, value)
	}
}

func (x264 *X264) getInitParams() []string {
	return x264.initParams
}

func (x264 *X264) setAddParams(params []string) {
	for _, value := range params {
		x264.initParams = append(x264.addParams, value)
	}
}

func (x264 *X264) getAddParmas() []string {
	return x264.addParams
}
