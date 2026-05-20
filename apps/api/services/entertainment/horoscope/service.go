package horoscope

import (
	"fmt"
	"hash/fnv"
	"strings"
	"time"
)

var signs = map[string]bool{
	"aries": true, "taurus": true, "gemini": true, "cancer": true,
	"leo": true, "virgo": true, "libra": true, "scorpio": true,
	"sagittarius": true, "capricorn": true, "aquarius": true, "pisces": true,
}

var readings = []string{
	"Today is a great day for new beginnings. Trust your instincts and take that first step toward your goals.",
	"Your patience will be rewarded. Take time to reflect before making important decisions today.",
	"Unexpected opportunities are on the horizon. Stay open-minded and embrace change with confidence.",
	"Focus on your relationships today. A heartfelt conversation could strengthen an important bond.",
	"Your creative energy is at a peak. Channel it into a project that truly matters to you.",
	"Financial matters deserve attention. Review your priorities and make thoughtful choices.",
	"A challenge you face today holds a valuable lesson. Approach it with curiosity rather than frustration.",
	"Your intuition is your greatest ally right now. Listen to that quiet inner voice guiding you.",
	"Collaboration will bring success. Reach out to someone whose strengths complement yours.",
	"Rest and self-care are essential today. Taking care of yourself enables you to take care of others.",
	"A long-held dream is closer than it appears. Keep moving forward with steady, deliberate action.",
	"Today brings clarity on a situation that has been confusing. Trust the process and stay grounded.",
}

var moods = []string{
	"energetic", "reflective", "optimistic", "calm", "adventurous",
	"focused", "creative", "grateful", "determined", "peaceful", "inspired", "curious",
}

type Horoscope struct {
	Sign        string `json:"sign"`
	Date        string `json:"date"`
	Horoscope   string `json:"horoscope"`
	LuckyNumber int    `json:"lucky_number"`
	Mood        string `json:"mood"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func IsValidSign(sign string) bool {
	return signs[strings.ToLower(sign)]
}

func (s *Service) Daily(sign string) (Horoscope, error) {
	sign = strings.ToLower(sign)
	if !IsValidSign(sign) {
		return Horoscope{}, fmt.Errorf("invalid zodiac sign: %s", sign)
	}

	today := time.Now().UTC().Format("2006-01-02")
	seed := hash(sign + today)

	return Horoscope{
		Sign:        sign,
		Date:        today,
		Horoscope:   readings[seed%uint64(len(readings))],
		LuckyNumber: int(seed%99) + 1,
		Mood:        moods[seed%uint64(len(moods))],
	}, nil
}

func hash(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}
