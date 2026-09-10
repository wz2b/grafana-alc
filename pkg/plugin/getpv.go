package plugin

import (
	"context"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (ds *AlcGrafanaDataSourceInstance) getPresentValue(ctx context.Context,
	refId string,
	query *AlcQuery) (*backend.DataResponse, error) {

	backend.Logger.Info("Get Present Value", "query", query.Metric)

	pvcache := ds.pvCache

	var strValue string
	var err error

	cachedValue, hit := pvcache.Get(query.Metric)
	if hit {
		backend.Logger.Info("Cache Hit")
		atomic.AddUint64(&ds.pvCacheHits, 1)

		strValue = cachedValue.(string)
	} else {
		backend.Logger.Info("Cache miss, fetching from SOAP")
		atomic.AddUint64(&ds.pvSoapRequests, 1)
		strValue, err = ds.service.Eval.GetValue(query.Metric)

		if err != nil {
			backend.Logger.Error("GetValue() failed", err)
			return nil, err
		}

		pvcache.Set(query.Metric, strValue, 0)
	}

	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		backend.Logger.Error("Unable to parse result as float")
		return nil, err
	}

	friendlyName := query.Metric
	if query.Alias != "" {
		friendlyName = query.Alias
	}

	//result := &datasource.QueryResult{
	//	RefId: refId,
	//	Series: []*datasource.TimeSeries{&datasource.TimeSeries{
	//		Name: friendlyName,
	//		Points: []*datasource.Point{&datasource.Point{
	//			Timestamp: time.Now().Unix() * 1000,
	//			Value:     value,
	//		}},
	//	}},
	//}

	frame := data.NewFrame("response")
	frame.Fields = append(frame.Fields,
		data.NewField("time", nil, []time.Time{time.Now()}),
		data.NewField(friendlyName, nil, []float64{value}),
	)

	response := backend.DataResponse{}
	response.Frames = append(response.Frames, frame)
	return &response, nil
}
