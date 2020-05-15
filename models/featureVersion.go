package models

import (
	"github.com/chenyu116/postgres-server/utils"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type NFeatureVersion struct {
	FeatureVersionId       int32  `xorm:"'feature_version_id' autoincr pk" json:"feature_version_id"`
	FeatureId              int32  `xorm:"'feature_id'" json:"feature_id"`
	FeatureVersionName     string `xorm:"'feature_version_name'" json:"feature_version_name"`
	FeatureVersionConfig   string `xorm:"'feature_version_config'" json:"feature_version_config"`
	FeatureVersionCreateAt int64  `xorm:"'feature_version_create_at'" json:"feature_version_create_at"`
}

func NewNFeatureVersion() *NFeatureVersion {
	return new(NFeatureVersion)
}

func (n *NFeatureVersion) GetFeatureVersion() ([]NFeatureVersion, error) {
	dbRead := utils.PgPoolRead.Get().(*xorm.Engine)
	defer func() {
		utils.PgPoolRead.Put(dbRead)
		dbRead = nil
	}()
	var cond []NFeatureVersion
	err := dbRead.Find(&cond)
	if err != nil {
		return nil, err
	}
	return cond, err
}
