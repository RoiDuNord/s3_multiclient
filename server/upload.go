package server

import (
	"log/slog"
	"net/http"
)

type objectResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Size   int    `json:"size_mb"`
}

type UploadRequestMetadata struct {
	ID          string
	FileName    string
	ContentType string
	Size        int64
}

// Upload godoc
//
//	@Summary		Upload file to storage
//	@Description	Upload a file to specified storage path
//	@Tags			storage
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			storage_name	path		string					true	"Storage name"
//	@Param			relative_path	path		string					true	"Relative path"
//	@Param			object_id		path		string					true	"Object ID"
//	@Param			file			formData	file					true	"File to upload"
//	@Success		200				{object}	objectResponse			"Upload successful"
//	@Failure		400				{object}	map[string]interface{}	"Bad Request"
//	@Failure		500				{object}	map[string]interface{}	"Internal Server Error"
//	@Router			/{storage_name}/{relative_path}/objects/{object_id}/content [post]
func (s *Server) Upload(w http.ResponseWriter, r *http.Request) {
	slog.Info("Начало обработки запроса на загрузку")

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "недопустимый метод запроса, нужен POST", http.StatusMethodNotAllowed)
		slog.Error("Недопустимый метод запроса")
		return
	}

	if errMsg := validateStorageName(r, "storage_name"); errMsg != "" {
		http.Error(w, errMsg, http.StatusBadRequest)
		slog.Error(errMsg)
		return
	}

	data, err := getUploadRequestData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := s.loadManager.Upload(r, s.ctx, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, data)
}
