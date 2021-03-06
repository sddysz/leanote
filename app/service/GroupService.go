package service

import (
	"github.com/sddysz/leanote/app/db"
	"github.com/sddysz/leanote/app/info"
	//	. "github.com/sddysz/leanote/app/lea"
	//	"strings"
)

// 用户组, 用户组用户管理

type GroupService struct {
}

// 添加分组
func (this *GroupService) AddGroup(userId int64, title string) (bool, info.Group) {
	group := info.Group{
		UserId: userId,
		Title:  title,
	}
	_, err := db.Engine.Insert(&group)
	return err == nil, group
}

// 删除分组
// 判断是否有好友
func (this *GroupService) DeleteGroup(userId, groupId string) (ok bool, msg string) {
	/*
		if db.Has(db.GroupUsers, bson.M{"GroupId": bson.ObjectIdHex(groupId)}) {
			return false, "groupHasUsers"
		}
	*/
	if !this.isMyGroup(userId, groupId) {
		return false, "notMyGroup"
	}

	// 删除分组后, 需要删除所有用户分享到该组的笔记本, 笔记

	// shareService.DeleteAllShareNotebookGroup(groupId)
	// shareService.DeleteAllShareNoteGroup(groupId)

	group := info.Group{}
	_, err := db.Engine.Id(groupId).Delete(group)
	return err == nil, ""

	// TODO 删除分组后, 在shareNote, shareNotebook中也要删除
}

// 修改group标题
func (this *GroupService) UpdateGroupTitle(userId, groupId int64, title string) (ok bool) {
	//return db.UpdateByIdAndUserIdField(db.Groups, groupId, userId, "Title", title)
	return true
}

// 得到用户的所有分组(包括下的所有用户)
func (this *GroupService) GetGroupsAndUsers(userId int64) []info.Group {
	/*
		// 得到我的分组
		groups := []info.Group{}
		db.ListByQ(db.Groups, bson.M{"UserId": bson.ObjectIdHex(userId)}, &groups)
	*/
	// 我的分组, 及我所属的分组
	groups := this.GetGroupsContainOf(userId)

	// 得到其下的用户
	for i, group := range groups {
		group.Users = this.GetUsers(group.GroupId)
		groups[i] = group
	}
	return groups
}

// 仅仅得到所有分组
func (this *GroupService) GetGroups(userId int64) []info.Group {
	// 得到分组s
	groups := []info.Group{}
	db.Engine.Where("UserId=?", userId).Find(&groups)
	return groups
}

// 得到我的和我所属组的ids
func (this *GroupService) GetMineAndBelongToGroupIds(userId int64) []int64 {
	// 所属组
	groupIds := this.GetBelongToGroupIds(userId)

	m := map[int64]bool{}
	for _, groupId := range groupIds {
		m[groupId] = true
	}

	// 我的组
	myGroups := this.GetGroups(userId)

	for _, group := range myGroups {
		if !m[group.GroupId] {
			groupIds = append(groupIds, group.GroupId)
		}
	}

	return groupIds
}

// 获取包含此用户的组对象数组
// 获取该用户所属组, 和我的组
func (this *GroupService) GetGroupsContainOf(userId int64) []info.Group {
	// 我的组
	myGroups := this.GetGroups(userId)
	myGroupMap := map[int64]bool{}

	for _, group := range myGroups {
		myGroupMap[group.GroupId] = true
	}

	// 所属组
	groupIds := this.GetBelongToGroupIds(userId)

	groups := []info.Group{}

	db.Engine.In("Id", groupIds).Find(&groups)
	for _, group := range groups {
		if !myGroupMap[group.GroupId] {
			myGroups = append(myGroups, group)
		}
	}

	return myGroups
}

// 得到分组, shareService用
func (this *GroupService) GetGroup(userId, groupId string) info.Group {
	// 得到分组s
	group := info.Group{}
	//db.GetByIdAndUserId(db.Groups, groupId, userId, &group)
	return group
}

// 得到某分组下的用户
func (this *GroupService) GetUsers(groupId int64) []info.User {
	// 得到UserIds
	// groupUsers := []info.GroupUser{}
	// db.ListByQWithFields(db.GroupUsers, bson.M{"GroupId": bson.ObjectIdHex(groupId)}, []string{"UserId"}, &groupUsers)
	// if len(groupUsers) == 0 {
	// 	return nil
	// }
	// userIds := make([]bson.ObjectId, len(groupUsers))
	// for i, each := range groupUsers {
	// 	userIds[i] = each.UserId
	// }
	// 得到userInfos
	// return userService.ListUserInfosByUserIds(userIds)
	return nil
}

// 得到我所属的所有分组ids
func (this *GroupService) GetBelongToGroupIds(userId int64) []int64 {
	// 得到UserIds
	groupUsers := []info.GroupUser{}
	db.Engine.Where("UserId=?", userId).Find(&groupUsers)
	if len(groupUsers) == 0 {
		return nil
	}
	groupIds := make([]int64, len(groupUsers))
	for i, each := range groupUsers {
		groupIds[i] = each.GroupId
	}
	return groupIds
}

func (this *GroupService) isMyGroup(ownUserId, groupId string) (ok bool) {
	// return db.Has(db.Groups, bson.M{"_id": bson.ObjectIdHex(groupId), "UserId": bson.ObjectIdHex(ownUserId)})
	return true
}

// 判断组中是否包含指定用户
func (this *GroupService) IsExistsGroupUser(userId, groupId string) (ok bool) {
	// 如果我拥有这个组, 那也行
	if this.isMyGroup(userId, groupId) {
		return true
	}
	// return db.Has(db.GroupUsers, bson.M{"UserId": bson.ObjectIdHex(userId), "GroupId": bson.ObjectIdHex(groupId)})
	return false
}

// 为group添加用户
// 用户是否已存在?
func (this *GroupService) AddUser(ownUserId, groupId, userId string) (ok bool, msg string) {
	// groupId是否是ownUserId的?
	/*
		if !this.IsExistsGroupUser(ownUserId, groupId) {
			return false, "forbiddenNotMyGroup"
		}
	*/
	if !this.isMyGroup(ownUserId, groupId) {
		return false, "forbiddenNotMyGroup"
	}

	// 是否已存在
	// if db.Has(db.GroupUsers, bson.M{"GroupId": bson.ObjectIdHex(groupId), "UserId": bson.ObjectIdHex(userId)}) {
	// 	return false, "userExistsInGroup"
	// }

	// return db.Insert(db.GroupUsers, info.GroupUser{
	// 	GroupUserId: bson.NewObjectId(),
	// 	GroupId:     bson.ObjectIdHex(groupId),
	// 	UserId:      bson.ObjectIdHex(userId),
	// 	CreatedTime: time.Now(),
	// }), ""
	return true, ""
}

// 删除用户
func (this *GroupService) DeleteUser(ownUserId, groupId, userId string) (ok bool, msg string) {
	// groupId是否是ownUserId的?
	/*
		if !this.IsExistsGroupUser(ownUserId, groupId) {
			return false, "forbiddenNotMyGroup"
		}
	*/
	if !this.isMyGroup(ownUserId, groupId) {
		return false, "forbiddenNotMyGroup"
	}

	// // 删除该用户分享到本组的笔记本, 笔记
	// shareService.DeleteShareNotebookGroupWhenDeleteGroupUser(userId, groupId)
	// shareService.DeleteShareNoteGroupWhenDeleteGroupUser(userId, groupId)

	groupUser := info.GroupUser{}
	db.Engine.Where("GroupId=?", groupId).And("UserId=?", userId).Delete(groupUser)
	return true, ""
}
