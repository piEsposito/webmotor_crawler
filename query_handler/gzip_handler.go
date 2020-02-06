package query_handler

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ConvertGzipToString(resp *http.Response) string {
	//convert gzip response to string

	reader, _ := gzip.NewReader(resp.Body)
	result, _ := ioutil.ReadAll(reader)
	return string(result)
}

func SaveJsonString(content string, path string) {
	//save our crawled json string
	f, _ := os.Create(path)
	fmt.Fprintln(f, content)
}
