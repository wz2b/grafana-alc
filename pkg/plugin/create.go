package plugin

import (
	"context"
	"encoding/json"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/patrickmn/go-cache"
	alcsoap "github.com/wz2b/webctrl-soap-go"
)

// NewDataSourceInstance
//
// This function is a factory that creates new datasource instances.  It will get called
// any time a grafana datasource instance for this plugin gets created or whenever its
// settings are changed.
//
// Since any change to any setting creates a new instance in this manner, this means that
// the old cache will be thrown away and new, empty caches will be created.
func NewDataSourceInstance(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	backend.Logger.Info("Creating a new ALC grafana datasource instance")

	/*
	 * Pull our custom settings out of JSONData
	 */
	config := CustomInstanceSettings{}
	json.Unmarshal([]byte(settings.JSONData), &config)
	duration, err := time.ParseDuration(config.MinPvCacheTimeStr)
	if err != nil {
		duration = 1 * time.Minute
	}
	username := config.SoapUser
	url := config.SoapUrl

	secrets := settings.DecryptedSecureJSONData
	password := secrets["soapPassword"]
	soap := alcsoap.NewSoapService(url, username, password)

	backend.Logger.Info("Successfully created a new ALC grafana datasource instance")
	return &AlcGrafanaDataSourceInstance{
		pvCache:       cache.New(duration, 1*time.Hour),
		locationCache: cache.New(12*time.Hour, 15*time.Minute),

		pvCacheHits:    0,
		pvSoapRequests: 0,

		treeCacheHits:    0,
		treeSoapRequests: 0,

		trendDsRequests:   0,
		trendSoapRequests: 0,

		service: soap,
	}, nil
}

/*
 * this creates an instance of the plugin - essentially our constructor.  Note
 * that only one plugin instance is created, even if the user defines multiple
 * datasource instances using that same plugin.
 */
//func CreateAlcDataSource(instMgr instancemgmt.InstanceManager) *AlcDatasource {
//	return &AlcDatasource{im: instMgr}
//}
