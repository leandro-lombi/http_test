package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	var payload bytes.Buffer

	nw := multipart.NewWriter(&payload)

	file, err := os.Open("./testdata/image.png")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	w, err := nw.CreateFormFile("file", "image.png")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		t.Error(err)
	}

	nw.Close()

	nr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upload", &payload)
	r.Header.Add("Content-Type", nw.FormDataContentType())

	UploadFile(nr, r)

	if http.StatusOK != nr.Code {
		t.Errorf("status code expected: %d got %d", http.StatusOK, nr.Code)
	}

	response, err := io.ReadAll(nr.Body)
	if err != nil {
		t.Error(err)
	}

	if "Sucessfully uploaded File\n" != string(response) {
		t.Errorf("response expected: 'Sucessfully uploaded File\n' got %s", string(response))
	}

}
