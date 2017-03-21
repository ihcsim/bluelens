package main

import (
	"flag"

	"github.com/ihcsim/bluelens/server/config"
)

func parseFlags(args []string) (*config.UserConfig, error) {
	c := &config.UserConfig{}

	flagSet := flag.NewFlagSet("bluelensFlags", flag.ContinueOnError)
	flagSet.StringVar(&c.MusicFile, "music", config.DefaultMusicFile, "Path to read music data from")
	flagSet.StringVar(&c.HistoryFile, "history", config.DefaultHistoryFile, "Path to read user's history data from")
	flagSet.StringVar(&c.FolloweesFile, "followees", config.DefaultFolloweesFile, "Path to read user's followees data from")

	if !flagSet.Parsed() {
		if err := flagSet.Parse(args); err != nil {
			return nil, err
		}
	}

	return c, nil
}
