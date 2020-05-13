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
	ins := models.NewNProjectFeatures()
	features, err := ins.GetProjectFeatures()
	if err != nil {
		return nil, err
	}
	var reply pb.GeneratorProjectFeaturedReply
	featuresMap := make(map[int32]*pb.ProjectFeaturedDetails)
	for _, v := range features {
		if _, ok := featuresMap[v.ProjectId]; !ok {
			featuresMap[v.ProjectId] = &pb.ProjectFeaturedDetails{
				ProjectId:            v.ProjectId,
				ProjectName:          v.ProjectName,
				ProjectContent:       v.NProject.ProjectContent,
			}
		}
		featuresMap[v.ProjectId].ProjectFeatures = append(featuresMap[v.ProjectId].ProjectFeatures, &pb.ProjectFeature{
			ProjectFeaturesId:   v.ProjectFeaturesId,
			ProjectFeaturesType: "",
		})
	}
	for _, v := range featuresMap {
		reply.Projects = append(reply.Projects, v)
	}
	return &reply, nil
}
