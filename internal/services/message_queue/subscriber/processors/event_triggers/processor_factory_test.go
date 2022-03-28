package event_triggers

import (
	"github.com/tanggalnya/queue-actor/internal/domain"
	"reflect"
	"testing"
)

func Test_factory_NewProcessor(t *testing.T) {
	type args struct {
		tableName domain.EventTriggerType
	}
	tests := []struct {
		name string
		args args
		want Processor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := factory{}
			if got := f.NewProcessor(tt.args.tableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
