package service

import (
	"github.com/sddysz/leanote/app/db"
	"github.com/sddysz/leanote/app/info"
	//	"time"
	//	"sort"
)

type SuggestionService struct {
}

// 得到某博客具体信息
func (this *SuggestionService) AddSuggestion(suggestion info.Suggestion) bool {

	_, err := db.Engine.Insert(&suggestion)
	return err == nil
}
