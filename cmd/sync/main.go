package main

import (
	"flag"
	"os"

	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/data"
	"github.com/adshao/ordinals-indexer/internal/ord"

	klogrus "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
	debug bool
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.BoolVar(&debug, "debug", false, "debug mode")
}

func newApp(c *conf.Ord, data *data.Data, collectionUc *biz.CollectionUsecase, tokenUc *biz.TokenUsecase, logger log.Logger) (*ord.Syncer, func(), error) {
	return ord.NewSyncer(c, data, collectionUc, tokenUc, logger)
}

func main() {
	flag.Parse()
	var logrusLog = logrus.New()
	logrusLog.Out = os.Stdout
	if debug {
		logrusLog.SetLevel(logrus.DebugLevel)
	} else {
		logrusLog.SetLevel(logrus.InfoLevel)
	}
	logrusLogger := klogrus.NewLogger(logrusLog)
	var logger log.Logger
	if debug {
		logger = log.With(logrusLogger,
			"caller", log.DefaultCaller,
		)
	} else {
		logger = log.With(logrusLogger)
	}
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Ord, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
