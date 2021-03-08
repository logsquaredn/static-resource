package commands

import (
	"fmt"

	resource "github.com/logsquaredn/static-resource"
)

// Check runs the in script which checks stdin for a JSON object of the form of a CheckRequest
// fetches and writes the all Versions that are newer than the provided Version to stdout
func (r *StaticResource) Check() error {
	var (
		req resource.CheckRequest
	    resp resource.CheckResponse
	)

	err := r.readInput(&req)
	if err != nil {
		return err
	}

	version, err := r.getVersion(&req.Source)
	if err != nil {
		return err
	}

	resp = append(resp, *version)

	r.writeOutput(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}
