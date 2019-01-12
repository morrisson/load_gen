package generator

import (
	"fmt"
	"sync"
	"time"

	"github.com/morrisson/load_gen/client"
)

type ServicesLoadGenerator struct {
	GohanClient *client.GohanClient
}

func (gen *ServicesLoadGenerator) Run() {
	t := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-t.C:
			go gen.run()
		}
	}
	t.Stop()
}

func (gen *ServicesLoadGenerator) run() {

	serviceFuncMap := map[string](func() (map[string]interface{}, error)){
		"service":          gen.GohanClient.GetServices,
		"service_template": gen.GohanClient.GetServiceTemplates,
	}

	policyFuncMap := map[string](func() (map[string]interface{}, error)){
		"application_gateway": gen.GohanClient.GetApplicationGatewayPolicies,
		"firewall":            gen.GohanClient.GetFirewallPolicies,
		"hybrid_wan":          gen.GohanClient.GetHybridWanPolicies,
		"microsegmentation":   gen.GohanClient.GetMicrosegmentationPolicies,
		"qos":                 gen.GohanClient.GetQosPolicies,
	}

	// Get services and service templates
	wg := &sync.WaitGroup{}
	for key, fn := range serviceFuncMap {
		go func(key string, fn func() (map[string]interface{}, error)) {
			wg.Add(1)

			startTime := time.Now()
			_, err := fn()
			endTime := time.Now()

			processSec := endTime.Sub(startTime).Seconds()
			if processSec > 2.0 {
				fmt.Printf("Took %g sec to get %ss in 'Services' page\n", processSec, key)
			}

			if err != nil {
				fmt.Println(err)
				return
			}
			wg.Done()
		}(key, fn)
	}
	wg.Wait()

	// Get policies
	wg = &sync.WaitGroup{}
	for key, fn := range policyFuncMap {
		go func(key string, fn func() (map[string]interface{}, error)) {
			wg.Add(1)

			startTime := time.Now()
			_, err := fn()
			endTime := time.Now()

			processSec := endTime.Sub(startTime).Seconds()
			if processSec > 2.0 {
				fmt.Printf("Took %g sec to get %ss in 'Services' page\n", processSec, key)
			}

			if err != nil {
				fmt.Println(err)
				return
			}
			wg.Done()
		}(key, fn)
	}
	wg.Wait()
}
