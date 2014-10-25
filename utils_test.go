package dockercommand

import "testing"

func TestConvertEnvMapToSlice(t *testing.T) {
	envMap := map[string]string{
		"TOTO1": "toto-1",
		"TOTO2": "toto-2",
	}
	slice := convertEnvMapToSlice(envMap)
	expectIsIn(t, slice, "TOTO1=toto-1")
	expectIsIn(t, slice, "TOTO2=toto-2")
}
