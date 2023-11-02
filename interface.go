// Copyright 2023 cputil Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudplatform

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	mu             = &sync.Mutex{}
	accessTokenMap = map[string]string{}
	cacheCtx       = map[string]*Context{}
	cacheCtxID     = map[uint]*Context{}
	refreshCh      = make(chan struct{}, 1)
	logger         = log.New(os.Stdout, "[cloudplatform] ", log.LstdFlags)
)

func AddCtx(ctx ...*Context) {
	for _, c := range ctx {
		c.refreshToken()
		cacheCtx[c.BaseUrl] = c
		cacheCtxID[c.ID] = c
	}
}

func init() { go startRefreshToken() }

func GetCtxByUrl(baseUrl string) *Context {
	mu.Lock()
	defer mu.Unlock()
	for k, v := range cacheCtx {
		if k == baseUrl {
			return v
		}
	}
	return nil
}

func GetCtxByID(id uint) *Context {
	mu.Lock()
	defer mu.Unlock()
	for _, v := range cacheCtx {
		if v.ID == id {
			return v
		}
	}
	return nil
}

type (
	Context struct {
		ID       uint
		BaseUrl  string
		Uid      int
		Username string
		Password string

		accessToken string
	}
	Request interface {
		Url() string
		Method() string
		Header() http.Header
		Values() (values url.Values)
		Form() (form url.Values)
		Body() any
	}
	Response[T any] struct {
		StatusCode int    `json:"-"`
		Error      string `json:"error"`
		Result     T      `json:"-"`
	}
)

func startRefreshToken() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			refreshToken()
		}
	}
}

func refreshToken() {
	if !mu.TryLock() {
		return
	}
	defer mu.Unlock()
	for _, v := range cacheCtx {
		v.refreshToken()
	}
	return
}

func (ctx *Context) refreshToken() {
	http.DefaultClient.Timeout = time.Second * 10
	logger.Println("start refresh token")
	if resp, err := http.Post(ctx.BaseUrl+"/v1/login", "application/json", bytes.NewBufferString(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", ctx.Username, ctx.Password))); err != nil {
		logger.Println("get token err:", err)
	} else if resp.StatusCode != http.StatusCreated {
		logger.Println("get token not 201")
	} else {
		if readAll, err := io.ReadAll(resp.Body); err != nil {
			logger.Println("get token read err:", err)
		} else {
			ctx.accessToken = string(readAll[1 : len(readAll)-1])
			logger.Println("get token:", ctx.BaseUrl, ctx.accessToken)
		}
	}
	logger.Println("end refresh token")
	return
}

// Execute 接口请求入口
func Execute[R Request, RR any](ctx *Context, request R) (resp Response[RR], err error) {
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, Timeout: time.Second * 10}
	var (
		body    io.Reader
		req     *http.Request
		rawResp *http.Response
	)
	if b := request.Body(); b != nil {
		typeOf := reflect.TypeOf(b)
		kind := typeOf.Kind()
		switch kind {
		case reflect.String:
			body = bytes.NewBufferString(b.(string))
			logger.Println("execute api body:", b.(string))
		case reflect.Map, reflect.Ptr, reflect.Struct:
			bs, _ := json.Marshal(b)
			body = bytes.NewReader(bs)
			logger.Println("execute api body:", string(bs))
		case reflect.Slice, reflect.Array:
			if typeOf.Elem().Kind() == reflect.Uint {
				body = bytes.NewReader(b.([]byte))
				logger.Println("execute api body:", string(b.([]byte)))
			}
		}
	}
	logger.Println("execute api url:", request.Url())
	logger.Println("execute req method:", request.Method())
	reqUrl := ctx.BaseUrl + request.Url()
	logger.Println("execute req url:", reqUrl)
	if v := request.Values(); v != nil {
		if vv := v.Encode(); len(vv) > 0 {
			if strings.LastIndexByte(reqUrl, '?') == -1 {
				reqUrl += "?" + vv
			} else {
				reqUrl += "&" + vv
			}
		}
	}
	logger.Println("execute req url with values:", reqUrl)
	req, err = http.NewRequest(request.Method(), reqUrl, body)
	if err != nil {
		return
	}
	if h := request.Header(); h != nil {
		logger.Println("execute header:", h)
		req.Header = h
	}
	if h := request.Form(); h != nil {
		logger.Println("execute form:", h)
		req.PostForm = h
	}
	req.Header.Set("Access-Token", ctx.accessToken)
	if req.Method != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
	if rawResp, err = client.Do(req); err != nil {
		logger.Println("execute api err:", err)
		return
	}
	readAll, err := io.ReadAll(rawResp.Body)
	resp.StatusCode = rawResp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed:
		if len(readAll) > 0 {
			err = json.Unmarshal(readAll, &resp.Result)
		}
	default:
		if len(readAll) > 0 {
			err = json.Unmarshal(readAll, &resp)
		}
	}
	if err != nil {
		logger.Println("execute err:", err)
	}
	return
}
