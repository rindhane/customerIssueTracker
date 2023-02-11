package azureOTP

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type bodyRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

func initiateClient() *http.Client {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // ref : https://stackoverflow.com/questions/23297520/how-can-i-make-the-go-http-client-not-follow-redirects-automatically
		},
	}
	return client
}

type EmailRequest struct {
	ServiceUrl string
	Email      string
	Message    string
}

func (reqEmail *EmailRequest) getRequestComponent() (string, string, string) {
	return reqEmail.ServiceUrl, reqEmail.Email, reqEmail.Message
}

func SendOTPRequestAzure(er *EmailRequest) (string, error) {
	fmt.Println(er)
	azureUrl, email, otp := er.getRequestComponent()
	client := initiateClient()
	//ref :https://stackoverflow.com/questions/24493116/how-to-send-a-post-request-in-go
	temp, _ := json.Marshal(bodyRequest{Email: email, Otp: otp})
	req, err := http.NewRequest("POST", azureUrl, strings.NewReader(string(temp)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data), email)
	return string(data), nil
}
