package data

type PipelineRequest struct {
	Source    string
	Input     string
	Entities  []string
	Sentiment float64
	Error     error
}
