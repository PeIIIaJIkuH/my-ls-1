package parse

import (
	"errors"
	"fmt"
	"my-ls-1/config"
	"os"
)

type Flags struct {
	All        bool
	Reverse    bool
	Recursive  bool
	Long       bool
	SortByTime bool
}

func Args(args []string) ([]string, *Flags, error) {
	flags := Flags{}
	directoryNames := make([]string, 0)
	for _, arg := range args {
		if arg[0] == '-' {
			if len(arg) == 1 {
				return []string{}, &Flags{}, errors.New(fmt.Sprintf(config.IncorrectFlags, ""))
			}
			for _, c := range arg[1:] {
				switch c {
				case 'a':
					flags.All = true
				case 'r':
					flags.Reverse = true
				case 'R':
					flags.Recursive = true
				case 'l':
					flags.Long = true
				case 't':
					flags.SortByTime = true
				default:
					return []string{}, &Flags{}, errors.New(fmt.Sprintf(config.IncorrectFlags, string(c)))
				}
			}
		} else {
			directoryNames = append(directoryNames, arg)
		}
	}
	return directoryNames, &flags, nil
}

func Directories(filenames []string) error {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); err != nil {
			return errors.New(fmt.Sprintf(config.DirectoryDoesntExist, filename))
		}
	}
	return nil
}
