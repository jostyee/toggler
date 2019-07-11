package httpapi

import (
	context2 "context"
	"net/http"

	"github.com/adamluzsi/toggler/extintf/httpintf/httputils"
	"github.com/adamluzsi/toggler/usecases"
)

func (sm *ServeMux) SetPilotEnrollmentForFeature(w http.ResponseWriter, r *http.Request) {

	pu := r.Context().Value(`ProtectedUsecases`).(*usecases.ProtectedUsecases)

	pilot, err := httputils.ParseFlagPilotFromForm(r)

	if err != nil || pilot.FeatureFlagID == `` || pilot.ExternalID == `` {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = pu.SetPilotEnrollmentForFeature(context2.TODO(), pilot.FeatureFlagID, pilot.ExternalID, pilot.Enrolled)

	if handleError(w, err, http.StatusInternalServerError) {
		return
	}

	serveJSON(w,  map[string]interface{}{})

}
