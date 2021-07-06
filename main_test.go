package main

import "testing"

func TestGenReport(t *testing.T) {
	if err := GenReport("test version"); err != nil {
		t.Errorf("%q", err)
	}
}
