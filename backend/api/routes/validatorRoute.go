package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski/validatorCpfCnpj/backend/pkg/cpfCnpj"
	"github.com/rkorpalski/validatorCpfCnpj/backend/pkg/util"
	"net/http"
	"time"
)

type ValidatorRoute struct {
	CpfCnpjService *cpfCnpj.CpfCnpjService
}

type ValidateRequest struct {
	Number string `json:"number"`
}

var startTime time.Time
var buscasRealizadas int

func NewValidatorRoute(CpfCnpjService *cpfCnpj.CpfCnpjService) *ValidatorRoute {
	startTime = time.Now()
	buscasRealizadas = 0
	return &ValidatorRoute{
		CpfCnpjService: CpfCnpjService,
	}
}

func (h *ValidatorRoute) BuildRoutes(router *gin.RouterGroup) {
	group := router.Group("/v1")
	{
		group.POST("/validate", h.validateCpfCnpj)
		group.POST("/save", h.saveCpfCnpj)
		group.POST("/find", h.FindByDocument)
		group.GET("/getDocuments", h.getAllDocuments)
		group.GET("/getBlacklist", h.getBlacklist)
		group.GET("/blacklist/add/:documentId", h.MoveToBlacklist)
		group.GET("/blacklist/remove/:documentId", h.RemoveToBlacklist)
		group.GET("/status", h.serverStatus)
		group.DELETE("/delete/:documentId", h.DeleteDocument)

	}
}

func (h *ValidatorRoute) validateCpfCnpj(c *gin.Context) {
	var request ValidateRequest
	err := c.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request body is invalid")
		return
	}

	isvalid, err := h.CpfCnpjService.Validate(request.Number)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, isvalid)
}

func (h *ValidatorRoute) saveCpfCnpj(c *gin.Context) {
	var request cpfCnpj.CpfCnpj
	err := c.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request body is invalid")
		return
	}

	err = h.CpfCnpjService.Save(request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *ValidatorRoute) getAllDocuments(c *gin.Context) {
	results, err := h.CpfCnpjService.GetAllDocuments(false)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *ValidatorRoute) MoveToBlacklist(c *gin.Context) {
	documentId := c.Param("documentId")
	err := h.CpfCnpjService.MoveToBlacklist(documentId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *ValidatorRoute) DeleteDocument(c *gin.Context) {
	documentId := c.Param("documentId")
	err := h.CpfCnpjService.DeleteDocument(documentId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *ValidatorRoute) getBlacklist(c *gin.Context) {
	results, err := h.CpfCnpjService.GetAllDocuments(true)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *ValidatorRoute) RemoveToBlacklist(c *gin.Context) {
	documentId := c.Param("documentId")
	err := h.CpfCnpjService.RemoveFromBlacklist(documentId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *ValidatorRoute) FindByDocument(c *gin.Context) {
	var request ValidateRequest
	err := c.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request body is invalid")
		return
	}

	buscasRealizadas++
	results, err := h.CpfCnpjService.FindByDocument(request.Number)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *ValidatorRoute) serverStatus(c *gin.Context) {
	startTime, _ := util.FormatTime("02-01-2006 15:04", startTime)

		serverStatus := gin.H{
		"Start time": startTime,
		"Uptime": uptime(),
		"Consultas realizadas": buscasRealizadas,
	}

	c.JSON(http.StatusOK, serverStatus)
}

func uptime() string {
	return time.Since(startTime).String()
}