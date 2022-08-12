package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}

func Log(c *gin.Context) {

	w := bufio.NewWriter(c.Writer)
	newWriter := &bufferedWriter{c.Writer, w, bytes.Buffer{}}
	oldWriter := c.Writer
	c.Writer = newWriter

	defer func() {
		w.Flush()
		c.Writer = oldWriter
	}()

	c.Next()

	if c.Writer.Status() >= http.StatusBadRequest {
		log.Printf("error status: %d, body: %s", c.Writer.Status(), newWriter.Buffer.String())
		return
	}
	log.Printf("ok status: %d", c.Writer.Status())
}

func main() {
	router := gin.New()
	router.Use(Log)
	router.GET("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen and serve failed: ", err)
	}
}
