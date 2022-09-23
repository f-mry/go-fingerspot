package fingerspot

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAttLog(request GetAttLogRequest) (GetAttLogResponse, error) {
	response := GetAttLogResponse{}
	data, err := json.Marshal(&request)
	req := bytes.NewBuffer(data)
	if err != nil {
		return response, err
	}
	err = c.Call(GetAttLogURL, req, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) GetUserInfo(request GetUserInfoRequest) (GetUserInfoResponse, error) {
	response := GetUserInfoResponse{}
	data, err := json.Marshal(&request)
	req := bytes.NewBuffer(data)
	if err != nil {
		return response, err
	}
	err = c.Call(GetUserInfoURL, req, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) SetUserInfo(request SetUserInfoRequest) (SetUserInfoResponse, error) {
	response := SetUserInfoResponse{}
	data, err := json.Marshal(&request)
	req := bytes.NewBuffer(data)
	if err != nil {
		return response, err
	}
	err = c.Call(SetUserInfoURL, req, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
