package util

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	urllib "net/url"
	"strings"
)

func GetInput(url string, session_cookie string) string {
	// dir, _ := os.Executable()
	// input_file := dir + "input"
	// if _, err := os.Stat(input_file); err == nil {
	// content, _ := os.ReadFile(input_file)
	// return string(content)
	// }
	content := string(querySite(url, session_cookie))
	content = strings.Trim(content, "\n")
	// os.WriteFile(input_file, content, os.FileMode(0700))
	// log.Println("input file: " + input_file)
	return content
}

func querySite(url string, session_cookie string) []byte {
	jar, _ := cookiejar.New(nil)
	u, _ := urllib.Parse(url)
	log.Println("using cookie:", session_cookie)
	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "session",
			Value: session_cookie,
		},
	})
	client := &http.Client{
		Jar: jar,
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// log.Println("Fetched input from site")
	return body
}
