package controller

import (
	"net/http"
	"strconv"

	"henotech.net/usecase"

	"github.com/labstack/echo/v4"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
	FindAll() echo.HandlerFunc
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

// NewTaskHandler task handlerのコンストラクタ
func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Post taskを保存するときのハンドラー
func (th *taskHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskUsecase.Create(req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      createdTask.ID,
			Title:   createdTask.Title,
			Content: createdTask.Content,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get taskを取得するときのハンドラー
func (th *taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      foundTask.ID,
			Title:   foundTask.Title,
			Content: foundTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		tasks, err := th.taskUsecase.Find()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := make([]responseTask, 0, 0)

		for _, v := range tasks {
			res = append(res, responseTask{
				ID:      v.ID,
				Title:   v.Title,
				Content: v.Content,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"data": res})
	}
}

// Put taskを更新するときのハンドラー
func (th *taskHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.taskUsecase.Update(id, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      updatedTask.ID,
			Title:   updatedTask.Title,
			Content: updatedTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete taskを削除するときのハンドラー
func (th *taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.taskUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
