package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// getResult returns the address of the request along with the
// MD5 hash of the response.
func getResult(addr string) (string, error) {
	url, err := parceUrl(addr)
	if err != nil {
		return "", err
	}
	body, err := sendRequest(url)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%s %x", url, md5.Sum(body))
	return result, nil
}

// parceUrl parces addr into a URL structure and reassembles
// it into a valid URL string.
//
//	- if u.Scheme is empty, scheme: is http.
func parceUrl(addr string) (string, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String(), nil
}

// sendRequest sends an http GET request to addr and returns the
// body response
func sendRequest(addr string) ([]byte, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func main() {
	parallel := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	addresses := flag.Args()

	ch := make(chan struct{}, *parallel)
	wg := sync.WaitGroup{}

	for _, address := range addresses {
		wg.Add(1)
		ch <- struct{}{}
		go func(addr string) {
			defer wg.Done()
			result, err := getResult(addr)
			if err != nil {
				log.Println(err)
			}
			if result != "" {
				fmt.Println(result)
			}
			<-ch
		}(address)
	}
	wg.Wait()
	close(ch)
}
