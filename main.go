package main

import (
    "io"
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

// Response represents the structure of the JSON response from GitLab
type Response []struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
    ExpiresAt any       `json:"expires_at"`
    Key       string    `json:"key"`
    UsageType string    `json:"usage_type"`
}


func main() {
    // Get the GitLab username from the user
    var username string
    fmt.Print("Enter your Gitlab username: ")
    _, err := fmt.Scanln(&username)
    if err != nil {
        fmt.Printf("error reading input: %v", err)
        return
    }

    // Construct the API URL
    url := fmt.Sprintf("https://gitlab.com/api/v4/users/%s/keys", username)

    // Fetch the SSH keys from Gitlab
    keys, err := fetchGitlabKeys(url)
    if err != nil {
        fmt.Printf("Error fetching keys: %v", err)
        return
    }

    // Print the fetched SSH keys
    for _, key := range keys {
        fmt.Println(key.Key)
    }
}

func fetchGitlabKeys(url string) (Response, error) {
    // Perform GET request
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("unable to complete GET request: %w", err)
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    // Unmarshal the JSON response
    var result Response
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("unable to unmarshal JSON: %w", err)
    }
    return result, nil
}
