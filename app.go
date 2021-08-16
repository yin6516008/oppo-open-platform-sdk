package oppo

import (
	"net/http"
)

func (c *OppoClient) GetAppDetail(query *GetAppDetailParams) (*GetAppDetailRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, "/resource/v1/app/info", query)
	if err != nil {
		return nil, nil, err
	}

	var getAppDetailRes *GetAppDetailRes
	resp, err := c.Do(req, &getAppDetailRes)
	if err != nil {
		return nil, resp, err
	}
	return getAppDetailRes, resp, err
}
