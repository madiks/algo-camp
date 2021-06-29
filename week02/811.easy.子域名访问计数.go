package week02

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/subdomain-visit-count/
// https://leetcode-cn.com/submissions/detail/190774264/

func subdomainVisits(cpdomains []string) []string {
	domainStat := map[string]int{}
	for _, v := range cpdomains {
		domains, stat := parseDomainVisit(v)
		for _, d := range domains {
			domainStat[d] += stat
		}
	}
	ans := []string{}
	for k, v := range domainStat {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return ans
}

func parseDomainVisit(in string) (domains []string, visits int) {
	ins := strings.Split(in, " ")
	if len(ins) != 2 {
		return
	}
	visits, _ = strconv.Atoi(ins[0])
	fragments := strings.Split(ins[1], ".")
	for l := len(fragments) - 1; l >= 0; l-- {
		domains = append(domains, strings.Join(fragments[l:], "."))
	}
	return
}
