package ner

import (
	"context"
	"log"

	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/data"
	"github.com/datravis/go-meetup-solution/src/go/protogen"
)

func ExtractEntities(ctx context.Context, nerClient protogen.NerServiceClient, req data.PipelineRequest) data.PipelineRequest {
	log.Println("Extracting Entities")
	resp, err := nerClient.ExtractSubjectOrgs(ctx, &protogen.NerRequest{
		Input: req.Input,
	})
	req.Error = err
	if err == nil {
		req.Entities = resp.Entities
	}

	return req
}
