package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NProjectFeatures struct {
	ProjectFeaturesId     int32  `xorm:"'project_features_id' autoincr pk" json:"project_features_id"`
	ProjectFeaturesType   string `xorm:"'project_features_type'" json:"project_features_type"`
	ProjectFeaturesConfig string `xorm:"'project_features_config'" json:"project_features_config"`
	NFeature              `xorm:"extends"`
	NFeatureVersion       `xorm:"extends"`
}

func NewNProjectFeatures() *NProjectFeatures {
	return new(NProjectFeatures)
}

func (n *NProjectFeatures) GetProjectFeaturesByProjectId(projectId int32) ([]NProjectFeatures, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NProjectFeatures
	err := dbRead.SQL("select * from n_project_features as pf LEFT JOIN n_feature as f ON pf.feature_id=f.feature_id LEFT JOIN n_feature_version as fv ON pf.feature_version_id=fv.feature_version_id WHERE pf.project_id=?", projectId).Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
