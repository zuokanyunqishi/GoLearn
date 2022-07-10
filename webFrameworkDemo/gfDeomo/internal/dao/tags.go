// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gfDeomo/internal/dao/internal"
)

// internalTagsDao is internal type for wrapping internal DAO implements.
type internalTagsDao = *internal.TagsDao

// tagsDao is the data access object for table tags.
// You can define custom methods on it to extend its functionality as you wish.
type tagsDao struct {
	internalTagsDao
}

var (
	// Tags is globally public accessible object for table tags operations.
	Tags = tagsDao{
		internal.NewTagsDao(),
	}
)

// Fill with you ideas below.
