package handler

import (
	"add/models"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSubscription godoc
// @ID create_Subscription
// @Router /subscription [POST]
// @Summary Create Subscription
// @Description Create Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param subscription body models.CreateSubscription true "CreateSubscriptionRequestBody"
// @Success 201 {object} models.Subscription "GetSubscriptionSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateSubscription(c *gin.Context) {

	var subscription models.CreateSubscription

	err := c.ShouldBindJSON(&subscription)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Subscription().Insert(context.Background(), &subscription)

	if err != nil {
		log.Println("error whiling create Subscription:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := h.storage.Subscription().GetByID(context.Background(), &models.SubscriptionPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Subscription:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIDSubscription godoc
// @ID get_by_id_Subscription
// @Router /subscription/{id} [GET]
// @Summary Get By ID Subscription
// @Description Get By ID Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Subscription "GetSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIdSubscription(c *gin.Context) {

	id := c.Param("id")

	res, err := h.storage.Subscription().GetByID(context.Background(), &models.SubscriptionPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Println("error whiling get by id Subscription:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListSubscription godoc
// @ID get_list_Subscription
// @Router /subscription [GET]
// @Summary Get List Subscription
// @Description Get List Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListSubscriptionResponse "GetSubscriptionListBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListSubscription(c *gin.Context) {
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

	res, err := h.storage.Subscription().GetList(context.Background(), &models.GetListSubscriptionRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateSubscription godoc
// @ID update_Subscription
// @Router /subscription/{id} [PUT]
// @Summary Update Subscription
// @Description Update Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param subscription body models.UpdateSubscriptionSwag true "UpdateSubscriptionRequestBody"
// @Success 202 {object} models.Subscription "UpdateSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateSubscription(c *gin.Context) {

	var (
		subscription models.UpdateSubscription
	)

	id := c.Param("id")

	err := c.ShouldBindJSON(&subscription)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.storage.Subscription().Update(context.Background(), &models.UpdateSubscription{
		Id:            id,
		Title_ru:      subscription.Title_ru,
		Title_en:      subscription.Title_en,
		Title_uz:      subscription.Title_uz,
		Price:         subscription.Price,
		Duration_type: subscription.Duration_type,
		Duration:      subscription.Duration,
		Free_trial:    subscription.Free_trial,
		UpdatedAt:     subscription.UpdatedAt,
	})
	fmt.Println(subscription.Duration)
	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	resp, err := h.storage.Subscription().GetByID(context.Background(), &models.SubscriptionPrimeryKey{
		Id: id,
	})
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusAccepted, resp)
}

// DeleteSubscription godoc
// @ID delete_Subscription
// @Router /subscription/{id} [DELETE]
// @Summary Delete Subscription
// @Description Delete Subscription
// @Tags Subscription
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} models.Empty "DeleteSubscriptionBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteSubscription(c *gin.Context) {

	id := c.Param("id")

	err := h.storage.Subscription().Delete(context.Background(), &models.SubscriptionPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Subscription:", err.Error())
		c.JSON(http.StatusNoContent, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "Deletet Subscription")
}
