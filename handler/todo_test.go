package handler_test

import (
	"app/domain/model"
	"app/handler"
	"app/handler/validator"
	"app/usecase"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockTodo struct {
	usecase.Todo
	mockCreate  func() error
	mockUpdate  func() error
	mockDelete  func() error
	mockFind    func() (*model.Todo, error)
	mockFindAll func() ([]*model.Todo, error)
}

func (m *mockTodo) Create(task string) error {
	return m.mockCreate()
}
func (m *mockTodo) Update(id int, task string, status model.TaskStatus) error {
	return m.mockUpdate()
}
func (m *mockTodo) Delete(id int) error {
	return m.mockDelete()
}

func (m *mockTodo) Find(id int) (*model.Todo, error) {
	return m.mockFind()
}
func (m *mockTodo) FindAll() ([]*model.Todo, error) {
	return m.mockFindAll()
}

func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		request          handler.CreateRequestParam
		usecase          usecase.Todo
		want_status_code int
	}{
		{
			name: "正常系_タスクの登録ができること",
			request: handler.CreateRequestParam{
				Task: "test",
			},
			usecase: &mockTodo{
				mockCreate: func() error {
					return nil
				},
			},
			want_status_code: http.StatusCreated,
		},
		{
			name:             "異常系_必須項目がなかった場合バリデーションエラーになること（task）",
			request:          handler.CreateRequestParam{},
			want_status_code: http.StatusBadRequest,
		},
		{
			name: "異常系_タスクの登録に失敗した場合",
			request: handler.CreateRequestParam{
				Task: "test",
			},
			usecase: &mockTodo{
				mockCreate: func() error {
					return errors.New("xxxx error")
				},
			},
			want_status_code: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := handler.NewTodo(tt.usecase)
			reqJSON, _ := json.Marshal(tt.request)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			validator.SetupValidator()

			r.POST("/", h.Create)
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqJSON))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want_status_code != rec.Code {
				t.Errorf("want = %v, got = %v", tt.want_status_code, rec.Code)
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		request          handler.UpdateRequestParam
		usecase          usecase.Todo
		want_status_code int
	}{
		{
			name: "正常系_タスクの更新ができること",
			request: handler.UpdateRequestParam{
				ID:     1,
				Task:   "test",
				Status: model.Created,
			},
			usecase: &mockTodo{
				mockUpdate: func() error {
					return nil
				},
			},
			want_status_code: http.StatusOK,
		},
		{
			name: "異常系_IDが指定されていなかった場合404エラーになること",
			request: handler.UpdateRequestParam{
				Task:   "test",
				Status: model.Created,
			},
			usecase: &mockTodo{
				mockUpdate: func() error {
					return nil
				},
			},
			want_status_code: http.StatusNotFound,
		},
		{
			name: "異常系_必須項目がなかった場合バリデーションエラーになること（task）",
			request: handler.UpdateRequestParam{
				ID:     1,
				Status: model.Created,
			},
			want_status_code: http.StatusBadRequest,
		},
		{
			name: "異常系_必須項目がなかった場合バリデーションエラーになること（status）",
			request: handler.UpdateRequestParam{
				ID:   1,
				Task: "task",
			},
			want_status_code: http.StatusBadRequest,
		},
		{
			name: "異常系_ステータスに不正な値が指定された場合バリデーションエラーになること（status）",
			request: handler.UpdateRequestParam{
				ID:     1,
				Task:   "task",
				Status: "sss",
			},
			want_status_code: http.StatusBadRequest,
		},
		{
			name: "異常系_タスクの更新に失敗した場合",
			request: handler.UpdateRequestParam{
				ID:     1,
				Task:   "test",
				Status: model.Created,
			},
			usecase: &mockTodo{
				mockUpdate: func() error {
					return errors.New("xxxx error")
				},
			},
			want_status_code: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := handler.NewTodo(tt.usecase)
			reqJSON, _ := json.Marshal(tt.request)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			validator.SetupValidator()

			r.PUT("/:id", h.Update)
			var id string
			if tt.request.ID != 0 {
				id = fmt.Sprintf("%d", tt.request.ID)
			}
			req := httptest.NewRequest("PUT", "/"+id, bytes.NewBuffer(reqJSON))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want_status_code != rec.Code {
				t.Errorf("want = %v, got = %v", tt.want_status_code, rec.Code)
			}
		})
	}

}

func TestDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		request          handler.DeleteRequestParam
		usecase          usecase.Todo
		want_status_code int
	}{
		{
			name: "正常系_タスクの削除ができること",
			request: handler.DeleteRequestParam{
				ID: 1,
			},
			usecase: &mockTodo{
				mockDelete: func() error {
					return nil
				},
			},
			want_status_code: http.StatusOK,
		},
		{
			name:    "異常系_IDが指定されていなかった場合404エラーになること",
			request: handler.DeleteRequestParam{},
			usecase: &mockTodo{
				mockDelete: func() error {
					return nil
				},
			},
			want_status_code: http.StatusNotFound,
		},
		{
			name: "異常系_タスクの削除に失敗した場合",
			request: handler.DeleteRequestParam{
				ID: 1,
			},
			usecase: &mockTodo{
				mockDelete: func() error {
					return errors.New("xxxx error")
				},
			},
			want_status_code: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := handler.NewTodo(tt.usecase)
			reqJSON, _ := json.Marshal(tt.request)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			validator.SetupValidator()

			r.DELETE("/:id", h.Delete)
			var id string
			if tt.request.ID != 0 {
				id = fmt.Sprintf("%d", tt.request.ID)
			}
			req := httptest.NewRequest("DELETE", "/"+id, bytes.NewBuffer(reqJSON))
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want_status_code != rec.Code {
				t.Errorf("want = %v, got = %v", tt.want_status_code, rec.Code)
			}
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	expected := model.Todo{
		ID:     1,
		Task:   "task",
		Status: model.Created,
	}

	tests := []struct {
		name             string
		request          handler.FindRequestParam
		usecase          usecase.Todo
		want_status_code int
		want_response    *model.Todo
	}{
		{
			name: "正常系_データの検索ができること",
			request: handler.FindRequestParam{
				ID: 1,
			},
			usecase: &mockTodo{
				mockFind: func() (*model.Todo, error) {
					return &expected, nil
				},
			},
			want_status_code: http.StatusOK,
			want_response:    &expected,
		},
		{
			name:    "異常系_IDが指定されていなかった場合404エラーになること",
			request: handler.FindRequestParam{},
			usecase: &mockTodo{
				mockFind: func() (*model.Todo, error) {
					return nil, nil
				},
			},
			want_status_code: http.StatusNotFound,
		},
		{
			name: "異常系_データの検索に失敗した場合",
			request: handler.FindRequestParam{
				ID: 1,
			},
			usecase: &mockTodo{
				mockFind: func() (*model.Todo, error) {
					return nil, errors.New("xxxx error")
				},
			},
			want_status_code: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := handler.NewTodo(tt.usecase)
			reqJSON, _ := json.Marshal(tt.request)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			validator.SetupValidator()

			r.GET("/:id", h.Find)
			var id string
			if tt.request.ID != 0 {
				id = fmt.Sprintf("%d", tt.request.ID)
			}
			req := httptest.NewRequest("GET", "/"+id, bytes.NewBuffer(reqJSON))
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want_status_code != rec.Code {
				t.Errorf("want = %v, got = %v", tt.want_status_code, rec.Code)
			}

			// statu code = 200 の場合だけレスポンスボディを返すので、
			// この時だけレスポンスボディのチェックを行う
			if rec.Code == http.StatusOK {
				wr, _ := json.Marshal(tt.want_response)
				if string(wr) != rec.Body.String() {
					t.Errorf("want = %v, got = %v", string(wr), rec.Body.String())
				}
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	t.Parallel()

	expected := model.Todo{
		ID:     1,
		Task:   "task",
		Status: model.Created,
	}

	tests := []struct {
		name             string
		usecase          usecase.Todo
		want_status_code int
		want_response    []*model.Todo
	}{
		{
			name: "正常系_タスクの検索ができること（1件）",
			usecase: &mockTodo{
				mockFindAll: func() ([]*model.Todo, error) {
					return []*model.Todo{
						&expected,
					}, nil
				},
			},
			want_status_code: http.StatusOK,
			want_response:    []*model.Todo{&expected},
		},
		{
			// 複数件検索の場合、検索結果が0件でも
			// statu code = 200 でレスポンスを返す
			name: "正常系_タスクの検索ができること（検索結果0件）",
			usecase: &mockTodo{
				mockFindAll: func() ([]*model.Todo, error) {
					return []*model.Todo{}, nil
				},
			},
			want_status_code: http.StatusOK,
			want_response:    []*model.Todo{},
		},
		{
			name: "異常系_タスクの検索に失敗した場合",
			usecase: &mockTodo{
				mockFindAll: func() ([]*model.Todo, error) {
					return nil, errors.New("xxxx error")
				},
			},
			want_status_code: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := handler.NewTodo(tt.usecase)

			gin.SetMode(gin.TestMode)
			r := gin.New()
			validator.SetupValidator()

			r.GET("/", h.FindAll)
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want_status_code != rec.Code {
				t.Errorf("want = %v, got = %v", tt.want_status_code, rec.Code)
			}

			// statu code = 200 の場合だけレスポンスボディを返すので、
			// この時だけレスポンスボディのチェックを行う
			if rec.Code == http.StatusOK {
				wr, _ := json.Marshal(tt.want_response)
				if string(wr) != rec.Body.String() {
					t.Errorf("want = %v, got = %v", string(wr), rec.Body.String())
				}
			}
		})
	}
}
