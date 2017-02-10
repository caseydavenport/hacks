package main

import (
	"fmt"
	"net"
)

func main() {
	// Build up some IP subnetworks.
	_, a, _ := net.ParseCIDR("192.168.254.3/32")
	_, b, _ := net.ParseCIDR("192.168.2.50/32")
	_, c, _ := net.ParseCIDR("192.168.200.1/32")
	_, d, _ := net.ParseCIDR("192.169.3.2/32")
	ips := []*net.IPNet{a, b, c, d}

	// Build the common supernet.
	common := SuperNet(ips)
	fmt.Printf("%s\n", common)

	// Assert that all the inputs are actually contained within the supernet.
	for _, ip := range ips {
		if !common.Contains(ip.IP) {
			fmt.Printf("DOESN'T CONTAIN %s\n", ip)
		} else {
			fmt.Printf("contains %s\n", ip)
		}
	}
}

// SuperNet returns the smaller supernet of the given subnets.
func SuperNet(ips []*net.IPNet) *net.IPNet {
	common := ips[0]
	for _, ip := range ips {
		fmt.Printf("Comparing %s and %s\n", common, ip)
		prefix := superNet(common, ip)
		_, common, _ = net.ParseCIDR(fmt.Sprintf("%s/%d", ip.IP, prefix))

	}
	return common
}

func superNet(a, b *net.IPNet) int {
	l, _ := a.Mask.Size()
	m, _ := b.Mask.Size()
	maxPre := max(l, m)
	var numMatches uint32 = 0
	for i, byte := range a.IP.To4() {
		byte2 := b.IP.To4()[i]

		for j := 7; j >= 0; j-- {
			// fmt.Printf("matches: %d\n", numMatches)
			fmt.Printf("New bit comparison, 1 << %d\n", (j))

			bit1 := byte & (1 << uint32(j))
			bit2 := byte2 & (1 << uint32(j))
			fmt.Printf("  Bytes: %8b, %8b\n", int32(byte), int32(byte2))
			fmt.Printf("  Bits : %8b, %8b\n", bit1, bit2)
			if bit1 == bit2 && int(numMatches+1) < maxPre {
				numMatches++
			} else {
				return int(numMatches)
			}

		}
	}
	return int(numMatches)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
