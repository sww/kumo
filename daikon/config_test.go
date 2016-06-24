package daikon

import (
	"reflect"
	"strings"
	"testing"
)

func Test_GetConfig(t *testing.T) {
	c := strings.NewReader(`
{
    "connections": 1,
    "host": "host",
	   "username": "user",
	   "password": "pass",
    "port": 119,
	   "temp": "tmp",
	   "download": "download"
}`)
	config, err := GetConfig(c)
	if err != nil {
		t.Fatalf("Config error %v", err)
	}

	want := Config{Connections: 1, Host: "host", Username: "user", Password: "pass", Port: 119, Temp: "tmp", Download: "download"}
	if !reflect.DeepEqual(config, &want) {
		t.Errorf("Returned %+v, want %+v", config, &want)
	}
}
