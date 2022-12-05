//  fetcher 爬虫相关的操作
package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetch 读取网页内容
func Fetch(url string, header map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %v", err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client do error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %v", err)
	}

	var respBody []byte
	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return respBody, fmt.Errorf("respBody read error: %v", err)
	}
	return respBody, nil
}
