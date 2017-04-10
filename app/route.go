package app

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Index ...
func (s *Service) Index(c *gin.Context) {
	tokens, err := s.TokenStore.All()
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"tokenCount":     len(tokens),
		"noneRegistered": true,
	})
}

// Token ...
func (s *Service) Token(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.AbortWithStatus(400)
	}

	err := s.TokenStore.Save(token)
	if err != nil {
		c.AbortWithError(500, err)
	}

	tokens, err := s.TokenStore.All()
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"tokenCount":     len(tokens),
		"noneRegistered": false,
	})
}
