package main

import (
	"log"
	"net/url"
	"strings"
)

func ArmUrl(azureConfiguration AzureConfiguration, rawUrl string, apiVersion string) string {
	var azUrl *url.URL
	var err error

	if strings.HasPrefix(rawUrl, "http") {
		azUrl, err = url.Parse(rawUrl + "?api-version=" + apiVersion)
	} else {
		armEndpoint := strings.TrimRight(azureConfiguration.EndpointResourceManager, "/")
		rawUrl = strings.TrimLeft(rawUrl, "/")
		azUrl, err = url.Parse(armEndpoint + "/" + rawUrl + "?api-version=" + apiVersion)
	}

	if err != nil {
		log.Fatal(err)
	}

	return azUrl.String()
}
