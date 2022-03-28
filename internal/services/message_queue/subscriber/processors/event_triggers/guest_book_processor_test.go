package event_triggers

import (
	"reflect"
	"testing"
)

func Test_guestBookProcessor_GuestBookDelete(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := guestBookProcessor{}
			if err := g.GuestBookDelete(); (err != nil) != tt.wantErr {
				t.Errorf("GuestBookDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_guestBookProcessor_GuestBookInsert(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := guestBookProcessor{}
			if err := g.GuestBookInsert(); (err != nil) != tt.wantErr {
				t.Errorf("GuestBookInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newGuestBookProcessor(t *testing.T) {
	tests := []struct {
		name string
		want guestBookProcessor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newGuestBookProcessor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGuestBookProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
