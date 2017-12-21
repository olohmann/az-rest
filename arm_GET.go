package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ArmGet(rawUrl string, apiVersion string, headers string) {
	azureConfiguration := GetAzureConfiguration()

	url := ArmUrl(azureConfiguration, rawUrl, apiVersion)

	log.Debug("GET ", url)
	log.Debug("GET HEADER ", headers)

	req, reqErr := http.NewRequest("GET", url, nil)
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
