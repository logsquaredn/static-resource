package commands

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	resource "github.com/logsquaredn/static-resource"
)

// StaticResource struct which has the Check, In, and Out methods on it which comprise
// the three scripts needed to implement a Concourse Resource Type
type StaticResource struct {
	stdin  io.Reader
	stderr io.Writer
	stdout io.Writer
	args   []string
}

// NewStaticResource creates a new StaticResource struct
func NewStaticResource(
	stdin io.Reader,
	stderr io.Writer,
	stdout io.Writer,
	args []string,
) *StaticResource {
	return &StaticResource{
		stdin,
		stderr,
		stdout,
		args,
	}
}

func (r *StaticResource) readInput(req interface{}) error {
	decoder := json.NewDecoder(r.stdin)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		return fmt.Errorf("invalid payload: %s", err)
	}

	return nil
}

func (r *StaticResource) writeOutput(resp interface{}) error {
	err := json.NewEncoder(r.stdout).Encode(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}

func (r *StaticResource) getSrc() (string, error) {
	if len(r.args) < 2 {
		return "", fmt.Errorf("destination path not specified")
	}
	return r.args[1], nil
}

func (r *StaticResource) getVersion(s *resource.Source) (*resource.Version, error) {
	json, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal json")
	}
	return &resource.Version{
		Hash: fmt.Sprintf("%x", md5.Sum(json)),
	}, nil
}

func (r *StaticResource) writeMetadata(mds []resource.Metadata) error {
	src, err := r.getSrc()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(src, ".metadata"), 0755)
	if err != nil {
		return fmt.Errorf("unable to make directory %s", filepath.Join(src, ".metadata"))
	}

	for _, md := range mds {
		err = ioutil.WriteFile(filepath.Join(src, ".metadata", md.Name), []byte(md.Value), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *StaticResource) expandEnv(s string) string {
	return os.Expand(s, func(v string) string {
		switch v {
		case "BUILD_ID", "BUILD_NAME", "BUILD_JOB_NAME", "BUILD_PIPELINE_NAME", "BUILD_TEAM_NAME", "ATC_EXTERNAL_URL":
			return os.Getenv(v)
		}
		return "$" + v
	})
}
