package config

const (
	// DefaultMusicFile is the path to the default music data file.
	DefaultMusicFile = "etc/music.json"

	// DefaultHistoryFile is the path to the default history data file.
	DefaultHistoryFile = "etc/history.json"

	// DefaultFolloweesFile is the path to the default followees data file.
	DefaultFolloweesFile = "etc/followees.json"
)

// UserConfig captures the user-provided runtime configurations.
type UserConfig struct {
	// MusicFile is the path to the music data file.
	MusicFile string

	// HistoryFile is the path to the history data file.
	HistoryFile string

	// FolloweesFile is the path to the followees data file.
	FolloweesFile string
}

// NewUserConfig returns a new instance of UserConfig.
func NewUserConfig() *UserConfig {
	return &UserConfig{
		MusicFile:     DefaultMusicFile,
		HistoryFile:   DefaultHistoryFile,
		FolloweesFile: DefaultFolloweesFile,
	}
}
