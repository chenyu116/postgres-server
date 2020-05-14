package server

import (
	"context"
	"github.com/chenyu116/postgres-server/config"
	"github.com/chenyu116/postgres-server/models"
	pb "github.com/chenyu116/postgres-server/proto"
	//"time"
)

type apiServer struct {
	cfg config.Config
}

func (s *apiServer) GeneratorProjects(ctx context.Context, req *pb.GeneratorProjectsRequest) (*pb.
GeneratorProjectsReply,
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
	var reply pb.GeneratorProjectsReply
	for _, v := range projects {
		reply.Projects = append(reply.Projects, &pb.ProjectDetails{
			ProjectId:   v.ProjectId,
			ProjectName: v.ProjectName,
		})
	}
	return &reply, nil
}
func (s *apiServer) GeneratorProjectFeatured(ctx context.Context, req *pb.GeneratorProjectFeaturedRequest) (*pb.
GeneratorProjectFeaturedReply,
	error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	ins := models.NewNProject()
	features, err := ins.GetProjectsFeatured()
	if err != nil {
		return nil, err
	}
	var reply pb.GeneratorProjectFeaturedReply
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

func (s *apiServer) GeneratorProjectFeaturesByProjectId(ctx context.Context, req *pb.GeneratorProjectFeaturesByProjectIdRequest) (*pb.
GeneratorProjectFeaturesByProjectIdReply,
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
	var reply pb.GeneratorProjectFeaturesByProjectIdReply
	//featuresMap := make(map[int32]*pb.ProjectFeaturedDetails)
	for _, v := range features {
		reply.Features = append(reply.Features, &pb.ProjectFeatureAll{
			ProjectFeaturesId:     v.ProjectFeaturesId,
			ProjectFeaturesType:   v.ProjectFeaturesType,
			ProjectFeaturesConfig: v.ProjectFeaturesConfig,
			FeatureId:             v.FeatureId,
			FeatureName:           v.FeatureName,
			FeatureLabels:         v.FeatureLabels,
			FeatureTypes:          v.FeatureTypes,
			FeatureVersionId:      v.FeatureVersionId,
			FeatureVersionName:    v.FeatureVersionName,
			FeatureIntro:          v.FeatureIntro,
		})
	}
	return &reply, nil
}
