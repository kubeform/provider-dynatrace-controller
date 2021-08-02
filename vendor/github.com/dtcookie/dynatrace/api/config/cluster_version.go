package api

// ClusterVersion is the version of the Dynatrace Cluster
type ClusterVersion struct {
	Version string `json:"version"`
}

// String returns the version of the cluster
func (cv ClusterVersion) String() string {
	return cv.Version
}
