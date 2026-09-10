package plugin

import (
	"github.com/patrickmn/go-cache"
	alcsoap "github.com/wz2b/webctrl-soap-go"
)

//type AlcDatasource struct {
//	im instancemgmt.InstanceManager
//}

/*
 * Instance variables.  Instance here means an instance of the datasource
 */
type AlcGrafanaDataSourceInstance struct {
	pvCache       *cache.Cache
	locationCache *cache.Cache

	pvCacheHits    uint64
	pvSoapRequests uint64

	treeCacheHits    uint64
	treeSoapRequests uint64

	trendDsRequests   uint64
	trendSoapRequests uint64

	service *alcsoap.SoapService
}

type CustomInstanceSettings struct {
	SoapUser          string `json:"soapUser"'`
	SoapUrl           string `json:"soapUrl"'`
	MinPvCacheTimeStr string `json:"minPvCache"'`
}

type AlcReqestContext struct {
	url      string
	username string
	password string
	soap     *alcsoap.SoapService
}
