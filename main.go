/*
	Created by Bontus Mayor
	Contact <bontus.doku@gmail.com>
*/

package main

import (
	"Doku/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/index.html")
	r.Static("/dist", "./templates/dist")

	// Load manifest file
	manifest, err := utils.LoadManifest("templates/dist/.vite/manifest.json")
	if err != nil {
		log.Fatalf("could not load manifest file: %v", err)
	}

	r.GET("/", func(ctx *gin.Context) {
		entry, ok := manifest["index.html"].(map[string]interface{})
		if !ok {
			log.Fatalf("could not find entry for index.html in the manifest")
		}

		script, ok := entry["file"].(string)
		if !ok {
			log.Fatalf("could not find the script file in the manifest entry")
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"script": script,
		})
	})

	log.Fatal(r.Run(":8080"))
}
