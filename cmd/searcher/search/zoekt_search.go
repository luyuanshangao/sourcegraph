package search

import (
	"sync"

	zoektquery "github.com/google/zoekt/query"

	"github.com/sourcegraph/sourcegraph/internal/endpoint"
	"github.com/sourcegraph/sourcegraph/internal/search"
	zoektutil "github.com/sourcegraph/sourcegraph/internal/search/zoekt"

	zoektrpc "github.com/google/zoekt/rpc"
	"github.com/sourcegraph/sourcegraph/internal/search/backend"
)

var zoektOnce sync.Once
var endpointMap *endpoint.Map
var zoektClient backend.StreamSearcher

func getZoektClient(indexerEndpoints []string) backend.StreamSearcher {
	zoektOnce.Do(func() {
		dial := func(endpoint string) backend.StreamSearcher {
			return backend.NewMeteredSearcher(endpoint, &backend.StreamSearchAdapter{zoektrpc.Client(endpoint)})
		}

		zoektClient = backend.NewMeteredSearcher(
			"", // no hostname means its the aggregator
			&backend.HorizontalSearcher{
				Map:  endpointMap,
				Dial: dial,
			},
		)
	})
	endpointMap = endpoint.Static(indexerEndpoints...)
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
