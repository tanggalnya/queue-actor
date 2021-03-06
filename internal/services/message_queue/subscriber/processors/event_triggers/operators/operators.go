package operators

import (
	"context"

	"github.com/tanggalnya/queue-actor/internal/domain"
)

var registry map[domain.EventTriggerType]Operator

type Operator interface {
	Process(ctx context.Context) error
	AfterProcess() error
}

func init() {
	registry = make(map[domain.EventTriggerType]Operator)
}

func RegisterEventOperator(eventType domain.EventTriggerType, op Operator) {
	registry[eventType] = op
}

func Get(eventType domain.EventTriggerType) Operator {
	return registry[eventType]
}
