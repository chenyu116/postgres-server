package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NFeature struct {
	FeatureId     int32  `xorm:"'feature_id' autoincr pk" json:"feature_id"`
	FeatureName   string `xorm:"'feature_name'" json:"feature_name"`
	FeatureLabels string `xorm:"'feature_labels'" json:"feature_labels"`
	FeatureTypes  string `xorm:"'feature_types'" json:"feature_types"`
	FeatureIntro  string `xorm:"'feature_intro'" json:"feature_intro"`
	FeatureOnBoot bool   `xorm:"'feature_onboot'" json:"feature_onboot"`
}

type NFeatureRelation struct {
	NFeature         `xorm:"extends"`
	NProjectFeatures `xorm:"extends"`
}

func NewNFeature() *NFeature {
	return new(NFeature)
}

func (n *NFeature) GetFeatures() ([]NFeature, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NFeature
	err := dbRead.Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}

func (n *NFeature) GetFeatureByFeatureId(featureId int32) ([]NFeatureRelation, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NFeatureRelation
	err := dbRead.SQL("select * from n_feature as f LEFT JOIN n_project_features as pf ON f.feature_id = pf.feature_id where f.feature_id=? limit 1", featureId).Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
