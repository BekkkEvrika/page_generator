package page_generator

import "net/url"

type QueryParams struct {
	Claims MapClaims
	QData  url.Values
	Token  string
}

func (c *QueryParams) GetQuery(key string) (string, bool) {
	if values, ok := c.GetQueryArray(key); ok {
		return values[0], ok
	}
	return "", false
}

func (c *QueryParams) GetQueryArray(key string) (values []string, ok bool) {
	values, ok = c.QData[key]
	return
}

func (c *QueryParams) Query(key string) (value string) {
	value, _ = c.GetQuery(key)
	return
}
