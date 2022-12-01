package GoSNMPServer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/slayercat/gosnmp"
)

var oidPattern = regexp.MustCompile(`^(?:\d+\.?)+\d+$`)

func getPktContextOrCommunity(i *gosnmp.SnmpPacket) string {
	if i.Version == gosnmp.Version3 {
		return i.ContextName
	} else {
		return i.Community
	}
}

func copySnmpPacket(i *gosnmp.SnmpPacket) gosnmp.SnmpPacket {
	var ret gosnmp.SnmpPacket = *i
	if i.SecurityParameters != nil {
		ret.SecurityParameters = i.SecurityParameters.Copy()
	}
	return ret
}

// oidLess returns true when the first OID (a) is smaller than the second OID (b)
// oidLess returns true when the first OID (a) is smaller than the second OID (b)
func oidLess(a, b string) bool {
	aa := strings.Split(a, ".")
	bb := strings.Split(b, ".")

	for i, sa := range aa {
		if i < len(bb) {
			sb := bb[i]
			if sa != sb {
				ia, err := strconv.Atoi(sa)
				if err != nil {
					panic(fmt.Errorf("invalid OID %v: %v", a, err))
				}
				ib, err := strconv.Atoi(sb)
				if err != nil {
					panic(fmt.Errorf("invalid OID %v: %v", b, err))
				}
				return ia < ib // Return which segment is the smaller one
			}
			continue // Continue with next OID segment
		}
		return false // b shorter than a (e.g. a = 1.2.3 and b = 1.2)
	}
	return true // a shorter than b (e.g. a = 1.2 and b = 1.2.3)
}

// IsValidObjectIdentifier will check a oid string is valid oid
func IsValidObjectIdentifier(oid string) (result bool) {
	return oidPattern.MatchString(oid)
}
