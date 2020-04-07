package main


var ApiServer string = "https://www2.deepl.com/jsonrpc"

var Headers map[string]string

type Job struct {
    Kind string `json:"kind"`
    RawEnSentence string `json:"raw_en_sentence"`
    RawEnContextBefore []string `json:"raw_en_context_before"`
    RawEnContextAfter []string `json:"raw_en_context_after"`
    PreferredNumBeams uint8 `json:"preferred_num_beams"`
    Quality string `json:"quality"`
}

type Lang struct {
    UserPreferredLangs []string `json:"user_preferred_langs"`
    SourceLangUserSelected string `json:"source_lang_user_selected"`
    TargetLang string `json:"target_lang"`
}

type Param struct {
    Jobs []Job `json:"jobs"`
    Lang Lang `json:"lang"`
    Priority int8 `json:"priority"`
    CommonJobParams map[string]string `json:"common_job_params"`
    Timestamp int64 `json:"timestamp"`
}

type DataBinary struct {
    JsonRpc string `json:"jsonrpc"`
    Method string `json:"method"`
    Params []Param `json:"params"`
    Id uint64 `json:"id"`
}

type Beam struct {
    NumSymbols int64 `json:"num_symbols"`
    PostProcessedSentence string `json:"postprocessed_sentence"`
    Score float64 `json:"score"`
    TotalLogProb float64 `json:"total_log_prob"`
}

type Translation struct {
    Beams []Beam `json:"beams"`
    Quality string `json:"quality"`
}

type Result struct {
    Date string `json:"date"`
    SourceLang string `json:"source_lang"`
    SourceLangIsConfident int8 `json:"source_lang_is_confident"`
    TargetLang string `json:"target_lang"`
    Timestamp int64 `json:"timestamp"`
    Translations []Translation `json:"translations"`
}

type Error struct {
    Message string
}

type ResponseResult struct {
    Id string `json:"id"`
    JsonRpc string `json:"jsonrpc"`
    Result  `json:"result"`
    Error Error
}
