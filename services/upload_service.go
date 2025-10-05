package services

import (
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

type UploadService struct {
    UploadDir string
}

func NewUploadService(uploadDir string) *UploadService {
    return &UploadService{UploadDir: uploadDir}
}

func (u *UploadService) SaveToLocal(file multipart.File, filename string) error {
    path := filepath.Join(u.UploadDir, filename)
    out, err := os.Create(path)
    if err != nil {
        return err
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    return err
}
