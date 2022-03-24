package common_test

import (
	"github.com/prongbang/localizegen/pkg/common"
	"testing"
)

func TestToCamel1(t *testing.T) {
	s := common.ToCamelClass("keys_localization.dart")
	if !(s == "KeysLocalization") {
		t.Error("Cannot convert to camel", s)
	}
}

func TestToCamel2(t *testing.T) {
	s := common.ToCamelClass("keys_localization")
	if !(s == "KeysLocalization") {
		t.Error("Cannot convert to camel", s)
	}
}

func TestToCamelVariable(t *testing.T) {
	s := common.ToCamelVariable("keys_localization")
	if !(s == "keysLocalization") {
		t.Error("Cannot convert to camel", s)
	}
}
