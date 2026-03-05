package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"app/model"
)

type bocchiResponse struct {
	ID []struct {
		Country  string `json:"country"`
		Email    string `json:"email"`
		ID       string `json:"id"`
		Password string `json:"password"`
		Status   string `json:"status"`
		Time     string `json:"time"`
	} `json:"id"`
}

func FetchBocchiAccounts(password string) ([]model.SharedAccount, error) {
	url := fmt.Sprintf(
		"https://id.bocchi.vip/api/list?password=%s",
		password,
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Origin", "https://id.bocchi2b.top")
	req.Header.Set("Referer", "https://id.bocchi2b.top/")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result bocchiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	accounts := make([]model.SharedAccount, 0, len(result.ID))
	for _, item := range result.ID {
		accounts = append(accounts, model.SharedAccount{
			Email:    item.Email,
			Password: item.Password,
			Country:  item.Country,
			Status:   item.Status,
			Time:     item.Time,
			ID:       item.ID,
			Type:     "shadowrocket",
		})
	}

	return accounts, nil
}
