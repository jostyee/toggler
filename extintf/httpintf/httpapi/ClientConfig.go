package httpapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/toggler-io/toggler/extintf/httpintf/httputils"
)

// ClientConfigRequest defines the parameters that
// swagger:parameters ClientConfig
type ClientConfigRequest struct {
	// in: body
	Body struct {
		// PilotExtID is the public uniq id that identify the caller pilot
		//
		// required: true
		// example: pilot-external-id-which-is-uniq-in-the-system
		PilotExtID string `json:"id"`
		// ReleaseFlags are the list of private release flag name that should be matched against the pilot and state the enrollment for each.
		//
		// required: true
		// example: ["my-release-flag"]
		ReleaseFlags []string `json:"release_flags"`
	}
}

// ClientConfigResponse returns information about the requester's rollout feature enrollment statuses.
// swagger:response clientConfigResponse
type ClientConfigResponse struct {
	// in: body
	Body ClientConfigResponseBody
}

// ClientConfigResponseBody will contain the requested feature flag states for a certain pilot.
// The content expected to be cached in some form of state container.
type ClientConfigResponseBody struct {
	// Release holds information related the release management
	Release struct {
		// Flags hold the states of the release flags of the client
		Flags map[string]bool `json:"flags"`
	} `json:"release"`
}

/*

	swagger:route GET /client/config.json release-flag pilot ClientConfig

	Return all the flag states that was requested in the favor of a Pilot.
	This endpoint especially useful for Mobile & SPA apps.
	The endpoint can be called with HTTP GET method as well,
	POST is used officially only to support most highly abstracted http clients,
	where using payload to upload cannot be completed with other http methods.

		Consumes:
		- application/json

		Produces:
		- application/json

		Schemes: http, https

		Responses:
		  200: clientConfigResponse
		  400: errorResponse
		  500: errorResponse

*/
func (sm *ServeMux) ClientConfigJSON(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	payloadDecoder := json.NewDecoder(r.Body)

	var request ClientConfigRequest
	parseErr := payloadDecoder.Decode(&request.Body)

	if parseErr != nil {
		parseErr = nil
		q := r.URL.Query()
		request.Body.PilotExtID = q.Get(`id`)
		request.Body.ReleaseFlags = append([]string{}, q[`release_flags`]...)
		request.Body.ReleaseFlags = append(request.Body.ReleaseFlags, q[`release_flags[]`]...)
	}

	ctx := context.WithValue(r.Context(), `pilot-ip-addr`, httputils.GetClientIP(r))

	states, err := sm.UseCases.GetReleaseFlagPilotEnrollmentStates(ctx, request.Body.PilotExtID, request.Body.ReleaseFlags...)

	if handleError(w, err, http.StatusInternalServerError) {
		return
	}

	var body ClientConfigResponseBody
	body.Release.Flags = states
	serveJSON(w, body)
}
