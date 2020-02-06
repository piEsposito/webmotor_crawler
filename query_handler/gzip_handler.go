package query_handler

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ConvertGzipToString(resp *http.Response) string {
	/*
		As our return from the request is compressed in gzip format,
		We return it as a string to be written in a file
	*/
	reader, _ := gzip.NewReader(resp.Body)
	result, _ := ioutil.ReadAll(reader)
	return string(result)
}

func SaveJsonString(content string, path string) {
	/*
		With this function we save our crawled Json to string
	*/
	f, _ := os.Create(path)
	fmt.Fprintln(f, content)
}
