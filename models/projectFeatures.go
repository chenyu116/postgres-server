package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NProjectFeatures struct {
	ProjectFeaturesId     int32  `xorm:"'project_features_id' autoincr pk" json:"project_features_id"`
	ProjectId             int32  `xorm:"'project_id'" json:"project_id"`
	ProjectFeaturesType   string `xorm:"'project_features_type'" json:"project_features_type"`
	ProjectFeaturesConfig string `xorm:"'project_features_config'" json:"project_features_config"`
	FeatureId             int32  `xorm:"'feature_id'" json:"feature_id"`
	FeatureVersionId      int32  `xorm:"'feature_version_id'" json:"feature_version_id"`
}

type NProjectFeaturesRelation struct {
	NProjectFeatures `xorm:"extends"`
	NFeature         `xorm:"extends"`
	NFeatureVersion  `xorm:"extends"`
}

func NewNProjectFeatures() *NProjectFeatures {
	return new(NProjectFeatures)
}

func (n *NProjectFeatures) Insert(cond *NProjectFeatures) error {
	dbWrite := utils.PgPoolWrite.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolWrite.Put(dbWrite)
		dbWrite = nil
	}()
	_, err := dbWrite.InsertOne(cond)
	if err != nil {
		return err
	}
	return nil
}
func (n *NProjectFeatures) GetProjectFeaturesByProjectId(projectId int32) ([]NProjectFeaturesRelation, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NProjectFeaturesRelation
	err := dbRead.SQL("select * from n_project_features as pf LEFT JOIN n_feature as f ON pf.feature_id=f.feature_id LEFT JOIN n_feature_version as fv ON pf.feature_version_id=fv.feature_version_id WHERE pf.project_id=?", projectId).Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
