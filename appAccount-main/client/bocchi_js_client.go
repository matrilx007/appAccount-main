package client

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"time"
)

var bocchiAPIRegex = regexp.MustCompile(
	`https://id\.bocchi\.vip/api/list\?password=([a-zA-Z0-9]+)`,
)

func FetchBocchiPassword() (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://id.bocchi2b.top/assets/index-BDMWxbVK.js",
		nil,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "https://id.bocchi2b.top/")
	req.Header.Set("Origin", "https://id.bocchi2b.top")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	matches := bocchiAPIRegex.FindSubmatch(body)
	if len(matches) < 2 {
		return "", errors.New("failed to extract bocchi password")
	}

	return string(matches[1]), nil
}
