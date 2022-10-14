package main

import (
	pb "analytics/server/analytics/pb"
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedRouteAnalyticsServiceServer
}

func (s *server) SessionStartedEvent(ctx context.Context, in *pb.SessionStartedEventRequest) (*pb.SessionStartedEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params IsNewUser: %v", in.NewUser)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("------- END MESSAGE")
	return &pb.SessionStartedEventResponse{}, nil
}

func (s *server) StageCompletedEvent(ctx context.Context, in *pb.StageCompletedEventRequest) (*pb.StageCompletedEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("-------params AllStars: %v", in.AllStars)
	log.Printf("-------params Attempts: %v", in.Attempts)
	log.Printf("-------params BasicStars: %v", in.BasicStars)
	log.Printf("-------params FirstSuccess: %v", in.FirstSuccess)
	log.Printf("-------params GameMode: %v", in.GameMode)
	log.Printf("-------params IsBonus: %v", in.IsBonus)
	log.Printf("-------params PrevAllStars: %v", in.PrevAllStars)
	log.Printf("-------params PrevBasicStars: %v", in.PrevBasicStars)
	log.Printf("-------params StageId: %v", in.StageId)
	log.Printf("------- END MESSAGE")

	return &pb.StageCompletedEventResponse{}, nil
}

func (s *server) ChapterCompletedEvent(ctx context.Context, in *pb.ChapterCompletedEventRequest) (*pb.ChapterCompletedEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params MaxChapter: %v", in.MaxChapter)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("------- END MESSAGE")

	return &pb.ChapterCompletedEventResponse{}, nil
}

func (s *server) BasicStageCompletedEvent(ctx context.Context, in *pb.BasicStageCompletedEventRequest) (*pb.BasicStageCompletedEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params BasicCompletion: %v", in.BasicCompletion)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("-------params IsBonus: %v", in.IsBonus)
	log.Printf("------- END MESSAGE")

	return &pb.BasicStageCompletedEventResponse{}, nil
}

func (s *server) EventStageCompletedEvent(ctx context.Context, in *pb.EventStageCompletedEventRequest) (*pb.EventStageCompletedEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params EventCompletion: %v", in.EventCompletion)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("------- END MESSAGE")

	return &pb.EventStageCompletedEventResponse{}, nil
}

func (s *server) BasicLevelNewScoreEvent(ctx context.Context, in *pb.BasicLevelNewScoreEventRequest) (*pb.BasicLevelNewScoreEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("-------params StarsCompletion: %v", in.StarsCompletion)
	log.Printf("-------params GameCompletion: %v", in.GameCompletion)
	log.Printf("-------params IsBonus: %v", in.IsBonus)
	log.Printf("-------params FirstSuccess: %v", in.FirstSuccess)
	log.Printf("-------params PrevBasicStars: %v", in.PrevBasicStars)
	log.Printf("-------params PrevAllStars: %v", in.PrevAllStars)
	log.Printf("-------params BasicStars: %v", in.BasicStars)
	log.Printf("-------params Allstars: %v", in.AllStars)
	log.Printf("------- END MESSAGE")

	return &pb.BasicLevelNewScoreEventResponse{}, nil
}

func (s *server) EventLevelNewScoreEvent(ctx context.Context, in *pb.EventLevelNewScoreEventRequest) (*pb.EventLevelNewScoreEventResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	log.Printf("-------params Datetime: %v", in.Datetime)
	log.Printf("-------params EventID: %v", in.EventId)
	log.Printf("-------params SessionCount: %v", in.SessionCount)
	log.Printf("-------params ActiveDays: %v", in.ActiveDays)
	log.Printf("-------params EventFullCompletion: %v", in.EventFullCompletion)
	log.Printf("-------params GameId: %v", in.GameEventId)
	log.Printf("-------params EventStamps: %v", in.EventStamps)
	log.Printf("-------params IsBonus: %v", in.IsBonus)
	log.Printf("-------params FirstSuccess: %v", in.FirstSuccess)
	log.Printf("-------params PrevBasicStars: %v", in.PrevBasicStars)
	log.Printf("-------params PrevAllStars: %v", in.PrevAllStars)
	log.Printf("-------params BasicStars: %v", in.BasicStars)
	log.Printf("-------params Allstars: %v", in.AllStars)
	log.Printf("------- END MESSAGE")

	return &pb.EventLevelNewScoreEventResponse{}, nil
}

func main() {
	// Heroku determines the port via env var
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	log.Printf("Started Analytics server DEV environment on port " + port)

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterRouteAnalyticsServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
