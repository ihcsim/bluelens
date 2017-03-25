package main

const (
	defaultMusicFile     = "etc/music.json"
	defaultHistoryFile   = "etc/history.json"
	defaultFolloweesFile = "etc/followees.json"
)

type userConfig struct {
	musicFile     string
	historyFile   string
	followeesFile string
	user          string
	password      string
	apiKey        string
}

func newUserConfig() *userConfig {
	return &userConfig{
		musicFile:     defaultMusicFile,
		historyFile:   defaultHistoryFile,
		followeesFile: defaultFolloweesFile,
	}
}
