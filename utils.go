package main

import (
    "fmt"
    "encoding/json"
    "io"
)

func GetJson(body io.ReadCloser, v interface{}) error {
    fmt.Printf("Body: %v", body)
    err := json.NewDecoder(body).Decode(v)
    fmt.Printf("json result: %v, type: %T", v, v)
    return err
}