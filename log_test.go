package log

import "testing"

func TestAddAppMsg(t *testing.T) {
	// Add without prefix
	l := Log{}
	l.AddAppMsg("ERR: This is an error message!")

	// Add with prefix
	l.Prefix = "MESSAGE"
	l.AddAppMsg("ERR: This is an error message that has prefix!")
	t.Logf("HasError: %t, Message: %s", l.HasErrors(), l.ToString())
}
