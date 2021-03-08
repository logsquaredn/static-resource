package commands

import (
	"fmt"

	resource "github.com/logsquaredn/static-resource"
)

// Out runs the in script which checks stdin for a JSON object of the form of an OutRequest
func (r *StaticResource) Out() error {
	var (
		req  resource.OutRequest
		resp resource.OutResponse
	)

	err := r.readInput(&req)
	if err != nil {
		return err
	}

	version, err := r.getVersion(&req.Source)
	if err != nil {
		return err
	}

	resp.Version = *version

	r.writeOutput(resp)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %s", err)
	}

	return nil
}
