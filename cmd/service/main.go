package main

import (
	"flag"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/pkg/setting"
	"github.com/vietpham1023/golang-uit-hackathon/server"
	"go.uber.org/zap"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.Parse()

	defer setting.WaitOSSignal()

	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	s, err := server.NewServer(cfg)
	if err != nil {
		zap.S().Errorf("create server fail with err: %v", err)
		panic(err)
	}

	s.Init()

	if err := s.Listen(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
