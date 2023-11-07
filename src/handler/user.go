package handler

import (
	"fmt"
	"net/http"
	entity "reflexscale/src/entitiy"
	"strconv"
	"time"

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
		User:  *convertResponse(c, user),
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
		usersResponse = append(usersResponse, convertResponse(c, userResponse))
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

	SuccessResponse(c, http.StatusOK, "Success fetch user", convertResponse(c, user))
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

	SuccessResponse(c, http.StatusOK, "Success Update User", convertResponse(c, update))
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

func convertResponse(c *gin.Context, user *entity.User) *entity.UserResponse {
	hasilTests := []*entity.UserHasilTestResponse{}
	for _, hasilTest := range user.HasilTests {
		// Parsing tanggal ke dalam objek time.Time
		parsedDate, err := time.Parse("2006-01-02 15:04:05.999 -0700 -07", hasilTest.CreatedAt.String())
		if err != nil {
			ErrorResponse(c, http.StatusUnprocessableEntity, err)
			return nil
		}

		// Mengubah format tanggal ke "07-Nov-23" dengan bulan dalam bahasa Indonesia
		formattedDate := fmt.Sprintf("%02d-%s-%02d", parsedDate.Day(), formatBulanIndonesia(parsedDate.Month()), parsedDate.Year()%100)

		hasilTests = append(hasilTests, &entity.UserHasilTestResponse{
			ID:        hasilTest.ID,
			Poin:      hasilTest.Poin,
			CreatedAt: formattedDate,
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

func formatBulanIndonesia(bulan time.Month) string {
	switch bulan {
	case time.January:
		return "Jan"
	case time.February:
		return "Feb"
	case time.March:
		return "Mar"
	case time.April:
		return "Apr"
	case time.May:
		return "Mei"
	case time.June:
		return "Jun"
	case time.July:
		return "Jul"
	case time.August:
		return "Ags"
	case time.September:
		return "Sep"
	case time.October:
		return "Okt"
	case time.November:
		return "Nov"
	case time.December:
		return "Des"
	default:
		return ""
	}
}
