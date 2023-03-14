package services

import (
	"context"
	"github.com/hovhannesyan/gsSportBot_EventService/pkg/db"
	"github.com/hovhannesyan/gsSportBot_EventService/pkg/models"
	"github.com/hovhannesyan/gsSportBot_EventService/pkg/pb"
)

type Server struct {
	DbHandler db.Handler
	pb.UnimplementedEventServer
}

func (s *Server) AddDiscipline(ctx context.Context, req *pb.AddDisciplineRequest) (*pb.AddDisciplineResponse, error) {
	discipline := models.Discipline{Name: req.Name}

	if err := s.DbHandler.DB.Create(&discipline).Error; err != nil {
		return nil, err
	}

	return &pb.AddDisciplineResponse{
		Id: discipline.Id,
	}, nil
}

func (s *Server) DeleteDiscipline(ctx context.Context, req *pb.DeleteDisciplineRequest) (*pb.DeleteDisciplineResponse, error) {
	if err := s.DbHandler.DB.Delete(&models.Discipline{}, req.Id).Error; err != nil {
		return &pb.DeleteDisciplineResponse{
			Success: false,
		}, err
	}
	return &pb.DeleteDisciplineResponse{
		Success: true,
	}, nil
}

func (s *Server) GetDisciplines(ctx context.Context, req *pb.GetDisciplinesRequest) (*pb.GetDisciplinesResponse, error) {
	var disciplines []models.Discipline
	if len(req.Id) != 0 {
		if err := s.DbHandler.DB.Where("id in (?)", req.Id).Find(&disciplines).Error; err != nil {
			return nil, err
		}
	} else {
		if err := s.DbHandler.DB.Find(&disciplines).Error; err != nil {
			return nil, err
		}
	}

	data := make([]*pb.DisciplineData, len(disciplines))
	for i, discipline := range disciplines {
		data[i] = &pb.DisciplineData{
			Id:   discipline.Id,
			Name: discipline.Name,
		}
	}

	return &pb.GetDisciplinesResponse{
		Disciplines: data,
	}, nil
}

func (s *Server) AddEvent(ctx context.Context, req *pb.AddEventRequest) (*pb.AddEventResponse, error) {
	event := models.Event{
		DisciplineId: req.DisciplineId,
		Place:        req.Place,
		Date:         req.Date,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Price:        req.Price,
		Limit:        req.Limit,
		Description:  req.Description,
	}

	if err := s.DbHandler.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return &pb.AddEventResponse{
		Id: event.Id,
	}, nil
}

func (s *Server) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	if err := s.DbHandler.DB.Delete(&models.Event{}, req.Id).Error; err != nil {
		return &pb.DeleteEventResponse{
			Success: false,
		}, err
	}
	return &pb.DeleteEventResponse{
		Success: true,
	}, nil
}

func (s *Server) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	var (
		events      []models.Event
		disciplines *pb.GetDisciplinesResponse
		err         error
	)
	dis := make(map[int64]string)

	if len(req.DisciplineId) != 0 {
		if err := s.DbHandler.DB.Where("discipline_id = ?", req.DisciplineId).Find(&events).Error; err != nil {
			return nil, err
		}
		disciplines, err = s.GetDisciplines(ctx, &pb.GetDisciplinesRequest{Id: req.DisciplineId})
		if err != nil {
			return nil, err
		}
	} else {
		if err := s.DbHandler.DB.Find(&events).Error; err != nil {
			return nil, err
		}
		disciplines, err = s.GetDisciplines(ctx, nil)
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(disciplines.Disciplines); i++ {
		dis[disciplines.Disciplines[i].Id] = disciplines.Disciplines[i].Name
	}

	data := make([]*pb.EventData, len(events))
	for i, event := range events {
		data[i] = &pb.EventData{
			Id:             event.Id,
			DisciplineName: dis[event.DisciplineId],
			Place:          event.Place,
			Date:           event.Date,
			StartTime:      event.StartTime,
			EndTime:        event.EndTime,
			Price:          event.Price,
			Limit:          event.Limit,
			Description:    event.Description,
		}
	}
	return &pb.GetEventsResponse{
		Events: data,
	}, nil
}
