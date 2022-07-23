package packer

import (
	"fmt"

	"github.com/JustinTimperio/osinfo"
)

func DetectManager() (Manager, error) {
	osversion := osinfo.GetVersion()
	//fmt.Println(osversion.String())
	opsystem := osversion.Runtime
	mgrs, ok := managers[opsystem]
	if !ok {
		return empty, fmt.Errorf("%s is %w", opsystem, ErrUnsuppored)
	}
	distro := osversion.Linux.Distro
	list, ok := mgrs[distro]
	if !ok {
		if len(distro) == 0 {
			return empty, fmt.Errorf("%s is %w", opsystem, ErrUnsuppored)
		}
		list, ok = mgrs[""]
		if !ok {
			return empty, fmt.Errorf("%s %s is %w", opsystem, distro, ErrUnsuppored)
		}
	}
	for _, mgr := range list {
		if Check(mgr.Name) {
			return mgr, nil
		}
	}
	return empty, ErrNotFound
}
