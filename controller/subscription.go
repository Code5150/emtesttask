package controller

import (
	"net/http"
	"strconv"

	"emtesttask/model"
	"emtesttask/service"

	"github.com/gin-gonic/gin"
)

type SubscriptionController struct {
	service service.SubscriptionService
}

// Провайдер для Wire
func NewSubscriptionController(service service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{service: service}
}

// GetById             godoc
// @Summary      Get subscription by id
// @Description  Responds with the subscription
// @Tags         Subscriptions
// @Produce      json
// @Param 		 id path uint64 true "Subscription ID"
// @Success      200  {object}  model.SubscriptionDTO
// @Router       /subscriptions/{id} [get]
func (c *SubscriptionController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	subscription, err := c.service.GetSubscriptionByID(ctx, id)
	if err != nil {
		if err.Error() == "subscription not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}

// GetPaged             godoc
// @Summary      Get subscription by page
// @Description  Responds with the subscriptions list
// @Tags         Subscriptions
// @Produce      json
// @Param 		 pagedRequest body model.PagedRequest true "Page filter"
// @Success      200  {object}  []model.SubscriptionDTO
// @Router       /subscriptions/page [post]
func (c *SubscriptionController) GetPaged(ctx *gin.Context) {

	var pagedRequest model.PagedRequest
	if err := ctx.BindJSON(&pagedRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.GetSubscriptionsPaged(ctx, &pagedRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// AddSubscription             godoc
// @Summary      Add new subscription
// @Description  Responds with the subscription
// @Tags         Subscriptions
// @Produce      json
// @Param 		 subscription body model.SubscriptionDTO true "Subscription data to insert"
// @Success      200  {object}  model.SubscriptionDTO
// @Router       /subscriptions/new [put]
func (c *SubscriptionController) AddSubscription(ctx *gin.Context) {

	var subscription model.SubscriptionDTO
	if err := ctx.BindJSON(&subscription); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.AddSubscription(ctx, &subscription)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *SubscriptionController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/subscriptions")
	{
		api.GET("/:id", c.GetByID)
		api.PUT("/new", c.AddSubscription)
		api.POST("/page", c.GetPaged)
		//api.POST("/users", c.Create)
	}
}
