package service

import (
	"github.com/sddysz/leanote/app/db"
	"github.com/sddysz/leanote/app/info"
	//	. "github.com/sddysz/leanote/app/lea"
	"time"
)

const IMAGE_TYPE = 0

type AlbumService struct {
}

// add album
func (this *AlbumService) AddAlbum(album info.Album) bool {
	album.CreatedTime = time.Now()
	album.Type = IMAGE_TYPE
	affected, err := db.Engine.Insert(&album)
	return err == nil
}

// get albums
func (this *AlbumService) GetAlbums(userId string) []info.Album {
	albums := []info.Album{}
	db.Engine.Where("UserId = ?", userId).Find(&albums)

	return albums
}

// delete album
// presupposition: has no images under this ablum
func (this *AlbumService) DeleteAlbum(userId, albumId string) (bool, string) {
	file := info.File{}
	total, err := db.Engine.Where("AlbumId=?", albumId).And("UserId=?", userId).Count(&file)
	if total == 0 {
		album := info.Album{}
		affected, err := db.Engine.Where("AlbumId=?", albumId).And("UserId=?", userId).Delete(album)
		return err == nil, ""
	}
	return false, "has images"
}

// update album name
func (this *AlbumService) UpdateAlbum(albumId, userId, name string) bool {
	album := info.Album{}
	affected, err := db.Engine.Where("AlbumId=?", albumId).And("UserId=?", userId).Cols("Name").Update(album)
	return err == nil
}
