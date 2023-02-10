package main

import (
	"errors"
	"fmt"
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
	fmt.Println(email, otp)
	return nil
}
func dispatchLoginPassword(email string) error {
	fmt.Println(createRandomString(6))
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

func validateOTP(email string, otp string) error {
	otpStored, ok := bufferedMap[email]
	if ok {
		if otpStored == otp {
			go dispatchLoginPassword(email)
			delete(bufferedMap, email)
			return nil
		}
		return errors.New("invalid OTP")
	}
	return errors.New("invalid email")
}
