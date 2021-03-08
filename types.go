package resource

// CheckRequest is the JSON object that Concourse passes to /opt/resource/check through stdin
type CheckRequest struct {
	Source  Source   `json:"source"`
	Version *Version `json:"version"`
}

// Version is the JSON object that is passed to and from Concourse
type Version struct {
	Hash string `json:"hash"`
}

// CheckResponse is the JSON object that we pass back to Concourse through stdout from /opt/resource/check
type CheckResponse []Version

// InRequest is the JSON object that Concourse passes to /opt/resource/in through stdin
type InRequest struct {
	Source  Source    `json:"source"`
	Params  GetParams `json:"params"`
	Version Version   `json:"version"`
}

// InResponse is the JSON object that we pass back to Concourse through stdout from /opt/resource/in
type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

// OutRequest is the JSON object that Concourse passes to /opt/resource/out through stdin
type OutRequest struct {
	Source Source    `json:"source"`
	Params PutParams `json:"params"`
}

// OutResponse is the JSON object that we pass back to Concourse through stdout from /opt/resource/out
type OutResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

// Source ...
type Source interface {}

// GetParams are additional parameters that can be passed to this Concourse Resource Type during a get step
type GetParams struct {
	Format string `json:"format"`
	Reveal bool   `json:"reveal"`
}

// PutParams are additional parameters that can be passed to this Concourse Resource Type during a put step
type PutParams interface {}

// Metadata is the object which is passed in array form to Concourse through stdout from /opt/resource/out and /opt/resource/in
// to provide additional information about the Version
type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}