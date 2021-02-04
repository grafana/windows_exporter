package collector

import (
	"reflect"
	"testing"
)

func TestExpandChildCollectors(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput []string
	}{
		{
			name:           "simple",
			input:          "testing1,testing2,testing3",
			expectedOutput: []string{"testing1", "testing2", "testing3"},
		},
		{
			name:           "duplicate",
			input:          "testing1,testing2,testing2,testing3",
			expectedOutput: []string{"testing1", "testing2", "testing3"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := expandEnabledChildCollectors(c.input)
			if !reflect.DeepEqual(output, c.expectedOutput) {
				t.Errorf("Output mismatch, expected %+v, got %+v", c.expectedOutput, output)
			}
		})
	}
}

func TestNewCollectors(t *testing.T) {
	var (
		set1 = NewCollectors()
		set2 = NewCollectors()

		iis1 = set1.builders["iis"].(*IISCollectorConfig)
		iis2 = set2.builders["iis"].(*IISCollectorConfig)
	)
	if iis1 == iis2 {
		t.Errorf("collector config structs not copied")
	}
}
