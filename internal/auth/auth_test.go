package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := make(http.Header)
	header2 := make(http.Header)

	header1.Set("Authorization", "ApiKey uahhe98hh98hjh")
	header2.Set("Authorization", "ApiKey isuehiuhsa auhsiuehiuahse")

	get1, _ := GetAPIKey(header1)
	get2, _ := GetAPIKey(header2)
	want1 := "uahhe98hh98hjh"
	want2 := "isuehiuhsa auhsiuehiuahse"

	if !reflect.DeepEqual(get1, want1) {
		t.Fatalf("error1")
	}

	if !reflect.DeepEqual(get2, want2) {
		t.Fatalf("error2")
	}
}
