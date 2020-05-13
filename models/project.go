package models

import (
	"fmt"
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
	err := dbRead.OrderBy("project_id DESC").Find(&cond)
	if err != nil {
		fmt.Println(dbRead)
		return nil, err
	}
	return cond, err
}
