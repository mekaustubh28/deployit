package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
)

func cloneRepo(c *gin.Context) {
	repo_username := c.Param("username")
	repo_name := c.Param("repo")
	repo_url := "https://github.com/" + repo_username + "/" + repo_name
	println(repo_url)

	id := redis_db_get(repo_url)

	if id == "" {
		id = give_id()
		redis_db_set(repo_url, id)
	}

	path := "../repos/" + id

	delete_cloned_folder(id)

	redis_push(id+": Repository cloning in progress...: "+repo_url, "log:"+id)

	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      repo_url,
		Progress: os.Stdout,
	})

	if err != nil {
		redis_push(id+": Failed to clone repository: "+err.Error(), "log:"+id)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error(), "id": id, "repo_url": repo_url})
		return
	}

	fmt.Printf("Repository cloned successfully! %s", repo)

	redis_push(id+": Repository cloned successfully: "+repo_url, "log:"+id)

	redis_push(id, "deploy")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "repository cloned", "id": id, "repo_url": repo_url})
}
