package plugin

import (
	"context"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func (ds *AlcGrafanaDataSourceInstance) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	backend.Logger.Info("Checking grafana-alc backend service health")

	//
	// Do a health check by getting the top of the geo tree, which should
	// contain at least one node.  This will test that we have connectivity
	// and that we have a working username and password.  We really don't
	// care about the actual result at this point.
	//

	if ds.service == nil {
		backend.Logger.Error("Backend's service provider is nil, this should not be!!!")
	} else {
		backend.Logger.Info(fmt.Sprintf("Launching request, host=%s user=%s\n",
			ds.service.Eval.Endpoint, ds.service.Eval.User))
	}

	children, err := ds.service.Eval.GetChildren("/Trees/geographic", nil)

	backend.Logger.Info(fmt.Sprintf("Got a response\n"))

	if err == nil {
		num := len(children)
		backend.Logger.Info(fmt.Sprintf("Found %d root nodes", num))
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusOk,
			Message: "Health check succeeded",
		}, nil
	} else {
		backend.Logger.Warn("Health check failure: %s", err)
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: fmt.Sprintf("Failed check with error '%s'\n", err.Error()),
		}, err
	}

}
