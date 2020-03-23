package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.Static("/tmp1", "/tmp")
	router.StaticFS("/tmp2", http.Dir("/tmp"))
	router.StaticFile("/tmp3", "/tmp/tmp3")

	router.GET("file2", func(c *gin.Context) {
		var start, end int64
		fmt.Sscanf(c.GetHeader("Range"), "bytes=%d-%d", &start, &end)

		path := "/tmp/gordondemo"
		file, err := os.Open(path)
		if err != nil {
			log.Println(err.Error())
			c.Status(500)
			return
		}

		info, err := file.Stat()
		if end == 0 {
			end = info.Size() - 1
		}

		c.Header("Accept-ranges", "bytes")
		c.Header("Content-Length", strconv.FormatInt(end-start+1, 10))
		c.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(info.Size()-start, 10))
		c.Header("Content-Disposition", "attachment; filename="+info.Name())
		_, err = file.Seek(start, 0)
		if err != nil {
			log.Println(err.Error())
			c.Status(500)
			return
		}
		_, err = io.CopyN(c.Writer, file, end-start+1)
		if err != nil {
			log.Println(err.Error())
			return
		}

	})
	router.GET("/file", func(c *gin.Context) {

		data := map[string]interface{}{
			"lang": "55555555555",
			"tag":  "9999999",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}