package models

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger("./", DebugLevel)

	if err != nil {
		t.Fatal(err)
	}

	logger.Infof("res")
}
