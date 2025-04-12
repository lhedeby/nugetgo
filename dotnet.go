package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func dotnetList() ([]Project, error) {
	cmd := exec.Command("dotnet", "list", "package", "--format", "json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var list DotnetList
	if err := json.Unmarshal(out, &list); err != nil {
		return nil, err
	}
	return list.Projects, nil
}

func dotnetUpdateProject(projects []Project) {
	for _, proj := range projects {
		printProjName := true
		for _, pkg := range proj.Frameworks[0].TopLevelPackages {
			if pkg.NewVersion != "" && pkg.NewVersion != pkg.RequestedVersion {
				if printProjName {
					fmt.Println("Updating packages in", proj.Path)
					printProjName = false
				}
				cmd := exec.Command("dotnet", "add", proj.Path, "package", pkg.Id, "--version", pkg.NewVersion)
				out, err := cmd.Output()
				if err != nil {
					fmt.Println("✖  Error ", pkg.Id, pkg.RequestedVersion, "->", pkg.NewVersion, out, err)
				} else {
					fmt.Println("✔  ", pkg.Id, pkg.RequestedVersion, "->", pkg.NewVersion)
				}
			}
		}
	}
}
