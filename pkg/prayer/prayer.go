package prayer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"islamicprojectapi/pkg/db"
)

// PrayerTimesHandler returns prayer times for a location
// URL: /api/prayer/times/:location
func PrayerTimesHandler(c *gin.Context) {
	location := c.Param("location")
	p, err := db.GetPrayerTimes(location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "prayer times not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// ListLocationsHandler returns available prayer location slugs
func ListLocationsHandler(c *gin.Context) {
	list, err := db.GetAllPrayerLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load locations"})
		return
	}
	c.JSON(http.StatusOK, list)
}
