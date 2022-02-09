package models

import "time"

type Flags struct {
	All        bool
	Reverse    bool
	Recursive  bool
	Long       bool
	SortByTime bool
}

type Entity struct {
	Name        string
	Permissions string
	HardLinks   uint64
	UserOwner   string
	GroupOwner  string
	ModTime     time.Time
	Children    []Entity
}

const (
	IncorrectFlags = "incorrect flag: \"%s\"\nsupported flags: -a, -r, -R, -l, -t"
	NotExist       = "\"%s\" does not exist"
	DotCharacter   = 46
)
