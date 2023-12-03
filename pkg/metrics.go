package pkg

import (
	"sort"
)

// updateDomainUsage update the number of times domain have been used for shortening
func updateDomainUsage(originalURL string) {
	domain := getDomainFromURL(originalURL)
	domainsUsageMap[domain]++
}

// getTopDomainsUsageMetrics return top 3 domains which have used the shortening service
// the most
func getTopDomainsUsageMetrics() TopDomainMetrics {
	// sorting the map based on usage count value and returning top 3

	domains := make([]string, 0, len(domainsUsageMap))
	for domain := range domainsUsageMap {
		domains = append(domains, domain)
	}

	sort.SliceStable(domains, func(i, j int) bool {
		return domainsUsageMap[domains[i]] > domainsUsageMap[domains[j]]
	})

	topDomainMetrics := TopDomainMetrics{}
	// Take top 3 used domains or less than that if total number of domains less then 3
	for i := 0; i < len(domains) && i < 3; i++ {
		domainUsage := DomainUsage{
			Domain: domains[i],
			Usage:  domainsUsageMap[domains[i]],
		}
		topDomainMetrics = append(topDomainMetrics, domainUsage)
	}
	return topDomainMetrics
}
