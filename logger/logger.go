package logger

import (
	"fmt"
	"os"

	"github.com/kgrvamsi/throne/conf"

	"github.com/sirupsen/logrus"
	"github.com/sohlich/elogrus"
	elastic "gopkg.in/olivere/elastic.v5"
)

//GetLogger ... Will return the logging object
func GetLogger(envType string, logType string) (*logrus.Logger, error) {
	var (
		logger    = logrus.New()
		config    = conf.GetConf()
		clientUrl string
	)

	if envType == "production" {
		if logType == "local" {
			// Log as JSON instead of the default ASCII formatter.
			logger.SetFormatter(&logrus.TextFormatter{
				TimestampFormat: "2006-01-02T15:04:05.000",
				FullTimestamp:   true,
			})

			// Output to stdout instead of the default stderr
			// Can be any io.Writer, see below for File example
			logger.SetOutput(os.Stdout)

			// Only log the warning severity or above.
			logger.SetLevel(logrus.InfoLevel)

			return logger, nil
		} else if logType == "elasticsearch" {

			if config.Log.ESHttps {
				clientUrl = fmt.Sprintf("https://%s:%s", config.Log.ESUrl, config.Log.ESPort)
			} else {
				clientUrl = fmt.Sprintf("http://%s:%s", config.Log.ESUrl, config.Log.ESPort)
			}

			client, err := elastic.NewClient(elastic.SetSniff(false),
				elastic.SetBasicAuth(config.Log.ESUsername, config.Log.ESPassword),
				elastic.SetURL(clientUrl))
			if err != nil {
				return nil, err
			}
			hook, err := elogrus.NewElasticHook(client, clientUrl, logrus.DebugLevel, config.Log.ESIndexName)
			if err != nil {
				return nil, err
			}
			logger.Hooks.Add(hook)

			return logger, nil
		}
	} else if envType == "development" {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		logger.SetOutput(os.Stdout)

		// Only log the warning severity or above.
		logger.SetLevel(logrus.DebugLevel)

		return logger, nil
	}
	return logger, nil
}
