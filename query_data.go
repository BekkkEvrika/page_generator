package page_generator

import "net/url"

type QueryParams struct {
	Claims MapClaims
	QData  url.Values
	Token  string
	Limit  int
	Offset int
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

func (c *QueryParams) GetPermissions() []string {
	ps := c.Claims["permissions"].([]interface{})
	var ms []string
	for _, val := range ps {
		ms = append(ms, val.(string))
	}
	return ms
}
func (c *QueryParams) ExistsAccess(access []string) bool {
	for _, val := range c.GetPermissions() {
		for _, acc := range access {
			if val == acc {
				return true
			}
		}
	}
	return false
}
