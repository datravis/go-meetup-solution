package snt

import (
	"context"
	"log"

	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/data"
	"github.com/datravis/go-meetup-solution/src/go/protogen"
)

func Analyze(ctx context.Context, sntClient protogen.SentimentServiceClient, req data.PipelineRequest) data.PipelineRequest {
	log.Println("Determining Sentiment")
	resp, err := sntClient.Analyze(ctx, &protogen.SentimentRequest{
		Input: req.Input,
	})
	req.Error = err
	if err == nil {
		req.Sentiment = float64(resp.Sentiment)
	}

	return req
}
