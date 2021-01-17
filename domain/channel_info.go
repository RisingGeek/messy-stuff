package domain

type ChannelEl struct {
	ID   int   `json:"id"`
	Time int64 `json:"time"`
}

type WorkerInfo struct {
	Workers   []ChannelEl `json:"workers"`
	TotalTime int64       `json:"totalTime"`
}
