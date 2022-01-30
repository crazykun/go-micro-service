package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 注册服务
func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServiceUrl, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Fail to retgister service: %s", res.Status)
	}
	defer res.Body.Close()
	return nil
}

// 取消注册服务
func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServiceUrl, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Fail to shutdown service: %s", res.StatusCode)
	}
	// defer req.Body.Close()
	return nil
}
