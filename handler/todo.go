package handler

import (
	"app/domain/model"
	"app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	FindAll(c *gin.Context)
}

type todoHandler struct {
	usecase usecase.Todo
}

func NewTodo(u usecase.Todo) Todo {
	return &todoHandler{u}
}

type CreateRequestParam struct {
	Task string `json:"task" binding:"required,max=60"`
}

func (t *todoHandler) Create(c *gin.Context) {
	var req CreateRequestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.usecase.Create(req.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, nil)

}

type UpdateRequestPathParam struct {
	ID int `uri:"id"`
}

type UpdateRequestBodyParam struct {
	Task   string           `json:"task" binding:"required,max=60"`
	Status model.TaskStatus `json:"status" binding:"required,task_status"`
}

func (t *todoHandler) Update(c *gin.Context) {
	var pathParam UpdateRequestPathParam
	var bodyParam UpdateRequestBodyParam

	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.usecase.Update(pathParam.ID, bodyParam.Task, bodyParam.Status); err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type DeleteRequestParam struct {
	ID int `uri:"id"`
}

func (t *todoHandler) Delete(c *gin.Context) {
	var req DeleteRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.usecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

type FindRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (t *todoHandler) Find(c *gin.Context) {
	var req FindRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := t.usecase.Find(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (t *todoHandler) FindAll(c *gin.Context) {
	res, err := t.usecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
