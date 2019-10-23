package main

import "fmt"

type IPAddr [4]byte

func (i *IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main(){

    /**

    hosts := map[string]IPAddr {
        "loopback": {127, 0, 0, 1},
        "fuzhouDNS": {218, 85, 157, 99},
        "googleDNS": {8, 8, 8, 8},
        "telecomDNS": {114, 114, 114, 114},
    }
    */

    var hosts = make(map[string]*IPAddr)
    hosts["loopback"] = &IPAddr{127, 0, 0, 1}
    hosts["fuzhouDNS"] = &IPAddr{218, 85, 157, 99}
    hosts["googleDNS"] = &IPAddr{8, 8, 8, 8}
    hosts["telecomDNS"] = &IPAddr{114, 114, 114, 114}

    for name, ptr := range hosts {
        fmt.Printf("%v: %v \r\n", name, ptr)
    }
}

