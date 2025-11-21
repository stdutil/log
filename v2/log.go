// Package log v2 is a lightweight direct logging library
//
// log is a derivative of LiveNote (github.com/narsilworks/livenote)
// to be used independently for stdutil
//
//	Author: Elizalde G. Baguinon
//	Created: 11/21/2025
package log

import (
	"fmt"
	"runtime"
	"strings"
)

type Log struct {
	Prefix  string
	ln      []LogInfo
	osIsWin bool
}

func NewLog(prefix string) *Log {
	return &Log{
		Prefix:  prefix,
		ln:      make([]LogInfo, 0),
		osIsWin: runtime.GOOS == "windows",
	}
}

// AddInfo adds an information message
func (r *Log) AddInfo(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	addMessage(&r.ln, r.Prefix, msg, Info, li.Format)
}

// AddWarning adds a warning message
func (r *Log) AddWarning(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	addMessage(&r.ln, r.Prefix, msg, Warn, li.Format)
}

// AddError adds an error message
func (r *Log) AddError(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	addMessage(&r.ln, r.Prefix, msg, Error, li.Format)
}

// AddFatal adds a fatal error message
func (r *Log) AddFatal(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	addMessage(&r.ln, r.Prefix, msg, Fatal, li.Format)
}

// AddSuccess adds a success message
func (r *Log) AddSuccess(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	addMessage(&r.ln, r.Prefix, msg, Success, li.Format)
}

// AddAppMsg adds an application message
func (r *Log) AddAppMsg(msg string, opts ...LogInfoOption) {
	li := LogInfo{
		Format: FormatNone,
	}
	for _, o := range opts {
		o(&li)
	}
	lpos := strings.Index(msg, "[")
	rpos := strings.Index(msg, "]")
	errT := ""
	ms := msg
	if (lpos > -1 && rpos > -1) && lpos < rpos {
		ms = msg[rpos+1:]
		errT = msg[0:lpos]
	}
	addMessage(&r.ln, r.Prefix, ms, LogType(errT), li.Format)
}

// Append adds a note object or more to the current list
func (r *Log) Append(ln ...LogInfo) {
	r.ln = append(r.ln, ln...)
}

// Clear live notes
func (r *Log) Clear() {
	r.ln = []LogInfo{}
}

// HasErrors checks if the message array has errors
func (r Log) HasErrors() bool {
	for _, ln := range r.ln {
		if ln.Type == Error {
			return true
		}
	}
	return false
}

// HasFatals checks if the message array has fatal errors
func (r Log) HasFatals() bool {
	for _, ln := range r.ln {
		if ln.Type == Fatal {
			return true
		}
	}
	return false
}

// HasWarnings checks if the message array has warnings
func (r Log) HasWarnings() bool {
	for _, ln := range r.ln {
		if ln.Type == Warn {
			return true
		}
	}
	return false
}

// HasInfos checks if the message array has information messages
func (r Log) HasInfos() bool {
	for _, ln := range r.ln {
		if ln.Type == Info {
			return true
		}
	}
	return false
}

// HasSuccess checks if the message array has success messages
func (r Log) HasSucceses() bool {
	for _, ln := range r.ln {
		if ln.Type == Success {
			return true
		}
	}
	return false
}

// Prevailing checks for a dominant message
func (r *Log) Prevailing() LogType {
	return getDominantNoteType(r.ln)
}

// Notes will list all notes
func (r *Log) Notes() []LogInfo {
	return r.ln
}

// String return the messages as a carriage/return delimited string
func (r *Log) String() string {
	lf := "\n"
	if r.osIsWin {
		lf = "\r\n"
	}
	sb := strings.Builder{}
	for _, v := range r.ln {
		sb.WriteString(v.String() + lf)
	}
	return sb.String()
}

// add new message to the message array
func addMessage(nt *[]LogInfo, prefix, msg string, typ LogType, styleFmt FontFormat) {
	msg = strings.TrimSpace(msg)
	*nt = append(*nt,
		LogInfo{
			Prefix:  prefix,
			Message: msg,
			Type:    typ,
			Format:  styleFmt,
		})
}

func getDominantNoteType(msgs []LogInfo) LogType {
	counts := map[LogType]int{
		Info:    0,
		Warn:    0,
		Error:   0,
		Success: 0,
		Fatal:   0,
	}

	// Count occurrences
	for _, msg := range msgs {
		if _, ok := counts[msg.Type]; ok {
			counts[msg.Type]++
		}
	}

	maxType := App
	maxCount := 0
	tie := false

	// Determine dominant type
	for t, c := range counts {
		if c > maxCount {
			maxCount = c
			maxType = t
			tie = false
		} else if c == maxCount && c != 0 {
			tie = true
		}
	}

	if tie {
		return App
	}
	return maxType
}

// Fmt accepts format and argument to return a string
func Fmt(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

// FormatHas checks if the format has one of the style
func FormatHas(sample, style FontFormat) bool {
	return sample&style != 0
}
