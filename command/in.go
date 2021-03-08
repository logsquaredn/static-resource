package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
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

	keys := reflect.ValueOf(req.Source)

	for i := 0; i < keys.NumField(); i++ {
		obj := keys.Field(i).Interface()
		var data []byte
		if strings.EqualFold(req.Params.Format, "yaml") || strings.EqualFold(req.Params.Format, "yml") {
			data, err = yaml.Marshal(obj)
			if err != nil && req.Params.Reveal {
				return fmt.Errorf("unable to marshal yml %s", obj)
			} else if err != nil {
				return fmt.Errorf("unable to marshal yml")
			}
		} else { // strings.EqualFold(req.Params.Format, "json"
			data, err = json.Marshal(obj)
			if err != nil {
				return fmt.Errorf("unable to marshal json %s", obj)
			} else if err != nil {
				return fmt.Errorf("unable to marshal json")
			}
		}
		ioutil.WriteFile(filepath.Join(src, keys.Type().Field(i).Name), data, 0644)
	}

	r.writeMetadata(resp.Metadata)

	r.writeOutput(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}
