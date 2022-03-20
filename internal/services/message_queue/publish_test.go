package message_queue

import "testing"

func Test_publisherClient_PublishEvent(t *testing.T) {
	type fields struct {
		QueueName string
		Reliable  bool
	}
	type args struct {
		body string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"success publish event",
			fields{
				QueueName: "test",
				Reliable:  false,
			},
			args{body: "test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := publisherClient{
				uri:       "amqp://guest:guest@localhost",
				queueName: tt.fields.QueueName,
				reliable:  tt.fields.Reliable,
			}
			if err := p.PublishEvent(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PublishEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
