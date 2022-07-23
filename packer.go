package packer

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/JustinTimperio/osinfo"
)

type Manager struct {
	Name       string
	InstallArg string
	UpdateArg  string
	RemoveArg  string
}

var (
	apk     = Manager{"apk", "add", "update", "del"}
	apt     = Manager{"apt", "-y install", "update", "remove"}
	brew    = Manager{"brew", "install", "update", "remove"}
	dnf     = Manager{"dnf", "install", "upgrade", "erase"}
	flatpak = Manager{"flatpak", "install", "update", "uninstall"}
	snap    = Manager{"snap", "install", "upgrade", "remove"}
	pacman  = Manager{"pacman", "--noconfirm -S", "--noconfirm -Syuu", "--noconfirm -Rscn"}
	paru    = Manager{"paru", "-S", "-Syuu", "-R"}
	yay     = Manager{"yay", "-S", "-Syuu", "-R"}
	zypper  = Manager{"zypper", "-n install", "update", "-n remove"}
)

func DetectManager() (Manager, error) {
	switch opsystem := osinfo.GetVersion().Runtime; opsystem {
	default:
		// windows, freebsd, plan9 ...
		return empty, fmt.Errorf("%s is %w", opsystem, ErrUnsuppored)
	case "darwin":
		return brew, nil
	case "linux":
		switch distro := osinfo.GetVersion().Linux.Distro; distro {
		case "arch":
			if Check("pacman") {
				return pacman, nil
			}
			if Check("yay") {
				return yay, nil
			}
			if Check("paru") {
				return paru, nil
			}
		case "alpine":
			if Check("apk") {
				return apk, nil
			}
		case "fedora":
			if Check("dnf") {
				return dnf, nil
			}
		case "opensuse":
			if Check("zypper") {
				return zypper, nil
			}
		case "debian":
			if Check("apt") {
				return apt, nil
			}
			if Check("snap") {
				return snap, nil
			}
		default:
			if Check("snap") {
				return snap, nil
			}
			if Check("flatpak") {
				return flatpak, nil
			}
			return empty, ErrNotFound
		}
		return empty, ErrNotFound
	}
}

func Check(packageName string) bool {
	_, err := exec.LookPath(packageName)
	return err == nil
}

func Install(packageName string) error {
	mngr, err := Default()
	if err != nil {
		return err
	}

	c := mngr.Name + " " + mngr.InstallArg + " " + packageName
	err = Command(c)
	return err
}

func Remove(packageName string) error {
	mngr, err := Default()
	if err != nil {
		return err
	}

	c := mngr.Name + " " + mngr.RemoveArg + " " + packageName
	err = Command(c)
	return err
}

func Update() error {
	mngr, err := Default()
	if err != nil {
		return err
	}

	c := mngr.Name + " " + mngr.UpdateArg
	err = Command(c)
	return err
}

func Command(command string) error {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	err := cmd.Run()
	return err
}
