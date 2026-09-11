package plugin

import (
	"context"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	alcsoap "github.com/wz2b/webctrl-soap-go"
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

	records, err := ds.getAllF1JTrendData(
		ctx,
		query.Metric,
		timeRange.From,
		timeRange.To,
	)
	if err != nil {
		return nil, err
	}

	x := make([]time.Time, 0, len(records))
	y := make([]float64, 0, len(records))

	for _, record := range records {
		// F1J also returns things like TimeChange, Failure,
		// LogStatus, Event, etc. Those are not graphable values.
		if !IsDataSample(record.ValueType) {
			logger.Info("NON DATA VALUE IGNORED")
			continue
		}

		value, err := strconv.ParseFloat(record.RawValue, 64)
		if err != nil {
			logger.Warn(
				"Unable to parse trend value",
				"metric", query.Metric,
				"timestamp", record.Timestamp,
				"valueType", record.ValueType,
				"rawValue", record.RawValue,
				"error", err,
			)
			continue
		}

		x = append(x, record.Timestamp)
		y = append(y, value)
	}

	logger.Info(
		"Trend fetch complete",
		"metric", query.Metric,
		"records", len(records),
		"points", len(x),
	)

	frame := data.NewFrame("response")
	frame.RefID = refID

	frame.Fields = append(
		frame.Fields,
		data.NewField("time", nil, x),
		data.NewField("value", nil, y),
	)

	response := backend.DataResponse{}
	response.Frames = append(response.Frames, frame)

	return &response, nil
}

func IsDataSample(t alcsoap.F1JValueType) bool {
	return (t >= alcsoap.F1JValueTypeBoolean && t <= alcsoap.F1JValueTypeNull) ||
		t == alcsoap.F1JValueTypeDouble
}

func (ds *AlcGrafanaDataSourceInstance) getAllF1JTrendData(
	ctx context.Context,
	metric string,
	from time.Time,
	to time.Time,
) ([]alcsoap.F1JTrendRecord, error) {
	const pageSize = 10000
	logger := backend.Logger.FromContext(ctx)

	logger.Info(
		"F1J request time range",
		"from", from,
		"fromUTC", from.UTC(),
		"fromUnixMilli", from.UnixMilli(),
		"to", to,
		"toUTC", to.UTC(),
		"toUnixMilli", to.UnixMilli(),
	)

	var records []alcsoap.F1JTrendRecord
	cursor := from

	for {
		atomic.AddUint64(&ds.trendSoapRequests, 1)
		//logger.Info("NEXT CHUNK")

		page, err := ds.service.F1JTrend.GetF1JTrendData(
			metric,
			cursor,
			to,
			true,
			pageSize,
		)
		if err != nil {
			return nil, err
		}

		records = append(records, page...)

		if len(page) < pageSize {
			break
		}

		last := page[len(page)-1].Timestamp

		// WebCTRL trend timestamps are millisecond resolution.
		next := last.Add(time.Millisecond)

		// Defensive guard against getting stuck on the same page.
		if !next.After(cursor) {
			break
		}

		cursor = next
	}

	return records, nil
}
