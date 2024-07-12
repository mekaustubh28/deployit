package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
)

const ShellToUse = "bash"

func cmd(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func contains(slice []fs.DirEntry, item string) bool {
	for _, str := range slice {
		if str.Name() == item {
			return true
		}
	}
	return false
}

func building(id string) {
	source_path := "../repos/" + id
	dest_path := "../build/" + id
	files, _ := os.ReadDir(source_path)

	redis_push(id+": Building in progress...", "log:"+id)

	if contains(files, "package.json") {
		out, errout, err := cmd("cd .. && cd repos/" + id + " && npm install && npm run build")

		redis_push(id+": Running npm install && npm run build", "log:"+id)

		if err != nil {
			log.Printf("error: %v\n", err)
			redis_push(id+": Error while building...: "+err.Error(), "log:"+id)
			return
		}
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		redis_push(id+" : "+out, "log:"+id)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
		redis_push(id+" : "+errout, "log:"+id)

		files_after_build, _ := os.ReadDir(source_path)

		if contains(files_after_build, "build") {
			source_path += "/build"
		} else {
			source_path += "/dist"
		}
	}
	redis_push(id+": Building done.", "log:"+id)
	redis_push(id+": Copying Built Folder in progress...", "log:"+id)
	copy_for_serving(source_path, dest_path, id)
}
