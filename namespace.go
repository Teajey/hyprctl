package hyprctl

import (
	"encoding/xml"
	"fmt"
	"runtime/debug"
)

type Namespace struct {
	HcXmlns string      `xml:"xmlns:c,attr" json:"-"`
	Docs    xml.Comment `xml:",comment" json:"-"`
}

var docs xml.Comment

const repo string = "github.com/Teajey/hyprctl"

func getVersion() string {
	version := "main"
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return version
	}

	// 1. Check if we are running tests/builds INSIDE the hyprctl repo
	if info.Main.Path == repo {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	// 2. Check if we are a dependency of another project
	for _, dep := range info.Deps {
		if dep.Path == repo {
			if dep.Version != "" && dep.Version != "(devel)" {
				return dep.Version
			}
		}
	}

	return version
}

func init() {
	version := getVersion()
	docs = xml.Comment(fmt.Sprintf("See documentation for this version of hyprctl at https://%s/blob/%s/README.md ", repo, version))
}

func SetNamespace() Namespace {
	return Namespace{
		HcXmlns: "https://github.com/Teajey/hyprctl",
		Docs:    docs,
	}
}
