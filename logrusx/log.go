package logrusx

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type Fields = logrus.Fields

func init() {
	Log.Out = os.Stdout
	Log.Formatter = &logrus.JSONFormatter{}
	Log.SetLevel(logrus.InfoLevel)
}
