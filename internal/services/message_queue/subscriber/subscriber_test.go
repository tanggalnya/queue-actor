package subscriber

import "testing"

//TODO: mock this test
func Test_subscriberClient_Consume(t *testing.T) {
	type fields struct {
		exchangeName string
		exchangeKind string
		consumerName string
		queueName    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"success consume event",
			fields{
				exchangeName: "events",
				exchangeKind: "topic",
				consumerName: "google_drive",
				queueName:    "insert_google_drive",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := subscriberClient{
				uri:          "amqp://guest:guest@localhost",
				exchangeName: tt.fields.exchangeName,
				exchangeKind: tt.fields.exchangeKind,
				consumerName: tt.fields.consumerName,
				queueName:    tt.fields.queueName,
			}
			if err := s.Consume(); (err != nil) != tt.wantErr {
				t.Errorf("Consume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
