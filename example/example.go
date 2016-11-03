//go:generate ./xmlrpcgen --file $GOFILE HelloService

package example

import (
	"github.com/beevik/etree"
)

/*
Service
*/
type HelloService struct {
	Config Config
}

/*
Search method
*/
func (h *HelloService) Search(query string, page int, isit bool) ([]string, error) {
	_ = etree.NewElement("asdf")
	return []string{}, nil
}
