package operators

import "github.com/tanggalnya/queue-actor/internal/domain"

var registry map[domain.EventTriggerType]Operator

type Operator interface {
	Process() error
	AfterProcess() error
}

func init() {
	registry = make(map[domain.EventTriggerType]Operator)
}

func Get(eventType domain.EventTriggerType) Operator {
	return registry[eventType]
}
