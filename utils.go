package main

import (
    "encoding/json"
)

func GetJson(body []byte, v interface{}) error {
    return json.Unmarshal(body, v)
}
