package golog

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	GetInstance().SetLogLevel(LevelString("DEBUG"))
	Debug("testing debug")
	Info("testing info")
	Warning("testing warning")
	Error("testing error")
}
