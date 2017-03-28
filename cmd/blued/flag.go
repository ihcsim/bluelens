package main

import "flag"

func parseFlags(args []string) (*userConfig, error) {
	c := newUserConfig()

	flagSet := flag.NewFlagSet("bluelensFlags", flag.ContinueOnError)
	flagSet.StringVar(&c.musicFile, "music", defaultMusicFile, "Path to read music data from")
	flagSet.StringVar(&c.historyFile, "history", defaultHistoryFile, "Path to read user's history data from")
	flagSet.StringVar(&c.followeesFile, "followees", defaultFolloweesFile, "Path to read user's followees data from")
	flagSet.StringVar(&c.user, "user", "", "Username used for HTTP Basic Authentication")
	flagSet.StringVar(&c.password, "password", "", "Password used for HTTP Basic Authentication")
	flagSet.StringVar(&c.apiKey, "apikey", "", "Key used for API key authentication")
	flagSet.StringVar(&c.certFile, "cert", defaultCertFile, "Path to the TLS cert file")
	flagSet.StringVar(&c.keyFile, "private", defaultKeyFile, "Path to the TLS private key file")
	flagSet.DurationVar(&c.timeout, "timeout", defaultTimeout, "Request timeout in seconds. Default to 10s.")

	if !flagSet.Parsed() {
		if err := flagSet.Parse(args); err != nil {
			return nil, err
		}
	}

	return c, nil
}
