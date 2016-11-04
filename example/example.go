//go:generate ./xmlrpcgen --file $GOFILE --debug SearchService
package core

import _ "github.com/beevik/etree"

/*
SearchService xml rpc service for searching packages
*/
type SearchService struct {
}

type SearchResult struct {
	_pypi_ordering int
	version        string
	name           string
	summary        string
}

/*
search xml rpc method
*/
func (h *SearchService) search(query string) (result []SearchResult, err error) {
	result = make([]SearchResult, 0)
	return
}
