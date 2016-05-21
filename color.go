package logplus

type Color int

const (
	ForegroundDefault Color = iota
	ForegroundBlack
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundMagenta
	ForegroundCyan
	ForegroundWhite

	BackgroundDefault
	BackgroundBlack
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

func (color Color) String() string {
	switch color {
	case ForegroundDefault:
		return "\033[39m"
	case ForegroundBlack:
		return "\033[30m"
	case ForegroundRed:
		return "\033[31m"
	case ForegroundGreen:
		return "\033[32m"
	case ForegroundYellow:
		return "\033[33m"
	case ForegroundBlue:
		return "\033[34m"
	case ForegroundMagenta:
		return "\033[35m"
	case ForegroundCyan:
		return "\033[36m"
	case ForegroundWhite:
		return "\033[97m"
	case BackgroundDefault:
		return "\033[49m"
	case BackgroundBlack:
		return "\033[40m"
	case BackgroundRed:
		return "\033[41m"
	case BackgroundGreen:
		return "\033[42m"
	case BackgroundYellow:
		return "\033[43m"
	case BackgroundBlue:
		return "\033[44m"
	case BackgroundMagenta:
		return "\033[45m"
	case BackgroundCyan:
		return "\033[46m"
	case BackgroundWhite:
		return "\033[107m"
	}
	panic("Unknown value")
}
