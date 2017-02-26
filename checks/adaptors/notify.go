package adaptors

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/monkeyherder/moirai/checks"
	"reflect"
)

func NewNotifierLogger(logger boshlog.Logger) checks.CheckAdaptor {
	return Notify(notifierLogger{
		Logger: logger,
	})
}

type Notifier interface {
	BeforeCheck(checks.Check)
	AfterCheck(checks.Check)
	OnError(checks.Check, error)
}

func Notify(notifier Notifier) checks.CheckAdaptor {
	return func(check checks.Check) checks.Check {
		return checks.CheckFunc(func() (string, string, error) {
			notifier.BeforeCheck(check)
			_, _, err := check.Run()
			if err != nil {
				notifier.OnError(check, err)
			}
			notifier.AfterCheck(check)
			return "", "", err
		})
	}
}

type notifierLogger struct {
	Logger boshlog.Logger
}

func (n notifierLogger) BeforeCheck(check checks.Check) {
	n.Logger.Debug(reflect.TypeOf(check).Name(), "Before Check ran")
}

func (n notifierLogger) AfterCheck(check checks.Check) {
	n.Logger.Debug(reflect.TypeOf(check).Name(), "After Check ran")
}

func (n notifierLogger) OnError(check checks.Check, err error) {
	n.Logger.Error(reflect.TypeOf(check).Name(), "Error occurred", err)
}
