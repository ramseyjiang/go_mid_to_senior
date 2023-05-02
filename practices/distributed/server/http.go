package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Log *Log
}

func newHTTPServer() *HttpServer {
	return &HttpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func (s *HttpServer) WriteLog(c *gin.Context) {
	req := &ProduceRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Record.ClientIP = c.ClientIP()
	off, err := s.Log.Store(req.Record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProduceResponse{Offset: off})
}

func (s *HttpServer) ReadLog(c *gin.Context) {
	req := &ConsumeRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := s.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ConsumeResponse{Record: record})
}

func NewHTTPServer(addr string) *http.Server {
	s := newHTTPServer()         // Create a new HttpServer instance
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode
	r := gin.Default()
	r.POST("/", s.WriteLog) // Use the instance methods
	r.GET("/", s.ReadLog)   // Use the instance methods

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
