package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Fetch(endpoint string) {
	request := NewRequest(endpoint)

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

	for _, v := range result.Response {
		fmt.Println(v)
	}
	fmt.Printf("Total: %d\n", len(result.Response))
}

func buildUrl(data *request) string {
	baseStripped := strings.Trim(data.baseUrl, "/")

	return fmt.Sprintf("%s/%s?api_key=%s", baseStripped, data.endpoint, data.apiKey)
}
