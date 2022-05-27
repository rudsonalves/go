package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
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

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	groups := mapGroup()

	user, _ := user.Current()
	groupsId, _ := user.GroupIds()

	groupsStr := []string{}
	for _, g := range groupsId {
		groupsStr = append(groupsStr, groups[toInt(g)])
	}

	fmt.Printf(" Username: %s\n", user.Username)
	fmt.Printf("    Group: %s\n", groups[toInt(user.Gid)])
	fmt.Printf("     Home: %s\n", user.HomeDir)
	fmt.Printf("User Name: %s\n", user.Name)
	fmt.Printf("  User Id: %s\n", user.Uid)
	fmt.Printf("   Groups: %v\n", groupsStr)
}
