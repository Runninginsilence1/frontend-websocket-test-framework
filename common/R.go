package common

type R struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data,omitempty"`
	Type    string `json:"type"`
}

type DD struct {
	Operation string `json:"operation"`
	Path      string `json:"path"`
	Device    string `json:"device"`
	Mod       string `json:"mod"`
	Ip        string `json:"ip"`
	Port      string `json:"port"`
}

type DeviceInfo struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Total      string `json:"total"`
	Free       string `json:"free"`
	Used       string `json:"used"`
}
