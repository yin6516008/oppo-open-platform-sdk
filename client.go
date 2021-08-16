package oppo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"
)

type OppoClient struct {
	baseURL       *url.URL
	clientId      string
	clientSecret  string
	client        *http.Client
	token         string
	tokenExpireIn int64
}

func NewOppoClient(clientId, clientSecret string) *OppoClient {
	return &OppoClient{
		clientId:     clientId,
		clientSecret: clientSecret,
		client:       &http.Client{Timeout: 15},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "oop-openapi-cn.heytapmobi.com",
		},
	}
}

func NewOppoClientWithEnv() (*OppoClient, error) {
	clientId := os.Getenv("OPPO_CLIENT_ID")
	clientSecret := os.Getenv("OPPO_CLIENT_SECRET")

	if clientId == "" {
		return nil, fmt.Errorf("OPPO_CLIENT_ID is empty")
	}

	if clientSecret == "" {
		return nil, fmt.Errorf("OPPO_CLIENT_SECRET is empty")
	}

	return &OppoClient{
		clientId:     clientId,
		clientSecret: clientSecret,
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "oop-openapi-cn.heytapmobi.com",
		},
	}, nil
}

func (c *OppoClient) NewRequest(method, path string, params interface{}) (*http.Request, error) {
	if c.token == "" || c.tokenExpireIn == 0 || time.Now().Unix() >= c.tokenExpireIn {
		err := refreshToken(c)
		if err != nil {
			panic(err)
		}
	}

	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	c.baseURL.Path = unescaped

	now := strconv.FormatInt(time.Now().Unix(), 10)
	query, err := ParseParamsToSortQeury(params, c.token, c.clientSecret, now)
	if err != nil {
		return nil, err
	}
	c.baseURL.RawQuery = query

	fmt.Println(c.baseURL.String())
	req, err := http.NewRequest(method, c.baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/jsonapplication/x-www-form-urlencoded")
	req.Header.Set("Charset", "UTF-8")
	return req, err
}

func (c *OppoClient) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("the oppo client received an unhealthy status code: %d, message: %s", resp.StatusCode, resp.Status)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func refreshToken(c *OppoClient) error {
	params := &RefreshTokenParams{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
	}

	paramsByte, err := json.Marshal(params)
	if err != nil {
		return err
	}

	var paramsMap map[string]string
	err = json.Unmarshal(paramsByte, &paramsMap)
	if err != nil {
		return err
	}

	query := MapToSortQeury(paramsMap)
	if err != nil {
		return err
	}

	path := "/developer/v1/token"
	c.baseURL.Path = path
	c.baseURL.RawQuery = query

	resp, err := http.Get(c.baseURL.String())
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var refreshTokenRes *RefreshTokenRes
	err = json.Unmarshal(content, &refreshTokenRes)
	if err != nil {
		return err
	}

	c.token = refreshTokenRes.Data.AccessToken
	c.tokenExpireIn = refreshTokenRes.Data.ExpireIn

	return err
}

func ParseParamsToSortQeury(params interface{}, token, clientSecret, now string) (string, error) {
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	var paramsMap map[string]string
	err = json.Unmarshal(paramsByte, &paramsMap)
	if err != nil {
		return "", err
	}
	paramsMap["access_token"] = token
	paramsMap["timestamp"] = now

	p := MapToSortQeury(paramsMap)
	sign := HmacSha256(p, clientSecret)
	paramsMap["api_sign"] = sign
	result := MapToSortQeury(paramsMap)

	return result, err

}

func MapToSortQeury(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	p := url.Values{}
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := params[key]
		p.Add(key, value)
	}

	return p.Encode()
}

// HmacSha256 加密
func HmacSha256(stringToSign string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
