package plugin

import (
	"context"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (ds *AlcGrafanaDataSourceInstance) getCacheStats(ctx context.Context,
	refId string) (*backend.DataResponse, error) {

	now := time.Now()

	response := backend.DataResponse{}
	// create data frame response

	frame := data.NewFrame(refId)
	frame.Fields = append(frame.Fields,
		data.NewField("time", nil, []time.Time{now}),
		data.NewField("PV Cache Hits", nil, []uint64{ds.pvCacheHits}),
		data.NewField("PV SOAP Requests", nil, []uint64{ds.pvSoapRequests}),
		data.NewField("Tree Cache Hits", nil, []uint64{ds.treeCacheHits}),
		data.NewField("Trend Data Requests", nil, []uint64{ds.trendDsRequests}),
		data.NewField("Trend SOAP Requests", nil, []uint64{ds.trendSoapRequests}),
	)

	response.Frames = append(response.Frames, frame)

	return &response, nil
}
