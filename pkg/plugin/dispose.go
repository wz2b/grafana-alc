package plugin

import "github.com/grafana/grafana-plugin-sdk-go/backend"

func (a *AlcGrafanaDataSourceInstance) Dispose() {
	backend.Logger.Info("ALC Data Source is cleaning up")
}
