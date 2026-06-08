package models

type AnalyzeResponse struct {
    Message  string          `json:"message"`
    Insights map[string]any  `json:"insights"`
}

type WaveformResponse struct {
    Message    string    `json:"message"`
    Samples    []float64 `json:"samples"`
    SampleRate int       `json:"sampleRate"`
}

type MetricsResponse struct {
    Message string             `json:"message"`
    Metrics map[string]float64 `json:"metrics"`
}
