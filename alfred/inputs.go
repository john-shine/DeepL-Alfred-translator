package alfred

import (
    "encoding/json"
    "fmt"
    "encoding/xml"
    "log"
)

type FilterItem struct {
   XMLName     xml.Name    `json:"-" xml:"item"`
   UID         string      `json:"uid,omitempty" xml:"uid,omitempty,attr"`
   Type        string      `json:"type,omitempty" xml:"type,omitempty"`
   Title       string      `json:"title" xml:"title"`
   Subtitle    string      `json:"subtitle" xml:"subtitle"`
   Arg         string      `json:"arg,omitempty" xml:"arg,attr"`
   Valid       bool        `json:"valid" xml:"valid"`
   Icon        string      `json:"icon,omitempty" xml:"icon,omitempty"`
   Extra       customMap
}

type FilterItems struct {
   XMLName xml.Name   `json:"-" xml:"items"`
   Items []FilterItem `json:"items"`
}

type Inputs interface {
    Error(string string)
    Success([]FilterItem)
}

type JsonFilter struct {}

type XMLFilter struct {}

func NewInput() *FilterItems {
    r := new(FilterItems)
    r.Items = []FilterItem{}

    return r
}

func (v JsonFilter) Error(title, subtitle string) {
    data := NewInput()
    item := FilterItem{
        Valid: false,
        Title: title,
        Subtitle: subtitle,
        Icon: "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
    }
    data.Items = append(data.Items, item)

    r, err := json.Marshal(data)
    if err != nil {
        log.Println("json encode error: ", err)
    }

    fmt.Printf(string(r))
}

func (v XMLFilter) Error(title, subtitle string) {
    data := NewInput()
    item := FilterItem{
        Valid: false,
        Title: title,
        Subtitle: subtitle,
        Icon: "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
    }
    data.Items = append(data.Items, item)

    r, err := xml.Marshal(data)
    if err != nil {
        log.Println("xml encode error: ", err)
    }

    fmt.Printf(xml.Header + string(r))
}

func (v JsonFilter) Success(items []FilterItem) {
    data := NewInput()
    data.Items = append(data.Items, items...)

    r, _ := json.Marshal(data)
    fmt.Printf(string(r))
}

func (v XMLFilter) Success(items []FilterItem) {
    data := NewInput()
    data.Items = append(data.Items, items...)

    r, _ := xml.Marshal(data)
    fmt.Printf(string(r))
}

// customMap is a map[string]string.
type customMap map[string]string

// customMap marshals into XML.
func (s customMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

    tokens := []xml.Token{start}

    for key, value := range s {
        t := xml.StartElement{Name: xml.Name{"", key}}
        tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
    }

    tokens = append(tokens, xml.EndElement{start.Name})

    for _, t := range tokens {
        err := e.EncodeToken(t)
        if err != nil {
            return err
        }
    }

    // flush to ensure tokens are written
    err := e.Flush()
    if err != nil {
        return err
    }

    return nil
}