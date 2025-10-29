package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"islamicprojectapi/pkg/hadith"
	"islamicprojectapi/pkg/prayer"
	"islamicprojectapi/pkg/quran"
)

func main() {
	router := gin.Default()

	// Quran routes
	q := router.Group("/api/quran")
	{
		q.GET("/surah/:number", quran.SurahHandler)
		q.GET("/surah", quran.SurahListHandler)
		q.GET("/verse/:surah/:number", quran.VerseHandler)
		q.GET("/verses/:surah", quran.VersesBySurahHandler)
		// root of /api/quran returns list of surahs
		q.GET("/", quran.SurahListHandler)
	}
	// also register plain /api/quran (no trailing slash) to return list
	router.GET("/api/quran", quran.SurahListHandler)

	// Hadith routes
	h := router.Group("/api/hadith")
	{
		h.GET("/:collection/:number", hadith.HadithHandler)
		h.GET("/list", hadith.ListHadithsHandler)
		// root of /api/hadith returns list of hadiths
		h.GET("/", hadith.ListHadithsHandler)
	}
	// also register plain /api/hadith
	router.GET("/api/hadith", hadith.ListHadithsHandler)

	// Prayer Times routes
	p := router.Group("/api/prayer")
	{
		p.GET("/times/:location", prayer.PrayerTimesHandler)
		p.GET("/locations", prayer.ListLocationsHandler)
		// root of /api/prayer returns available prayer locations (list)
		p.GET("/", prayer.ListLocationsHandler)
	}
	// also register plain /api/prayer
	router.GET("/api/prayer", prayer.ListLocationsHandler)

	log.Println("Starting server on :8080")
	router.Run("0.0.0.0:8080")
}
