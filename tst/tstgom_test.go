package tst

import "testing"

func TestTstGoM(t *testing.T) {

	got := TstGoM()
	expected := "hello tstgom"

	if got != expected {
		t.Errorf("got '%q' want '%q'", got, expected)
	}

}
