package cwb

type Result struct {
	ResourceId *string  `json:"resource_id,omitempty"`
	Fields     []Fields `json:"fields"`
}

type Fields struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
