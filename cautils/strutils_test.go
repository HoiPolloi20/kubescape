package cautils

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertLabelsToString(t *testing.T) {
	str := "a=b;c=d"
	strMap := map[string]string{"a": "b", "c": "d"}
	rsrt := ConvertLabelsToString(strMap)
	spilltedA := strings.Split(rsrt, ";")
	spilltedB := strings.Split(str, ";")
	for i := range spilltedA {
		if spilltedA[i] != spilltedB[i] {
			t.Errorf("%s != %s", spilltedA[i], spilltedB[i])
		}
	}
}

func TestConvertStringToLabels(t *testing.T) {
	str := "a=b;c=d"
	strMap := map[string]string{"a": "b", "c": "d"}
	rstrMap := ConvertStringToLabels(str)
	if fmt.Sprintf("%v", rstrMap) != fmt.Sprintf("%v", strMap) {
		t.Errorf("%s != %s", fmt.Sprintf("%v", rstrMap), fmt.Sprintf("%v", strMap))
	}
}
