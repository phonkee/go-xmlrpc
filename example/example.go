//go:generate ./xmlrpcgen --file $GOFILE --debug SearchService
package core

import _ "github.com/beevik/etree"

/*
SearchService xml rpc service for searching packages
*/
type SearchService struct {
}


type Result struct {
	Hello string
}

/*
search xml rpc method
*/
func (h *SearchService) search(what Result) (string, error) {
	return "", nil
}
