package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

func ping(url string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s is down\n", url)
	} else {
		c <- fmt.Sprintf("%s is up\n", url)
		resp.Body.Close()
	}
}

func getTitle(content string) (title string) {
	re := regexp.MustCompile("<title>(.*)</title>")

	parts := re.FindStringSubmatch(content)
	if len(parts) > 0 {
		return parts[1]
	} else {
		return "no title"
	}
}

func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://golangr.com/pointers",
		"https://sixtusdevportfolio.netlify.app",
		"https://nicepage.com",
		"https://www.tutorialspoint.com",
		"https://go.dev",
		"https://blog.logrocket.com",
		"https://www.freecodecamp.org",
		"http://webcode.me",
		"https://example.com",
		"http://httpbin.org",
		"https://www.perl.org",
		"https://www.php.net",
		"https://www.python.org",
		"https://code.visualstudio.com",
		"https://clojure.org",
		"https://chat.openai.com",
	}

	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(urls))

	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("output_%s.txt", timestamp)

	for _, url := range urls {
		go ping(url, &wg, c)

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("%s is down\n", url)
				return
			}
			defer resp.Body.Close()

			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Failed to read response from %s: %v\n", url, err)
				return
			}

			title := getTitle(string(content))
			fmt.Println(title)

			result := fmt.Sprintf("%s - %s\n", url, title)
			err = writeToFile(filename, result)
			if err != nil {
				fmt.Printf("Failed to write to file: %v\n", err)
			}
		}(url)

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

}
