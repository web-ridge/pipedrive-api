package pipedrive

import (
	"fmt"
	"net/url"
)

func LimitOpts(limit int) url.Values {
	v := url.Values{}
	v.Set("limit", fmt.Sprintf("%v", limit))
	return v
}
