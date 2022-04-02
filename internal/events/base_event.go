package events

type BaseEvent struct {
	Attributes map[string]interface{} `json:"attributes"`
}
