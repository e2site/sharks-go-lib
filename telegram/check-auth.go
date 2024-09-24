package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/e2site/sharks-go-lib/log"
	"net/url"
	"sort"
)

type TelegramAuth struct {
	initData string
	token    string

	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
	AuthDate     string
	Hash         string
}

func NewTelegramAuth(initData string, token string) *TelegramAuth {
	telegramAuth := &TelegramAuth{initData: initData, token: token}

	params, _ := url.ParseQuery(telegramAuth.initData)
	for k, v := range params {
		if k == "hash" {
			telegramAuth.Hash = v[0]
			continue
		}
		if k == "user" {
			err := json.Unmarshal([]byte(v[0]), telegramAuth)
			if err != nil {
				log.Log(err)
			}
		}

		if k == "auth_date" {
			telegramAuth.AuthDate = v[0]
		}
	}

	return telegramAuth
}

func (telegramAuth *TelegramAuth) CheckAuth() bool {
	params, _ := url.ParseQuery(telegramAuth.initData)
	strs := []string{}
	for k, v := range params {
		if k == "hash" {
			telegramAuth.Hash = v[0]
			continue
		}

		strs = append(strs, k+"="+v[0])
	}
	sort.Strings(strs)
	var imploded = ""
	for _, s := range strs {
		if imploded != "" {
			imploded += "\n"
		}
		imploded += s
	}
	hKey := hmac.New(sha256.New, []byte("WebAppData"))
	hKey.Write([]byte(telegramAuth.token))
	secKey := hKey.Sum(nil)

	dKey := hmac.New(sha256.New, []byte(secKey))
	dKey.Write([]byte(imploded))
	dataKey := dKey.Sum(nil)
	ss := hex.EncodeToString(dataKey)

	return ss == telegramAuth.Hash
}
