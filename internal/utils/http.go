package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	overTime = time.Minute
)

// Post post请求
func Post(url string, body io.Reader, header map[string]string, inOverTime ...time.Duration) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(url), body)
	if err != nil {
		return nil, err
	}

	content := make([]byte, 0)
	_, _ = body.Read(content)
	println("开始输出请求体")
	println(string(content))

	for k, v := range header {
		req.Header.Set(k, v)
	}
	timeOut := overTime
	if len(inOverTime) > 0 {
		timeOut = inOverTime[0]
	}
	cl := http.Client{
		Timeout: timeOut,
	}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
