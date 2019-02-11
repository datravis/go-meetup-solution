package db

import (
	"log"
	"strings"

	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/data"
	"github.com/influxdata/influxdb/client/v2"
	influx "github.com/influxdata/influxdb/client/v2"
)

func NewInfluxDBClient(url string) (influx.Client, error) {
	return influx.NewHTTPClient(client.HTTPConfig{
		Addr: url,
	})
}

func WriteSentiment(client influx.Client, measurement string, database string, req data.PipelineRequest) data.PipelineRequest {
	log.Println("Writing to InfluxDB")

	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	if err != nil {
		req.Error = err
		return req
	}

	for _, entity := range req.Entities {
		if entity == "" {
			continue
		}
		entity = strings.Trim(entity, " ")
		entity = strings.Replace(entity, "\n", "", -1)
		pt, err := influx.NewPoint(
			measurement,
			map[string]string{
				"entity": entity,
				"source": req.Source,
			},
			map[string]interface{}{
				"sentiment": req.Sentiment,
			},
		)
		if err != nil {
			req.Error = err
			return req
		}
		bp.AddPoint(pt)
	}

	if len(bp.Points()) == 0 {
		log.Println("Skipping write, no points")
		return req
	}

	req.Error = client.Write(bp)
	return req
}
