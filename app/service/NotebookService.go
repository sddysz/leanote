package service

import (
	//	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sddysz/leanote/app/db"
	"github.com/sddysz/leanote/app/info"
	//	"html"
)

// 笔记本

type NotebookService struct {
}

// 排序
func sortSubNotebooks(eachNotebooks info.SubNotebooks) info.SubNotebooks {
	// 遍历子, 则子往上进行排序
	for _, eachNotebook := range eachNotebooks {
		if eachNotebook.Subs != nil && len(eachNotebook.Subs) > 0 {
			eachNotebook.Subs = sortSubNotebooks(eachNotebook.Subs)
		}
	}

	// 子排完了, 本层排
	sort.Sort(&eachNotebooks)
	return eachNotebooks
}

// 整理(成有关系)并排序
// GetNotebooks()调用
// ShareService调用
func ParseAndSortNotebooks(userNotebooks []info.Notebook, noParentDelete, needSort bool) info.SubNotebooks {
	// 整理成info.Notebooks
	// 第一遍, 建map
	// notebookId => info.Notebooks
	userNotebooksMap := make(map[int64]*info.Notebooks, len(userNotebooks))
	for _, each := range userNotebooks {
		newNotebooks := info.Notebooks{Subs: info.SubNotebooks{}}
		newNotebooks.NotebookId = each.NotebookId
		newNotebooks.Title = each.Title
		//		newNotebooks.Title = html.EscapeString(each.Title)
		newNotebooks.Title = strings.Replace(strings.Replace(each.Title, "<script>", "", -1), "</script", "", -1)
		newNotebooks.Seq = each.Seq
		newNotebooks.UserId = each.UserId
		newNotebooks.ParentNotebookId = each.ParentNotebookId
		newNotebooks.NumberNotes = each.NumberNotes
		newNotebooks.IsTrash = each.IsTrash
		newNotebooks.IsBlog = each.IsBlog

		// 存地址
		userNotebooksMap[each.NotebookId] = &newNotebooks
	}

	// 第二遍, 追加到父下

	// 需要删除的id
	needDeleteNotebookId := map[int64]bool{}
	for id, each := range userNotebooksMap {
		// 如果有父, 那么追加到父下, 并剪掉当前, 那么最后就只有根的元素
		if each.ParentNotebookId != 0 {
			if userNotebooksMap[each.ParentNotebookId] != nil {
				userNotebooksMap[each.ParentNotebookId].Subs = append(userNotebooksMap[each.ParentNotebookId].Subs, each) // Subs是存地址
				// 并剪掉
				// bug
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			} else if noParentDelete {
				// 没有父, 且设置了要删除
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			}
		}
	}

	// 第三遍, 得到所有根
	final := make(info.SubNotebooks, len(userNotebooksMap)-len(needDeleteNotebookId))
	i := 0
	for id, each := range userNotebooksMap {
		if !needDeleteNotebookId[id] {
			final[i] = each
			i++
		}
	}

	// 最后排序
	if needSort {
		return sortSubNotebooks(final)
	}
	return final
}

// 得到某notebook
func (this *NotebookService) GetNotebook(notebookId, userId int64) info.Notebook {
	notebook := info.Notebook{}
	db.Engine.Where("NotebookId=？", notebookId).And("UserId=?", userId).Get(&notebook)
	return notebook
}
func (this *NotebookService) GetNotebookById(notebookId int64) info.Notebook {
	notebook := info.Notebook{}
	db.Engine.Id(notebookId).Get(&notebook)
	return notebook
}
func (this *NotebookService) GetNotebookByUserIdAndUrlTitle(userId int64, notebookIdOrUrlTitle string) info.Notebook {
	notebook := info.Notebook{}
	db.Engine.Where("NotebookId=?", notebookIdOrUrlTitle).Or("UrlTitle=?", notebookIdOrUrlTitle).And("UserId=?", userId).Get(&notebook)
	return notebook
}

// 同步的方法
func (this *NotebookService) GeSyncNotebooks(userId int64, afterUsn, maxEntry int) []info.Notebook {
	notebooks := []info.Notebook{}

	db.Engine.Where("UserId=?", userId).And("Usn>=?", afterUsn).Asc("Usn").Limit(0, maxEntry).Find(&notebooks)
	return notebooks
}

// 得到用户下所有的notebook
// 排序好之后返回
// [ok]
func (this *NotebookService) GetNotebooks(userId int64) info.SubNotebooks {
	userNotebooks := []info.Notebook{}

	db.Engine.Where("UserId=?", userId).And("IsDeleted <>?", true).Find(&userNotebooks)
	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, true, true)
}

// share调用, 不需要删除没有父的notebook
// 不需要排序, 因为会重新排序
// 通过notebookIds得到notebooks, 并转成层次有序
func (this *NotebookService) GetNotebooksByNotebookIds(notebookIds []int64) info.SubNotebooks {
	userNotebooks := []info.Notebook{}
	db.Engine.In("NotebookId", &notebookIds).Find(userNotebooks)
	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, false, false)
}

// 添加
func (this *NotebookService) AddNotebook(notebook info.Notebook) (bool, info.Notebook) {

	notebook.UrlTitle = GetUrTitle(notebook.UserId, notebook.Title, "notebook", notebook.NotebookId)
	notebook.Usn = userService.IncrUsn(notebook.UserId)
	now := time.Now()
	notebook.CreatedTime = now
	notebook.UpdatedTime = now
	_, err := db.Engine.Insert(&notebook)
	if err != nil {
		return false, notebook
	}
	return true, notebook
}

// 更新笔记, api
func (this *NotebookService) UpdateNotebookApi(userId, notebookId int64, title, parentNotebookId string, seq int, usn int64) (bool, string, info.Notebook) {
	if notebookId == 0 {
		return false, "notebookIdNotExists", info.Notebook{}
	}

	// 先判断usn是否和数据库的一样, 如果不一样, 则冲突, 不保存
	notebook := this.GetNotebookById(notebookId)
	// 不存在
	if notebook.NotebookId == 0 {
		return false, "notExists", notebook
	} else if notebook.Usn != usn {
		return false, "conflict", notebook
	}
	notebook.Usn = userService.IncrUsn(userId)
	notebook.Title = title

	// updates := bson.M{"Title": title, "Usn": notebook.Usn, "Seq": seq, "UpdatedTime": time.Now()}
	// if parentNotebookId != "" && bson.IsObjectIdHex(parentNotebookId) {
	// 	updates["ParentNotebookId"] = bson.ObjectIdHex(parentNotebookId)
	// } else {
	// 	updates["ParentNotebookId"] = ""
	// }
	// ok := db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, updates)
	// if ok {
	// 	return ok, "", this.GetNotebookById(notebookId)
	// }
	return false, "", notebook
}

// 判断是否是blog
func (this *NotebookService) IsBlog(notebookId int64) bool {
	notebook := info.Notebook{}
	db.Engine.Id(notebookId).Get(&notebook)
	return notebook.IsBlog
}

// 判断是否是我的notebook
func (this *NotebookService) IsMyNotebook(notebookId, userId int64) bool {
	notebook := info.Notebook{}
	db.Engine.Id(notebookId).Get(&notebook)
	return notebook.UserId == userId
}

// 更新笔记本信息
// 太广, 不用
/*
func (this *NotebookService) UpdateNotebook(notebook info.Notebook) bool {
	return db.UpdateByIdAndUserId2(db.Notebooks, notebook.NotebookId, notebook.UserId, notebook)
}
*/

// 更新笔记本标题
// [ok]
func (this *NotebookService) UpdateNotebookTitle(notebookId, userId int64, title string) bool {
	usn := userService.IncrUsn(userId)
	notebook := info.Notebook{}
	notebook.Usn = usn
	notebook.Title = title
	_, err := db.Engine.Id(notebookId).Cols("Title", "Usn").Update(&notebook)
	return err == nil
}

// 更新notebook
// func (this *NotebookService) UpdateNotebook(userId, notebookId string, needUpdate bson.M) bool {

// 	notebook := info.Notebook{}
// 	notebook.Usn = userService.IncrUsn(userId)
// 	notebook.UpdatedTime = time.Now()
// 	affected, err := db.Engine.Id(notebookId).Cols("UpdatedTime", "Usn").Update(&notebook)
// 	return err == nil
// }

// ToBlog or Not
func (this *NotebookService) ToBlog(userId, notebookId int64, isBlog bool) bool {
	// updates := bson.M{"IsBlog": isBlog, "Usn": userService.IncrUsn(userId)}
	// // 笔记本
	// db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, updates)

	// // 更新笔记
	// q := bson.M{"UserId": bson.ObjectIdHex(userId),
	// 	"NotebookId": bson.ObjectIdHex(notebookId)}
	// data := bson.M{"IsBlog": isBlog}
	// if isBlog {
	// 	data["PublicTime"] = time.Now()
	// } else {
	// 	data["HasSelfDefined"] = false
	// }
	// // usn
	// data["Usn"] = userService.IncrUsn(userId)
	// db.UpdateByQMap(db.Notes, q, data)

	// // noteContents也更新, 这个就麻烦了, noteContents表没有NotebookId
	// // 先查该notebook下所有notes, 得到id
	// notes := []info.Note{}
	// db.ListByQWithFields(db.Notes, q, []string{"_id"}, &notes)
	// if len(notes) > 0 {
	// 	noteIds := make([]bson.ObjectId, len(notes))
	// 	for i, each := range notes {
	// 		noteIds[i] = each.NoteId
	// 	}
	// 	db.UpdateByQMap(db.NoteContents, bson.M{"_id": bson.M{"$in": noteIds}}, bson.M{"IsBlog": isBlog})
	// }

	// // 重新计算tags
	// go (func() {
	// 	blogService.ReCountBlogTags(userId)
	// })()

	return true
}

// 查看是否有子notebook
// 先查看该notebookId下是否有notes, 没有则删除
func (this *NotebookService) DeleteNotebook(userId, notebookId int64) (bool, string) {
	// if db.Count(db.Notebooks, bson.M{
	// 	"ParentNotebookId": bson.ObjectIdHex(notebookId),
	// 	"UserId":           bson.ObjectIdHex(userId),
	// 	"IsDeleted":        false,
	// }) == 0 { // 无
	// 	if db.Count(db.Notes, bson.M{"NotebookId": bson.ObjectIdHex(notebookId),
	// 		"UserId":    bson.ObjectIdHex(userId),
	// 		"IsTrash":   false,
	// 		"IsDeleted": false}) == 0 { // 不包含trash
	// 		// 不是真删除 1/20, 为了同步笔记本
	// 		ok := db.UpdateByQMap(db.Notebooks, bson.M{"_id": bson.ObjectIdHex(notebookId)}, bson.M{"IsDeleted": true, "Usn": userService.IncrUsn(userId)})
	// 		return ok, ""
	// 		//			return db.DeleteByIdAndUserId(db.Notebooks, notebookId, userId), ""
	// 	}
	// 	return false, "笔记本下有笔记"
	// } else {
	return false, "笔记本下有子笔记本"
	// }

}

// API调用, 删除笔记本, 不作笔记控制
func (this *NotebookService) DeleteNotebookForce(userId, notebookId int64, usn int64) (bool, string) {
	notebook := this.GetNotebookById(notebookId)
	// 不存在
	if notebook.NotebookId == 0 {
		return false, "notExists"
	} else if notebook.Usn != usn {
		return false, "conflict"
	}
	_, err := db.Engine.Id(notebookId).Delete(&notebook)
	return err == nil, ""
}

// 排序
// 传入 notebookId => Seq
// 为什么要传入userId, 防止修改其它用户的信息 (恶意)
// [ok]
func (this *NotebookService) SortNotebooks(userId int64, notebookId2Seqs map[string]int) bool {
	if len(notebookId2Seqs) == 0 {
		return false
	}

	for notebookId, seq := range notebookId2Seqs {
		notebook := info.Notebook{}
		notebook.Seq = seq
		notebook.Usn = userService.IncrUsn(userId)

		_, err := db.Engine.Id(notebookId).Update(&notebook)
		if err != nil {
			return false
		}
	}

	return true
}

// 排序和设置父
func (this *NotebookService) DragNotebooks(userId int64, curNotebookId int64, parentNotebookId int64, siblings []string) bool {
	// ok := false
	// // 如果没parentNotebookId, 则parentNotebookId设空
	// if parentNotebookId == "" {
	// 	ok = db.UpdateByIdAndUserIdMap(db.Notebooks, curNotebookId, userId, bson.M{"ParentNotebookId": "", "Usn": userService.IncrUsn(userId)})
	// } else {
	// 	ok = db.UpdateByIdAndUserIdMap(db.Notebooks, curNotebookId, userId, bson.M{"ParentNotebookId": bson.ObjectIdHex(parentNotebookId), "Usn": userService.IncrUsn(userId)})
	// }

	// if !ok {
	// 	return false
	// }

	// // 排序
	// for seq, notebookId := range siblings {
	// 	if !db.UpdateByIdAndUserIdMap(db.Notebooks, notebookId, userId, bson.M{"Seq": seq, "Usn": userService.IncrUsn(userId)}) {
	// 		return false
	// 	}
	// }

	return true
}

// 重新统计笔记本下的笔记数目
// noteSevice: AddNote, CopyNote, CopySharedNote, MoveNote
// trashService: DeleteNote (recove不用, 都统一在MoveNote里了)
func (this *NotebookService) ReCountNotebookNumberNotes(notebookId int64) bool {
	// notebookIdO := notebookId
	// notebook := info.Notebook{}
	// count, _ := db.Engine.Where("NotebookId=?", notebookId).And("IsTrash=?", false).And("IsDeleted=?", false).Count(&notebook)
	// Log(count)
	// Log(notebookId)
	// return db.UpdateByQField(db.Notebooks, bson.M{"_id": notebookIdO}, "NumberNotes", count)
	return true
}

func (this *NotebookService) ReCountAll() {
	/*
		// 得到所有笔记本
		notebooks := []info.Notebook{}
		db.ListByQWithFields(db.Notebooks, bson.M{}, []string{"NotebookId"}, &notebooks)

		for _, each := range notebooks {
			this.ReCountNotebookNumberNotes(each.NotebookId )
		}
	*/
}
