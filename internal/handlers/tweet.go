package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/usescase"
)

type TweetHandler struct {
	TweetHandlerUsesCase usescase.CreateTweet
}

func NewTweetHandler(tweetHandlerUsesCase usescase.CreateTweet) *TweetHandler {
	return &TweetHandler{TweetHandlerUsesCase: tweetHandlerUsesCase}
}

func (t *TweetHandler) CreateTweet(c echo.Context) error {
	fmt.Println("CreateTweet")
	var tweet domain.Tweet
	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := t.TweetHandlerUsesCase.CreateTweet(tweet); err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusCreated, tweet)
}
