package web

import (
	"io"

	"github.com/davecgh/go-spew/spew"
	glog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type EchoLogger struct {
	*logrus.Entry
}

func (l *EchoLogger) SetPrefix(string) {
}
func (l *EchoLogger) Prefix() string {
	return ""
}
func (l *EchoLogger) SetLevel(level glog.Lvl) {
	l.Entry.Level = logrus.Level(level)
}
func (l *EchoLogger) Level() glog.Lvl {
	return glog.Lvl(l.Entry.Level)
}
func (l *EchoLogger) Output() io.Writer {
	return l.Entry.Writer()
}
func (l *EchoLogger) SetOutput(w io.Writer) {
	l.Entry.Logger.Out = w
}
func (l *EchoLogger) Printj(msg glog.JSON) {
	l.Entry.Print(spew.Sprint(msg))
}
func (l *EchoLogger) Debugj(msg glog.JSON) {
	l.Entry.Debug(spew.Sprint(msg))
}
func (l *EchoLogger) Infoj(msg glog.JSON) {
	l.Entry.Info(spew.Sprint(msg))
}
func (l *EchoLogger) Errorj(msg glog.JSON) {
	l.Entry.Error(spew.Sprint(msg))
}
func (l *EchoLogger) Warnj(msg glog.JSON) {
	l.Entry.Warn(spew.Sprint(msg))
}
func (l *EchoLogger) Fatalj(msg glog.JSON) {
	l.Entry.Fatal(spew.Sprint(msg))
}
func (l *EchoLogger) Panicj(msg glog.JSON) {
	l.Entry.Panic(spew.Sprint(msg))
}
