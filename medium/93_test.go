package medium

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
	ip := "25525511135"
	t.Logf("%+v", RestoreIpAddresses(ip))
}
