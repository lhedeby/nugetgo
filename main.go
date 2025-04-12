package main

import (
	"fmt"
)

func main() {
	projects, err := dotnetList()
	if err != nil {
		fmt.Println(err)
		return
	}
    updatedProjects, err := menu(projects)
	if err != nil {
		fmt.Println(err)
		return
	}
    dotnetUpdateProject(updatedProjects)
}
