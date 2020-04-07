package main

import (
    "flag"
    "net/http"
    "log"
    "fmt"
    "os"
    "time"
    "encoding/json"
    "bytes"
    "io"
)

func main() {
    var query string
    flag.StringVar(&query, "q", "", "query string")
    flag.Parse()

    if query == "" {
        panic("must specify query string by ./DeepL-Alfred-translator -q ${query}")
    }
    var jobs []Job

    jobs = append(jobs, Job {
        Kind: "default",
        RawEnSentence: "b",
        RawEnContextBefore: []string{},
        RawEnContextAfter: []string{},
        PreferredNumBeams: 4,
        Quality: "fast",
    })

    var params []Param

    params = append(params, Param {
        Jobs: jobs,
        Lang: Lang{
            UserPreferredLangs: []string{"EN"},
            SourceLangUserSelected: "auto",
            TargetLang: "ZH",
        },
        Priority: -1,
        CommonJobParams: make(map[string]string),
        Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
    })

    bodyData := DataBinary {
        JsonRpc: "2.0",
        Method: "LMT_handle_jobs",
        Params: params,
    }

    body, err := json.Marshal(bodyData)
    if err != nil {
        panic(fmt.Sprintf("assemble request data error: %v", err))
    }

    fmt.Println("body:", string(body))
    requestMethod := http.MethodPost

    request, err := http.NewRequest(requestMethod, ApiServer, bytes.NewBuffer(body))
    if err != nil {
        log.Println(err)
        panic("new request is fail.")
    }

    Headers = map[string]string {
        "authority": "www2.deepl.com",
        "origin": "https://www.deepl.com",
        "sec-fetch-dest": "empty",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36 Edg/80.0.361.62' -H 'content-type: text/plain",
        "accept": "*/*",
        "sec-fetch-site": "same-site",
        "sec-fetch-mode": "cors",
        "referer": "https://www.deepl.com/translator",
        "accept-language": "zh-Hans-CN,zh-CN;q=0.9,zh;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5",
        "cookie": "LMTBID=9cb3eebb-9f13-4936-b22c-ab72656d1f9f|09ef6a01196c9faa4abebc4d6e4b7f6d",
    }

    // addr request headers
    if Headers != nil {
        for key, val := range Headers {
            fmt.Printf("key: %v, value: %v\n", key, val)
            request.Header.Add(key, val)
        }
    }

    // http client
    client := &http.Client{Timeout: 10 * time.Second}
    log.Printf("%s URL: %s \n", requestMethod, request.URL.String())
    response, err := client.Do(request)
    if err != nil {
        panic(fmt.Sprintf("request to: %v with error: %v", request.URL.String(), err))
    }

    defer response.Body.Close()

    result := ResponseResult{}
    err = GetJson(response.Body, result)
    if err != nil {
        panic(err)
    }

    if response.StatusCode != http.StatusOK {
        fmt.Printf("request response: ")
        panic(fmt.Sprintf("request to: %v, result with error status: %v", request.URL.String(), response.StatusCode))
    }
    fmt.Printf("request to: %v ok!\n", request.URL.String())

    fmt.Printf("request response: ")
    io.Copy(os.Stdout, response.Body)

}