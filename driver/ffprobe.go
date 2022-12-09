package driver

import "fmt"

type FFprobe struct {
	logPath         string
	execLibraryPath string
	threads         string
	timeout         string
}

func (ffprobe *FFprobe) GetName() string {
	return "ffprobe"
}

func (ffprobe *FFprobe) Create(configtion map[string]string) FFprobe {
	if libraraPath, ok := configtion["exec_library_path"]; ok {
		ffprobe.execLibraryPath = libraraPath
	}
	if logPath, ok := configtion["log_path"]; ok {
		ffprobe.logPath = logPath
	}
	if threads, ok := configtion["threads"]; ok {
		ffprobe.threads = threads
	}
	if timeout, ok := configtion["timeout"]; ok {
		ffprobe.timeout = timeout
	}
	return *ffprobe

}
func (FFprobe *FFprobe) Exec(command string) {
	fmt.Println("")
}
