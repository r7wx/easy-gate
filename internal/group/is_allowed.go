package group

import "net"

// IsAllowed - Return true if incoming addr is allowed in one of the groups
func IsAllowed(groups []Group, allowedGroups []string, addr string) bool {
	if len(allowedGroups) == 0 {
		return true
	}

	for _, allowedGroup := range allowedGroups {
		for _, group := range groups {
			if group.Name != allowedGroup {
				continue
			}

			_, groupNet, err := net.ParseCIDR(group.Subnet)
			if err != nil {
				continue
			}

			ipAddr := net.ParseIP(addr)
			if groupNet.Contains(ipAddr) {
				return true
			}
		}
	}

	return false
}
