package trivia

import (
	"fmt"
	"math/rand/v2"
)

// Question is a trivia question with multiple-choice answers.
type Question struct {
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	Answer     string   `json:"answer"`
	Category   string   `json:"category"`
	Difficulty string   `json:"difficulty"`
}

// questions is the in-memory trivia question database.
var questions = []Question{
	// --- science ---
	{
		Question:   "What is the largest planet in our solar system?",
		Options:    []string{"Earth", "Jupiter", "Saturn", "Mars"},
		Answer:     "Jupiter",
		Category:   "science",
		Difficulty: "easy",
	},
	{
		Question:   "What is the chemical symbol for water?",
		Options:    []string{"O2", "H2O", "CO2", "HO"},
		Answer:     "H2O",
		Category:   "science",
		Difficulty: "easy",
	},
	{
		Question:   "What force keeps planets in orbit around the sun?",
		Options:    []string{"Magnetism", "Friction", "Gravity", "Inertia"},
		Answer:     "Gravity",
		Category:   "science",
		Difficulty: "easy",
	},
	{
		Question:   "What is the speed of light in a vacuum (approximately)?",
		Options:    []string{"300,000 km/s", "150,000 km/s", "450,000 km/s", "100,000 km/s"},
		Answer:     "300,000 km/s",
		Category:   "science",
		Difficulty: "medium",
	},
	{
		Question:   "Which element has the atomic number 79?",
		Options:    []string{"Silver", "Platinum", "Gold", "Copper"},
		Answer:     "Gold",
		Category:   "science",
		Difficulty: "medium",
	},
	{
		Question:   "What is the powerhouse of the cell?",
		Options:    []string{"Nucleus", "Ribosome", "Mitochondria", "Golgi apparatus"},
		Answer:     "Mitochondria",
		Category:   "science",
		Difficulty: "easy",
	},
	{
		Question:   "What particle has a negative charge?",
		Options:    []string{"Proton", "Neutron", "Electron", "Photon"},
		Answer:     "Electron",
		Category:   "science",
		Difficulty: "easy",
	},
	{
		Question:   "What is the half-life of Carbon-14 (approximately)?",
		Options:    []string{"1,000 years", "5,730 years", "10,000 years", "500 years"},
		Answer:     "5,730 years",
		Category:   "science",
		Difficulty: "hard",
	},
	{
		Question:   "What is the Heisenberg Uncertainty Principle about?",
		Options:    []string{"The speed of light", "The position and momentum of a particle", "The charge of electrons", "The mass of atoms"},
		Answer:     "The position and momentum of a particle",
		Category:   "science",
		Difficulty: "hard",
	},

	// --- history ---
	{
		Question:   "In which year did World War II end?",
		Options:    []string{"1943", "1944", "1945", "1946"},
		Answer:     "1945",
		Category:   "history",
		Difficulty: "easy",
	},
	{
		Question:   "Who was the first President of the United States?",
		Options:    []string{"John Adams", "Thomas Jefferson", "George Washington", "Benjamin Franklin"},
		Answer:     "George Washington",
		Category:   "history",
		Difficulty: "easy",
	},
	{
		Question:   "In which year did the French Revolution begin?",
		Options:    []string{"1776", "1789", "1799", "1804"},
		Answer:     "1789",
		Category:   "history",
		Difficulty: "medium",
	},
	{
		Question:   "Which ancient wonder was located in Alexandria?",
		Options:    []string{"The Colossus of Rhodes", "The Lighthouse of Alexandria", "The Hanging Gardens", "The Temple of Artemis"},
		Answer:     "The Lighthouse of Alexandria",
		Category:   "history",
		Difficulty: "medium",
	},
	{
		Question:   "Who was the last pharaoh of ancient Egypt?",
		Options:    []string{"Nefertiti", "Cleopatra VII", "Hatshepsut", "Ramesses II"},
		Answer:     "Cleopatra VII",
		Category:   "history",
		Difficulty: "hard",
	},
	{
		Question:   "In what year was the Magna Carta signed?",
		Options:    []string{"1215", "1066", "1415", "1305"},
		Answer:     "1215",
		Category:   "history",
		Difficulty: "medium",
	},

	// --- geography ---
	{
		Question:   "What is the capital of France?",
		Options:    []string{"London", "Berlin", "Paris", "Madrid"},
		Answer:     "Paris",
		Category:   "geography",
		Difficulty: "easy",
	},
	{
		Question:   "Which is the longest river in the world?",
		Options:    []string{"Amazon", "Yangtze", "Mississippi", "Nile"},
		Answer:     "Nile",
		Category:   "geography",
		Difficulty: "easy",
	},
	{
		Question:   "What is the smallest country in the world?",
		Options:    []string{"Monaco", "San Marino", "Vatican City", "Liechtenstein"},
		Answer:     "Vatican City",
		Category:   "geography",
		Difficulty: "medium",
	},
	{
		Question:   "Which country has the most natural lakes?",
		Options:    []string{"Russia", "USA", "Canada", "Brazil"},
		Answer:     "Canada",
		Category:   "geography",
		Difficulty: "medium",
	},
	{
		Question:   "What is the capital of Burkina Faso?",
		Options:    []string{"Dakar", "Ouagadougou", "Bamako", "Niamey"},
		Answer:     "Ouagadougou",
		Category:   "geography",
		Difficulty: "hard",
	},
	{
		Question:   "On which continent is the Atacama Desert located?",
		Options:    []string{"Africa", "Australia", "South America", "Asia"},
		Answer:     "South America",
		Category:   "geography",
		Difficulty: "medium",
	},

	// --- sports ---
	{
		Question:   "How many players are on a standard soccer (football) team?",
		Options:    []string{"9", "10", "11", "12"},
		Answer:     "11",
		Category:   "sports",
		Difficulty: "easy",
	},
	{
		Question:   "In which sport would you perform a 'slam dunk'?",
		Options:    []string{"Volleyball", "Basketball", "Baseball", "Tennis"},
		Answer:     "Basketball",
		Category:   "sports",
		Difficulty: "easy",
	},
	{
		Question:   "How many Grand Slam tournaments are there in tennis?",
		Options:    []string{"2", "3", "4", "5"},
		Answer:     "4",
		Category:   "sports",
		Difficulty: "medium",
	},
	{
		Question:   "In which country did the ancient Olympic Games originate?",
		Options:    []string{"Italy", "Turkey", "Greece", "Egypt"},
		Answer:     "Greece",
		Category:   "sports",
		Difficulty: "easy",
	},
	{
		Question:   "What is the maximum break score in snooker?",
		Options:    []string{"147", "155", "132", "168"},
		Answer:     "147",
		Category:   "sports",
		Difficulty: "medium",
	},
	{
		Question:   "What distance is a marathon?",
		Options:    []string{"26.2 miles", "25 miles", "30 km", "24.8 miles"},
		Answer:     "26.2 miles",
		Category:   "sports",
		Difficulty: "medium",
	},

	// --- music ---
	{
		Question:   "How many strings does a standard guitar have?",
		Options:    []string{"4", "5", "6", "7"},
		Answer:     "6",
		Category:   "music",
		Difficulty: "easy",
	},
	{
		Question:   "Which band released the album 'Abbey Road'?",
		Options:    []string{"The Rolling Stones", "The Beatles", "Led Zeppelin", "Pink Floyd"},
		Answer:     "The Beatles",
		Category:   "music",
		Difficulty: "easy",
	},
	{
		Question:   "What is the term for the speed of a piece of music?",
		Options:    []string{"Dynamics", "Tempo", "Pitch", "Timbre"},
		Answer:     "Tempo",
		Category:   "music",
		Difficulty: "medium",
	},
	{
		Question:   "In music notation, how many beats does a whole note receive in 4/4 time?",
		Options:    []string{"2", "3", "4", "8"},
		Answer:     "4",
		Category:   "music",
		Difficulty: "medium",
	},
	{
		Question:   "What is the lowest male vocal range?",
		Options:    []string{"Tenor", "Baritone", "Bass", "Counter-tenor"},
		Answer:     "Bass",
		Category:   "music",
		Difficulty: "medium",
	},

	// --- movies ---
	{
		Question:   "Who directed the movie 'Jaws' (1975)?",
		Options:    []string{"Francis Ford Coppola", "Martin Scorsese", "Steven Spielberg", "George Lucas"},
		Answer:     "Steven Spielberg",
		Category:   "movies",
		Difficulty: "medium",
	},
	{
		Question:   "Which film won the first Academy Award for Best Picture?",
		Options:    []string{"Gone with the Wind", "Casablanca", "Wings", "It Happened One Night"},
		Answer:     "Wings",
		Category:   "movies",
		Difficulty: "hard",
	},
	{
		Question:   "What is the name of the toy cowboy in 'Toy Story'?",
		Options:    []string{"Buzz", "Woody", "Rex", "Hamm"},
		Answer:     "Woody",
		Category:   "movies",
		Difficulty: "easy",
	},
	{
		Question:   "In 'The Matrix', what color pill does Neo take?",
		Options:    []string{"Blue", "Green", "Red", "Yellow"},
		Answer:     "Red",
		Category:   "movies",
		Difficulty: "easy",
	},
	{
		Question:   "Which 1994 film features a character played by Tom Hanks who says 'Life is like a box of chocolates'?",
		Options:    []string{"Philadelphia", "Cast Away", "Forrest Gump", "The Green Mile"},
		Answer:     "Forrest Gump",
		Category:   "movies",
		Difficulty: "easy",
	},

	// --- literature ---
	{
		Question:   "Who wrote 'Romeo and Juliet'?",
		Options:    []string{"Charles Dickens", "Jane Austen", "William Shakespeare", "Geoffrey Chaucer"},
		Answer:     "William Shakespeare",
		Category:   "literature",
		Difficulty: "easy",
	},
	{
		Question:   "What is the first book of the Bible?",
		Options:    []string{"Exodus", "Genesis", "Psalms", "Matthew"},
		Answer:     "Genesis",
		Category:   "literature",
		Difficulty: "easy",
	},
	{
		Question:   "Who wrote 'Pride and Prejudice'?",
		Options:    []string{"Charlotte Brontë", "Emily Brontë", "George Eliot", "Jane Austen"},
		Answer:     "Jane Austen",
		Category:   "literature",
		Difficulty: "easy",
	},
	{
		Question:   "In which Dickens novel does the character Ebenezer Scrooge appear?",
		Options:    []string{"Oliver Twist", "Great Expectations", "A Christmas Carol", "A Tale of Two Cities"},
		Answer:     "A Christmas Carol",
		Category:   "literature",
		Difficulty: "medium",
	},
	{
		Question:   "Who wrote '1984'?",
		Options:    []string{"Aldous Huxley", "Ray Bradbury", "George Orwell", "H.G. Wells"},
		Answer:     "George Orwell",
		Category:   "literature",
		Difficulty: "medium",
	},
	{
		Question:   "What is the name of the whale in 'Moby-Dick'?",
		Options:    []string{"Leviathan", "Moby Dick", "The White Whale", "Queequeg"},
		Answer:     "Moby Dick",
		Category:   "literature",
		Difficulty: "medium",
	},

	// --- math ---
	{
		Question:   "What is 7 × 8?",
		Options:    []string{"54", "56", "48", "64"},
		Answer:     "56",
		Category:   "math",
		Difficulty: "easy",
	},
	{
		Question:   "What is the value of π (pi) to two decimal places?",
		Options:    []string{"3.12", "3.14", "3.16", "3.18"},
		Answer:     "3.14",
		Category:   "math",
		Difficulty: "easy",
	},
	{
		Question:   "What is the square root of 144?",
		Options:    []string{"11", "12", "13", "14"},
		Answer:     "12",
		Category:   "math",
		Difficulty: "easy",
	},
	{
		Question:   "What is the sum of angles in a triangle?",
		Options:    []string{"90°", "180°", "270°", "360°"},
		Answer:     "180°",
		Category:   "math",
		Difficulty: "easy",
	},
	{
		Question:   "What is the derivative of sin(x)?",
		Options:    []string{"-sin(x)", "sin(x)", "-cos(x)", "cos(x)"},
		Answer:     "cos(x)",
		Category:   "math",
		Difficulty: "medium",
	},
	{
		Question:   "How many prime numbers are less than 20?",
		Options:    []string{"6", "7", "8", "9"},
		Answer:     "8",
		Category:   "math",
		Difficulty: "medium",
	},
	{
		Question:   "What is Euler's number (e) to two decimal places?",
		Options:    []string{"2.71", "2.73", "2.69", "2.75"},
		Answer:     "2.71",
		Category:   "math",
		Difficulty: "hard",
	},

	// --- technology ---
	{
		Question:   "What does 'CPU' stand for?",
		Options:    []string{"Central Processing Unit", "Computer Personal Unit", "Central Program Unit", "Core Processing Utility"},
		Answer:     "Central Processing Unit",
		Category:   "technology",
		Difficulty: "easy",
	},
	{
		Question:   "What does 'HTTP' stand for?",
		Options:    []string{"HyperText Transfer Protocol", "HyperText Transmission Program", "High Transfer Text Protocol", "Hyper Transfer Text Process"},
		Answer:     "HyperText Transfer Protocol",
		Category:   "technology",
		Difficulty: "easy",
	},
	{
		Question:   "In which year was the World Wide Web invented by Tim Berners-Lee?",
		Options:    []string{"1985", "1989", "1991", "1994"},
		Answer:     "1989",
		Category:   "technology",
		Difficulty: "medium",
	},
	{
		Question:   "What is the base of the binary number system?",
		Options:    []string{"2", "8", "10", "16"},
		Answer:     "2",
		Category:   "technology",
		Difficulty: "easy",
	},
	{
		Question:   "Which programming language is known as the 'backbone of the web'?",
		Options:    []string{"Python", "Java", "JavaScript", "Ruby"},
		Answer:     "JavaScript",
		Category:   "technology",
		Difficulty: "medium",
	},
	{
		Question:   "What does 'SQL' stand for?",
		Options:    []string{"Structured Query Language", "Simple Query Logic", "Standard Query List", "Sequential Query Language"},
		Answer:     "Structured Query Language",
		Category:   "technology",
		Difficulty: "medium",
	},
	{
		Question:   "What is the time complexity of a binary search algorithm?",
		Options:    []string{"O(n)", "O(n²)", "O(log n)", "O(1)"},
		Answer:     "O(log n)",
		Category:   "technology",
		Difficulty: "hard",
	},

	// --- nature ---
	{
		Question:   "What is the fastest land animal?",
		Options:    []string{"Lion", "Cheetah", "Leopard", "Greyhound"},
		Answer:     "Cheetah",
		Category:   "nature",
		Difficulty: "easy",
	},
	{
		Question:   "What gas do plants absorb during photosynthesis?",
		Options:    []string{"Oxygen", "Nitrogen", "Carbon dioxide", "Hydrogen"},
		Answer:     "Carbon dioxide",
		Category:   "nature",
		Difficulty: "easy",
	},
	{
		Question:   "How many legs does a spider have?",
		Options:    []string{"6", "8", "10", "12"},
		Answer:     "8",
		Category:   "nature",
		Difficulty: "easy",
	},
	{
		Question:   "What is the term for animals that eat only plants?",
		Options:    []string{"Carnivore", "Omnivore", "Herbivore", "Insectivore"},
		Answer:     "Herbivore",
		Category:   "nature",
		Difficulty: "easy",
	},
	{
		Question:   "Which ocean is the largest?",
		Options:    []string{"Atlantic", "Indian", "Arctic", "Pacific"},
		Answer:     "Pacific",
		Category:   "nature",
		Difficulty: "easy",
	},
	{
		Question:   "What is the gestation period of an elephant (approximately)?",
		Options:    []string{"9 months", "14 months", "22 months", "30 months"},
		Answer:     "22 months",
		Category:   "nature",
		Difficulty: "hard",
	},
	{
		Question:   "What type of rock is formed from cooled lava?",
		Options:    []string{"Sedimentary", "Metamorphic", "Igneous", "Limestone"},
		Answer:     "Igneous",
		Category:   "nature",
		Difficulty: "medium",
	},
}

// Service provides trivia question operations.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}

// Random returns a randomly selected trivia question, optionally filtered by
// category and/or difficulty. Returns an error if no questions match the
// given filters.
func (s *Service) Random(category, difficulty string) (Question, error) {
	pool := make([]Question, 0, len(questions))
	for _, q := range questions {
		if category != "" && q.Category != category {
			continue
		}
		if difficulty != "" && q.Difficulty != difficulty {
			continue
		}
		pool = append(pool, q)
	}

	if len(pool) == 0 {
		return Question{}, fmt.Errorf("no trivia questions found for the given filters")
	}

	return pool[rand.IntN(len(pool))], nil //nolint:gosec // Good enough for trivia selection.
}
