package log

import "testing"

func TestString(t *testing.T) {
	l := NewLog(Prefix("LOGGER"))
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

func TestStringRaw(t *testing.T) {
	l := NewLog(Prefix("LOGGER"))
	l.AddAppMsg("Application message")
	l.AddAppMsg("Application message", Format(FormatBold))
	l.AddError("Error writing message (bold)", Format(FormatBold))
	l.AddError("Error writing message (italic)", Format(FormatItalic))
	l.AddError("Error writing message (underline)", Format(FormatUnderline))
	l.AddFatal("Fatal error")
	l.AddInfo("More information")
	l.AddSuccess("Successful!")
	l.AddAppMsg("Anything goes!", Format(FormatBold|FormatItalic|FormatUnderline))
	t.Log(l.StringRaw())
}

func TestAddAppMsg(t *testing.T) {
	// Add without prefix
	l := Log{}
	l.AddAppMsg("ERR: This is an error message!")

	// Add with prefix
	l.prefix = "MESSAGE"
	l.AddAppMsg("ERR: This is an error message that has prefix!")
	t.Logf("HasError: %t, Message: %s", l.HasErrors(), l.String())
}
