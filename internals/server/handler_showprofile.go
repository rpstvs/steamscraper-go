package server

import (
	"fmt"
	"net/http"

	"github.com/rpstvs/steamscraper-go/internals/database"
)

func (srv *Server) ShowProfile(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("estamos in")
	users, err := srv.DB.GetUserbyId(r.Context(), "76561198000318676")
	fmt.Println(users)
	userBags, err := srv.DB.GetBagsByUser(r.Context(), user.ID)

	fmt.Println(user.ID)

	if err != nil {
		fmt.Println("User doesn't have any bags created")
		return
	}

	RespondWithJson(w, http.StatusOK, userBags)

}
