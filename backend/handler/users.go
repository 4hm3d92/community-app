package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/4hm3d92/community-app/backend/db"
	"github.com/4hm3d92/community-app/backend/models"
	"github.com/4hm3d92/community-app/backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var userIDKey = "userID"

func users(router chi.Router) {
	router.Get("/", getAllUsers)
	router.Post("/", createUser)
	router.Post("/login", loginUser)
	router.Post("/logout", logoutUser)
	router.Post("/setPassword", setPassword)
	//router.Options("/", corsFix)

	router.Route("/{userId}", func(router chi.Router) {
		router.Use(UserContext)
		//router.Get("/", getUser)
		router.Put("/", updateUser)
		//router.Delete("/", deleteUser)
	})
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		if userId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("user ID is required")))
			return
		}
		id, err := strconv.Atoi(userId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid user ID")))
		}
		ctx := context.WithValue(r.Context(), userIDKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

/*
func corsFix(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
*/

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddUser(user); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

/*
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(int)
	user, err := dbInstance.GetUserById(userID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(int)
	err := dbInstance.DeleteUser(userId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
*/

func updateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userIDKey).(int)
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	err := dbInstance.UpdateUser(userId, userData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func setPassword(w http.ResponseWriter, r *http.Request) {
	//userId := r.Context().Value(userIDKey).(int)
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	err := dbInstance.SetPassword(userData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	/*if err := render.Render(w, r, ); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}*/
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	userLoginRequest := &models.UserLoginRequest{}
	//userData := models.User{}
	if err := render.Bind(r, userLoginRequest); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	/*
		if err := dbInstance.AddUser(user); err != nil {
			render.Render(w, r, ErrorRenderer(err))
			return
		}
		if err := render.Render(w, r, user); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	*/

	userData, err := dbInstance.GetUserByUsername(userLoginRequest)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	valid, err := utils.ComparePasswordAndHash(userLoginRequest.Password, userData.Password)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
	}
	/*
		if err := render.Render(w, r, &userData); err != nil {
			render.Render(w, r, ServerErrorRenderer(err))
			return
		}
	*/
	if valid {

		err = sm.RenewToken(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		sm.Put(r.Context(), "userId", userData.ID)
		sm.Put(r.Context(), "role", userData.Role)

	} else {
		render.Render(w, r, ErrUnauthorized)
		return
	}

	userResponse := &models.UserLoginResponse{}
	userResponse.Username = userLoginRequest.Username
	userResponse.Role = userData.Role

	if err := render.Render(w, r, userResponse); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func logoutUser(w http.ResponseWriter, r *http.Request) {

	sm.Destroy(r.Context())
	err := sm.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
