package main

import (
    "flag"
    "net/http"
    "log"
    "fmt"
    "time"
    "encoding/json"
    "bytes"
    "github.com/john-shine/DeepL-Alfred-translator/alfred"
    "os"
    "io/ioutil"
)

func main() {
    var query string
    var isDebug bool

    flag.StringVar(&query, "q", "", "query string to deepl")
    flag.BoolVar(&isDebug, "debug", false, "debug flag")
    flag.Parse()

    xmlFilter := alfred.XMLFilter{}

    if query == "" {
        xmlFilter.Error("出了点问题", "必须指定查询参数，例如：./DeepL-Alfred-translator -q hello")
        os.Exit(0)
    }

    if !isDebug {
        log.SetOutput(ioutil.Discard)
    } else {
        log.SetOutput(os.Stdout)
    }

    var jobs []Job

    jobs = append(jobs, Job {
        Kind: "default",
        RawEnSentence: query,
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
        log.Println(fmt.Sprintf("assemble request data error: %v", err))
        xmlFilter.Error("出了点问题", "准备请求失败")
        os.Exit(0)
    }

    log.Println("body:", string(body))
    requestMethod := http.MethodPost

    request, err := http.NewRequest(requestMethod, ApiServer, bytes.NewBuffer(body))
    if err != nil {
        log.Println(fmt.Sprintf("request to server error: %v", err))
        xmlFilter.Error("出了点问题", "请求失败")
        os.Exit(0)
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
    }

    // addr request headers
    if Headers != nil {
        for key, val := range Headers {
            log.Printf("header key: %v, value: %v\n", key, val)
            request.Header.Add(key, val)
        }
    }

    // http client
    client := &http.Client{Timeout: 10 * time.Second}
    log.Printf("%s URL: %s \n", requestMethod, request.URL.String())
    resp, err := client.Do(request)
    if err != nil {
        log.Printf(fmt.Sprintf("request to: %v error: %v", request.URL.String(), err))
        xmlFilter.Error("出了点问题", "发送请求失败")
        os.Exit(0)
    }

    defer func() {
        if resp != nil {
            resp.Body.Close()
        }
    }()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf(fmt.Sprintf("read body error: %v", err))
        xmlFilter.Error("出了点问题", "读取请求失败")
        os.Exit(0)
    }

    result := ServerResponse{}
    err = GetJson(bodyBytes, &result)
    if err != nil {
        log.Printf(fmt.Sprintf("decode body: %v error: %v", string(bodyBytes), err))
        xmlFilter.Error("出了点问题", "解析请求失败")
        os.Exit(0)
    }

    if resp.StatusCode != http.StatusOK {
        log.Printf(fmt.Sprintf("reponse not ok but status: %v", resp.StatusCode))
        log.Printf("json: %+v\n", result)
        if result.Error.Message != "" {
            xmlFilter.Error("出了点问题", result.Error.Message)
            os.Exit(0)
        }
        xmlFilter.Error("出了点问题", "DeepL服务器异常")
        os.Exit(0)
    }
    log.Printf("request ok!\n")

    log.Printf("request response: %v", string(bodyBytes))
    // alfred.JsonSuccess()
    os.Exit(0)

}