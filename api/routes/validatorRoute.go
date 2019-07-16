package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/cpfCnpj"
	"net/http"
)

type ValidatorRoute struct {
	CpfCnpjService *cpfCnpj.CpfCnpjService
}

type ValidateRequest struct {
	Number string `json:"number"`
}

func NewValidatorRoute(CpfCnpjService *cpfCnpj.CpfCnpjService) *ValidatorRoute {
	return &ValidatorRoute{
		CpfCnpjService: CpfCnpjService,
	}
}

func (h *ValidatorRoute) BuildRoutes(router *gin.RouterGroup) {
	group := router.Group("/v1")
	{
		group.POST("/validate", h.validateCpfCnpj)
		group.POST("/save", h.saveCpfCnpj)
		group.GET("/getAll", h.getAllDocuments)
	}
}

func (h *ValidatorRoute) validateCpfCnpj(c *gin.Context) {
	var request ValidateRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request body is invalid")
		return
	}

	isvalid, err := h.CpfCnpjService.Validate(request.Number)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, isvalid)
	return
}

func (h *ValidatorRoute) saveCpfCnpj(c *gin.Context) {
	var request cpfCnpj.CpfCnpj
	err := c.ShouldBind(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request body is invalid")
		return
	}

	err = h.CpfCnpjService.Save(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
	return
}

func (h *ValidatorRoute) getAllDocuments(c *gin.Context) {
	results, err := h.CpfCnpjService.GetAllDocuments()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
	return
}
