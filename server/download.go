package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type DownloadRequestMetadata struct {
	ID    string
	CRC32 uint32
}

// Download godoc
//
//	@Summary		Download file from storage
//	@Description	Download a file from specified storage path
//	@Tags			storage
//	@Produce		application/octet-stream
//	@Param			storage_name	path		string					true	"Storage name"
//	@Param			relative_path	path		string					true	"Relative path"
//	@Param			object_id		path		string					true	"Object ID"
//	@Success		200				{file}		file					"File content"
//	@Failure		400				{object}	map[string]interface{}	"Bad Request"
//	@Failure		404				{object}	map[string]interface{}	"Not Found"
//	@Failure		500				{object}	map[string]interface{}	"Internal Server Error"
//	@Router			/{storage_name}/{relative_path}/objects/{object_id}/content [get]
func (s *Server) Download(w http.ResponseWriter, r *http.Request) {
	slog.Info("Начало обработки запроса на скачивание")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "используешь неразрешенный метод, нужен GET", http.StatusMethodNotAllowed)
		slog.Error("Недопустимый метод запроса: %s, необходим: %s", r.Method, http.MethodGet)
		return
	}

	if errMsg := validateStorageName(r, "storage_name"); errMsg != "" {
		http.Error(w, errMsg, http.StatusBadRequest)
		slog.Error(errMsg)
		return
	}

	// 	func checkStorageName(r *http.Request, storageName string) bool {
	// 	nameInReqPes := chi.URLParam(r, storageName)
	// 	nameInCfg := os.Getenv("S3_STORAGE_NAME")
	// 	return nameInReqPes == nameInCfg
	// }

	// 	S3_STORAGE_NAME="FORTRESS"

	// 	router.Post("/{storage_name}/{relative_path}/objects/{object_id}/content", s.Upload)

	parsedData := chi.URLParam(r, "object_id")
	downloadData, err := getIDandCRC(parsedData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("Не удалось извлечь object_id и crc32", "error", err)
		return
	}

	if err := s.loadManager.Download(w, s.ctx, downloadData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// slog.Error("Ошибка загрузки файла", "error", err)
		return
	}

	slog.Info("Файл успешно скачан", "object_id", downloadData.ID)
}
