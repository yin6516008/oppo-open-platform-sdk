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
	"reflect"
	"sort"
	"strconv"
	"strings"
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
		client:       &http.Client{},
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
		client:       &http.Client{},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "oop-openapi-cn.heytapmobi.com",
		},
	}, nil
}

func (c *OppoClient) NewBaseParams() map[string]string {
	if c.token == "" || c.tokenExpireIn == 0 || time.Now().Unix() >= c.tokenExpireIn {
		err := refreshToken(c)
		if err != nil {
			panic(err)
		}
	}

	p := map[string]string{
		"access_token": c.token,
		"timestamp":    strconv.FormatInt(time.Now().Unix(), 10),
	}

	return p

}

func (c *OppoClient) Get(path string, params map[string]string, v interface{}) (*http.Response, error) {
	c.baseURL.Path = path

	query := ParamsToSortQuery(params)
	c.baseURL.RawQuery = query

	resp, err := c.client.Get(c.baseURL.String())
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

func (c *OppoClient) Post(path string, body map[string]string, v interface{}) (*http.Response, error) {
	c.baseURL.Path = path

	p := url.Values{}

	for k, v := range body {
		p.Set(k, v)
	}

	resp, err := c.client.PostForm(c.baseURL.String(), p)
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
	path := "/developer/v1/token"
	c.baseURL.Path = path
	c.baseURL.RawQuery = fmt.Sprintf("client_id=%s&client_secret=%s", c.clientId, c.clientSecret)

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

func HandleParams(params interface{}, p map[string]string, clientSecret string) (map[string]string, error) {
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var mi map[string]interface{}
	err = json.Unmarshal(paramsByte, &mi)
	if err != nil {
		return nil, err
	}

	for k, v := range mi {
		vt := reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Map, reflect.Array, reflect.Slice:
			value, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			p[k] = string(value)
		default:
			p[k] = v.(string)
		}
	}

	sign := Signature(p, clientSecret)
	p["api_sign"] = sign

	return p, err
}

func ParamsToSortQuery(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultList []string
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := params[key]
		resultList = append(resultList, fmt.Sprintf("%s=%s", key, value))
	}

	result := strings.Join(resultList, "&")
	return result
}

func Signature(params map[string]string, clientSecret string) string {
	query := ParamsToSortQuery(params)
	sign := HmacSha256(query, clientSecret)
	return sign
}

// HmacSha256 加密
func HmacSha256(stringToSign string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
