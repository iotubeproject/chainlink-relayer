package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iotexproject/chainlink-relayer/relayer"
	"go.uber.org/config"
)

type Configuration struct {
	Interval            time.Duration                `yaml:"interval"`
	StartHeight         uint64                       `yaml:"startHeight"`
	DatabaseURL         string                       `yaml:"databaseURL"`
	SourceClientURL     string                       `yaml:"sourceClientURL"`
	TargetClientURL     string                       `yaml:"targetClientURL"`
	TargetChainID       uint32                       `yaml:"targetChainID"`
	PrivateKey          string                       `yaml:"privateKey"`
	AggregatorPairs     map[string]string            `yaml:"aggregatorPairs"`
	ExchangeAggregators map[string]map[string]string `yaml:"exchangeAggregators"`
}

var defaultConfig = Configuration{}

var configFile = flag.String("config", "", "path of config file")

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "-config <filename>")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	opts := []config.YAMLOption{config.Static(defaultConfig), config.Expand(os.LookupEnv)}
	if *configFile != "" {
		opts = append(opts, config.File(*configFile))
	}
	if cf, exists := os.LookupEnv("OVERWRITE_CONFIG_FILE"); exists {
		opts = append(opts, config.File(cf))
	}
	yaml, err := config.NewYAML(opts...)
	if err != nil {
		log.Fatalln(err)
	}
	var cfg Configuration
	if err := yaml.Get(config.Root).Populate(&cfg); err != nil {
		log.Fatalln(err)
	}
	if pk, exists := os.LookupEnv("RELAYER_PRIVATE_KEY"); exists {
		cfg.PrivateKey = pk
	}
	if url, exists := os.LookupEnv("DATABASE_URL"); exists {
		cfg.DatabaseURL = url
	}
	if url, exists := os.LookupEnv("SOURCE_CLIENT_URL"); exists {
		cfg.SourceClientURL = url
	}
	if url, exists := os.LookupEnv("TARGET_CLIENT_URL"); exists {
		cfg.TargetClientURL = url
	}

	service, err := relayer.NewService(
		cfg.Interval,
		cfg.StartHeight,
		cfg.DatabaseURL,
		cfg.SourceClientURL,
		cfg.TargetChainID,
		cfg.TargetClientURL,
		cfg.PrivateKey,
		cfg.AggregatorPairs,
		cfg.ExchangeAggregators,
	)
	if err != nil {
		log.Fatalln(err)
	}
	if err := service.Start(context.Background()); err != nil {
		log.Fatalln(err)
	}
	select {}
}
