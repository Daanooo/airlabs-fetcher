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

func Fetch(endpoint string, params map[string]any) {
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

	writeToFile(result.Response)

	fmt.Printf("Done. Total records written: %d\n", len(result.Response))
}

func buildUrl(data *request) string {
	baseStripped := strings.Trim(data.baseUrl, "/")

	// Create string from the map of parameters
	paramStr := ""
	for k, v := range data.params {
		paramStr += fmt.Sprintf("&%s=%s", k, v)
	}

	return fmt.Sprintf("%s/%s?api_key=%s%s", baseStripped, data.endpoint, data.apiKey, paramStr)
}

func writeToFile(airports []Airport) {
	currentTime := time.Now().Format("02-01-2006-150405")

	file, err := os.Create(fmt.Sprintf("export_%s.csv", currentTime))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, v := range airports {
		writer.Write(v.toStringSlice())
	}

	fmt.Printf("Writing to file %s\n", file.Name())
}
