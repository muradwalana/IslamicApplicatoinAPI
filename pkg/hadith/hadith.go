package hadith

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"

	"islamicprojectapi/pkg/db"
)

// HadithHandler returns a hadith from a collection
// URL: /api/hadith/:collection/:number
func HadithHandler(c *gin.Context) {
	collection := c.Param("collection")
	numStr := c.Param("number")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hadith number"})
		return
	}
	h, err := db.GetHadith(collection, num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hadith not found"})
		return
	}
	c.JSON(http.StatusOK, h)
}

// ListHadithsHandler returns all available hadiths
func ListHadithsHandler(c *gin.Context) {
	list, err := db.GetAllHadiths()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load hadiths"})
		return
	}
	// defensive sort: by collection then hadith number
	sort.Slice(list, func(i, j int) bool {
		if list[i].Collection == list[j].Collection {
			return list[i].HadithNumber < list[j].HadithNumber
		}
		return list[i].Collection < list[j].Collection
	})
	c.JSON(http.StatusOK, list)
}
