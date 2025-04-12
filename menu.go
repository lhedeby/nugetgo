package main

import (
	"fmt"
)

func menu(projects []Project) ([]Project, error) {
	enterAlternativeScreen()
	defer exitAlternativeScreen()

	selection := Selection{0, 0, 0}
	includePreRelease := false

	var project *Project
	var pkg *TopLevelPackage

	for {
		clear()
		if project == nil {
			drawMenu("Project", projects, selection.Project)
		} else if pkg == nil {
			drawSelected("Project", project)
			drawMenu("Package", project.Frameworks[0].TopLevelPackages, selection.Pkg)
		} else {
			drawSelected("Project", project)
			drawSelected("Package", pkg)
			drawMenu("Version", pkg.Versions, selection.Version)
			box := "☐"
			if includePreRelease {
				box = "☑"
			}
			fmt.Println("\n", box, "Include (P)rereleases")
		}

		key, err := readKey()
		if err != nil {
			panic("key error")
		}

		switch key {
		case 'j', '↓':
			if project == nil {
				move(&selection.Project, 1, 0, len(projects)-1)
			} else if pkg == nil {
				move(&selection.Pkg, 1, 0, len(project.Frameworks[0].TopLevelPackages)-1)
			} else {
				move(&selection.Version, 1, 0, len(pkg.Versions)-1)
			}
		case 'k', '↑':
			if project == nil {
				move(&selection.Project, -1, 0, len(projects)-1)
			} else if pkg == nil {
				move(&selection.Pkg, -1, 0, len(project.Frameworks[0].TopLevelPackages)-1)
			} else {
				move(&selection.Version, -1, 0, len(pkg.Versions)-1)
			}
		case 'l', '\n':
			if project == nil {
				project = &projects[selection.Project]
				selection.Pkg = 0
			} else if pkg == nil {
				pkg = &project.Frameworks[0].TopLevelPackages[selection.Pkg]
				if len(pkg.AllVersions) == 0 {
					pkg.AllVersions = getVersions(pkg.Id)
					pkg.FilteredVersions = filterContains(pkg.AllVersions, "-")
				}
				if includePreRelease {
					pkg.Versions = pkg.AllVersions
				} else {
					pkg.Versions = pkg.FilteredVersions
				}
				selection.Version = 0
			} else {
				pkg.NewVersion = pkg.Versions[selection.Version].String()
			}
		case 'h':
			if pkg == nil {
				project = nil
			}
			pkg = nil
		case 'p':
			if pkg != nil {
				includePreRelease = !includePreRelease
				selected := pkg.Versions[selection.Version]
				if includePreRelease {
					pkg.Versions = pkg.AllVersions
				} else {
					pkg.Versions = pkg.FilteredVersions
				}

				selection.Version = 0
				for i, v := range pkg.Versions {
					if v == selected {
						selection.Version = i
						break
					}
				}
			}
		case 'q':
			return projects, nil
		}
	}
}
