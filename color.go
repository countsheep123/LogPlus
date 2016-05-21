package logplus

type Color int

const (
	TextDefault Color = iota
	TextBlack
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite

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

const (
	resetAll = "\033[0m"
)

func (color Color) String() string {
	switch color {
	case TextDefault:
		return "\033[39m"
	case TextBlack:
		return "\033[30m"
	case TextRed:
		return "\033[31m"
	case TextGreen:
		return "\033[32m"
	case TextYellow:
		return "\033[33m"
	case TextBlue:
		return "\033[34m"
	case TextMagenta:
		return "\033[35m"
	case TextCyan:
		return "\033[36m"
	case TextWhite:
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
