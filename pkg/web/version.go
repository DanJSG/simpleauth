package web

import "fmt"

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v *Version) MajorMinorPatch() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) MajorMinor() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}
