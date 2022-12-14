package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CreateExec(opts types.CreateExecRequest) (string, error) {

	url := c.url + "/v1/exec"

	body, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.CreateExecResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil

}

func (c *Client) GetExecList() ([]types.Exec, error) {

	url := c.url + "/v1/exec"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []types.Exec{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.GetExecListResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return []types.Exec{}, err
	}

	if response.Status != "success" {
		return []types.Exec{}, errors.New("unexpected error")
	}
	return response.Data, nil

}
