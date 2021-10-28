package monitors

type Monitors struct {
	Monitors []*MonitorCollectionElement `json:"monitors"` // The list of synthetic monitors
}
