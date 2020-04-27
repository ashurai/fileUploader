package handlres_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func uploadFileTest(t *testing.T) {
	Path := "temp/files/test-data/test.go"//The path to upload the file
    file, err := os.Open(path)
    if err != nil {
        t.Error(err)
    }
 
    defer file.Close()
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", filepath.Base(path))
    if err != nil {
                writer.Close()
        t.Error(err)
    }
    io.Copy(part, file)
    writer.Close()
 
    req := httptest.NewRequest("POST", "/upload", body)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    res := httptest.NewRecorder()
 
    //UploadFile(res, req)
 
    if res.Code != http.StatusOK {
        t.Error("not 200")
    }
 
    t.Log(res.Body.String())
}