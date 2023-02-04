package cronjob

import (
	"add/storage"
	"context"
	"log"
	"time"
)

type CronJobService struct {
	storage storage.StorageI
	logger  log.Logger
}

func NewCronJobService(s storage.StorageI, logger log.Logger) *CronJobService {

	return &CronJobService{
		storage: s,
		logger:  logger,
	}

}

func CronJob(c *CronJobService) {

	for {

		err:=c.storage.UserSubscription().UpdateUserSubscriptionStatus(context.Background())
		if err!=nil{
		log.Panicln("Error on cron job service : ",err)
		}

		time.Sleep(time.Second * 10)

	}
}
