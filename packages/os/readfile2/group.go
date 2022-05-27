package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mapGroup() map[int]string {
	data, err := os.ReadFile("/etc/group")
	if err != nil {
		log.Fatal(err)
	}
	etcContents := strings.Split(string(data), "\n")

	groups := map[int]string{}
	for _, line := range etcContents {
		split := strings.Split(line, ":")
		if len(split) > 2 {
			index, err := strconv.Atoi(split[2])
			if err == nil {
				groups[index] = split[0]
			}
		}
	}

	return groups
}

func main() {
	groups := mapGroup()

	userId := os.Getuid()
	groupId := os.Getgid()

	groupsId, _ := os.Getgroups()
	groupsStr := []string{}
	for _, g := range groupsId {
		groupsStr = append(groupsStr, groups[g])
	}

	fmt.Printf(" Username: %s\n", groups[userId])
	fmt.Printf("    Group: %s\n", groups[groupId])
	fmt.Printf("   Groups: %v\n", groupsStr)
}
