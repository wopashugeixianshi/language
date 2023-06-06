package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	oracleReleaseRegexp = regexp.MustCompile(`(?P<os>[^\s]*) (Linux Server release) (?P<version>[\d]+)`)
	centosReleaseRegexp = regexp.MustCompile(`(?P<os>[^\s]*) (Linux release|release) (?P<version>[\d]+)`)
	//centosReleaseRegexp = regexp.MustCompile(`(?P<os>[^\s]*) (Linux release) (?P<version>[\d]+)`)
	redhatReleaseRegexp = regexp.MustCompile(`(?P<os>Red Hat Enterprise Linux) (Client release|Server release|Workstation release) (?P<version>[\d]+)`)

	osReleaseOSRegexp      = regexp.MustCompile(`^ID=(.*)`)
	osReleaseVersionRegexp = regexp.MustCompile(`^VERSION_ID=(.*)`)
)

func main() {

	var str = "  BigCloud Enterprise Linux For LDK release 7.6.1906 (core)"
	//var str1 = "CentOS Linux release 7.6.1810 (Core)"
	r := redhatReleaseRegexp.FindStringSubmatch(str)
	fmt.Println(len(r))

	r = oracleReleaseRegexp.FindStringSubmatch(str)
	fmt.Println(len(r))
	r = centosReleaseRegexp.FindStringSubmatch(str)
	fmt.Println(len(r))
	if len(r) == 4 {
		if strings.Contains(str, "BigCloud") {
			fmt.Println("bclinux", r[3])
			fmt.Println(str)
		}
	}
	fmt.Println("----------------------------------------")
	for i, v := range r {
		fmt.Println(i, v)
	}

	var str2 = "VERSION_ID=\"8.2\""
	r = osReleaseVersionRegexp.FindStringSubmatch(str2)
	fmt.Println("----------------------------------------")
	for i, v := range r {
		fmt.Println(i, v)
	}

}
