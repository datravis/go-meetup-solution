package srv

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	influx "github.com/influxdata/influxdb/client/v2"

	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/data"
	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/db"
	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/ner"
	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/snt"
	"github.com/datravis/go-meetup-solution/src/go/protogen"
)

type PipelineService struct {
	nerClient       protogen.NerServiceClient
	sentimentClient protogen.SentimentServiceClient
	influxClient    influx.Client
}

func (p PipelineService) Ingest(stream protogen.PipelineService_IngestServer) error {
	ctx := stream.Context()

	for in := range p.Write(ctx, p.Sentiment(ctx, p.Extract(ctx, p.Recv(ctx, stream)))) {
		if in.Error != nil {
			log.Println(in.Error)
			return in.Error
		}
	}

	stream.SendAndClose(&protogen.IngestResponse{
		Message: "OK",
	})
	return nil
}

func (p PipelineService) Recv(ctx context.Context, stream protogen.PipelineService_IngestServer) <-chan data.PipelineRequest {
	outChan := make(chan data.PipelineRequest, 0)

	go func() {
		defer close(outChan)
		for {
			request := data.PipelineRequest{}

			in, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				request.Error = err
				outChan <- request
				break
			}
			log.Println("Receiving message")
			request.Input = in.Input
			request.Source = in.Source
			outChan <- request
		}
	}()

	return outChan
}

func (p PipelineService) Extract(ctx context.Context, inChan <-chan data.PipelineRequest) <-chan data.PipelineRequest {
	outChan := make(chan data.PipelineRequest, 0)
	go func() {
		for req := range inChan {
			if req.Error != nil {
				outChan <- req
				continue
			}

			outChan <- ner.ExtractEntities(ctx, p.nerClient, req)
		}
	}()

	return outChan
}

func (p PipelineService) Sentiment(ctx context.Context, inChan <-chan data.PipelineRequest) <-chan data.PipelineRequest {
	outChan := make(chan data.PipelineRequest, 0)
	go func() {
		for req := range inChan {
			if req.Error != nil {
				outChan <- req
				continue
			}

			outChan <- snt.Analyze(ctx, p.sentimentClient, req)
		}
	}()

	return outChan
}

func (p PipelineService) Write(ctx context.Context, inChan <-chan data.PipelineRequest) <-chan data.PipelineRequest {
	outChan := make(chan data.PipelineRequest, 0)
	go func() {
		for req := range inChan {
			if req.Error != nil {
				outChan <- req
				continue
			}

			outChan <- db.WriteSentiment(p.influxClient, "entity_sentiment", "meetup", req)
		}
	}()

	return outChan
}

func NewServer() *PipelineService {
	nerConn, err := grpc.Dial("ner:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	sentimentConn, err := grpc.Dial("sentiment:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	influxClient, err := db.NewInfluxDBClient("http://influxdb:8086")
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &PipelineService{
		nerClient:       protogen.NewNerServiceClient(nerConn),
		sentimentClient: protogen.NewSentimentServiceClient(sentimentConn),
		influxClient:    influxClient,
	}
}
