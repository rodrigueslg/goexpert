package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"

	"github.com/rodrigueslg/fc-goexpert/api/internal/dto"
	"github.com/rodrigueslg/fc-goexpert/api/internal/entity"
	"github.com/rodrigueslg/fc-goexpert/api/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Auth godoc
// @Summary 	Get a JWT token
// @Description Get a JWT token
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request  	body	dto.GetJWTOutput	true	"user credentials"
// @Success 	200
// @Failure 	404
// @Failure 	500 	 	{object} Error
// @Router 		/users/auth [post]
func (h *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		log.Println(err)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !u.ComparePassword(user.Password) {
		error := Error{Message: "user doens't exist"}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary 	Create a new user
// @Description Create a new user
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request  body	dto.CreateUserRequest	true	"user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		log.Println(err)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
