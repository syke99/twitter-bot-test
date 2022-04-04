package main

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func downloadImage(URL string) (string, []byte, error) {
	fileName := createFileName(URL)

	res, err := http.Get(URL)
	if err != nil {
		return "", nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", nil, errors.New("received non 200 response code")
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	return fileName, bytes, nil
}

func createFileName(URL string) string {
	uniqueId := uuid.New()

	splits := strings.Split(URL, ".")

	fileExt := "." + splits[len(splits)-1]

	return uniqueId.String() + fileExt
}
