package fingerspot

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	ApiURL            = "http://developer.fingerspot.io/"
	GetAttLogURL      = ApiURL + "api/get_attlog"
	GetUserInfoURL    = ApiURL + "api/get_userinfo"
	SetUserInfoURL    = ApiURL + "api/set_userinfo"
	DeleteUserInfoURL = ApiURL + "api/delete_userinfo"
	SetTimeURL        = ApiURL + "api/set_time"
	RegisterOnlineURL = ApiURL + "api/reg_online"
	RestartMesinURL   = ApiURL + "api/restart_device"
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(token string) Client {
	return Client{
		client: &http.Client{},
		token:  token,
	}

}

func (c *Client) Call(url string, body io.Reader, response interface{}) error {

	// build request
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var respError Err
	json.Unmarshal(respBody, &respError)
	if !respError.Success {
		return respError
	}
	err = json.Unmarshal(respBody, response)
	if err != nil {
		return err
	}
	return nil
}
