package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// required to translate language flags
var (
	// source language
	term = flag.String("term", "", "a term to search")
	// target language
	lang = flag.String("lang", "en-us", "a used language to search")
	// source language text
	// Use the Oxford dictionary API
	endpoint = flag.String("endpoint", "https://od-api.oxforddictionaries.com/api/v2/entries", "dictionary endpoint")
)

func searchUrl(endpoint, lang, term string) string {
	return endpoint + "/" + lang + "/" + term
}

// translate language
func search(term, lang, app_id, app_key string) (string, error) {
	url := searchUrl(*endpoint, lang, term)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("app_id", app_id)
	req.Header.Set("app_key", app_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func run(args []string) int {
	envEndpoint := os.Getenv("GTRAN_ENDPOINT")
	envAppId := os.Getenv("OXDIC_APP_ID")
	envAppKey := os.Getenv("OXDIC_APP_KEY")

	if envAppId == "" || envAppKey == "" {
		fmt.Fprintf(os.Stderr, "OXDIC_APP_ID or OXDIC_APP_KEY not set\n")
		return -1
	}

	if envEndpoint != "" {
		*endpoint = envEndpoint
	}

	if len(args) == 0 && *term == "" {
		flag.Usage()
		return -1
	}

	if *term == "" && args[0] != "" {
		*term = args[0]
	}

	result, err := search(*term, *lang, envAppId, envAppKey)
	if err != nil {
		return -1
	}
	fmt.Println(result)
	return 0
}

func main() {
	flag.Parse()
	os.Exit(run(flag.Args()))
}
