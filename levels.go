package golog

type Level int

const (
	LevelOff Level = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)

var level_to_string = map[Level]string{
	LevelOff:     "OFF",
	LevelFatal:   "FATAL",
	LevelError:   "ERROR",
	LevelWarning: "WARNING",
	LevelInfo:    "INFO",
	LevelDebug:   "DEBUG",
	LevelTrace:   "TRACE",
	LevelAll:     "ALL",
}

func (l Level) String() string {
	s := level_to_string[l]
	return s
}

func LevelString(s string) Level {

	for l, v := range level_to_string {
		if s == v {
			return l
		}
	}

	panic("Invalid level string")
}

var (
	// Normal colors
	nBlack   = []byte{'\033', '[', '3', '0', 'm'}
	nRed     = []byte{'\033', '[', '3', '1', 'm'}
	nGreen   = []byte{'\033', '[', '3', '2', 'm'}
	nYellow  = []byte{'\033', '[', '3', '3', 'm'}
	nBlue    = []byte{'\033', '[', '3', '4', 'm'}
	nMagenta = []byte{'\033', '[', '3', '5', 'm'}
	nCyan    = []byte{'\033', '[', '3', '6', 'm'}
	nWhite   = []byte{'\033', '[', '3', '7', 'm'}
	// Bright colors
	bBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	bRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	bGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	bYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	bBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	bMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	bCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	bWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	reset = []byte{'\033', '[', '0', 'm'}
)

var Level2Color = map[Level][]byte{
	LevelFatal:   bRed,
	LevelError:   bRed,
	LevelWarning: bBlue,
	LevelInfo:    bGreen,
	LevelDebug:   bCyan,
	LevelTrace:   bWhite,
}
