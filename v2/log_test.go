package log

import "testing"

func TestString(t *testing.T) {
	l := NewLog("LOGGER")
	l.AddAppMsg("Application message")
	l.AddAppMsg("Application message", Format(FormatBold))
	l.AddError("Error writing message (bold)", Format(FormatBold))
	l.AddError("Error writing message (italic)", Format(FormatItalic))
	l.AddError("Error writing message (underline)", Format(FormatUnderline))
	l.AddFatal("Fatal error")
	l.AddInfo("More information")
	l.AddSuccess("Successful!")
	l.AddAppMsg("Anything goes!", Format(FormatBold|FormatItalic|FormatUnderline))
	t.Log(l.String())
}
