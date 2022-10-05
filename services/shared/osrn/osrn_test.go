package osrn_test

import (
	"testing"

	"go.h4n.io/openschool/services/shared/osrn"
)

func TestOsrnParse(t *testing.T) {
	classOsrn := osrn.OSRN{}
	classStr := "osrn:class::ch72gsb320000udocl363eofy"

	teacherOsrn := osrn.OSRN{}
	teacherStr := "osrn:people:teacher:ch72gsb320000udocl363eofy"

	if err := classOsrn.Parse(classStr); err != nil {
		t.Error(err.Error())
	} else {
		if classOsrn.Service != "class" {
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
		if teacherOsrn.Service != "people" {
			t.Error("osrn parsed invalid service")
		} else if teacherOsrn.Type != "teacher" {
			t.Error("osrn parsed invalid type")
		} else if teacherOsrn.Id != "ch72gsb320000udocl363eofy" {
			t.Error("osrn parsed invalid id")
		}
	}
}

func TestOsrnString(t *testing.T) {
	classOsrn := osrn.OSRN{
		Service: `class`,
		Id:      `ch72gsb320000udocl363eofy`,
	}
	classExpected := "osrn:class::ch72gsb320000udocl363eofy"

	teacherOsrn := osrn.OSRN{
		Service: `people`,
		Type:    `teacher`,
		Id:      `ch72gsb320000udocl363eofy`,
	}
	teacherExpected := "osrn:people:teacher:ch72gsb320000udocl363eofy"

	classOut := classOsrn.String()
	teacherOut := teacherOsrn.String()

	if classOut != classExpected {
		t.Error("invalid osrn output")
	}

	if teacherOut != teacherExpected {
		t.Error("invalid osrn output")
	}
}
