package handlers

import (
	"encoding/json"
	"io/ioutil"
	//"html/template"

	"net/http"

	model "../models"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	page := model.Page{ID: 3, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/register/users"}
	users := loadUsers()
	interests := loadInterests()
	interestMappings := loadInterestMappings()

	var newUsers []model.User

	for _, user := range users {

		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := model.UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
	//t, _ := template.ParseFiles("template/page.html")
	//t.Execute(w, viewModel)
}

func loadUsers() []model.User {
	bytes, _ := ioutil.ReadFile("json/users.json")
	var users []model.User
	json.Unmarshal(bytes, &users)
	return users
}

func loadInterests() []model.Interest {
	bytes, _ := ioutil.ReadFile("json/interests.json")
	var interests []model.Interest
	json.Unmarshal(bytes, &interests)
	return interests
}

func loadInterestMappings() []model.InterestMapping {
	bytes, _ := ioutil.ReadFile("json/userInterestMappings.json")
	var interestMappings []model.InterestMapping
	json.Unmarshal(bytes, &interestMappings)
	return interestMappings
}
