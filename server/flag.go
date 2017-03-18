package main

import "flag"

const (
	defaultMusicFile     = "etc/music.json"
	defaultHistoryFile   = "etc/history.json"
	defaultFolloweesFile = "etc/followees.json"
)

type userConfig struct {
	musicFile     string
	historyFile   string
	followeesFile string
}

func newUserConfig() *userConfig {
	return &userConfig{
		musicFile:     defaultMusicFile,
		historyFile:   defaultHistoryFile,
		followeesFile: defaultFolloweesFile,
	}
}

func parseFlags(args []string) (*userConfig, error) {
	c := &userConfig{}

	flagSet := flag.NewFlagSet("bluelensFlags", flag.ContinueOnError)
	flagSet.StringVar(&c.musicFile, "music", defaultMusicFile, "Path to read music data from")
	flagSet.StringVar(&c.historyFile, "history", defaultHistoryFile, "Path to read user's history data from")
	flagSet.StringVar(&c.followeesFile, "followees", defaultFolloweesFile, "Path to read user's followees data from")

	if !flagSet.Parsed() {
		if err := flagSet.Parse(args); err != nil {
			return nil, err
		}
	}

	return c, nil
}
