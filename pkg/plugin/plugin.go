package plugin

import (
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
)

// Compile-time checks to make sure we have implemented the necessary interfaces
var (
	_ backend.QueryDataHandler   = (*AlcGrafanaDataSourceInstance)(nil)
	_ backend.CheckHealthHandler = (*AlcGrafanaDataSourceInstance)(nil)
	//_ backend.StreamHandler         = (*AlcDatasource)(nil)
	_ instancemgmt.InstanceDisposer = (*AlcGrafanaDataSourceInstance)(nil)
)
