package utils

import "fmt"

func AddWww(sites []string) (updatedsites []string) {
	updatedsites = sites
	for _, site := range sites {
		prefixedSite := fmt.Sprintf("www.%s", site)
		updatedsites = append(updatedsites, prefixedSite)
	}
	return
}
