/*
Xpath function helpers
*/
package xmlrpc

import (
	"strconv"

	"github.com/beevik/etree"
	"strings"
)

/*
XPathValueGetInt Returns int from value
*/
func XPathValueGetInt(element *etree.Element, name string) (result int, err error) {

	var tmp *etree.Element

	if tmp = element.FindElement("int"); tmp == nil {
		tmp = element.FindElement("i4")
	}

	if tmp == nil {
		err = Errorf(400, "Not found %v")
		return
	}

	result, err = strconv.Atoi(tmp.Text())

	return
}

/*
XPathValueGetInt64 Returns int64 from value
*/
func XPathValueGetInt64(element *etree.Element, name string) (result int64, err error) {
	var i int

	if i, err = XPathValueGetInt(element, name); err != nil {
		return
	}

	result = int64(i)

	return
}

/*
XPathValueGetInt64 Returns int64 from value
*/
func XPathValueGetInt32(element *etree.Element, name string) (result int32, err error) {
	var i int

	if i, err = XPathValueGetInt(element, name); err != nil {
		return
	}

	result = int32(i)

	return
}

/*
XPathValueGetString Returns bool from value
*/
func XPathValueGetString(element *etree.Element, name string) (result string, err error) {
	var tmp *etree.Element

	if tmp = element.FindElement("string"); tmp == nil {
		err = Errorf(400, "not found %v", name)
	}

	result = tmp.Text()

	return
}

/*
XPathValueGetBool Returns bool from value
*/
func XPathValueGetBool(element *etree.Element, name string) (result bool, err error) {
	var tmp *etree.Element

	if tmp = element.FindElement("boolean"); tmp == nil {
		err = Errorf(400, "not found %v", name)
	}

	result = strings.TrimSpace(tmp.Text()) == "1"

	return
}
