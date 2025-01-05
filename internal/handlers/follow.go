package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/usescase"
)

type FollowHandler struct {
	FollowUserUsesCase usescase.FollowUser
}

func NewFollowHandler(followUser usescase.FollowUser) *FollowHandler {
	return &FollowHandler{FollowUserUsesCase: followUser}
}

func (f *FollowHandler) FollowUser(c echo.Context) error {
	var follow domain.Follow
	if err := c.Bind(&follow); err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}
	err := f.FollowUserUsesCase.FollowUser(follow)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusCreated, follow)

}

func (f *FollowHandler) UnfollowUser(c echo.Context) error {
	var follow domain.Follow
	if err := c.Bind(&follow); err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}
	err := f.FollowUserUsesCase.UnfollowUser(follow)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusCreated, follow)
}
