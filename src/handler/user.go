package handler

import (
	"net/http"
	"reflexscale/sdk/custome_time"
	"reflexscale/src/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *rest) CreateUser(c *gin.Context) {

	var userInput entity.UserRegisterInputParam

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createUser, err := r.uc.User.Register(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Create user success", createUser)
}

func (r *rest) UserLogin(c *gin.Context) {
	var userInput entity.UserLoginInputParam

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, tokenJwt, err := r.uc.User.Login(&userInput)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	userResponse := entity.UserLoginResponse{
		User:  *convertUserResponse(c, user),
		Token: tokenJwt,
	}

	SuccessResponse(c, http.StatusOK, "Login success", userResponse)
}

func (r *rest) ReadAll(c *gin.Context) {
	users, err := r.uc.User.ReadAll()

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	usersResponse := []*entity.UserResponse{}

	for _, userResponse := range users {
		usersResponse = append(usersResponse, convertUserResponse(c, userResponse))
	}

	SuccessResponse(c, http.StatusOK, "Success Fetch Users", usersResponse)
}

func (r *rest) ReadByID(c *gin.Context) {
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 32)

	user, err := r.uc.User.ReadByID(int(idUint))
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	SuccessResponse(c, http.StatusOK, "Success fetch user", convertUserResponse(c, user))
}

func (r *rest) Update(c *gin.Context) {
	var UserRequest entity.UserUpdateInputParam

	err := c.ShouldBindJSON(&UserRequest)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	update, err := r.uc.User.Update(id, &UserRequest)
	if err != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	SuccessResponse(c, http.StatusOK, "Success Update User", convertUserResponse(c, update))
}

func (r *rest) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	_, err := r.uc.User.Delete(idInt)
	if err != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, err)
	}

	SuccessResponse(c, http.StatusOK, "delete success", nil)
}

func convertUserResponse(c *gin.Context, user *entity.User) *entity.UserResponse {
	hasilTests := []*entity.UserHasilTestResponse{}
	for _, hasilTest := range user.HasilTests {
		hasilTests = append(hasilTests, &entity.UserHasilTestResponse{
			ID:        hasilTest.ID,
			Poin:      hasilTest.Poin,
			CreatedAt: custome_time.ConvertTimeFormat(hasilTest.CreatedAt.String()),
		})
	}

	return &entity.UserResponse{
		ID:           user.ID,
		NamaLengkap:  user.NamaLengkap,
		Email:        user.Email,
		Usia:         user.Usia,
		JenisKelamin: user.JenisKelamin,
		Alamat:       user.Alamat,
		Password:     user.Password,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		HasilTests:   hasilTests,
	}
}
