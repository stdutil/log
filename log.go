// Package log is a lightweight direct logging library
//
// log is a derivative of LiveNote (github.com/narsilworks/livenote) to be used
// independently for stdutil
//
//	Author: Elizalde G. Baguinon
//	Created: May 28, 2022
package log

import (
	"fmt"
	"runtime"
	"strings"
)

// LogType
type LogType string

// LogType constants
const (
	Info    LogType = "INF"
	Warn    LogType = "WRN"
	Error   LogType = "ERR"
	Fatal   LogType = "FTL"
	Success LogType = "SUC"
	App     LogType = ""
)

const DelimMsgType string = `: `

type Log struct {
	Prefix  string // Prefix
	ln      []LogInfo
	osIsWin bool
}

type LogInfo struct {
	Type    LogType
	Prefix  string
	Message string
}

func NewLog(prefix string) *Log {
	return &Log{
		Prefix:  prefix,
		ln:      make([]LogInfo, 0),
		osIsWin: runtime.GOOS == "windows",
	}
}

// Fmt accepts format and argument to return a string
func Fmt(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

// AddInfo adds an information message
func (r *Log) AddInfo(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Info)
	}
}

// AddWarning adds a warning message
func (r *Log) AddWarning(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Warn)
	}
}

// AddError adds an error message
func (r *Log) AddError(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Error)
	}
}

// AddSuccess adds a success message
func (r *Log) AddSuccess(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, Success)
	}
}

// AddAppMsg adds an application message
func (r *Log) AddAppMsg(msg ...string) {
	for _, m := range msg {
		addMessage(&r.ln, r.Prefix, m, App)
	}
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
	return getDominantNoteType(&r.ln)
}

// Notes will list all notes
func (r *Log) Notes() []LogInfo {
	return r.ln
}

// ToString return the messages as a carriage/return delimited string
func (r *Log) ToString() string {
	lf := "\n"
	if r.osIsWin {
		lf = "\r\n"
	}
	sb := strings.Builder{}
	for _, v := range r.ln {
		sb.Write([]byte(v.ToString() + lf))
	}
	return sb.String()
}

// ToString return the messages as a carriage/return delimited string
func (lni *LogInfo) ToString() string {
	td := ""
	if lni.Type != App {
		td += string(lni.Type)
		if lni.Prefix != "" {
			td += "[" + lni.Prefix + "]"
		}
		td += DelimMsgType
	}
	td += lni.Message
	return td
}

// add new message to the message array
func addMessage(nt *[]LogInfo, prefix, msg string, typ LogType) {
	msg = strings.TrimSpace(msg)
	*nt = append(*nt, LogInfo{
		Prefix:  prefix,
		Message: msg,
		Type:    typ,
	})
}

// get dominant message
func getDominantNoteType(msgs *[]LogInfo) LogType {
	var (
		nfo, wrn, err, suc int
	)

	for _, msg := range *msgs {
		switch msg.Type {
		case Info:
			nfo++
		case Warn:
			wrn++
		case Error:
			err++
		case Success:
			suc++
		}
	}
	if nfo > wrn && nfo > err && nfo > suc {
		return Info
	}
	if wrn > nfo && wrn > err && wrn > suc {
		return Warn
	}
	if err > nfo && err > wrn && err > suc {
		return Error
	}
	if suc > nfo && suc > wrn && suc > err {
		return Success
	}
	return App
}
