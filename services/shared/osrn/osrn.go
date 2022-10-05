package osrn

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// OSRNVersion represents the format of the OSRN
type OSRNVersion string

const (
	// OSRNVersion1 is an OSRN v1, which explicitly does not include a resource
	// type.
	OSRNVersion1 OSRNVersion = "1"

	// OSRNVersion2 is an OSRN v2, which explicitly includes a resource type.
	OSRNVersion2 OSRNVersion = "2"
)

var (
	OSRNVersionRegex *regexp.Regexp = regexp.MustCompile(string(OSRNVersion1) + "|" + string(OSRNVersion2))
)

// OSRN represents an OpenSchool Resource Name, documented in the architecture
// document.
type OSRN struct {
	// Version is the format to use for the OSRN.
	Version OSRNVersion

	// Service is the name of the service this OSRN belongs to.
	Service string

	// Type is the type of resource this OSRN represents on the service.
	// Only found on OSRN v2
	Type string

	// Id is the unique ID of the resource this OSRN defines.
	Id string
}

func (o OSRN) String() (string, error) {
	if o.Service == "" {
		return "", errors.New("invalid osrn: no Service set")
	}
	if o.Id == "" {
		return "", errors.New("invalid osrn: no Id set")
	}

	var osrn string
	base := fmt.Sprintf("osrn:%v", o.Version)

	if o.Version == OSRNVersion1 {
		osrn = fmt.Sprintf("%v:%v", base, o.Id)
	} else if o.Version == OSRNVersion2 {
		osrn = fmt.Sprintf("%v:%v:%v", base, o.Type, o.Id)
	} else {
		return "", errors.New("invalid osrn: no Version set")
	}

	return osrn, nil
}

// Parse loads an OSRN into the struct given a string.
//
//	o := OSRN{}
//	s := "osrn:1:class:ch72gsb320000udocl363eofy"
//	o.Parse(s)
func (o *OSRN) Parse(input string) error {
	parts := strings.Split(input, ":")

	if len(parts) != 4 {
		return errors.New("invalid osrn: must specify osrn, version, service, and id")
	}

	if parts[0] != "osrn" {
		return errors.New("invalid osrn: must specify osrn")
	}

	if !OSRNVersionRegex.Match([]byte(parts[1])) {
		return errors.New("invalid osrn: invalid version")
	}

	if parts[2] == "" {
		return errors.New("invalid osrn: no service specified")
	}

	if parts[1] == "1" {
		if len(parts) != 4 {
			return errors.New("invalid osrn: too many parts specified, must only specify osrn, version, service and id")
		} else if parts[3] == "" {
			return errors.New("invalid osrn: no id specified")
		}

		o.Version = OSRNVersion1
		o.Id = parts[3]
	} else if parts[1] == "2" {
		subParts := strings.Split(parts[3], "/")
		if len(subParts) != 2 {
			return errors.New("invalid osrn: must specify osrn, version, service, resource type, and id")
		} else if subParts[0] == "" {
			return errors.New("invalid osrn: no resource type specified")
		} else if subParts[1] == "" {
			return errors.New("invalid osrn: no id specified")
		}

		o.Version = OSRNVersion2
		o.Type = subParts[0]
		o.Id = subParts[1]
	}

	o.Service = parts[2]

	return nil
}
