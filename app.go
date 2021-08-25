package oppo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *OppoClient) GetAppDetail(query *GetAppDetailParams) (*GetAppDetailRes, *http.Response, error) {
	p := c.NewBaseParams()

	data, err := HandleParams(query, p, c.clientSecret)
	if err != nil {
		return nil, nil, err
	}

	var getAppDetailRes *GetAppDetailRes
	resp, err := c.Get("/resource/v1/app/info", data, &getAppDetailRes)

	return getAppDetailRes, resp, err
}

func (c *OppoClient) PublishVersion(params *PublishVersionParams) (*PublishVersionRes, *http.Response, error) {
	p := c.NewBaseParams()

	data, err := HandleParams(params, p, c.clientSecret)
	if err != nil {
		return nil, nil, err
	}

	var res *PublishVersionRes
	resp, err := c.Post("/resource/v1/app/upd", data, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

func (c *OppoClient) GetUploadFileConfig() (*GetUploadFileConfigRes, *http.Response, error) {
	p := c.NewBaseParams()

	data, err := HandleParams(nil, p, c.clientSecret)
	if err != nil {
		return nil, nil, err
	}

	var res *GetUploadFileConfigRes
	resp, err := c.Get("/resource/v1/upload/get-upload-url", data, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

func (c *OppoClient) UploadFile(params *UploadFileParams) (*UploadFileRes, *http.Response, error) {

	file, err := os.Open(params.FilePath)
	if err != nil {
		return nil, nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(params.FilePath))
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, nil, err
	}

	writer.WriteField("type", params.Type)
	writer.WriteField("sign", params.Sign)

	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	resp, err := http.Post(params.UploadURL, writer.FormDataContentType(), body)
	if err != nil {
		return nil, resp, err
	}

	var res *UploadFileRes
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res, resp, err

}
