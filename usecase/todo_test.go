package usecase_test

import (
	"app/domain/model"
	"app/domain/repository"
	"app/usecase"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mockTodo struct {
	repository.Todo
	mockCreate  func() error
	mockDelete  func() error
	mockUpdate  func() error
	mockFind    func() (*model.Todo, error)
	mockFindAll func() ([]*model.Todo, error)
}

func (m *mockTodo) Create(t *model.Todo) error {
	return m.mockCreate()
}
func (m *mockTodo) Delete(id int) error {
	return m.mockDelete()
}
func (m *mockTodo) Update(t *model.Todo) error {
	return m.mockUpdate()
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
		name       string
		task       string
		repository repository.Todo
		err        error
	}{
		{
			name: "正常系_タスクの登録ができること",
			task: "task",
			repository: &mockTodo{
				mockCreate: func() error {
					return nil
				},
			},
			err: nil,
		},
		{
			name: "異常系_タスクの登録に失敗した場合エラーが返ること",
			task: "task",
			repository: &mockTodo{
				mockCreate: func() error {
					return errors.New("xxxx error")
				},
			},
			err: errors.New("xxxx error"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := usecase.NewTodo(tt.repository)

			got := u.Create(tt.task)
			if !equalError(got, tt.err) {
				t.Errorf("different than expected...")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		id         int
		task       string
		status     model.TaskStatus
		repository repository.Todo
		err        error
	}{
		{
			name:   "正常系_タスクの更新ができること",
			id:     1,
			task:   "task",
			status: model.Created,
			repository: &mockTodo{
				mockUpdate: func() error {
					return nil
				},
			},
			err: nil,
		},
		{
			name:   "異常系_タスクの更新に失敗した場合エラーが返ること",
			id:     1,
			task:   "task",
			status: model.Created,
			repository: &mockTodo{
				mockUpdate: func() error {
					return errors.New("xxxx error")
				},
			},
			err: errors.New("xxxx error"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := usecase.NewTodo(tt.repository)

			got := u.Update(tt.id, tt.task, tt.status)
			if !equalError(got, tt.err) {
				t.Errorf("different than expected...")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		id         int
		repository repository.Todo
		err        error
	}{
		{
			name: "正常系_タスクの削除ができること",
			id:   1,
			repository: &mockTodo{
				mockDelete: func() error {
					return nil
				},
			},
			err: nil,
		},
		{
			name: "異常系_タスクの削除に失敗した場合エラーが返ること",
			id:   1,
			repository: &mockTodo{
				mockDelete: func() error {
					return errors.New("xxxx error")
				},
			},
			err: errors.New("xxxx error"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := usecase.NewTodo(tt.repository)

			got := u.Delete(tt.id)
			if !equalError(got, tt.err) {
				t.Errorf("different than expected...")
			}
		})
	}
}

func TestFind(t *testing.T) {
	t.Parallel()
	td := model.Todo{
		ID:     1,
		Task:   "task",
		Status: model.Created,
	}
	tests := []struct {
		name       string
		id         int
		repository repository.Todo
		expected   *model.Todo
		err        error
	}{
		{
			name: "正常系_タスクの検索ができること",
			id:   1,
			repository: &mockTodo{
				mockFind: func() (*model.Todo, error) {
					return &model.Todo{
						ID:     td.ID,
						Task:   td.Task,
						Status: td.Status,
					}, nil
				},
			},
			expected: &model.Todo{
				ID:     td.ID,
				Task:   td.Task,
				Status: td.Status,
			},
			err: nil,
		},
		{
			name: "異常系_タスクの検索に失敗した場合エラーが返ること",
			id:   1,
			repository: &mockTodo{
				mockFind: func() (*model.Todo, error) {
					return nil, errors.New("xxxx error")
				},
			},
			expected: nil,
			err:      errors.New("xxxx error"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := usecase.NewTodo(tt.repository)

			got, err := u.Find(tt.id)
			if !cmp.Equal(got, tt.expected) {
				t.Errorf("diff %s", cmp.Diff(got, tt.expected))
			}
			if !equalError(err, tt.err) {
				t.Errorf("different than expected...")
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	t.Parallel()
	td := model.Todo{
		ID:     1,
		Task:   "task",
		Status: model.Created,
	}
	tests := []struct {
		name       string
		repository repository.Todo
		expected   []*model.Todo
		err        error
	}{
		{
			name: "正常系_タスクの検索ができること",
			repository: &mockTodo{
				mockFindAll: func() ([]*model.Todo, error) {
					return []*model.Todo{
						&td,
					}, nil
				},
			},
			expected: []*model.Todo{&td},
			err:      nil,
		},
		{
			name: "異常系_タスクの検索に失敗した場合エラーが返ること",
			repository: &mockTodo{
				mockFindAll: func() ([]*model.Todo, error) {
					return nil, errors.New("xxxx error")
				},
			},
			expected: nil,
			err:      errors.New("xxxx error"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := usecase.NewTodo(tt.repository)

			got, err := u.FindAll()
			if !cmp.Equal(got, tt.expected) {
				t.Errorf("diff %s", cmp.Diff(got, tt.expected))
			}
			if !equalError(err, tt.err) {
				t.Errorf("different than expected...")
			}
		})
	}
}

func equalError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}
