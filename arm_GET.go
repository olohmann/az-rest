package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ArmGet(rawUrl string, apiVersion string, jmespathQuery string) {
	azureConfiguration := GetAzureConfiguration()

	url := ArmUrl(azureConfiguration, rawUrl, apiVersion)

	log.Debug("GET ", url)

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

	if resp.StatusCode > 399 {
		log.Fatal("HTTP StatusCode ", resp.StatusCode, "\n", jsonPrettyPrint(respBody))
	} else {
		fmt.Print(applyJmespathToJson(respBody, jmespathQuery))
	}
}
