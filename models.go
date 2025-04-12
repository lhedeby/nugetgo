package main

import "strings"

type DotnetList struct {
	Version  int       `json:"version"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Path       string      `json:"path"`
	Frameworks []Framework `json:"frameworks"`
}

type Framework struct {
	Framework        string            `json:"framework"`
	TopLevelPackages []TopLevelPackage `json:"topLevelPackages"`
}

type TopLevelPackage struct {
	Id               string `json:"id"`
	RequestedVersion string `json:"requestedVersion"`
	NewVersion       string
	Versions         []Version
	AllVersions      []Version
	FilteredVersions []Version
}

type NuGetVersions struct {
	Versions []Version `json:"versions"`
}

type Selection struct {
	Project int
	Pkg     int
	Version int
}

type Version string

func (v Version) String() string {
	return string(v)
}

func (p Project) String() string {
	split := strings.Split(p.Path, "/")
	return split[len(split)-1]
}

func (t TopLevelPackage) String() string {
	if t.NewVersion != "" && t.NewVersion != t.RequestedVersion {
		return t.Id + " (" + t.RequestedVersion + " -> " + t.NewVersion + ")"
	}
	return t.Id + " (" + t.RequestedVersion + ")"
}
