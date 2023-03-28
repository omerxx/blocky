package block

import "github.com/txn2/txeh"

func CleanBlocks(hosts *txeh.Hosts, blacklistConfiguredSites []string) (cleanedTargets []string) {
	// TODO: restore state
	// stateSites := state.ListSites()
	stateSites := blacklistConfiguredSites

	for _, stateSite := range stateSites {
		// TODO: restore
		// exists := isInList(stateSite.Url, blacklistConfiguredSites)
		// if !exists {
		// hosts.RemoveHost(stateSite.Url)
		hosts.RemoveHost(stateSite)
		hosts.Save()
		// state.Remove(stateSite.Url)
		// cleanedTargets = append(cleanedTargets, stateSite.Url)
		cleanedTargets = append(cleanedTargets, stateSite)
		// }
	}
	return
}
