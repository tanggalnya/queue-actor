package apiModels

type HasuraEvent struct {
	ID      string `json:"id"`
	Event   `json:"event"`
	Table   `json:"table"`
	Trigger `json:"trigger"`
}

type Event struct {
	Op   string `json:"op"`
	Data `json:"data"`
}

type Data struct {
	Old map[string]interface{} `json:"old"`
	New map[string]interface{} `json:"new"`
}

type Table struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type Trigger struct {
	Name string `json:"name"`
}
