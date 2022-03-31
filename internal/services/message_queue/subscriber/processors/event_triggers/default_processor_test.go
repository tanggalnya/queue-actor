package event_triggers

import (
	"reflect"
	"testing"
)

func Test_newDefaultProcessor(t *testing.T) {
	tests := []struct {
		name string
		want defaultProcessor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultProcessor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
