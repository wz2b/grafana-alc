package plugin

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/wz2b/webctrl-soap-go"
)

type AlcFindChildrenQuery struct {
	StartingNode string   `json:"gql"'`
	Types        []string `json:"types"'`
}

func (ds *AlcGrafanaDataSourceInstance) findChildren(ctx context.Context,
	refId string, parameters *AlcFindChildrenQuery) (*backend.DataResponse, error) {

	backend.Logger.Info("Fetching children",
		"path", parameters.StartingNode)

	var err error
	var nodes []alcsoap.GqlNode
	cache := ds.locationCache

	var filter func(node alcsoap.GqlNode) bool

	if parameters.Types == nil || len(parameters.Types) == 0 {
		filter = nil
	} else {
		filter = func(node alcsoap.GqlNode) bool {
			for _, t := range parameters.Types {
				if t == "ALL" {
					return true
				} else {
					return node.Type == t
				}
			}
			return false
		}
	}

	cachedValue, hit := cache.Get(parameters.StartingNode)
	if hit {
		backend.Logger.Info("Cache hit!")
		atomic.AddUint64(&ds.treeCacheHits, 1)

		nodes = cachedValue.([]alcsoap.GqlNode)
	} else {
		atomic.AddUint64(&ds.treeSoapRequests, 1)

		/*
		 * Fetch all nodes, so that the complete set of children ends up in the cache.
		 * If the user wants a filter, it will be supplied later.  This means an extra
		 * copy, but that downside is outweighed by the benefit of the cache.
		 */
		nodes, err = ds.service.Eval.GetChildren(parameters.StartingNode, nil)
		if err != nil {
			return nil, err
		}
		cache.Set(parameters.StartingNode, nodes, 1*time.Hour)
	}

	if filter != nil {
		n := 0
		for _, x := range nodes {
			if filter(x) {
				nodes[n] = x
				n++
			}
		}
		nodes = nodes[:n]
	}

	frame := data.NewFrame("Children",
		data.NewField("gql", nil, []string{}),
		data.NewField("displayName", nil, []string{}),
		data.NewField("type", nil, []string{}),
	)

	for _, node := range nodes {
		backend.Logger.Info("Child", "gql", node.ReferenceName)
		frame.AppendRow(
			node.ReferenceName,
			node.DisplayName,
			node.Type)
	}

	frames := data.Frames{frame}
	qr := &backend.DataResponse{
		Frames: frames,
		Error:  nil,
	}

	return qr, nil
}
