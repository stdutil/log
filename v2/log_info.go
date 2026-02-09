package log

import "strings"

// LogType
type LogType string
type LogMedia int8
type FontFormat uint16

type LogInfo struct {
	Type    LogType
	Format  FontFormat
	Prefix  string
	Message string
	Media   LogMedia
}

// Formats
const (
	FormatNone      FontFormat = 0
	FormatBold      FontFormat = 1 << iota // 1  -> 0001
	FormatItalic                           // 2  -> 0010
	FormatUnderline                        // 4  -> 0100
)

// LogType constants
const (
	Info    LogType = "INF"
	Warn    LogType = "WRN"
	Error   LogType = "ERR"
	Fatal   LogType = "FTL"
	Success LogType = "SUC"
	App     LogType = ""
)

const (
	escStart string = "\033["
	escEnd   string = "m"
	rst      string = "0"
	fmtBold  string = "1"
	fmtItal  string = "3"
	fmtUndr  string = "4"
	errClr   string = "91"
	ftlClr   string = "31"
	infClr   string = "96"
	wrnClr   string = "33"
	sucClr   string = "32"
)

const (
	MediaConsole LogMedia = iota
	MediaWeb
)

// String returns the message as string
func (lni LogInfo) String() string {
	sb := strings.Builder{}

	shouldHaveStyle := false
	switch lni.Type {
	case Error:
		shouldHaveStyle = true
	case Fatal:
		shouldHaveStyle = true
	case Warn:
		shouldHaveStyle = true
	case Info:
		shouldHaveStyle = true
	case Success:
		shouldHaveStyle = true
	}

	if lni.Type != App {
		if shouldHaveStyle && lni.Media == MediaConsole {
			switch lni.Type {
			case Error:
				sb.WriteString(sgr(errClr))
			case Fatal:
				sb.WriteString(sgr(ftlClr))
			case Warn:
				sb.WriteString(sgr(wrnClr))
			case Info:
				sb.WriteString(sgr(infClr))
			case Success:
				sb.WriteString(sgr(sucClr))
			}
		}

		sb.WriteString(string(lni.Type))

		// Reset
		if shouldHaveStyle && lni.Media == MediaConsole {
			sb.WriteString(sgr(rst))
		}

		if lni.Prefix != "" {
			sb.WriteString("[")
			sb.WriteString(lni.Prefix)
			sb.WriteString("]")
		}
		sb.WriteString(`: `)
	}

	if lni.Format != FormatNone && lni.Media == MediaConsole {
		var codes []string
		if lni.Format.Has(FormatBold) {
			codes = append(codes, fmtBold)
		}
		if lni.Format.Has(FormatItalic) {
			codes = append(codes, fmtItal)
		}
		if lni.Format.Has(FormatUnderline) {
			codes = append(codes, fmtUndr)
		}
		if len(codes) > 0 {
			sb.WriteString(sgr(codes...))
		}
	}

	sb.WriteString(lni.Message)

	if lni.Format != FormatNone && lni.Media == MediaConsole {
		sb.WriteString(sgr(rst))
	}

	return sb.String()
}

func (f FontFormat) Has(flag FontFormat) bool {
	return f&flag != 0
}

func sgr(codes ...string) string {
	if len(codes) == 0 {
		return escStart + escEnd // ESC[m
	}
	return escStart + strings.Join(codes, ";") + escEnd
}
