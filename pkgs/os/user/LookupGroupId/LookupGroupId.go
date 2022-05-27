package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

var sysUsers map[string]string

func init() {
	sysUsers = mapUsersId()
}

// mapFile retorna um map[int]string mapeando o arquivo de nome "fileName"
// com o separador de campos ":". iIndex e iName são os índices dos elementos
// string:string do map no arquivo em filename.
func mapFile(fileName, sep string, iIndex, iName int) map[string]string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	strData := strings.Split(string(data), "\n")
	outMap := map[string]string{}

	for _, line := range strData {
		split := strings.Split(line, sep)
		if len(split) > 2 {
			id := split[iIndex]
			outMap[id] = split[iName]
		}
	}
	return outMap
}

func mapUsersId() map[string]string {
	return mapFile("/etc/passwd", ":", 2, 0)
}

func id2GroupName(ids []string) (gNames []string) {
	for _, gid := range ids {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			log.Fatal(err)
		}
		gNames = append(gNames, group.Name)
	}
	return
}

func main() {
	for uid, username := range sysUsers {
		if len(uid) == 4 && uid >= "1000" && uid < "2000" {
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
