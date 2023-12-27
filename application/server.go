package server

import (
	"distributed_database_server/config"
	"os"
	"time"

	"bitbucket.org/hasaki-tech/zeus/package/elastic"
	"bitbucket.org/hasaki-tech/zeus/package/kafka"
	"bitbucket.org/hasaki-tech/zeus/package/sql/mysql"
	"bitbucket.org/hasaki-tech/zeus/package/transaction"
	"go.etcd.io/etcd/proxy/grpcproxy/cache"
	"go.uber.org/zap"
	"google.golang.org/api/translate/v2"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg *config.Config
	lib *models.Lib
}

func NewServer(cfg *config.Config) *Server {
	redisCli := redisPkg.InitConnection(&cfg.Redis)
	cache := cache.NewCache(redisCli, 30*time.Second)

	db, _, err := mysql.InitConnection(&cfg.Mysql)
	if err != nil {
		zap.S().Fatalf("Mysql Init error %v", err)
	}

	cert, err := os.ReadFile("cert/ca.pem")
	if err != nil {
		panic(err)
	}

	esClient, err := elastic.NewClient(cfg.ES, cert)
	if err != nil {
		zap.S().Fatalf("init es client failed with err=%v", err)
	}
	initTools(cfg)

	rpcServer := grpcPkg.CreateServer()
	if os.Getenv("ENVIRONMENT") == "local" {
		reflection.Register(rpcServer)
	}

	kafkaPublisher := kafka.NewPublisher(cfg.Kafka)

	translate, err := translate.New("././application/translate/language")
	if err != nil {
		zap.S().Fatalf("init translate client failed with err=%v", err)
	}

	ts := transaction.InitTransactionRedis(cache, nil)

	return &Server{
		cfg:     cfg,
		grpcSvc: rpcServer,
		lib: &models.Lib{
			Es:             esClient,
			Db:             db,
			RedisCli:       redisCli,
			Trans:          translate,
			Cache:          cache,
			KafkaPublisher: kafkaPublisher,
			Transaction:    ts,
		},
	}
}

func (s *Server) Start() {
	s.run()
}


func (s *Server) run() {
	
	listen, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		panic(err)
	}
	
}
