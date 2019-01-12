package client

import (
	"encoding/json"
	"net/http"
)

type GohanClient struct {
	Endpoint   string
	Token      string
	HttpClient *http.Client
}

func (g *GohanClient) GetSiteGroups() (res map[string]interface{}, err error) {
	return g.get("/v1.0/site_groups", nil)
}

func (g *GohanClient) GetSitesOfSiteGroup(siteGroupID string) (res map[string]interface{}, err error) {
	queryStrings := map[string]string{
		"site_group_id": siteGroupID,
	}
	return g.get("/v1.0/sites", queryStrings)
}

func (g *GohanClient) GetServices() (res map[string]interface{}, err error) {
	return g.get("/v1.0/services", nil)
}

func (g *GohanClient) GetServiceTemplates() (res map[string]interface{}, err error) {
	return g.get("/v1.0/service_templates", nil)
}

func (g *GohanClient) GetApplicationGatewayPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/application_gateway_policies", nil)
}

func (g *GohanClient) GetFirewallPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/firewall_policies", nil)
}

func (g *GohanClient) GetHybridWanPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/hybrid_wan_policies", nil)
}

func (g *GohanClient) GetMicrosegmentationPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/microsegmentation_policies", nil)
}

func (g *GohanClient) GetQosPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/qos_policies", nil)
}

func (g *GohanClient) GetSnmpPolicies() (res map[string]interface{}, err error) {
	return g.get("/v1.0/snmp_policies", nil)
}

func (g *GohanClient) GetTrafficClasses() (res map[string]interface{}, err error) {
	return g.get("/v1.0/traffic_classes", nil)
}

func (g *GohanClient) GetTrafficApplications() (res map[string]interface{}, err error) {
	return g.get("/v1.0/traffic_applications", nil)
}

func (g *GohanClient) GetPathSelectionProfiles() (res map[string]interface{}, err error) {
	return g.get("/v1.0/path_selection_profiles", nil)
}

func (g *GohanClient) GetIPFilters() (res map[string]interface{}, err error) {
	return g.get("/v1.0/ipfilters", nil)
}

func (g *GohanClient) GetWebFilters() (res map[string]interface{}, err error) {
	return g.get("/v1.0/webfilters", nil)
}

func (g *GohanClient) get(path string, queryStrings map[string]string) (res map[string]interface{}, err error) {
	reqURL := g.Endpoint + path
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("X-Auth-Token", g.Token)
	if queryStrings != nil {
		q := req.URL.Query()
		for key, value := range queryStrings {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	return
}
