package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func ArmPost(rawUrl string, apiVersion string, headers string, reqBody string) {
	azureConfiguration := GetAzureConfiguration()

	url := ArmUrl(azureConfiguration, rawUrl, apiVersion)

	log.Debug("POST ", url)
	log.Debug("POST HEADER ", headers)
	log.Debug("POST BODY ", reqBody)
	req, reqErr := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	req.Header.Set("Authorization", "BEARER "+azureConfiguration.AccessToken)

	resp, httpGetErr := http.DefaultClient.Do(req)
	if httpGetErr != nil {
		log.Fatal(reqErr)
	}

	defer resp.Body.Close()
	respBody, readErr := ioutil.ReadAll(resp.Body)
	if httpGetErr != nil {
		log.Fatal(readErr)
	}

	var prettyJson bytes.Buffer
	indentErr := json.Indent(&prettyJson, respBody, "", "  ")
	if indentErr != nil {
		log.Fatal(indentErr)
	}

	if resp.StatusCode > 399 {
		log.Fatal("HTTP StatusCode ", resp.StatusCode, "\n", string(prettyJson.Bytes()))
	} else {
		fmt.Print(string(prettyJson.Bytes()))
	}
}
