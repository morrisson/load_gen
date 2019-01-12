package main

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/morrisson/load_gen/client"
	"github.com/morrisson/load_gen/common"
	"github.com/morrisson/load_gen/generator"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	token := client.GetV2Token()

	sitesLoadGen := &generator.SitesLoadGenerator{
		GohanClient: &client.GohanClient{
			Endpoint:   common.GohanEndpointURL,
			Token:      token.ID,
			HttpClient: &http.Client{},
		},
	}
	servicesLoadGen := &generator.ServicesLoadGenerator{
		GohanClient: &client.GohanClient{
			Endpoint:   common.GohanEndpointURL,
			Token:      token.ID,
			HttpClient: &http.Client{},
		},
	}
	policiesLoadGen := &generator.PoliciesLoadGenerator{
		GohanClient: &client.GohanClient{
			Endpoint:   common.GohanEndpointURL,
			Token:      token.ID,
			HttpClient: &http.Client{},
		},
	}

	loadGens := []generator.LoadGenerator{
		sitesLoadGen, servicesLoadGen, policiesLoadGen,
	}

	wg := &sync.WaitGroup{}
	for _, loadGen := range loadGens {
		go func(loadGen generator.LoadGenerator) {
			wg.Add(1)
			loadGen.Run()
			wg.Done()
		}(loadGen)
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	}
	wg.Wait()

	for {
	}
}
