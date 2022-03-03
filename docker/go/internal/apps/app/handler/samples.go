package handler

// GetOneSample は災害通知単体情報を取得します
func GetOneNotification(c *gin.Context) {
	serviceMaker := c.MustGet(registry.ServiceKey).(registry.ServiceMaker)
	notificationsService := serviceMaker.NewNotifications()

	cond := &model.NotificationSearchCondition{}
	err := c.ShouldBindQuery(cond)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	output, err := notificationsService.GetOne(cond)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, output)
}
