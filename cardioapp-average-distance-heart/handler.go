package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cast"
)

// Datas This is response struct from create
type Datas struct {
    Data struct {
        Data struct {
            Data map[string]interface{} `json:"data"`
        } `json:"data"`
    } `json:"data"`
}

// ClientApiResponse This is get single api response
type ClientApiResponse struct {
    Data ClientApiData `json:"data"`
}

type ClientApiData struct {
    Data ClientApiResp `json:"data"`
}

type ClientApiResp struct {
    Response map[string]interface{} `json:"response"`
}

type HttpRequest struct {
    Method  string      `json:"method"`
    Path    string      `json:"path"`
    Headers http.Header `json:"headers"`
    Params  url.Values  `json:"params"`
    Body    []byte      `json:"body"`
}

type AuthData struct {
    Type string                 `json:"type"`
    Data map[string]interface{} `json:"data"`
}

type NewRequestBody struct {
    RequestData HttpRequest            `json:"request_data"`
    Auth        AuthData               `json:"auth"`
    Data        map[string]interface{} `json:"data"`
}

type Request struct {
    Data map[string]interface{} `json:"data"`
}

type GetListClientApiResponse struct {
    Data GetListClientApiData `json:"data"`
}

type GetListClientApiData struct {
    Data GetListClientApiResp `json:"data"`
}

type GetListClientApiResp struct {
    Response []map[string]interface{} `json:"response"`
}

func DoRequest(url string, method string, body interface{}, appId string) ([]byte, error) {
    data, err := json.Marshal(&body)
    if err != nil {
        return nil, err
    }
    client := &http.Client{
        Timeout: time.Duration(5 * time.Second),
    }

    request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
    if err != nil {
        return nil, err
    }
    request.Header.Add("authorization", "API-KEY")
    request.Header.Add("X-API-KEY", appId)

    resp, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    respByte, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return respByte, nil
}
