package datasource

import (
	"henotech.net/domain/model"
	"henotech.net/domain/repository"

	"gorm.io/gorm"
)

// TaskRepository task repositoryの構造体
type TaskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}

// Create taskの保存
func (tr *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (tr *TaskRepository) Find() (tasks []*model.Task, err error) {
	if err := tr.Conn.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// FindByID taskをIDで取得
func (tr *TaskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskの更新
func (tr *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(&task).Updates(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskの削除
func (tr *TaskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
