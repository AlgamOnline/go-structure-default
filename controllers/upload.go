package controllers

import (
    "net/http"
    "golang-default/services"
)

func UploadImageHandler(uploadService *services.UploadService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(10 << 20)
        file, handler, err := r.FormFile("image")
        if err != nil {
            http.Error(w, "Error retrieving file", http.StatusBadRequest)
            return
        }
        defer file.Close()

        err = uploadService.SaveToLocal(file, handler.Filename)
        if err != nil {
            http.Error(w, "Failed to save file", http.StatusInternalServerError)
            return
        }
        w.Write([]byte("Upload successful"))
    }
}
