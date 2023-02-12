package main

import (
	"azureComm/azureOTP"
	"context"
	"errors"
	"time"
)

var bufferedMap = make(map[string]string)

func initiateOTP(email string) error {
	if email != "" {
		err := dispatchOTP(email, generateOTP(email))
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("not a valid email")
}

func eliminateEmailOTPKey(store map[string]string, email string) {
	time.Sleep(5 * 60 * time.Second)
	delete(store, email)
}

func dispatchOTP(email string, otp string) error {
	var em azureOTP.EmailRequest
	em.Email = email
	em.Message = otp
	em.ServiceUrl = ENV_INPUTS.Secrets.OtpServiceEndpoint
	_, err := azureOTP.SendOTPRequestAzure(&em)
	if err != nil {
		return err
	}
	return nil
}
func dispatchLoginPassword(email string, password string) error {
	var em azureOTP.EmailRequest
	em.Email = email
	em.Message = password
	em.ServiceUrl = ENV_INPUTS.Secrets.OtpServiceEndpoint
	_, err := azureOTP.SendOTPRequestAzure(&em)
	if err != nil {
		return err
	}
	return nil
}

func generateOTP(email string) string {
	otp := randomGenerator()
	bufferedMap[email] = otp
	go eliminateEmailOTPKey(bufferedMap, email)
	return otp
}

func randomGenerator() string {
	return createRandomNumericalString(6)
}

func validateOTPNewAccount(ct *Controller, ctx context.Context, email string, otp string) error {
	otpStored, ok := bufferedMap[email]
	if ok {
		if otpStored == otp {
			newPassword := createRandomString(6)
			setPasswordToUser(ctx, ct, email, createHashString(newPassword))
			go dispatchLoginPassword(email, newPassword)
			delete(bufferedMap, email)
			return nil
		}
		return errors.New("invalid OTP")
	}
	return errors.New("invalid email")
}
