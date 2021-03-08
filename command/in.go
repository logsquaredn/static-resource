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
		switch t := value.(type) {
		case int, string, bool:
			if strings.EqualFold(req.Params.Format, "raw") {
				ioutil.WriteFile(filepath.Join(src, key), []byte(fmt.Sprintf("%s", t)), 0644)
			} else if strings.EqualFold(req.Params.Format, "trim") {
				ioutil.WriteFile(filepath.Join(src, key), []byte(strings.Trim(fmt.Sprintf("%s", t), " \t\n")), 0644)
			} else { // strings.EqualFold(req.Params.Format, "json") or "yaml" or "yml"
				return fmt.Errorf("format %s doesn't apply to raw types", req.Params.Format)
			}

		default:
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
			} else { // strings.EqualFold(req.Params.Format, "raw") or "trim"
				return fmt.Errorf("format %s doesn't apply to objects or arrays", req.Params.Format)
			}

			ioutil.WriteFile(filepath.Join(src, key + "." + req.Params.Format), data, 0644)
		}
	}

	r.writeOutput(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}
