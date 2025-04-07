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
