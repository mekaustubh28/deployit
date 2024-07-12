package main

import (
	"fmt"
	"os"
)

func delete_cloned_folder(id string) {
	path := "../repos/" + id
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error  while deleting : ", err)
	} else {
		fmt.Println("Directory", path, "removed successfully")
	}
}
