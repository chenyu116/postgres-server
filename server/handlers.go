package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/chenyu116/postgres-server/config"
	"github.com/chenyu116/postgres-server/models"
	pb "github.com/chenyu116/postgres-server/proto"
)

type apiServer struct {
	cfg config.Config
}

func (s *apiServer) Projects(ctx context.Context, req *pb.ProjectsRequest) (*pb.
ProjectsReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	nProject := models.NewNProject()
	projects, err := nProject.GetProjects()

	if err != nil {
		return nil, err
	}
	var reply pb.ProjectsReply
	for _, v := range projects {
		reply.Projects = append(reply.Projects, &pb.ProjectDetails{
			ProjectId:   v.ProjectId,
			ProjectName: v.ProjectName,
		})
	}
	return &reply, nil
}

func (s *apiServer) Features(ctx context.Context, req *pb.FeaturesRequest) (*pb.
FeaturesReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	f := models.NewNFeature()
	features, err := f.GetFeatures()
	if err != nil {
		return nil, err
	}
	fv := models.NewNFeatureVersion()
	featureVersion, err := fv.GetFeatureVersion()
	if err != nil {
		return nil, err
	}
	var reply pb.FeaturesReply
	featureVersionMap := make(map[int32][]*pb.FeatureVersion)
	for _, v := range featureVersion {
		if _, ok := featureVersionMap[v.FeatureId]; !ok {
			featureVersionMap[v.FeatureId] = []*pb.FeatureVersion{}
		}
		featureVersionMap[v.FeatureId] = append(featureVersionMap[v.FeatureId], &pb.FeatureVersion{
			FeatureVersionId:     v.FeatureVersionId,
			FeatureVersionName:   v.FeatureVersionName,
			FeatureVersionConfig: v.FeatureVersionConfig,
		})
	}
	for _, v := range features {
		if getFeatureVersion, ok := featureVersionMap[v.FeatureId]; ok {
			reply.Feature = append(reply.Feature, &pb.Feature{
				FeatureId:      v.FeatureId,
				FeatureName:    v.FeatureName,
				FeatureLabels:  v.FeatureLabels,
				FeatureTypes:   v.FeatureTypes,
				FeatureIntro:   v.FeatureIntro,
				FeatureVersion: getFeatureVersion,
				FeatureOnboot:  v.FeatureOnBoot,
			})
		}
	}
	return &reply, nil
}

func (s *apiServer) Feature(ctx context.Context, req *pb.FeatureRequest) (*pb.
FeatureReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	f := models.NewNFeature()
	features, err := f.GetFeatureByFeatureId(req.FeatureId)
	if err != nil {
		return nil, err
	}
	if len(features) == 0 {
		return nil, errors.New("no result")
	}
	fv := models.NewNFeatureVersion()
	featureVersion, err := fv.GetFeatureVersionByFeatureId(req.FeatureId)
	if err != nil {
		return nil, err
	}
	var reply pb.FeatureReply
	var featureVersionMap []*pb.FeatureVersion
	for _, v := range featureVersion {
		featureVersionMap = append(featureVersionMap, &pb.FeatureVersion{
			FeatureVersionId:     v.FeatureVersionId,
			FeatureVersionName:   v.FeatureVersionName,
			FeatureVersionConfig: v.FeatureVersionConfig,
			FeatureVersionIntro:  v.FeatureVersionIntro,
		})
	}
	v := features[0]
	reply.Feature = &pb.Feature{
		FeatureId:             v.NFeature.FeatureId,
		FeatureName:           v.FeatureName,
		FeatureLabels:         v.FeatureLabels,
		FeatureTypes:          v.FeatureTypes,
		FeatureIntro:          v.FeatureIntro,
		FeatureVersion:        featureVersionMap,
		ProjectFeaturesId:     v.ProjectFeaturesId,
		ProjectFeaturesType:   v.ProjectFeaturesType,
		FeatureVersionId:      v.FeatureVersionId,
		FeatureOnboot:         v.NFeature.FeatureOnBoot,
		ProjectFeaturesConfig: v.ProjectFeaturesConfig,
		FeatureReuse:          v.FeatureReUse,
	}
	return &reply, nil
}

func (s *apiServer) ProjectInitialized(ctx context.Context, req *pb.ProjectInitializedRequest) (*pb.
ProjectInitializedReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	ins := models.NewNProject()
	features, err := ins.GetProjectsInitialized(req.Projects)
	if err != nil {
		return nil, err
	}
	var reply pb.ProjectInitializedReply
	for _, v := range features {
		reply.Projects = append(reply.Projects, &pb.ProjectDetails{
			ProjectId:      v.ProjectId,
			ProjectName:    v.ProjectName,
			ProjectContent: v.ProjectContent,
		})
	}
	//featuresMap := make(map[int32]*pb.ProjectFeaturedDetails)
	//for _, v := range featuresMap {
	//	reply.Projects = append(reply.Projects, v)
	//}
	return &reply, nil
}

func (s *apiServer) ProjectFeaturesByProjectId(ctx context.Context, req *pb.ProjectFeaturesByProjectIdRequest) (*pb.
ProjectFeaturesByProjectIdReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	ins := models.NewNProjectFeatures()
	features, err := ins.GetProjectFeaturesByProjectId(req.ProjectId)
	if err != nil {
		return nil, err
	}
	var reply pb.ProjectFeaturesByProjectIdReply
	//featuresMap := make(map[int32]*pb.ProjectFeaturedDetails)
	for _, v := range features {
		reply.Features = append(reply.Features, &pb.ProjectFeatureAll{
			ProjectFeaturesId:          v.ProjectFeaturesId,
			ProjectFeaturesType:        v.ProjectFeaturesType,
			ProjectFeaturesConfig:      v.ProjectFeaturesConfig,
			FeatureId:                  v.NFeature.FeatureId,
			FeatureName:                v.FeatureName,
			FeatureLabels:              v.FeatureLabels,
			FeatureTypes:               v.FeatureTypes,
			FeatureVersionId:           v.NProjectFeatures.FeatureVersionId,
			FeatureVersionName:         v.FeatureVersionName,
			FeatureIntro:               v.FeatureIntro,
			FeatureOnboot:              v.FeatureOnBoot,
			ProjectFeaturesInstallName: v.ProjectFeaturesInstallName,
			ProjectFeaturesRoutePath:   v.ProjectFeaturesRoutePath,
			ProjectFeaturesDeployTo:    v.ProjectFeaturesDeployTo,
			ProjectFeaturesSortOrder:   v.ProjectFeaturesSortOrder,
			FeatureReuse:               v.FeatureReUse,
			ProjectFeaturesName:        v.ProjectFeaturesName,
		})
	}
	return &reply, nil
}

func (s *apiServer) CreateProjectFeature(ctx context.Context, req *pb.CreateProjectFeatureRequest) (*pb.CreateProjectFeatureReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	instance := models.NewNProjectFeatures()
	cond := &models.NProjectFeatures{
		ProjectFeaturesType:        req.GetProjectFeaturesType(),
		ProjectFeaturesConfig:      req.GetProjectFeaturesConfig(),
		ProjectId:                  req.GetProjectId(),
		FeatureId:                  req.GetFeatureId(),
		FeatureVersionId:           req.GetFeatureVersionId(),
		ProjectFeaturesInstallName: req.GetProjectFeaturesInstallName(),
		ProjectFeaturesName:        req.GetProjectFeaturesName(),
	}
	fmt.Printf("%+v", cond)
	err = instance.Insert(cond)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProjectFeatureReply{}, nil
}
func (s *apiServer) UpdateProjectFeature(ctx context.Context, req *pb.UpdateProjectFeatureRequest) (*pb.UpdateProjectFeatureReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	instance := models.NewNProjectFeatures()
	cond := &models.NProjectFeatures{
		ProjectFeaturesType:        req.GetProjectFeaturesType(),
		ProjectFeaturesConfig:      req.GetProjectFeaturesConfig(),
		ProjectFeaturesId:          req.GetProjectFeaturesId(),
		FeatureVersionId:           req.GetFeatureVersionId(),
		ProjectFeaturesName:        req.GetProjectFeaturesName(),
	}
	fmt.Printf("%+v", cond)
	err = instance.Update(cond)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProjectFeatureReply{}, nil
}
