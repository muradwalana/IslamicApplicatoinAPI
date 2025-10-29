package db

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Embedded sample data (packaged with the binary)
//
//go:embed data/quran.json
//go:embed data/hadith.json
//go:embed data/prayer_times.json
var embedFiles embed.FS

// Surah represents a chapter of the Quran
type Surah struct {
	Number      int    `json:"number"`
	Name        string `json:"name"`
	EnglishName string `json:"englishName"`
	Verses      int    `json:"numberOfVerses"`
	Type        string `json:"revelationType"`
}

// Verse represents a single verse from the Quran
type Verse struct {
	SurahNumber int    `json:"surahNumber"`
	VerseNumber int    `json:"verseNumber"`
	ArabicText  string `json:"arabicText"`
	EnglishText string `json:"englishText"`
	Translation string `json:"translation"`
	AudioURL    string `json:"audioUrl,omitempty"`
}

// Hadith represents a prophetic tradition
type Hadith struct {
	ID           string `json:"id"`
	Collection   string `json:"collection"`
	BookNumber   int    `json:"bookNumber"`
	HadithNumber int    `json:"hadithNumber"`
	ArabicText   string `json:"arabicText"`
	EnglishText  string `json:"englishText"`
	Grade        string `json:"grade"`
	Reference    string `json:"reference"`
}

// PrayerTime represents prayer times for a location
type PrayerTime struct {
	Date     string `json:"date"`
	Fajr     string `json:"fajr"`
	Sunrise  string `json:"sunrise"`
	Dhuhr    string `json:"dhuhr"`
	Asr      string `json:"asr"`
	Maghrib  string `json:"maghrib"`
	Isha     string `json:"isha"`
	Location string `json:"location"`
}

var (
	loadOnce  sync.Once
	loadErr   error
	surahMap  map[int]Surah
	verseMap  map[string]Verse
	hadithMap map[string]Hadith
	prayerMap map[string]PrayerTime
)

func loadData() error {
	// load quran
	qdata, err := embedFiles.ReadFile("data/quran.json")
	if err != nil {
		return fmt.Errorf("read quran data: %w", err)
	}
	var qb struct {
		Surahs []Surah `json:"surahs"`
		Verses []Verse `json:"verses"`
	}
	if err := json.Unmarshal(qdata, &qb); err != nil {
		return fmt.Errorf("unmarshal quran: %w", err)
	}
	surahMap = make(map[int]Surah, len(qb.Surahs))
	verseMap = make(map[string]Verse, len(qb.Verses))
	for _, s := range qb.Surahs {
		surahMap[s.Number] = s
	}
	for _, v := range qb.Verses {
		k := fmt.Sprintf("%d:%d", v.SurahNumber, v.VerseNumber)
		verseMap[k] = v
	}

	// load hadith
	hdata, err := embedFiles.ReadFile("data/hadith.json")
	if err != nil {
		return fmt.Errorf("read hadith data: %w", err)
	}
	var hlist []Hadith
	if err := json.Unmarshal(hdata, &hlist); err != nil {
		return fmt.Errorf("unmarshal hadith: %w", err)
	}
	hadithMap = make(map[string]Hadith, len(hlist))
	for _, h := range hlist {
		k := fmt.Sprintf("%s:%d", h.Collection, h.HadithNumber)
		hadithMap[k] = h
	}

	// load prayer times
	pdata, err := embedFiles.ReadFile("data/prayer_times.json")
	if err != nil {
		return fmt.Errorf("read prayer data: %w", err)
	}
	var plist map[string]PrayerTime
	if err := json.Unmarshal(pdata, &plist); err != nil {
		return fmt.Errorf("unmarshal prayer: %w", err)
	}
	prayerMap = plist

	return nil
}

func ensureLoaded() error {
	loadOnce.Do(func() { loadErr = loadData() })
	return loadErr
}

// GetSurah returns a Surah by its number
func GetSurah(number int) (*Surah, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	if s, ok := surahMap[number]; ok {
		return &s, nil
	}
	return nil, errors.New("surah not found")
}

// GetVerse returns a specific verse
func GetVerse(surah, verse int) (*Verse, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	k := fmt.Sprintf("%d:%d", surah, verse)
	if v, ok := verseMap[k]; ok {
		return &v, nil
	}
	return nil, errors.New("verse not found")
}

// GetHadith returns a hadith by collection and number
func GetHadith(collection string, number int) (*Hadith, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	k := fmt.Sprintf("%s:%d", collection, number)
	if h, ok := hadithMap[k]; ok {
		return &h, nil
	}
	return nil, errors.New("hadith not found")
}

// GetPrayerTimes returns prayer times for a location slug (e.g., "new-york")
func GetPrayerTimes(location string) (*PrayerTime, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	if p, ok := prayerMap[location]; ok {
		return &p, nil
	}
	return nil, errors.New("prayer times not found")
}

// GetAllSurahs returns all surah metadata in numeric order (map iteration order is not guaranteed)
func GetAllSurahs() ([]Surah, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	// produce deterministic ascending list by iterating numeric keys
	if len(surahMap) == 0 {
		return []Surah{}, nil
	}
	max := 0
	for k := range surahMap {
		if k > max {
			max = k
		}
	}
	res := make([]Surah, 0, len(surahMap))
	for i := 1; i <= max; i++ {
		if s, ok := surahMap[i]; ok {
			res = append(res, s)
		}
	}
	return res, nil
}

// GetVersesBySurah returns all verses for a given surah
func GetVersesBySurah(surah int) ([]Verse, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	res := []Verse{}
	prefix := fmt.Sprintf("%d:", surah)
	for k, v := range verseMap {
		if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
			res = append(res, v)
		}
	}
	// sort by verse number
	sort.Slice(res, func(i, j int) bool { return res[i].VerseNumber < res[j].VerseNumber })
	return res, nil
}

// GetAllHadiths returns all hadiths
func GetAllHadiths() ([]Hadith, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	res := make([]Hadith, 0, len(hadithMap))
	for _, h := range hadithMap {
		res = append(res, h)
	}
	// sort by collection then hadith number
	sort.Slice(res, func(i, j int) bool {
		if res[i].Collection == res[j].Collection {
			return res[i].HadithNumber < res[j].HadithNumber
		}
		return res[i].Collection < res[j].Collection
	})
	return res, nil
}

// GetAllPrayerLocations returns slugs of available prayer locations
func GetAllPrayerLocations() ([]string, error) {
	if err := ensureLoaded(); err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(prayerMap))
	for k := range prayerMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, nil
}
