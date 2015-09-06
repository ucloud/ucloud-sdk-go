package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/xiaohui/goucloud/ucloud"
	"github.com/xiaohui/goucloud/ucloud/uclouderr"
	"github.com/xiaohui/goucloud/ucloud/utils"
)

type Service struct {
	Config      *ucloud.Config
	ServiceName string
	APIVersion  string

	BaseUrl    string
	HttpClient *http.Client
}

func (s *Service) DoRequest(action string, params interface{}, response interface{}) error {
	requestURL, err := s.RequestURL(action, params)
	if err != nil {
		return fmt.Errorf("build request url failed, error: %s", err)
	}

	httpReq, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return fmt.Errorf("new request url failed, error: %s", err)
	}

	httpResp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("do request url failed, error: %s", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return fmt.Errorf("do request url failed, error: %s", err)
	}

	statusCode := httpResp.StatusCode
	if statusCode >= 400 && statusCode <= 599 {

		uerr := uclouderr.UcloudError{}
		err = json.Unmarshal(body, &uerr)
		return &uclouderr.RequestFailed{
			UcloudError: uerr,
			StatusCode:  statusCode,
		}
	}

	err = json.Unmarshal(body, &response)
	fmt.Println("%+v", response)

	if err != nil {
		return fmt.Errorf("unmarshal url failed, error: %s", err)
	}

	return nil
}

// RequestURL is fully url of api request
func (s *Service) RequestURL(action string, params interface{}) (string, error) {
	if len(s.BaseUrl) == 0 {
		return "", errors.New("baseUrl is not set")
	}

	commonRequest := ucloud.CommonRequest{
		Action:    action,
		PublicKey: s.Config.Credentials.PublicKey,
		ProjectId: s.Config.ProjectID,
	}

	values := url.Values{}
	utils.ConvertParamsToValues(commonRequest, &values)
	utils.ConvertParamsToValues(params, &values)

	url, err := utils.UrlWithSignature(&values, s.BaseUrl, s.Config.Credentials.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("convert params error: %s", err)
	}

	return url, nil
}
