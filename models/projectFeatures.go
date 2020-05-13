package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NProjectFeatures struct {
	ProjectFeaturesId   int32  `xorm:"'project_features_id' autoincr pk" json:"project_features_id"`
	ProjectFeaturesType string `xorm:"'project_features_type'" json:"project_features_type"`
	NProject            `xorm:"extends"`
}

func NewNProjectFeatures() *NProjectFeatures {
	return new(NProjectFeatures)
}

func (n *NProjectFeatures) GetProjectFeatures() ([]NProjectFeatures, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NProjectFeatures
	err := dbRead.Join("LEFT", "n_project", "n_project.project_id = n_project_features.project_id").OrderBy("n_project_features.project_features_id DESC").Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
