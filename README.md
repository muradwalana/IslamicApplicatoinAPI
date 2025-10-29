# Islamic Application API

A comprehensive Go-based REST API service providing Islamic content including Quran, Hadith, and Prayer Times, similar to Quran.com.

## Description

This project is a RESTful API service built with Go and the Gin framework that provides various Islamic features including:
- Quran verses with translations and audio
- Hadith collections
- Prayer times by location

## Prerequisites

- Go 1.x or higher
- Git
- [Gin Framework](https://github.com/gin-gonic/gin)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/muradwalana/IslamicApplicatoinAPI.git
cd IslamicApplicatoinAPI
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

To start the server locally (development):
```bash
go run main.go
```

The server listens on `:8080` by default.

Quick smoke checks (examples):
```bash
# Get surah metadata
curl -sS http://localhost:8080/api/quran/surah/1

# Get a verse
curl -sS http://localhost:8080/api/quran/verse/1/1

# Get a hadith (collection slug + number)
curl -sS http://localhost:8080/api/hadith/bukhari/1

# Get prayer times for a supported location slug
curl -sS http://localhost:8080/api/prayer/times/new-york
```

Notes:
- Data is embedded in the binary (sample JSON under `pkg/db/data/`). Replace or extend those files to add real content.
- For production use build the binary: `go build -o islamicapi .` and run the produced executable.
- Consider setting `GIN_MODE=release` before running in production to disable debug logging.

## API Endpoints

### Data Structures

#### Surah Structure
```go
type Surah struct {
    Number       int    `json:"number"`
    Name         string `json:"name"`
    EnglishName string `json:"englishName"`
    Verses      int    `json:"numberOfVerses"`
    Type        string `json:"revelationType"`
}
```

#### Verse Structure
```go
type Verse struct {
    SurahNumber  int    `json:"surahNumber"`
    VerseNumber  int    `json:"verseNumber"`
    ArabicText   string `json:"arabicText"`
    EnglishText  string `json:"englishText"`
    Translation  string `json:"translation"`
    AudioURL     string `json:"audioUrl,omitempty"`
}
```

#### Hadith Structure
```go
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
```

#### Prayer Time Structure
```go
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
```

### Available Endpoints

1. **Quran Endpoints**
   - Get Surah Information
     ```bash
     GET /api/quran/surah/:number
     # Example: curl http://localhost:8080/api/quran/surah/1
     ```
   - Get Specific Verse
     ```bash
     GET /api/quran/verse/:surah/:number
     # Example: curl http://localhost:8080/api/quran/verse/1/1
     ```

2. **Hadith Endpoints**
   - Get Specific Hadith
     ```bash
     GET /api/hadith/:collection/:number
     # Example: curl http://localhost:8080/api/hadith/bukhari/1
     ```

3. **Prayer Times Endpoints**
   - Get Prayer Times by Location
     ```bash
     GET /api/prayer/times/:location
     # Example: curl http://localhost:8080/api/prayer/times/new-york
     ```

## Sample Responses

### Surah Response
```json
{
    "number": 1,
    "name": "الفاتحة",
    "englishName": "Al-Fatiha",
    "numberOfVerses": 7,
    "revelationType": "Meccan"
}
```

### Verse Response
```json
{
    "surahNumber": 1,
    "verseNumber": 1,
    "arabicText": "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ",
    "englishText": "In the name of Allah, the Entirely Merciful, the Especially Merciful.",
    "translation": "Sahih International",
    "audioUrl": "https://verses.quran.com/1/1.mp3"
}
```

### Hadith Response
```json
{
    "id": "1",
    "collection": "Sahih al-Bukhari",
    "bookNumber": 1,
    "hadithNumber": 1,
    "arabicText": "إِنَّمَا الأَعْمَالُ بِالنِّيَّاتِ",
    "englishText": "Actions are but by intentions",
    "grade": "Sahih",
    "reference": "Bukhari 1:1"
}
```

### Prayer Times Response
```json
{
    "date": "2025-10-29",
    "fajr": "05:30",
    "sunrise": "06:45",
    "dhuhr": "12:30",
    "asr": "15:45",
    "maghrib": "18:15",
    "isha": "19:30",
    "location": "New York, USA"
}
```

## Error Handling

The API implements basic error handling:
- Returns 404 status for not found resources
- Returns 201 status for successful creation
- Returns 200 status for successful operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Development Notes

- Built using the Gin framework for efficient routing and middleware support
- Currently uses placeholder data (TODO: Implement database integration)
- Implements RESTful conventions for endpoint design
- Provides JSON responses with proper status codes and error handling
- Supports Arabic text with proper UTF-8 encoding
- Includes audio URLs for Quranic verses
- Provides multiple translations support
- Implements prayer time calculations based on geographical location

## Planned Features

1. **Quran Features**
   - Complete Quran text in Arabic
   - Multiple translations
   - Word-by-word analysis
   - Verse tafsir (commentary)
   - Audio recitations from various reciters
   - Verse topics and themes

2. **Hadith Features**
   - Multiple hadith collections (Bukhari, Muslim, etc.)
   - Chain of narration (isnad)
   - Hadith commentary
   - Topic-based browsing
   - Advanced search functionality

3. **Prayer Times Features**
   - Multiple calculation methods
   - Qibla direction
   - Hijri calendar conversion
   - Special prayer times (Eid, Friday prayer, etc.)
   - Location auto-detection

4. **Additional Features**
   - Islamic calendar events
   - Duas (supplications)
   - Islamic articles and resources
   - Multi-language support
   - User preferences and settings

## Technical Roadmap

1. **Phase 1: Core Implementation**
   - Database integration (PostgreSQL/MongoDB)
   - Complete Quran text and translation import
   - Basic hadith collection implementation
   - Prayer time calculation engine

2. **Phase 2: Enhanced Features**
   - Audio files integration
   - Advanced search functionality
   - Caching layer implementation
   - API rate limiting and security

3. **Phase 3: Optimization**
   - Performance optimization
   - Analytics integration
   - Documentation enhancement
   - Community contribution guidelines
