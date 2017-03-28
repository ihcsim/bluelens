package main

import "time"

const (
	defaultMusicFile     = "etc/music.json"
	defaultHistoryFile   = "etc/history.json"
	defaultFolloweesFile = "etc/followees.json"
	defaultCertFile      = "tls/cert.pem"
	defaultKeyFile       = "tls/key.pem"
	defaultTimeout       = time.Second * 10
)

type userConfig struct {
	musicFile     string
	historyFile   string
	followeesFile string
	user          string
	password      string
	apiKey        string
	certFile      string
	keyFile       string
	timeout       time.Duration
}

func newUserConfig() *userConfig {
	return &userConfig{
		musicFile:     defaultMusicFile,
		historyFile:   defaultHistoryFile,
		followeesFile: defaultFolloweesFile,
		certFile:      defaultCertFile,
		keyFile:       defaultKeyFile,
		timeout:       defaultTimeout,
	}
}
