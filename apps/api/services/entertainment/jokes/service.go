package jokes

import (
	"fmt"
	"math/rand/v2"
)

var dadJokes = []string{
	"Why don't scientists trust atoms? Because they make up everything!",
	"I'm reading a book about anti-gravity. It's impossible to put down.",
	"Did you hear about the mathematician who's afraid of negative numbers? He'll stop at nothing to avoid them.",
	"Why do cows wear bells? Because their horns don't work.",
	"I told my wife she was drawing her eyebrows too high. She looked surprised.",
	"What do you call fake spaghetti? An impasta.",
	"Why can't you give Elsa a balloon? Because she'll let it go.",
	"I used to hate facial hair, but then it grew on me.",
	"What do you call a fish without eyes? A fsh.",
	"Why did the scarecrow win an award? Because he was outstanding in his field.",
	"I'm on a seafood diet. I see food and I eat it.",
	"What do you call cheese that isn't yours? Nacho cheese.",
	"Why did the bicycle fall over? Because it was two-tired.",
	"What do you call a sleeping dinosaur? A dino-snore.",
	"How does a penguin build its house? Igloos it together.",
	"I would tell you a construction joke, but I'm still working on it.",
	"Why don't eggs tell jokes? They'd crack each other up.",
	"What do you call a bear with no teeth? A gummy bear.",
	"Why did the golfer bring extra pants? In case he got a hole in one.",
	"I told my doctor I broke my arm in two places. He told me to stop going to those places.",
	"What do you call a factory that makes okay products? A satisfactory.",
	"Did you hear about the claustrophobic astronaut? He just needed a little space.",
	"Why do fish swim in salt water? Because pepper makes them sneeze.",
	"What do you call a pile of cats? A meow-ntain.",
	"I asked my dog what two minus two is. He said nothing.",
	"What do you call a sad cup of coffee? Depresso.", //nolint:misspell // intentional coffee pun, not a typo
	"Why did the coffee file a police report? It got mugged.",
	"I only know 25 letters of the alphabet. I don't know y.",
	"What do you call a dinosaur that crashes their car? Tyrannosaurus wrecks.",
	"Why did the math book look so sad? Because it had too many problems.",
}

// DadJoke represents a single dad joke with its identifier.
type DadJoke struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}


// Service provides access to a collection of dad jokes.
type Service struct{}

// NewService returns a new jokes Service.
func NewService() *Service {
	return &Service{}
}

// Random returns a random dad joke from the collection.
func (s *Service) Random() DadJoke {
	idx := rand.IntN(len(dadJokes)) //nolint:gosec // Good enough for joke selection.
	return DadJoke{
		ID:   fmt.Sprintf("joke_%d", idx+1),
		Joke: dadJokes[idx],
	}
}
