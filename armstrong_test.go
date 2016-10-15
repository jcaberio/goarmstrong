package goarmstrong_test

import (
	"github.com/jcaberio/goarmstrong"
	"testing"
)

func TestArmstrong(t *testing.T) {
	v := goarmstrong.IsArmstrong(912985153)
	if !v {
		t.Error("Expected true, got ", v)
	}
}
