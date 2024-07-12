package main

import (
	"fmt"

	cp "github.com/otiai10/copy"
)

func copy_for_serving(source string, dest string, id string) {

	err := cp.Copy(source, dest)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(id, " : Folder moved.")

	delete_cloned_folder(id)

	redis_push(id+": Copying Built Folder done.", "log:"+id)
}
