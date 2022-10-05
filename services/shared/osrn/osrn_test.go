package osrn_test

import (
	"testing"

	"go.h4n.io/openschool/services/shared/osrn"
)

func TestOsrnParse(t *testing.T) {
	classOsrn := osrn.OSRN{}
	classStr := "osrn:1:class:ch72gsb320000udocl363eofy"

	teacherOsrn := osrn.OSRN{}
	teacherStr := "osrn:2:people:teacher/ch72gsb320000udocl363eofy"

	if err := classOsrn.Parse(classStr); err != nil {
		t.Error(err.Error())
	} else {
		if classOsrn.Version != osrn.OSRNVersion1 {
			t.Error("osrn parsed invalid version")
		} else if classOsrn.Service != "class" {
			t.Error("osrn parsed invalid service")
		} else if classOsrn.Type != "" {
			t.Error("osrn parsed a type")
		} else if classOsrn.Id != "ch72gsb320000udocl363eofy" {
			t.Error("osrn parsed invalid id")
		}
	}

	if err := teacherOsrn.Parse(teacherStr); err != nil {
		t.Error(err.Error())
	} else {
		if teacherOsrn.Version != osrn.OSRNVersion2 {
			t.Error("osrn parsed invalid version")
		} else if teacherOsrn.Service != "people" {
			t.Error("osrn parsed invalid service")
		} else if teacherOsrn.Type != "teacher" {
			t.Error("osrn parsed invalid type")
		} else if teacherOsrn.Id != "ch72gsb320000udocl363eofy" {
			t.Error("osrn parsed invalid id")
		}
	}
}
