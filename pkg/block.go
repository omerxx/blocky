package block

import (
	"fmt"
	"net/http"

	"github.com/txn2/txeh"
)

func BlockSites(hosts *txeh.Hosts, sites []string) (blacklistConfiguredSites []string) {
	for _, s := range sites {
		blacklistConfiguredSites = append(blacklistConfiguredSites, s)
	}
	// TODO: this should be restored with state managment later!
	// sites = state.FilterStateListedSites(sites)

	for _, site := range sites {
		if isTargetAlive(site) {
			blockSite(hosts, site)
		} else {
			// TODO: restore state managment
			// blacklistConfiguredSites, _ = state.RemoveUrlFromList(blacklistConfiguredSites, site)
		}
	}
	return blacklistConfiguredSites
}

func isTargetAlive(url string) bool {
	resp, err := http.Get(fmt.Sprintf("https://%s", url))
	if err != nil {
		// fmt.Printf("Can't probe: %s\n", url)
		// TODO switch back --->
		// TODO
		return true
		// return false
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 500 {
		return true
	} else {
		fmt.Printf("Status code for %s is %d, Ignoring\n", url, resp.StatusCode)
		return false
	}
}

func blockSite(hosts *txeh.Hosts, site string) {
	// TODO: integrate Viper!
	// target := viper.GetString("app.blockRoute")
	target := "0.0.0.0"
	hosts.AddHost(target, site)
	hosts.Save()
}
