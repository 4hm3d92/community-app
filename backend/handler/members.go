package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/4hm3d92/community-app/backend/db"
	"github.com/4hm3d92/community-app/backend/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var memberIDKey = "memberID"

func members(router chi.Router) {
	router.Get("/", getAllMembers)
	router.Post("/", createMember)
	//router.Options("/", corsFix)

	router.Route("/{memberId}", func(router chi.Router) {
		router.Use(MemberContext)
		//router.Get("/", getMember)
		router.Put("/", updateMember)
		//router.Delete("/", deleteMember)
	})
}

func MemberContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		memberId := chi.URLParam(r, "memberId")
		if memberId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("member ID is required")))
			return
		}
		id, err := strconv.Atoi(memberId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid member ID")))
		}
		ctx := context.WithValue(r.Context(), memberIDKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllMembers(w http.ResponseWriter, r *http.Request) {

	members, err := dbInstance.GetAllMembers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, members); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createMember(w http.ResponseWriter, r *http.Request) {
	member := &models.Member{}
	if err := render.Bind(r, member); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddMember(member); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, member); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

/*
func getMember(w http.ResponseWriter, r *http.Request) {
	memberID := r.Context().Value(memberIDKey).(int)
	member, err := dbInstance.GetMemberById(memberID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &member); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	memberId := r.Context().Value(memberIDKey).(int)
	err := dbInstance.DeleteMember(memberId)
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

func updateMember(w http.ResponseWriter, r *http.Request) {
	memberId := r.Context().Value(memberIDKey).(int)
	memberData := models.Member{}
	if err := render.Bind(r, &memberData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	member, err := dbInstance.UpdateMember(memberId, memberData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &member); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
