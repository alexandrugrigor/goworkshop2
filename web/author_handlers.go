package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
)

func GetAllAuthors(w http.ResponseWriter, _ *http.Request) error {
	authors, err := persistence.Store.GetAuthors()
	if err != nil {
		return err
	}
	WriteJson(w, authors)
	return nil
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	authorUUID := mux.Vars(r)["uuid"]
	if author, err := persistence.Store.GetAuthor(authorUUID); err != nil {
		return err
	} else {
		WriteJson(w, author)
		return nil
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) error {
	var authorUUID = mux.Vars(r)["uuid"]
	if err := persistence.Store.DeleteAuthor(authorUUID); err != nil {
		return err
	} else {
		WriteJson(w, struct{ Message string }{Message: "Deleted"})
		return nil
	}
}

func AddAuthor(w http.ResponseWriter, r *http.Request) error {
	var author model.Author
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &author); err != nil {
		return err
	} else {
		persistence.Store.CreateAuthor(&author)
		WriteJson(w, author)
		return nil
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) error {
	var author model.Author
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &author); err != nil {
		return err
	}

	if err := persistence.Store.UpdateAuthor(&author); err != nil {
		return err
	} else {
		WriteJson(w, author)
	}
	return nil
}
