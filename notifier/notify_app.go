package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run notify_app.go <changed_file>")
	}
	changedFile := os.Args[1]

	// Normalize the file path to match the map keys
	if !strings.HasPrefix(changedFile, "api/") {
		changedFile = "api/" + changedFile
	}

	// Map file names to App Insights queries
	queries := map[string]string{
		"api/get_handler.go":  `customEvents | where name == \"GET /v1/profiles/phones-and-email\" | extend appName = tostring(customDimensions[\"application-name\"]) | summarize by appName`,
		"api/post_handler.go": `customEvents | where name == \"POST /v1/profiles/phones-and-email\" | extend appName = tostring(customDimensions[\"application-name\"]) | summarize by appName`,
	}
	query, ok := queries[changedFile]
	if !ok {
		log.Fatalf("No query mapped for file: %s", changedFile)
	}

	// Use provided App ID and API Key
	appID := "29293bd0-493f-4f08-aef2-eba210b93093"
	apiKey := "k780qwk3kly6pixcagzfdb3ib3kwn1b2xbcb4jb4"
	url := fmt.Sprintf("https://api.applicationinsights.io/v1/apps/%s/query", appID)
	body := fmt.Sprintf(`{"query": %q}`, query)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request error: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	type Table struct {
		Name    string          `json:"name"`
		Columns []interface{}   `json:"columns"`
		Rows    [][]interface{} `json:"rows"`
	}
	type Response struct {
		Tables []Table `json:"tables"`
	}

	var result Response
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Map app names to Slack channel IDs
	appSlackChannels := map[string]string{
		"membercore":   "C09G161KD0Q",
		"member funds": "C09F8HM77L6",
	}
	slackToken := os.Getenv("SLACK_TOKEN") // Pass as env variable

	for _, table := range result.Tables {
		for _, row := range table.Rows {
			if appName, ok := row[0].(string); ok && appName != "" {
				fmt.Println("App Name:", appName)
				channelID, exists := appSlackChannels[appName]
				if exists {
					message := fmt.Sprintf("API changed: %s. Please review!", appName)
					slackURL := "https://slack.com/api/chat.postMessage"
					payload := map[string]interface{}{
						"channel": channelID,
						"text":    message,
					}
					payloadBytes, _ := json.Marshal(payload)
					req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(payloadBytes))
					if err != nil {
						log.Printf("Slack request error: %v", err)
						continue
					}
					req.Header.Set("Authorization", "Bearer "+slackToken)
					req.Header.Set("Content-Type", "application/json")
					resp, err := client.Do(req)
					if err != nil {
						log.Printf("Slack notification error: %v", err)
						continue
					}
					defer resp.Body.Close()
					respBody, _ := ioutil.ReadAll(resp.Body)
					fmt.Printf("Slack response: %s\n", string(respBody))
				}
			}
		}
	}
}
