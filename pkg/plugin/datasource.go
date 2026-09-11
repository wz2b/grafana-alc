package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

//func (ds *AlcGrafanaDataSourceInstance) getInstance(ctx backend.PluginContext) (*AlcGrafanaDataSourceInstance, error) {
//	s, err := ds.im.Get(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return s.(*AlcGrafanaDataSourceInstance), nil
//}

func (ds *AlcGrafanaDataSourceInstance) QueryData(ctx context.Context,
	req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {

	var response = backend.NewQueryDataResponse()

	var err error
	err = nil

	for _, query := range req.Queries {
		backend.Logger.Info("Processing query", "query", query)
		switch query.QueryType {

		case "findChildren":
			var fq AlcFindChildrenQuery
			json.Unmarshal([]byte(query.JSON), &fq)

			backend.Logger.Info(fmt.Sprintf("Find children request: %v+", fq))

			r, err := ds.findChildren(ctx, query.RefID, &fq)

			if err == nil {
				response.Responses[query.RefID] = *r
			}

		case "Trend Data":
			alcQuery := &AlcQuery{}
			json.Unmarshal([]byte(query.JSON), alcQuery)

			r, err := ds.getTrendData(ctx, query.RefID, query.TimeRange, alcQuery)

			backend.Logger.Info("F1J SOAP request is finished")

			if err == nil {
				backend.Logger.Info("Returning result")
				response.Responses[query.RefID] = *r
			} else {
				backend.Logger.Error("Unable to get trend: ", err)
			}

		case "Present Value":
			alcQuery := &AlcQuery{}
			json.Unmarshal([]byte(query.JSON), alcQuery)
			r, err := ds.getPresentValue(ctx, query.RefID, alcQuery)

			if err == nil {
				response.Responses[query.RefID] = *r
			} else {
				backend.Logger.Error("Unable to get present value: ", err)
			}

		case "Cache Stats":
			alcQuery := &AlcQuery{}
			json.Unmarshal([]byte(query.JSON), alcQuery)
			r, err := ds.getCacheStats(ctx, query.RefID)
			if err == nil {
				response.Responses[query.RefID] = *r
			} else {
				backend.Logger.Error("Unable to get cache statistics ", err)
			}

		default:
			backend.Logger.Error("Unknown query type", "type", query.QueryType)
		}
	}

	return response, err
}
