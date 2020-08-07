package persist

import (
	"log"

	"github.com/olivere/elastic"
	"learngo/crawler/engine"
	"learngo/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("item:%s save error:%s", item, err)
	}

	return err
}
