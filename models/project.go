package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NProject struct {
	ProjectId   int32  `xorm:"'project_id' autoincr pk" json:"project_id"`
	ProjectName string `xorm:"'project_name'" json:"project_name"`
	ProjectContent string `xorm:"'project_content'" json:"project_content"`
}

func NewNProject() *NProject {
	nProject := new(NProject)
	return nProject
}

func (n *NProject) GetProjects() ([]NProject, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NProject
	err := dbRead.Cols("project_id,project_name").OrderBy("project_id DESC").Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
func (n *NProject) GetProjectsFeatured() ([]NProject, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NProject
	err := dbRead.Join("INNER" , "n_project_features" , "n_project.project_id=n_project_features.project_id").GroupBy("n_project.project_id,n_project.project_name,n_project.project_content").OrderBy("n_project.project_id DESC").Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
