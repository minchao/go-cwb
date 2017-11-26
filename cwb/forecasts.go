package cwb

type ForecastsService service

type Parameter struct {
	ParameterName  string  `json:"parameterName"`
	ParameterValue *string `json:"parameterValue,omitempty"`
	ParameterUnit  *string `json:"parameterUnit,omitempty"`
}
