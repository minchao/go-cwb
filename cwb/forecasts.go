package cwb

type ForecastsService service

type Result struct {
	ResourceId string   `json:"resource_id"`
	Fields     []Fields `json:"fields"`
}

type Fields struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Parameter struct {
	ParameterName  string  `json:"parameterName"`
	ParameterValue *string `json:"parameterValue,omitempty"`
	ParameterUnit  *string `json:"parameterUnit,omitempty"`
}
