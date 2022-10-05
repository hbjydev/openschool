package osrn

import (
	"fmt"
	"regexp"
)

var (
	OSRNRegex *regexp.Regexp = regexp.MustCompile(
		`^osrn:(?P<service>[a-z]+):(?P<resource>[a-z0-9]+)?:(?P<resourceId>[a-z0-9]+)$`,
	)
)

// OSRN represents an OpenSchool Resource Name, documented in the architecture
// document.
type OSRN struct {
	// Service is the name of the service this OSRN belongs to.
	Service string

	// Type is the type of resource this OSRN represents on the service.
	// Only found on OSRN v2
	Type string

	// Id is the unique ID of the resource this OSRN defines.
	Id string
}

func (o OSRN) String() string {
	osrn := fmt.Sprintf("osrn:%v:%v:%v", o.Service, o.Type, o.Id)
	return osrn
}

// Parse loads an OSRN into the struct given a string.
//
//	o := OSRN{}
//	s := "osrn:class:ch72gsb320000udocl363eofy"
//	o.Parse(s)
func (o *OSRN) Parse(input string) error {
	parts := OSRNRegex.FindAllStringSubmatch(input, -1)
	o.Service = parts[0][1]
	o.Type = parts[0][2]
	o.Id = parts[0][3]
	return nil
}
