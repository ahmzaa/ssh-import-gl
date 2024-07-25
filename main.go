package main

import (
    "io"
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

// array for the returned data marshaled from the GET request
type Response []struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    CreatedAt time.Time `json:"created_at"`
    ExpiresAt any       `json:"expires_at"`
    Key       string    `json:"key"`
    UsageType string    `json:"usage_type"`
}


func main() {
    // Ask user for their gitlab username
    var usrName string
    fmt.Println("Enter your Gitlab username:")
    _, err := fmt.Scanln(&usrName)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Concatenate the base URL with the username and specific API endpoint
    url := "https://gitlab.com/api/v4/users/"+usrName+"/keys"

    // DEBUG: Test if the url is concatenated correctly
    fmt.Println(url)

    // Perform the GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Unable to complete GET request: ", err)
        return
    }
    
    // Close the get request
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Unmarshal the JSON data in the Response struct
    var result Response
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Unable to unmarshal JSON", err)
    }

    // For loop allows us to print each specific line from the struct
    for _, rec := range result {
        fmt.Println(rec.Key)
    }
}
