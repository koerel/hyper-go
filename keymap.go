package main

func lookup(c string) string {
	switch c {
	case "$":
		return "shift+4"
	case ">":
		return "shift+period"
	case "-":
		return "minus"
	case "=":
		return "equal"
	}
	return c
}
