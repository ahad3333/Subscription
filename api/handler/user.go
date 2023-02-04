package handler

import (
	"add/models"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID CreateUser
// @Router /user [POST]
// @Summary CreateUser
// @Description CreateUser
// @Tags User
// @Accept json
// @Produce json
// @Param book body models.CreateUser true "CreateUserRequestBody"
// @Success 201 {object} models.User "GetBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateUser(c *gin.Context) {

	var user models.CreateUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.User().Insert(context.Background(), &user)
	if err != nil {
		log.Println("error whiling create user:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.User().GetByID(context.Background(), &models.UserPrimarKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id User:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIDUser godoc
// @ID Get_By_IDUser
// @Router /user/{id} [GET]
// @Summary GetByID User
// @Description GetByID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.User "GetByIDBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDUser(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.User().GetByID(context.Background(), &models.UserPrimarKey{
		Id: id,
	})

	if err != nil {
		log.Println("error whiling get by id user:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListUser godoc
// @ID UserPrimarKey
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListUserResponse "GetUserListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListUser(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := h.storage.User().GetList(context.Background(), &models.GetListUserRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		log.Println("error whiling get list user:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateUser godoc
// @ID UpdateUser
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.UpdateUserSwag true "UpdateUserRequestBody"
// @Success 202 {object} models.User "UpdateBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user.Id = c.Param("id")
	err = h.storage.User().Update(context.Background(), &models.UpdateUser{
		Id: user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	})
	if err != nil {
		log.Printf("error whiling update 2: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.User().GetByID(context.Background(), &models.UserPrimarKey{Id: user.Id})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// DeleteUser godoc
// @ID DeleteUser
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.User().Delete(context.Background(), &models.UserPrimarKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  user:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "delete user")
}
