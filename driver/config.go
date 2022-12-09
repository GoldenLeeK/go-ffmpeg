package driver

type Config struct {
	execLibraryPath string
	logPath         string
	threads         string
	timeout         string
}

func (config *Config) NewDrvier(drvier Driverable) Driverable {
	return drvier.Create(map[string]string{
		"exec_library_path": config.execLibraryPath,
		"log_path":          config.logPath,
		"threads":           config.threads,
		"timeout":           config.timeout,
	})
}
