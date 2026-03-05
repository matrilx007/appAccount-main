package service

import (
	"app/client"
	"app/model"
)

func GetSharedAccounts() ([]model.SharedAccount, error) {
	password, err := client.FetchBocchiPassword()
	if err != nil {
		return nil, err
	}

	accounts, err := client.FetchBocchiAccounts(password)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
