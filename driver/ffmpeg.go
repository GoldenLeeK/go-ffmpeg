package driver

type FFmpeg struct {
	logPath         string
	execLibraryPath string
	threads         string
	timeout         string
}

func (ffmpeg *FFmpeg) GetName() string {
	return "ffpmeg"
}

func (ffmpeg *FFmpeg) Create(configtion map[string]string) Driverable {
	if libraraPath, ok := configtion["exec_library_path"]; ok {
		ffmpeg.execLibraryPath = libraraPath
	}
	if logPath, ok := configtion["log_path"]; ok {
		ffmpeg.logPath = logPath
	}
	if threads, ok := configtion["threads"]; ok {
		ffmpeg.threads = threads
	}
	if timeout, ok := configtion["timeout"]; ok {
		ffmpeg.timeout = timeout
	}
	return ffmpeg

}

func (ffmpeg *FFmpeg) Exec(command string) {

}
