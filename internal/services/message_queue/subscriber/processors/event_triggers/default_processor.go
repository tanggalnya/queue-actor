package event_triggers

type defaultProcessor struct{}

func (g defaultProcessor) Insert() error {
	//TODO implement me
	panic("implement me")
}

func (g defaultProcessor) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (g defaultProcessor) Update() error {
	//TODO implement me
	panic("implement me")
}

func newDefaultProcessor() defaultProcessor {
	return defaultProcessor{}
}
