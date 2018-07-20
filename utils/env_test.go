package utils

import "testing"

func Test_GetEnv(t *testing.T) {
	str := "APP_PORT"

	env := GetEnv(str)

	if env == "7447" {
		t.Log("one test passed")
	} else {
		t.Error("One test fail")
	}
}
