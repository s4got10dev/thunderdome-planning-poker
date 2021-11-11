package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var ActiveAlerts []interface{}

// handleGetAlerts gets a list of alerts
// @Summary Get Alerts
// @Description get list of alerts (global notices)
// @Tags alert
// @Produce  json
// @Param limit query int true "Max number of results to return"
// @Param offset query int true "Starting point to return rows from, should be multiplied by limit or 0"
// @Success 200 object standardJsonResponse{data=[]model.Alert}
// @Failure 500 object standardJsonResponse{}
// @Router /alerts [get]
func (a *api) handleGetAlerts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Limit, Offset := getLimitOffsetFromRequest(r, w)
		Alerts, Count, err := a.db.AlertsList(Limit, Offset)
		if err != nil {
			Failure(w, r, http.StatusInternalServerError, err)
			return
		}

		Meta := &pagination{
			Count:  Count,
			Offset: Offset,
			Limit:  Limit,
		}

		Success(w, r, http.StatusOK, Alerts, Meta)
	}
}

// handleAlertCreate creates a new alert
// @Summary Create Alert
// @Description Creates an alert (global notice)
// @Tags alert
// @Produce  json
// @Param name body string false "Name of the alert"
// @Param type body string false "Type of alert" Enums(ERROR, INFO, NEW, SUCCESS, WARNING)
// @Param content body string false "Alert content"
// @Param active body boolean false "Whether alert should be displayed or not"
// @Param allowDismiss body boolean false "Whether or not to allow users to dismiss the alert"
// @Param registeredOnly body boolean false "Whether or not to only show to users with an active session"
// @Success 200 object standardJsonResponse{data=[]model.Alert} "returns active alerts"
// @Failure 500 object standardJsonResponse{}
// @Router /alerts [post]
func (a *api) handleAlertCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyVal := getJSONRequestBody(r, w)

		Name := keyVal["name"].(string)
		Type := keyVal["type"].(string)
		Content := keyVal["content"].(string)
		Active := keyVal["active"].(bool)
		AllowDismiss := keyVal["allowDismiss"].(bool)
		RegisteredOnly := keyVal["registeredOnly"].(bool)

		err := a.db.AlertsCreate(Name, Type, Content, Active, AllowDismiss, RegisteredOnly)
		if err != nil {
			Failure(w, r, http.StatusInternalServerError, err)
			return
		}

		ActiveAlerts = a.db.GetActiveAlerts()

		Success(w, r, http.StatusOK, ActiveAlerts, nil)
	}
}

// handleAlertUpdate updates an alert
// @Summary Update Alert
// @Description Updates an Alert
// @Tags alert
// @Produce  json
// @Param alertId path int false "the alert ID to update"
// @Param name body string false "Name of the alert"
// @Param type body string false "Type of alert" Enums(ERROR, INFO, NEW, SUCCESS, WARNING)
// @Param content body string false "Alert content"
// @Param active body boolean false "Whether alert should be displayed or not"
// @Param allowDismiss body boolean false "Whether or not to allow users to dismiss the alert"
// @Param registeredOnly body boolean false "Whether or not to only show to users with an active session"
// @Success 200 object standardJsonResponse{data=[]model.Alert} "returns active alerts"
// @Failure 500 object standardJsonResponse{}
// @Router /alerts/{alertId} [put]
func (a *api) handleAlertUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyVal := getJSONRequestBody(r, w)
		vars := mux.Vars(r)

		ID := vars["alertId"]
		Name := keyVal["name"].(string)
		Type := keyVal["type"].(string)
		Content := keyVal["content"].(string)
		Active := keyVal["active"].(bool)
		AllowDismiss := keyVal["allowDismiss"].(bool)
		RegisteredOnly := keyVal["registeredOnly"].(bool)

		err := a.db.AlertsUpdate(ID, Name, Type, Content, Active, AllowDismiss, RegisteredOnly)
		if err != nil {
			Failure(w, r, http.StatusInternalServerError, err)
			return
		}

		ActiveAlerts = a.db.GetActiveAlerts()

		Success(w, r, http.StatusOK, ActiveAlerts, nil)
	}
}

// handleAlertDelete handles deleting an alert
// @Summary Delete Alert
// @Description Deletes an Alert
// @Tags alert
// @Produce  json
// @Param alertId path int false "the alert ID to delete"
// @Success 200 object standardJsonResponse{data=[]model.Alert} "returns active alerts"
// @Failure 500 object standardJsonResponse{}
// @Router /alerts/{alertId} [delete]
func (a *api) handleAlertDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		AlertID := vars["alertId"]

		err := a.db.AlertDelete(AlertID)
		if err != nil {
			Failure(w, r, http.StatusInternalServerError, err)
			return
		}

		ActiveAlerts = a.db.GetActiveAlerts()

		Success(w, r, http.StatusOK, ActiveAlerts, nil)
	}
}