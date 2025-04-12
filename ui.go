package main

import (
	"fmt"
	"strings"
)

func padRight(str string, length int) string {
	format := fmt.Sprintf("%%-%ds", length)
	return fmt.Sprintf(format, str)
}

func drawSelected(title string, item fmt.Stringer) {
	l := max(len(item.String()), 5)
	fmt.Printf("┌%s%s┐\n", title, strings.Repeat("─", l-5))
	fmt.Printf("│ %s │\n", item)
	fmt.Printf("└%s┘\n", strings.Repeat("─", l+2))
}

func drawMenu[T fmt.Stringer](title string, items []T, selected int) {
	boxWidth := 5
	for _, item := range items {
		if len(item.String()) > boxWidth {
			boxWidth = len(item.String())
		}
	}
	boxWidth += 5

	fmt.Printf("┌%s%s┐\n", title, strings.Repeat("─", boxWidth-2-len(title)))

	start := max(selected-5, 0)
	end := min(start+10, len(items)-1)

	if end == len(items)-1 {
		start = max(end-10, 0)
	}

	if start != 0 {
		fmt.Printf("│ %s▲%s│\n", strings.Repeat(" ", (boxWidth-4)/2), strings.Repeat(" ", (boxWidth-3)/2))
	} else {
		fmt.Printf("│%s│\n", strings.Repeat(" ", boxWidth-2))
	}

	for i := start; i <= end; i++ {
		item := padRight(items[i].String(), boxWidth-4)
		fmt.Printf("│ ")
		if i == selected {
			swapColors()
		}
		fmt.Printf("%s", item)
		if i == selected {
			resetColors()
		}
		fmt.Printf(" │\n")
	}

	if end != len(items)-1 {
		fmt.Printf("│ %s▼%s│\n", strings.Repeat(" ", (boxWidth-4)/2), strings.Repeat(" ", (boxWidth-3)/2))
	} else {
		fmt.Printf("│%s│\n", strings.Repeat(" ", boxWidth-2))
	}

	fmt.Printf("└%s┘\n", strings.Repeat("─", boxWidth-2))
}
