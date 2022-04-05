# packer

[![Go Report Card](https://goreportcard.com/badge/github.com/makifdb/packer)](https://goreportcard.com/report/github.com/makifdb/packer)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/makifdb/packer.svg)](https://github.com/makifdb/packer)
[![License: Apache-2.0](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/makifdb/packer/blob/main/LICENSE)

<img align="right" width="328" alt="icon" src="https://user-images.githubusercontent.com/31243845/161521421-ca0328fd-9395-47c2-8f0d-b348a89c09db.png">

Packer is a simple package management tool for Go. Packer, helps you detect the system package manager and automate install your package or dependencies inside Go.

| Operation Systems   | Package Managers |
|---------------------|------------------|
| Ubuntu              | apk 		 	 | 
| Debian              | apt 		     | 
| MXLinux             | brew         	 | 
| Mint                | dnf        		 | 
| Kali                | flatpak          | 
| ParrotOS            | snap             |
| OpenSUSE Leap       | pacman           | 
| OpenSUSE TumbleWeed | paru             |
| OpenSUSE SLES       | yay              |
| Arch                | zypper           |
| Manjaro             |                  |
| Alpine              |                  |
| Fedora              |                  |
| RHEL                |                  |
| CentOS              |                  |
| Oracle              |                  |
| MacOS               |                  |

## Example Usage `import "github.com/makifdb/packer"`

1. Check package installation

```go
func main() {
	p:= packer.Check("curl")
	fmt.Println(p)
}
// output: true
```

2. Install package

```go
func main() {
	packer.Install("curl")
}
```

3. Remove package

```go
func main() {
	packer.Remove("curl")
}
```


4. Update system

```go
func main() {	
	packer.Update()
}
```

5. Dedect Package Manager

```go
func main() {	
	mngr, _ := DedectManager()
	fmt.Println(mngr.Name)
}
// output: yay

```
