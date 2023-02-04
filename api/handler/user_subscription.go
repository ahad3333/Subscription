package handler

import (
	"add/models"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSubscriptionUser godoc
// @ID create_SubscriptionUser
// @Router /subscription-user [POST]
// @Summary Create Subscription User
// @Description Create Subscription User
// @Tags Subscription User
// @Accept json
// @Produce json
// @Param subscription body models.CreateUserSubscription true "CreateSubscriptionUserRequestBody"
// @Success 201 {object} models.SubscriptionUser "GetUserSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateSubscriptionUser(c *gin.Context) {

	var subscriptionuser models.CreateUserSubscription

	err := c.ShouldBindJSON(&subscriptionuser)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.UserSubscription().Insert(context.Background(), &subscriptionuser)

	if err != nil {
		log.Println("error whiling create SubscriptionUser:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.UserSubscription().GetById(context.Background(), &models.SubscriptionUserPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Subscription User:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetUserSubscriptionsByUserId godoc
// @ID get_by_id_SubscriptionUser
// @Router /subscription-user/{id} [GET]
// @Summary Get By ID SubscriptionUser
// @Description GetUserSubscriptionsByUserId
// @Tags Subscription User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.SubscriptionUser "GetSubscriptionUserBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdSubscriptionUser(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.UserSubscription().GetByIdUser(context.Background(), &models.SubscriptionUserPrimeryKey{
		User_id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Subscription User:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateSubscriptionUser godoc
// @ID update_Subscription_User
// @Router /subscription-user/{id} [PUT]
// @Summary Update Subscription User
// @Description Update Subscription User
// @Tags Subscription User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param subscription body models.UpdateSubscriptionUserSwag true "UpdateSubscriptionUserRequestBody"
// @Success 202 {object} models.SubscriptionUser "UpdateSubscriptioUserBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateSubscriptionUser(c *gin.Context) {

	var (
		subscriptionuser models.UpdateSubscriptionUser
	)

	id := c.Param("id")

	err := c.ShouldBindJSON(&subscriptionuser)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.UserSubscription().Update(context.Background(), &models.UpdateSubscriptionUser{
		Id:              id,
		User_id:         subscriptionuser.User_id,
		Subscription_id: subscriptionuser.Subscription_id,
	})
	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.UserSubscription().GetById(context.Background(), &models.SubscriptionUserPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// GetUserSubscription godoc
// @ID GetUserSubscription
// @Router /subscription-get/{id} [GET]
// @Summary GetUserSubscription
// @Description GetUserSubscription
// @Tags GetUserSubscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Subscription "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetUserSubscription(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.UserSubscription().GetUserSubscription(context.Background(), &models.SubscriptionUserPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Subscription:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}


// Check-access godoc
// @ID CheckAccess
// @Router /Check-access/{id} [GET]
// @Summary Check-access
// @Description Check-access
// @Tags CheckAccess
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Checkaccess "CheckAccessBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CheckAccess(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.UserSubscription().CheckAccess(context.Background(), &models.SubscriptionUserPrimeryKey{
		User_id: id,
	})
	if err != nil {
		log.Println("error whiling get Checkaccess:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}


// MakeSubscriptionActive godoc
// @ID MakeSubscriptionActive
// @Router /make-active/{id} [GET]
// @Summary MakeSubscriptionActive
// @Description MakeSubscriptionActive
// @Tags MakeSubscriptionActive
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.MakeSubscriptionActive "MakeSubscriptionActiveBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) MakeSubscriptionActive(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.UserSubscription().MakeSubscriptionActive(context.Background(), &models.SubscriptionUserPrimeryKey{
		User_id: id,
	})

	if err != nil {
		log.Println("error whiling get MakeSubscriptionActive:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}