package main

import "testing"

func Test_selectChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test select chan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectChan()
		})
	}
}
