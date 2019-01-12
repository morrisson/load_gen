package generator

import (
	"fmt"
	"sync"
	"time"

	"github.com/morrisson/load_gen/client"
)

type PoliciesLoadGenerator struct {
	GohanClient *client.GohanClient
}

func (gen *PoliciesLoadGenerator) Run() {
	t := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-t.C:
			go gen.run()
		}
	}
	t.Stop()
}

func (gen *PoliciesLoadGenerator) run() {
	funcMap := map[string](func() (map[string]interface{}, error)){
		"application_gateway_policy": gen.GohanClient.GetApplicationGatewayPolicies,
		"firewall_policy":            gen.GohanClient.GetFirewallPolicies,
		"hybrid_wan_policy":          gen.GohanClient.GetHybridWanPolicies,
		"microsegmentation_policy":   gen.GohanClient.GetMicrosegmentationPolicies,
		"qos_policy":                 gen.GohanClient.GetQosPolicies,
		"snmp_policy":                gen.GohanClient.GetSnmpPolicies,
		"traffic_classes":            gen.GohanClient.GetTrafficClasses,
		"traffic_applications":       gen.GohanClient.GetTrafficApplications,
		"path_selection_profiles":    gen.GohanClient.GetPathSelectionProfiles,
		"ipfilters":                  gen.GohanClient.GetIPFilters,
		"webfilters":                 gen.GohanClient.GetWebFilters,
	}

	// Get policies
	wg := &sync.WaitGroup{}
	for key, fn := range funcMap {
		go func(key string, fn func() (map[string]interface{}, error)) {
			wg.Add(1)

			startTime := time.Now()
			_, err := fn()
			endTime := time.Now()

			processSec := endTime.Sub(startTime).Seconds()
			if processSec > 2.0 {
				fmt.Printf("Took %g sec to get %ss in 'Service Policies' page\n", processSec, key)
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
