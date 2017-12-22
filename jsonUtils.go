package main

import (
	"bytes"
	"encoding/json"
	"github.com/jmespath/go-jmespath"
	log "github.com/sirupsen/logrus"
)

func jsonPrettyPrint(jsonRaw []byte) string {
	var prettyJson bytes.Buffer
	indentErr := json.Indent(&prettyJson, jsonRaw, "", "  ")
	if indentErr != nil {
		log.Fatal(indentErr)
	}

	return string(prettyJson.Bytes())
}

func applyJmespathToJson(jsonRaw []byte, jmespathExpression string) string {
	if jmespathExpression == "" {
		return jsonPrettyPrint(jsonRaw)
	}

	jmespathParser := jmespath.NewParser()

	parsed, parseErr := jmespathParser.Parse(jmespathExpression)
	if parseErr != nil {
		if syntaxError, ok := parseErr.(jmespath.SyntaxError); ok {
			log.Fatal("jmespath error ", syntaxError, " ", syntaxError.HighlightLocation())
		}

		log.Fatal(parseErr)
	} else {
		log.Debug("jmespath expression\n", parsed)
	}

	var jsonData interface{}
	jsonUnmarshalErr := json.Unmarshal(jsonRaw, &jsonData)
	if jsonUnmarshalErr != nil {
		log.Fatal(jsonUnmarshalErr)
	}

	jmespathResult, jmespathErr := jmespath.Search(jmespathExpression, jsonData)
	if jmespathErr != nil {
		log.Fatal(jmespathResult)
	}

	prettyJson, indentErr := json.MarshalIndent(jmespathResult, "", "  ")
	if indentErr != nil {
		log.Fatal(indentErr)
	}

	log.Debug("JSON\n", jsonPrettyPrint(jsonRaw), "\n")
	log.Debug("jmespath(JSON)\n", string(prettyJson), "\n")
	return string(prettyJson)
}
