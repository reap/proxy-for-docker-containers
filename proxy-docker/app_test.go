package main

import "testing"

func TestSplitServiceUrlReturnsServicePart(t *testing.T) {
	service, _ := SplitServiceURL("/service-part/rest/of/url")
	if service != "service-part" {
		t.Errorf("Expected 'service-part', but was %s", service)
	}
}

func TestSplitServiceUrlReturnsUrlPart(t *testing.T) {
	_, url := SplitServiceURL("/service-part/rest/of/url")
	if url != "rest/of/url" {
		t.Errorf("Expected 'rest/of/url', but was %s", url)
	}
}
