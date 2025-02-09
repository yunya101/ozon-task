package cheker

import (
	"time"

	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/redis"
)

type Cheker struct {
	redis *data.RedisRepo
}

func (c *Cheker) SetRedis(db *data.RedisRepo) {
	c.redis = db
}

func (c *Cheker) CheckPopularity() {

	for {
		config.InfoLog("starting popularity checking")

		posts, err := c.redis.GetAllPosts()

		if err != nil {
			config.InfoLog("stoping popularity checking")
			return
		}

		for _, p := range posts {
			p.CalcPopularity()

			if p.Popularity < config.PopularityThreshold {
				if err = c.redis.DeletePost(p.ID); err != nil {
					config.ErrorLog(err)
					return
				}
			}
		}
		config.InfoLog("popularity calculate successfully")
		time.Sleep(time.Hour * 3)
	}
}
