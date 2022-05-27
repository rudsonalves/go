package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

var groups map[int]string
var sysUsers map[int]string

func init() {
	groups = mapGroupId()
	sysUsers = mapUsersId()
}

// mapFile retorna um map[int]string mapeando o arquivo de nome "fileName"
// com o separador de campos ":". iInt e iStr são os índices dos elementos
// int:string do map no arquivo em filename.
func mapFile(fileName string, sep string, iInt int, iStr int) map[int]string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	strData := strings.Split(string(data), "\n")
	outMap := map[int]string{}

	for _, line := range strData {
		split := strings.Split(line, sep)
		if len(split) > 2 {
			id, err := strconv.Atoi(split[iInt])
			if err == nil {
				outMap[id] = split[iStr]
			}
		}
	}
	return outMap
}

func mapGroupId() map[int]string {
	return mapFile("/etc/group", ":", 2, 0)
}

func mapUsersId() map[int]string {
	return mapFile("/etc/passwd", ":", 2, 0)
}

func id2GroupName(ids []string) (gNames []string) {
	for _, sgid := range ids {
		gid, _ := strconv.Atoi(sgid)
		gNames = append(gNames, groups[gid])
	}
	return
}

func main() {
	for uid, username := range sysUsers {
		if uid >= 1000 && uid < 2000 {
			u, err := user.Lookup(username)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("    Name: %v\n", u.Name)
			fmt.Printf(" HomeDir: %v\n", u.HomeDir)
			fmt.Printf("     Uid: %v\n", u.Uid)
			fmt.Printf("Username: %v\n", u.Username)
			sgid, _ := u.GroupIds()
			myGroups := id2GroupName(sgid)
			fmt.Printf("GroupIds: %v\n\n", myGroups)
		}
	}
}
