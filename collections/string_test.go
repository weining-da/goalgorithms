package collections

import (
	"github.com/weining-da/goalgorithms/resources"
	"testing"
)

func TestStringCreation(t *testing.T) {
	if CreateString() != "asd" {
		t.Errorf("not working\n")
	}
}

func TestReadEmbeddedFiles(t *testing.T) {
	if _, ok := resources.ReadConfFile("config/conf.json"); !ok {
		t.Errorf("not working\n")
	}
	if _, ok := resources.ReadTestFile("testdata/base/assets/model.mdl"); !ok {
		t.Errorf("not working\n")
	}
}
