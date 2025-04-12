package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getVersions(id string) []Version {
	url := fmt.Sprintf("https://api.nuget.org/v3-flatcontainer/%s/index.json", strings.ToLower(id))
	resp, err := http.Get(url)
	if err != nil {
		panic("sad")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var data NuGetVersions
	if err := json.Unmarshal(body, &data); err != nil {
		panic("error parsing json")
	}
	reverse(data.Versions)
	return data.Versions
}
