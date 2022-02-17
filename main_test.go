package main

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func Test(t *testing.T) {
	actual := emoji.Sprint("Hello:world_map:!")
	expected := "Hello🗺️ !"
	
	if expected != actual {

		t.Errorf("got %q, wanted %q", actual, expected)
	}
}
