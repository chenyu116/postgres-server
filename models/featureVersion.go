package models

import (
	_ "github.com/lib/pq"
)

type NFeatureVersion struct {
	FeatureVersionId       int32  `xorm:"'feature_version_id' autoincr pk" json:"feature_version_id"`
	FeatureVersionName     string `xorm:"'feature_version_name'" json:"feature_version_name"`
	FeatureVersionConfig   string `xorm:"'feature_version_config'" json:"feature_version_config"`
	FeatureVersionCreateAt int64  `xorm:"'feature_version_create_at'" json:"feature_version_create_at"`
}

func NewNFeatureVersion() *NFeatureVersion {
	return new(NFeatureVersion)
}

