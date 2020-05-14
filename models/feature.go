package models

import (
	_ "github.com/lib/pq"
)

type NFeature struct {
	FeatureId     int32  `xorm:"'feature_id' autoincr pk" json:"feature_id"`
	FeatureName   string `xorm:"'feature_name'" json:"feature_name"`
	FeatureLabels string `xorm:"'feature_labels'" json:"feature_labels"`
	FeatureTypes  string `xorm:"'feature_types'" json:"feature_types"`
	FeatureIntro  string `xorm:"'feature_intro'" json:"feature_intro"`
}

func NewNFeature() *NFeature {
	return new(NFeature)
}
