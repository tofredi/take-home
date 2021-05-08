package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestParseUrl(t *testing.T) {
	got, _ := parceUrl("http://www.google.com")
	expected := "http://www.google.com"
	if got != expected {
		t.Errorf("parceUrl(www.google.com) = %s; want %s", got, expected)
	}
}

func TestParseUrlWithoutScheme(t *testing.T) {
	got, _ := parceUrl("www.google.com")
	expected := "http://www.google.com"
	if got != expected {
		t.Errorf("parceUrl(www.google.com) = %s; want %s", got, expected)
	}
}

func TestParseUrlInvalidAddr(t *testing.T) {
	got, err := parceUrl(".%()=asd")
	if err == nil && got != "" {
		t.Errorf("url parsing did not fail")
	}
}

func TestSendRequest(t *testing.T) {
	body, err := sendRequest("http://www.google.com")
	if err != nil && body == nil {
		t.Errorf("Error sending request")
	}
}

func TestSendRequestWithoutScheme(t *testing.T) {
	body, err := sendRequest("www.google.com")
	if err == nil && body != nil {
		t.Errorf("Request did not fail")
	}
}

func TestSendRequestInvalidAddr(t *testing.T) {
	body, err := sendRequest(".%()=asd")
	if err == nil && body != nil {
		t.Errorf("Request did not fail")
	}
}

func TestSendGetResult(t *testing.T) {
	got, err := getResult("google.com")
	if err != nil {
		t.Errorf("Error getting result")
	}
	result := strings.Split(got, " ")
	addr := result[0]
	md5Hash := result[1]
	if addr != "http://google.com" {
		t.Errorf("Incorrect address")
	}
	match, _ := regexp.MatchString("^[a-fA-F0-9]{32}$", md5Hash)
	if !match {
		t.Errorf("Bad md5 hash")
	}
}

func TestSendGetResultWithScheme(t *testing.T) {
	got, err := getResult("http://google.com")
	if err != nil {
		t.Errorf("Error getting result")
	}
	result := strings.Split(got, " ")
	addr, md5Hash := result[0], result[1]
	if addr != "http://google.com" {
		t.Errorf("Incorrect address")
	}
	match, _ := regexp.MatchString("^[a-fA-F0-9]{32}$", md5Hash)
	if !match {
		t.Errorf("Bad md5 hash")
	}
}

func TestSendGetResultInvalidAddress(t *testing.T) {
	got, err := getResult(".%()=asd")
	if err == nil && got != "" {
		t.Errorf("Get result did not fail")
	}
}
