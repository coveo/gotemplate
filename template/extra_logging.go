package template

import (
	"os"
	"sync"

	"github.com/coveooss/multilogger"
	multicolor "github.com/coveooss/multilogger/color"
	"github.com/coveooss/multilogger/errors"
)

const (
	loggingBase = "Logging"
)

var (
	// TemplateLog is the logger used to log message during template processing
	TemplateLog = multilogger.New("gotemplate")
	// InternalLog is application logger used to follow the behaviour of the application
	InternalLog = multilogger.New("gotemplate-internal")
)

var loggingFuncs = dictionary{
	"trace":   func(args ...interface{}) string { return logBase(TemplateLog.Trace, args...) },
	"debug":   func(args ...interface{}) string { return logBase(TemplateLog.Debug, args...) },
	"info":    func(args ...interface{}) string { return logBase(TemplateLog.Info, args...) },
	"warning": func(args ...interface{}) string { return logBase(TemplateLog.Warning, args...) },
	"error":   func(args ...interface{}) string { return logBase(TemplateLog.Error, args...) },
	"fatal":   func(args ...interface{}) string { return logBase(TemplateLog.Fatal, args...) },
	"panic":   func(args ...interface{}) string { return logBase(TemplateLog.Panic, args...) },

	// Deprecated
	"notice": func(args ...interface{}) string {
		noticeWarning.Do(func() {
			InternalLog.Warning("Please note that notice is deprecated, use info instead.")
		})
		return logBase(TemplateLog.Info, args...)
	},
	"critical": func(args ...interface{}) string {
		criticalWarning.Do(func() {
			InternalLog.Warning("Please note that critical is deprecated, use error instead.")
		})
		return logBase(TemplateLog.Error, args...)
	},
}

var noticeWarning, criticalWarning sync.Once

var loggingFuncsAliases = aliases{
	"debug":   {"debugf"},
	"error":   {"errorf"},
	"fatal":   {"fatalf"},
	"info":    {"infof"},
	"panic":   {"panicf"},
	"trace":   {"tracef"},
	"warning": {"warn", "warnf", "warningf"},

	// Deprecated
	"notice":   {"noticef"},
	"critical": {"criticalf"},
}

var loggingFuncsHelp = descriptions{
	"panic":   "Logs a message using PANIC as log level (0) followed by a call to panic.",
	"fatal":   "Logs a message using FATAL as log level (1) followed by a call to os.Exit(1).",
	"error":   "Logs a message using ERROR as log level (2).",
	"warning": "Logs a message using WARNING as log level (3).",
	"info":    "Logs a message using INFO as log level (4).",
	"debug":   "Logs a message using DEBUG as log level (5).",
	"trace":   "Logs a message using TRACE as log level (6).",

	// Deprecated
	"critical": "Deprecated: Use error instead. Logs a message using ERROR log level (2).",
	"notice":   "Deprecated: Use info instead. Logs a message using INFO log level (4).",
}

func (t *Template) addLoggingFuncs() {
	t.AddFunctions(loggingFuncs, loggingBase, FuncOptions{
		FuncHelp:    loggingFuncsHelp,
		FuncAliases: loggingFuncsAliases,
	})
}

func logBase(f func(...interface{}), args ...interface{}) string {
	f(multicolor.FormatMessage(args...))
	return ""
}

func init() {
	if level := os.Getenv(EnvLogLevel); level != "" {
		if err := TemplateLog.SetHookLevel("", level); err != nil {
			errors.Printf("Unable to set logging level for templates: %v", err)
		}
	}
	if level := os.Getenv(EnvInternalLogLevel); level != "" {
		if err := InternalLog.SetHookLevel("", level); err != nil {
			errors.Printf("Unable to set logging level for internal logs: %v", err)
		}
	}
}
