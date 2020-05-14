// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package postgresServer

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GeneratorProjectFeaturedRequest) Validate() error {
	return nil
}
func (this *GeneratorProjectsRequest) Validate() error {
	return nil
}
func (this *ProjectFeature) Validate() error {
	return nil
}
func (this *ProjectFeaturedDetails) Validate() error {
	for _, item := range this.ProjectFeatures {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ProjectFeatures", err)
			}
		}
	}
	return nil
}
func (this *ProjectDetails) Validate() error {
	return nil
}
func (this *GeneratorProjectsReply) Validate() error {
	for _, item := range this.Projects {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Projects", err)
			}
		}
	}
	return nil
}
func (this *GeneratorProjectFeaturedReply) Validate() error {
	for _, item := range this.Projects {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Projects", err)
			}
		}
	}
	return nil
}
func (this *ProjectFeatureAll) Validate() error {
	return nil
}
func (this *GeneratorProjectFeaturesByProjectIdRequest) Validate() error {
	return nil
}
func (this *GeneratorProjectFeaturesByProjectIdReply) Validate() error {
	for _, item := range this.Features {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Features", err)
			}
		}
	}
	return nil
}
