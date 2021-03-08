package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	resource "github.com/logsquaredn/static-resource"
	"gopkg.in/yaml.v2"
)

// In runs the in script which checks stdin for a JSON object of the form of an InRequest
// fetches and writes the requested Version as well as Metadata about it to stdout
func (r *StaticResource) In() error {
	var (
		req  resource.InRequest
		resp resource.InResponse
	)

	err := r.readInput(&req)
	if err != nil {
		return err
	}

	src, err := r.getSrc()
	if err != nil {
		return err
	}

	err = os.MkdirAll(src, 0755)
	if err != nil {
		return fmt.Errorf("unable to make directory %s", src)
	}

	version, err := r.getVersion(&req.Source)
	if err != nil {
		return err
	}

	resp.Version = *version

	for key, value := range req.Source {
		var data []byte
		if strings.EqualFold(req.Params.Format, "yaml") || strings.EqualFold(req.Params.Format, "yml") {
			data, err = yaml.Marshal(value)
			if err != nil && req.Params.Reveal {
				return fmt.Errorf("unable to marshal yml %s", value)
			} else if err != nil {
				return fmt.Errorf("unable to marshal yml")
			}
		} else if strings.EqualFold(req.Params.Format, "json") {
			data, err = json.MarshalIndent(value, "", "	")
			if err != nil {
				return fmt.Errorf("unable to marshal json %s", value)
			} else if err != nil {
				return fmt.Errorf("unable to marshal json")
			}
		} // else { // strings.EqualFold(req.Params.Format, "json") }
		ioutil.WriteFile(filepath.Join(src, key), data, 0644)
	}

	r.writeMetadata(resp.Metadata)

	r.writeOutput(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}
