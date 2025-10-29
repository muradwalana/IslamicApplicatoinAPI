package quran

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"

	"islamicprojectapi/pkg/db"
)

// SurahHandler returns metadata for a surah
func SurahHandler(c *gin.Context) {
	numStr := c.Param("number")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid surah number"})
		return
	}
	s, err := db.GetSurah(num)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "surah not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

// VerseHandler returns a specific verse
func VerseHandler(c *gin.Context) {
	surahStr := c.Param("surah")
	verseStr := c.Param("number")
	surahNum, err := strconv.Atoi(surahStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid surah number"})
		return
	}
	verseNum, err := strconv.Atoi(verseStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verse number"})
		return
	}
	v, err := db.GetVerse(surahNum, verseNum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "verse not found"})
		return
	}
	c.JSON(http.StatusOK, v)
}

// SurahListHandler returns all surahs
func SurahListHandler(c *gin.Context) {
	list, err := db.GetAllSurahs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load surahs"})
		return
	}
	// ensure sorted by number (defensive)
	sort.Slice(list, func(i, j int) bool { return list[i].Number < list[j].Number })
	c.JSON(http.StatusOK, list)
}

// VersesBySurahHandler returns all verses for a surah
func VersesBySurahHandler(c *gin.Context) {
	surahStr := c.Param("surah")
	surahNum, err := strconv.Atoi(surahStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid surah number"})
		return
	}
	verses, err := db.GetVersesBySurah(surahNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load verses"})
		return
	}
	c.JSON(http.StatusOK, verses)
}
