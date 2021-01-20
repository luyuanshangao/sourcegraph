package search

import (
	"errors"
	"sync"
	"sync/atomic"

	zoektquery "github.com/google/zoekt/query"
	zoektrpc "github.com/google/zoekt/rpc"
	"github.com/sourcegraph/sourcegraph/internal/search"
	"github.com/sourcegraph/sourcegraph/internal/search/backend"
	zoektutil "github.com/sourcegraph/sourcegraph/internal/search/zoekt"
)

var zoektOnce sync.Once
var endpointMap atomicEndpoints
var zoektClient backend.StreamSearcher

func getZoektClient(indexerEndpoints []string) backend.StreamSearcher {
	zoektOnce.Do(func() {
		dial := func(endpoint string) backend.StreamSearcher {
			return backend.NewMeteredSearcher(endpoint, &backend.StreamSearchAdapter{zoektrpc.Client(endpoint)})
		}
		zoektClient = backend.NewMeteredSearcher(
			"", // no hostname means its the aggregator
			&backend.HorizontalSearcher{
				Map:  &endpointMap,
				Dial: dial,
			},
		)
	})
	endpointMap.Set(indexerEndpoints)
	return zoektClient
}

func HandleFilePathPatterns(query *search.TextPatternInfo) (zoektquery.Q, error) {
	var and []zoektquery.Q

	// Zoekt uses regular expressions for file paths.
	// Unhandled cases: PathPatternsAreCaseSensitive and whitespace in file path patterns.
	for _, p := range query.IncludePatterns {
		q, err := zoektutil.FileRe(p, query.IsCaseSensitive)
		if err != nil {
			return nil, err
		}
		and = append(and, q)
	}
	if query.ExcludePattern != "" {
		q, err := zoektutil.FileRe(query.ExcludePattern, query.IsCaseSensitive)
		if err != nil {
			return nil, err
		}
		and = append(and, &zoektquery.Not{Child: q})
	}

	// For conditionals that happen on a repo we can use type:repo queries. eg
	// (type:repo file:foo) (type:repo file:bar) will match all repos which
	// contain a filename matching "foo" and a filename matchinb "bar".
	//
	// Note: (type:repo file:foo file:bar) will only find repos with a
	// filename containing both "foo" and "bar".
	for _, p := range query.FilePatternsReposMustInclude {
		q, err := zoektutil.FileRe(p, query.IsCaseSensitive)
		if err != nil {
			return nil, err
		}
		and = append(and, &zoektquery.Type{Type: zoektquery.TypeRepo, Child: q})
	}
	for _, p := range query.FilePatternsReposMustExclude {
		q, err := zoektutil.FileRe(p, query.IsCaseSensitive)
		if err != nil {
			return nil, err
		}
		and = append(and, &zoektquery.Not{Child: &zoektquery.Type{Type: zoektquery.TypeRepo, Child: q}})
	}

	return zoektquery.NewAnd(and...), nil
}

// atomicEndpoints allows us to update the endpoints used by our zoekt client.
type atomicEndpoints struct {
	endpoints atomic.Value
}

func (a *atomicEndpoints) Endpoints() (map[string]struct{}, error) {
	eps := a.endpoints.Load()
	if eps == nil {
		return nil, errors.New("endpoints have not been set")
	}
	return eps.(map[string]struct{}), nil
}

func (a *atomicEndpoints) Set(endpoints []string) {
	if !a.needsUpdate(endpoints) {
		return
	}

	eps := make(map[string]struct{}, len(endpoints))
	for _, addr := range endpoints {
		eps[addr] = struct{}{}
	}
	a.endpoints.Store(eps)
}

func (a *atomicEndpoints) needsUpdate(endpoints []string) bool {
	old, err := a.Endpoints()
	if err != nil {
		return true
	}
	if len(old) != len(endpoints) {
		return true
	}

	for _, addr := range endpoints {
		if _, ok := old[addr]; !ok {
			return true
		}
	}

	return false
}
