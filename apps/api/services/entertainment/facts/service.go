package facts

import (
	"fmt"
	"math/rand/v2"
)

// Fact is the response payload for a single random fact.
type Fact struct {
	Fact     string `json:"fact"`
	Category string `json:"category"`
	Source   string `json:"source"`
}

func (Fact) IsData() {}

type entry struct {
	fact     string
	category string
	source   string
}

var database = []entry{
	// science
	{"Honey never spoils. Archaeologists have found 3,000-year-old honey in Egyptian tombs that was still perfectly edible.", "science", "National Geographic"},
	{"A day on Venus is longer than a year on Venus.", "science", "NASA"},
	{"Bananas are berries, but strawberries are not.", "science", "Britannica"},
	{"Hot water can freeze faster than cold water under certain conditions — this is known as the Mpemba effect.", "science", "Scientific American"},
	{"There are more atoms in a glass of water than glasses of water in all the world's oceans.", "science", "Physics Today"},
	{"The human body contains enough iron to make a nail about 3 inches long.", "science", "Smithsonian Magazine"},
	{"Octopuses have three hearts and blue blood.", "science", "National Geographic"},
	{"A bolt of lightning is five times hotter than the surface of the sun.", "science", "NOAA"},
	{"The moon is slowly drifting away from Earth at about 3.8 cm per year.", "science", "NASA"},
	{"Cats have a specialized collarbone that allows them to always land on their feet.", "science", "Cornell Feline Health Center"},
	// history
	{"Cleopatra lived closer in time to the Moon landing than to the construction of the Great Pyramid.", "history", "Smithsonian Magazine"},
	{"Oxford University is older than the Aztec Empire.", "history", "Oxford University"},
	{"Vikings used to give kittens to new brides as essential household gifts.", "history", "History Channel"},
	{"The shortest war in history lasted 38 to 45 minutes — the Anglo-Zanzibar War of 1896.", "history", "Guinness World Records"},
	{"Nintendo was founded in 1889, originally as a playing-card company.", "history", "Nintendo"},
	// technology
	{"The first computer bug was an actual bug — a moth found trapped in a Harvard Mark II computer in 1947.", "technology", "Smithsonian National Museum of American History"},
	{"More than 90% of the world's currency exists only digitally.", "technology", "Forbes"},
	{"The QWERTY keyboard layout was designed in the 1870s to slow down typists and prevent typewriter jams.", "technology", "Mental Floss"},
	{"The first text message was sent on December 3, 1992, and it said 'Merry Christmas'.", "technology", "Guinness World Records"},
	{"Google's founders originally planned to sell it for $1 million in 1999, but the offer was rejected.", "technology", "Wired"},
	// nature
	{"A group of flamingos is called a flamboyance.", "nature", "Merriam-Webster"},
	{"Trees can communicate and share nutrients through underground fungal networks known as the Wood Wide Web.", "nature", "BBC Earth"},
	{"Elephants are the only animals that cannot jump.", "nature", "National Geographic"},
	{"Sharks are older than trees — they've existed for around 400 million years.", "nature", "Smithsonian Ocean"},
	{"A single cloud can weigh more than a million pounds.", "nature", "USGS"},
	{"Sea otters hold hands while sleeping so they don't drift apart.", "nature", "Monterey Bay Aquarium"},
	// space
	{"One million Earths could fit inside the Sun.", "space", "NASA"},
	{"There is a giant cloud of alcohol in outer space, 1,000 times larger than our solar system's diameter.", "space", "Royal Astronomical Society"},
	{"Neutron stars are so dense that a teaspoon of one would weigh about a billion tons.", "space", "NASA"},
	{"The footprints on the Moon will last for about 100 million years because there is no wind.", "space", "NASA"},
	{"It takes light about 8 minutes and 20 seconds to travel from the Sun to Earth.", "space", "NASA"},
	// food
	{"Carrots were originally purple, not orange.", "food", "World Carrot Museum"},
	{"Apples, pears, and roses are all part of the same plant family.", "food", "Britannica"},
	{"Cashews grow on trees as the bottom of a fruit called a cashew apple.", "food", "Britannica"},
	{"The fear of running out of coffee is called 'coffeephobia'.", "food", "Merriam-Webster"},
	{"Pistachios are technically fruits, not nuts.", "food", "Britannica"},
}

// validCategories holds all known fact categories for validation.
var validCategories = map[string]bool{
	"science":    true,
	"history":    true,
	"technology": true,
	"nature":     true,
	"space":      true,
	"food":       true,
}

// Service provides random fact retrieval.
type Service struct{}

// NewService creates a new facts Service.
func NewService() *Service {
	return &Service{}
}

// IsValidCategory reports whether the given category is known.
func IsValidCategory(category string) bool {
	return validCategories[category]
}

// Random returns a random fact, optionally filtered by category.
func (s *Service) Random(category string) (Fact, error) {
	pool := database
	if category != "" {
		pool = make([]entry, 0, len(database))
		for _, e := range database {
			if e.category == category {
				pool = append(pool, e)
			}
		}
		if len(pool) == 0 {
			return Fact{}, fmt.Errorf("no facts found for category: %s", category)
		}
	}

	e := pool[rand.IntN(len(pool))] //nolint:gosec // non-security random selection; crypto/rand not needed here
	return Fact{
		Fact:     e.fact,
		Category: e.category,
		Source:   e.source,
	}, nil
}
