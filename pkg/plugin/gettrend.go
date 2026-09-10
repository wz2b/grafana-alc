package plugin

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (ds *AlcGrafanaDataSourceInstance) getTrendData(
	ctx context.Context,
	refID string,
	timeRange backend.TimeRange,
	query *AlcQuery,
) (*backend.DataResponse, error) {

	logger := backend.Logger.FromContext(ctx)

	logger.Info(
		"Fetching trend",
		"metric", query.Metric,
		"from", timeRange.From,
		"to", timeRange.To,
	)

	atomic.AddUint64(&ds.trendDsRequests, 1)
	atomic.AddUint64(&ds.trendSoapRequests, 1)

	x := make([]time.Time, 0)
	y := make([]float64, 0)

	trendData := ds.service.Trend.GetTrendData(
		query.Metric,
		timeRange.From,
		timeRange.To,
	)

	for point := range trendData {
		t := point.Data.Time
		v := point.Data.Value
		x = append(x, t)
		y = append(y, v)
	}

	logger.Info(
		"Trend fetch complete",
		"metric", query.Metric,
		"points", len(x),
	)

	frame := data.NewFrame("response")
	frame.Fields = append(
		frame.Fields,
		data.NewField("time", nil, x),
		data.NewField("value", nil, y),
	)

	response := backend.DataResponse{}
	response.Frames = append(response.Frames, frame)

	return &response, nil
}
