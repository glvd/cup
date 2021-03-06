package service

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	c "github.com/glvd/cup/config"
	"log"
	"sync"
)

// Service ...
type Service struct {
	Name string
	Task int
	err  chan error
	serv *machinery.Server
}

var _service *Service
var _once = sync.Once{}

// NewService ...
func NewService(cpuCfg c.Config) *Service {
	_once.Do(func() {
		var cfg = &config.Config{
			Broker:        cpuCfg.Broker,
			DefaultQueue:  cpuCfg.QueueName,
			ResultBackend: cpuCfg.ResultBackend,
			AMQP:          cpuCfg.AMQP,
		}

		server, err := machinery.NewServer(cfg)
		if err != nil {
			log.Fatal(err)
			return
		}
		_service = &Service{
			Name: cpuCfg.Name,
			Task: cpuCfg.Task,
			err:  make(chan error),
			serv: server,
		}
	})

	return _service
}

// NewWorker ...
func (s *Service) NewWorker() {
	worker := s.serv.NewWorker(s.Name, s.Task)
	worker.LaunchAsync(s.err)
	return
}

// HandleWorker ...
func (s *Service) HandleWorker() error {
	return <-s.err
}

// AddTask ...
func (s *Service) Register(name string, val interface{}) error {
	return s.serv.RegisterTask(name, val)
}

// Send ...
func (s *Service) Send(signature *tasks.Signature) (result *result.AsyncResult, e error) {
	return s.serv.SendTask(signature)
}
