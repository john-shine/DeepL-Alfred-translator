package alfred

import (
    "encoding/json"
    "fmt"
    "encoding/xml"
)

type JsonItems struct {
    XMLName xml.Name `xml:"items"`
    Items []JsonItem `json:"items",xml:"items"`
}

type IconItem struct {
    Type string `json:"type"`
    Path string `json:"path"`
}

type JsonItem struct {
    UID           string      `json:"uid,omitempty"`
    Type          string      `json:"type,omitempty"`
    Title         string      `json:"title"`
    Subtitle      string      `json:"subtitle"`
    Arg           string      `json:"arg,omitempty"`
    AutoComplete  bool        `json:"autocomplete"`
    Icon          IconItem    `json:"icon,omitempty"`
}

func NewInput() *JsonItems {
    r := new(JsonItems)
    r.Items = []JsonItem{}

    return r
}

func JsonError(title, subtitle string) {
    data := NewInput()
    item := JsonItem{
        AutoComplete: false,
        Title: title,
        Subtitle: subtitle,
        Icon: IconItem {
            Type: "fileicon",
            Path: "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
        },
    }
    data.Items = append(data.Items, item)

    result, _ := json.Marshal(data)
    fmt.Printf(string(result))
}

func XMLError(title, subtitle string) {
    data := NewInput()
    item := JsonItem{
        AutoComplete: false,
        Title: title,
        Subtitle: subtitle,
        Icon: IconItem {
            Type: "fileicon",
            Path: "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
        },
    }
    data.Items = append(data.Items, item)

    r, _ := xml.Marshal(data)

    fmt.Printf(xml.Header + string(r))
}

func JsonSuccess(items []JsonItem) {
    data := NewInput()
    data.Items = append(data.Items, items...)

    r, _ := json.Marshal(data)
    fmt.Printf(string(r))
}