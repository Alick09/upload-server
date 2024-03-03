package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
	maxMb int64
	token string
}

func loadConfig() Config {
	result := Config{maxMb: 32, token: ""};

	maxMbStr, exists := os.LookupEnv("MAX_MB")
	if exists {
		if maxMb, err := strconv.ParseInt(maxMbStr, 10, 64); err == nil && maxMb > 0 {
			result.maxMb = maxMb;
			log.Printf("File size limit was set to %d MB", maxMb)
		} else {
			log.Printf("Incorrect MAX_MB param, should be positive integer, found %s", maxMbStr)
		}
	}

	token, exists := os.LookupEnv("TOKEN")
	if exists {
		result.token = token;
		log.Print("Token was set")
	} else {
		log.Print("Token wasn't set")
	}
	
	return result
}

var config Config

func main() {
    router := gin.Default()
	config = loadConfig()
	router.MaxMultipartMemory = config.maxMb << 20
    router.POST("/upload", uploadFiles)
    router.Run()
}

func fetchPathPrefix(pathParam []string) string {
	if len(pathParam) == 0 {
		return ""
	} else {
		return pathParam[0] + "/"
	}
}

func uploadFiles(c *gin.Context) {
	if config.token != "" {
		token := c.Request.Header["Token"]
		if (len(token) != 1 || token[0] != config.token){
			c.String(http.StatusForbidden, "Not allowed")
			return
		}
	}
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	path := fetchPathPrefix(form.Value["path"])
	for _, file := range files {
		fullPath := fmt.Sprintf("/data/%s%s", path, file.Filename)
		err := c.SaveUploadedFile(file, fullPath)
		if err != nil {
			log.Printf("Couldn't upload file to %s", fullPath)
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}