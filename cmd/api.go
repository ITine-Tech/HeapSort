package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"heapsort/internal"

)

type Request struct {
	Numbers []int `json:"numbers"`
}

func postHeapSort(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	if len(request.Numbers) == 0 {
		c.JSON(400, gin.H{"error": "numbers array cannot be empty"})
		return
	}

	if len(request.Numbers) == 10000 {
		c.JSON(400, gin.H{"error": "numbers array cannot exceed 10000 elements"})
		return
	}

	if err := internal.HeapSort(request.Numbers); err != nil {
		c.JSON(500, gin.H{"error": fmt.Errorf("heap sort failed: %v", err)})
		return
	}
	c.JSON(200, gin.H{"sorted_numbers": request.Numbers})
}
