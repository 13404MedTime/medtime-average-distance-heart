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
