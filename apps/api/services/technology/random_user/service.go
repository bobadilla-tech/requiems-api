package randomuser

import (
	"net/url"

	faker "github.com/jaswdr/faker/v2"
)

// Generates random fake user data.
type Service struct{}

// NewService returns a new instance of Service.
func NewService() *Service {
	return &Service{}
}

// GenerateBatch generates count random users and returns them in order.
// The caller is responsible for passing a valid count; no bounds checking is performed here.
func (s *Service) GenerateBatch(count int) ([]User, error) {
	results := make([]User, count)
	for i := range count {
		results[i] = s.Generate()
	}
	return results, nil
}

// Returns a randomly generated User.
func (s *Service) Generate() User {
	f := faker.New()

	name := f.Person().Name()
	email := f.Internet().SafeEmail()
	phone := f.Phone().Number()

	fakerAddress := f.Address()

	address := Address{
		Street:  fakerAddress.StreetAddress(),
		City:    fakerAddress.City(),
		State:   fakerAddress.State(),
		Zip:     fakerAddress.PostCode(),
		Country: fakerAddress.Country(),
	}

	avatar := "https://api.dicebear.com/9.x/identicon/svg?seed=" + url.QueryEscape(name)

	return User{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
		Avatar:  avatar,
	}
}
