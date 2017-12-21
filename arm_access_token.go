package main

import (
	"bytes"
	"encoding/json"
	"github.com/SermoDigital/jose/jws"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

var cloudConfig = map[string]string{
	// Azure Cloud
	"https://management.core.windows.net/": "https://management.azure.com/",
	// Azure China Cloud
	"https://management.core.chinacloudapi.cn/": "https://management.chinacloudapi.cn",
	// Azure US Government
	"https://management.core.usgovcloudapi.net/": "https://management.usgovcloudapi.net/",
	// Azure German Cloud
	"https://management.core.cloudapi.de/": "https://management.microsoftazure.de",
}

type azureAccountAccessTokenInfo struct {
	AccessToken  string `json:"accessToken"`
	ExpiresOn    string `json:"expiresOn"`
	Subscription string `json:"subscription"`
	Tenant       string `json:"tenant"`
	TokenType    string `json:"tokenType"`
}

type AzureConfiguration struct {
	EndpointResourceManager string
	AccessToken             string
	SubscriptionId          string
	TenantId                string
}

func getAzureAccountAccessTokenInfo() azureAccountAccessTokenInfo {
	log.Debug("Getting AccessToken via az CLI.")
	cmd := exec.Command("az", "account", "get-access-token")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()
	if err != nil {
		log.Error("Ensure Azure CLI 2.0 (az) is installed and registered in PATH and that you are logged in via 'az login'.")
		log.Fatal(err)
	}

	log.Debug(cmdOutput.String())
	accessToken := azureAccountAccessTokenInfo{}
	json.Unmarshal(cmdOutput.Bytes(), &accessToken)

	return accessToken
}

func GetAzureConfiguration() AzureConfiguration {
	accessTokenInfo := getAzureAccountAccessTokenInfo()

	jwt, err := jws.ParseJWT([]byte(accessTokenInfo.AccessToken))
	if err != nil {
		log.Fatal(err)
	}

	endpointActiveDirectoryResourceId := jwt.Claims().Get("aud").(string)
	log.Debug("endpoint_active_directory_resource_id (via JWT aud): ", endpointActiveDirectoryResourceId)

	endpointResourceManager := cloudConfig[endpointActiveDirectoryResourceId]
	log.Debug("endpoint_resource_manager (via translation table): ", endpointResourceManager)

	return AzureConfiguration{
		EndpointResourceManager: endpointResourceManager,
		SubscriptionId:          accessTokenInfo.Subscription,
		AccessToken:             accessTokenInfo.AccessToken,
		TenantId:                accessTokenInfo.Tenant,
	}
}
