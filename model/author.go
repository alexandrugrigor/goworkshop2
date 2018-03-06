package model

import "fmt"

// Authors - the list of available authors
var Authors AuthorsList

type Entity struct{
	UUID      string `json:"uuid" gorm:"primary_key"`
}
//Author - The DTO used to access authors
type Author struct {
	Entity
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (author Author) String() string {
	return fmt.Sprintf("Author{UUID='%s', FirstName='%s', LastName='%s', Birthday='%s', Death='%s'}", author.UUID,
		author.FirstName, author.LastName, author.Birthday, author.Death)
}

type AuthorsList []Author

//deletes the given author from the list
func (a *AuthorsList) Delete(authorUUID string) error {
	var err = fmt.Errorf("could not find author by uuid %s", authorUUID)
	var updatedAuthors AuthorsList
	for _, author := range *a {
		if author.UUID == authorUUID {
			err = nil
		} else {
			updatedAuthors = append(updatedAuthors, author)
		}
	}
	if err == nil {
		*a = updatedAuthors
	}
	return err
}

//updates the given author
func (a *AuthorsList) Update(updatedAuthor Author) (Author, error) {
	var err = fmt.Errorf("could not find author by uuid %s", updatedAuthor.UUID)
	var newAuthors AuthorsList
	for _, author := range *a {
		if author.UUID == updatedAuthor.UUID {
			newAuthors = append(newAuthors, updatedAuthor)
			err = nil
		} else {
			newAuthors = append(newAuthors, author)
		}
	}
	if err == nil {
		*a = newAuthors
	}
	return updatedAuthor, err
}

//searches for the author with the given uuid
func (a *AuthorsList) Get(authorUUID string) (Author, error) {
	err := fmt.Errorf("could not find author by uuid %s", authorUUID)
	for _, author := range *a {
		if author.UUID == authorUUID {
			return author, nil
		}
	}
	return Author{}, err
}

//Adds the given author into the list
func (a *AuthorsList) Add(author Author) {
	*a = append(*a, author)
}
