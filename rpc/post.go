package rpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type PostErrHandler func(statusCode int, body []byte, err error) error

func DefaultPostErrHandler(statusCode int, body []byte, err error) error {
	return fmt.Errorf("golkrpc: status code: %d, body: %s, err: %v", statusCode, string(body), err)
}

func Post[D any](ctx context.Context, client *Client, rpcMethod string, params ...any) (D, error) {
	var _params []any
	for _, param := range params {
		val := reflect.ValueOf(param)
		if val.Kind() == reflect.Ptr && val.IsNil() {
			continue
		}
		_params = append(_params, val.Interface())
	}
	reqData := NewReqData(rpcMethod, _params...)
	if client.Limiter != nil {
		client.Limiter.Wait(ctx)
	}
	reader, err := reqData.ToReader()
	if err != nil {
		return *new(D), client.PostErrHandler(0, nil, err)
	}
	req, err := http.NewRequest(http.MethodPost, client.URL, reader)
	if err != nil {
		return *new(D), client.PostErrHandler(0, nil, err)
	}
	req.Header = client.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return *new(D), client.PostErrHandler(0, nil, err)
	}
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	if statusCode != http.StatusOK {
		var body []byte
		if resp.Body != nil {
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				return *new(D), client.PostErrHandler(statusCode, nil, err)
			}
		}
		return *new(D), client.PostErrHandler(statusCode, body, nil)
	}
	respData := &Resp[D]{}
	if err := respData.FromReader(resp.Body); err != nil {
		if errors.Is(err, io.EOF) {
			return *new(D), client.PostErrHandler(statusCode, nil, err)
		}
		var body []byte
		var er error
		if resp.Body != nil {
			body, er = io.ReadAll(resp.Body)
			if er != nil {
				return *new(D), client.PostErrHandler(statusCode, nil, fmt.Errorf("%w, %w", err, er))
			}
		}
		return *new(D), client.PostErrHandler(statusCode, body, err)
	}
	return respData.Result, nil
}
