package migrate

import "flag"

func parseFlag() (command string, filename string) {
	flag.StringVar(&command, "c", "", helpCommand)
	flag.StringVar(&filename, "n", "", "migrate file name")

	flag.Parse()

	return command, filename
}
