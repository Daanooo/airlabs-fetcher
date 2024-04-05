package fetcher

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Fetch(endpoint string, params []string) []Airport {
	request := NewRequest(endpoint, params)

	res, err := http.Get(buildUrl(request))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Convert response into auto-generated response struct
	result := Response{}
	json.Unmarshal(body, &result)

	return result.Response
}

func MakeExport(airports []Airport, fileName string) error {
	currentTime := time.Now().Format("02-01-2006-150405")

	if fileName == "" {
		fileName = fmt.Sprintf("export_%s.csv", currentTime)
	}

	// Don't proceed with the export if the file already exists to avoid losing data unintentionally
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("file %s already exists, not overwriting", fileName)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("MakeExport(), os.Create(%s) failed with %s", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, v := range airports {
		writer.Write(v.toStringSlice())
	}

	fmt.Printf("%d lines written to %s\n", len(airports), file.Name())

	return nil
}

func buildUrl(data *request) string {
	baseStripped := strings.Trim(data.baseUrl, "/")
	paramStr := strings.Join(data.params, "&")

	return fmt.Sprintf("%s/%s?api_key=%s&%s", baseStripped, data.endpoint, data.apiKey, paramStr)
}
