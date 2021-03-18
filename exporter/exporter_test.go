// +build windows

package exporter

import (
	"sort"
	"strings"
	"testing"
)

type expansionTestCase struct {
	input          string
	expectedOutput []string
}

func TestExpandEnabled(t *testing.T) {
	expansionTests := []expansionTestCase{
		{"", []string{}},
		// Default case
		{"cs,os", []string{"cs", "os"}},
		// Placeholder expansion
		{defaultCollectorsPlaceholder, strings.Split(defaultCollectors, ",")},
		// De-duplication
		{"cs,cs", []string{"cs"}},
		// De-duplicate placeholder
		{defaultCollectorsPlaceholder + "," + defaultCollectorsPlaceholder, strings.Split(defaultCollectors, ",")},
		// Composite case
		{"foo," + defaultCollectorsPlaceholder + ",bar", append(strings.Split(defaultCollectors, ","), "foo", "bar")},
	}

	for _, testCase := range expansionTests {
		output := expandEnabledCollectors(testCase.input)
		sort.Strings(output)

		success := true
		if len(output) != len(testCase.expectedOutput) {
			success = false
		} else {
			sort.Strings(testCase.expectedOutput)
			for idx := range output {
				if output[idx] != testCase.expectedOutput[idx] {
					success = false
					break
				}
			}
		}
		if !success {
			t.Error("For", testCase.input, "expected", testCase.expectedOutput, "got", output)
		}
	}
}

func TestNewCollector(t *testing.T) {
	//collector.iis.site-whitelist
	config :=
		`---
collector:
  iis:
    site-whitelist: test `
	_, err := NewWindowsCollector("iis", config)
	if err != nil {
		t.Error("Error in TestNewcollector creating collector with error ", err)
	}
}
