package generator

import (
	"fmt"
	"sync"
	"time"

	"github.com/morrisson/load_gen/client"
)

type SitesLoadGenerator struct {
	GohanClient *client.GohanClient
}

func (gen *SitesLoadGenerator) Run() {
	t := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-t.C:
			go gen.run()
		}
	}
	t.Stop()
}

func (gen *SitesLoadGenerator) run() {
	siteGroups, err := gen.GohanClient.GetSiteGroups()
	if err != nil {
		fmt.Println(err)
		return
	}

	siteGroupList, ok := siteGroups["site_groups"].([]interface{})

	if ok {
		wg := &sync.WaitGroup{}
		for _, siteGroup := range siteGroupList {
			sg := siteGroup.(map[string]interface{})
			go func(siteGroupID string) {
				wg.Add(1)

				startTime := time.Now()
				_, err := gen.GohanClient.GetSitesOfSiteGroup(siteGroupID)
				endTime := time.Now()

				processSec := endTime.Sub(startTime).Seconds()
				if processSec > 2.0 {
					fmt.Printf("Took %g sec to get Sites from site group %s\n", processSec, siteGroupID)
				}

				if err != nil {
					fmt.Println(err)
					return
				}
				wg.Done()
			}(sg["id"].(string))
		}
		wg.Wait()
	}
}
