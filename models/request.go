package models

type AnalyzeRequest struct {
    AudioURL string `json:"audioUrl"`
}

type WaveformRequest struct {
    AudioURL string `json:"audioUrl"`
}

type MetricsRequest struct {
    AudioURL string `json:"audioUrl"`
}
