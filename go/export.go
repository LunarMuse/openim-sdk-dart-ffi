package main

/*
#include <stdio.h>
#include <stdbool.h>
#include "openim_sdk_ffi.h"
*/
import "C"

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
	"sync"
)

type SDKInstance struct {
	listener C.Openim_Listener
	port     C.Dart_Port_DL
}

var (
	instance     *SDKInstance
	instanceOnce sync.Once
	instanceLock sync.RWMutex
)

// 初始化SDK实例 (线程安全)
func initSDKInstance(listener C.Openim_Listener, port C.Dart_Port_DL) {
	instanceOnce.Do(func() {
		instance = &SDKInstance{
			listener: listener,
			port:     port,
		}
	})
}

// 安全获取实例 (带检查)
func getSDKInstance() *SDKInstance {
	instanceLock.RLock()
	defer instanceLock.RUnlock()

	if instance == nil {
		return nil
	}

	return instance
}

// =====================================================Conversation===============================================
//
//export GetAllConversationList
func GetAllConversationList(operationID *C.char) {
	baseCallback := NewBaseCallback("GetAllConversationList", operationID)
	open_im_sdk.GetAllConversationList(baseCallback, C.GoString(operationID))
}

//export GetConversationListSplit
func GetConversationListSplit(operationID *C.char, offset C.int, count C.int) {
	baseCallback := NewBaseCallback("GetConversationListSplit", operationID)
	open_im_sdk.GetConversationListSplit(baseCallback, C.GoString(operationID), int(offset), int(count))
}

//export GetOneConversation
func GetOneConversation(operationID *C.char, sessionType C.int32_t, sourceID *C.char) {
	baseCallback := NewBaseCallback("GetOneConversation", operationID)
	open_im_sdk.GetOneConversation(baseCallback, C.GoString(operationID), int32(sessionType), C.GoString(sourceID))
}

//export GetMultipleConversation
func GetMultipleConversation(operationID *C.char, conversationIDList *C.char) {
	baseCallback := NewBaseCallback("GetMultipleConversation", operationID)
	open_im_sdk.GetMultipleConversation(baseCallback, C.GoString(operationID), C.GoString(conversationIDList))
}

//export GetConversationIDBySessionType
func GetConversationIDBySessionType(operationID *C.char, sourceID *C.char, sessionType C.int) *C.char {
	return C.CString(open_im_sdk.GetConversationIDBySessionType(C.GoString(operationID), C.GoString(sourceID), int(sessionType)))
}

//export GetTotalUnreadMsgCount
func GetTotalUnreadMsgCount(operationID *C.char) {
	baseCallback := NewBaseCallback("GetTotalUnreadMsgCount", operationID)
	open_im_sdk.GetTotalUnreadMsgCount(baseCallback, C.GoString(operationID))
}

//export MarkConversationMessageAsRead
func MarkConversationMessageAsRead(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback("MarkConversationMessageAsRead", operationID)
	open_im_sdk.MarkConversationMessageAsRead(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export MarkAllConversationMessageAsRead
func MarkAllConversationMessageAsRead(operationID *C.char) {
	baseCallback := NewBaseCallback("MarkConversationMessageAsRead", operationID)
	open_im_sdk.MarkAllConversationMessageAsRead(baseCallback, C.GoString(operationID))
}

//export SetConversation
func SetConversation(operationID *C.char, conversationID *C.char, draftText *C.char) {
	baseCallback := NewBaseCallback("SetConversation", operationID)
	open_im_sdk.SetConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}

//export SetConversationDraft
func SetConversationDraft(operationID *C.char, conversationID *C.char, draftText *C.char) {
	baseCallback := NewBaseCallback("SetConversationDraft", operationID)
	open_im_sdk.SetConversationDraft(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}

//export HideConversation
func HideConversation(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback("HideConversation", operationID)
	open_im_sdk.HideConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export ChangeInputStates
func ChangeInputStates(operationID *C.char, conversationID *C.char, focus C._Bool) {
	baseCallback := NewBaseCallback("ChangeInputStates", operationID)
	open_im_sdk.ChangeInputStates(baseCallback, C.GoString(operationID), C.GoString(conversationID), bool(focus))
}

//export HideAllConversations
func HideAllConversations(operationID *C.char) {
	baseCallback := NewBaseCallback("HideAllConversations", operationID)
	open_im_sdk.HideAllConversations(baseCallback, C.GoString(operationID))
}

//export ClearConversationAndDeleteAllMsg
func ClearConversationAndDeleteAllMsg(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback("ClearConversationAndDeleteAllMsg", operationID)
	open_im_sdk.ClearConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export GetInputStates
func GetInputStates(operationID *C.char, conversationID *C.char, userID *C.char) {
	baseCallback := NewBaseCallback("GetInputStates", operationID)
	open_im_sdk.GetInputStates(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(userID))
}

//export DeleteConversationAndDeleteAllMsg
func DeleteConversationAndDeleteAllMsg(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback("DeleteConversationAndDeleteAllMsg", operationID)
	open_im_sdk.DeleteConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

// =====================================================conversation_msg===============================================
//
//export CreateTextMessage
func CreateTextMessage(operationID, text *C.char) *C.char {
	message := C.CString(open_im_sdk.CreateTextMessage(C.GoString(operationID), C.GoString(text)))
	return message
}

//export CreateTextAtMessage
func CreateTextAtMessage(operationID, text, atUserList, atUsersInfo, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList),
		C.GoString(atUsersInfo), C.GoString(message)))
}

//export CreateImageMessageFromFullPath
func CreateImageMessageFromFullPath(operationID, imageFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath)))
}

//export CreateImageMessageByURL
func CreateImageMessageByURL(operationID, sourcePath, sourcePicture, bigPicture, snapshotPicture *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePath),
		C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture)))
}

//export CreateForwardMessage
func CreateForwardMessage(operationID, m *C.char) *C.char {
	return C.CString(open_im_sdk.CreateForwardMessage(C.GoString(operationID), C.GoString(m)))
}

//export CreateLocationMessage
func CreateLocationMessage(operationID, description *C.char, longitude, latitude C.double) *C.char {
	return C.CString(open_im_sdk.CreateLocationMessage(C.GoString(operationID), C.GoString(description),
		float64(longitude), float64(latitude)))
}

//export CreateQuoteMessage
func CreateQuoteMessage(operationID, text, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message)))
}

//export CreateCardMessage
func CreateCardMessage(operationID, cardInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo)))
}

//export CreateCustomMessage
func CreateCustomMessage(operationID, data, extension, description *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension),
		C.GoString(description)))
}

//export SendMessage
func SendMessage(operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	sendMsgCallback := NewSendMessageCallback("SendMessage", operationID)
	// isOnlineOnly C.int -> parseBool(int(isOnlineOnly))
	open_im_sdk.SendMessage(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo), true)
}

//export SendMessageNotOss
func SendMessageNotOss(operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	sendMsgCallback := NewSendMessageCallback("SendMessageNotOss", operationID)
	// isOnlineOnly C.int -> parseBool(int(isOnlineOnly))
	open_im_sdk.SendMessageNotOss(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo), true)
}

//export TypingStatusUpdate
func TypingStatusUpdate(operationID *C.char, recvID *C.char, msgTip *C.char) {
	baseCallback := NewBaseCallback("TypingStatusUpdate", operationID)
	open_im_sdk.TypingStatusUpdate(baseCallback, C.GoString(operationID), C.GoString(recvID), C.GoString(msgTip))
}

//export RevokeMessage
func RevokeMessage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback("RevokeMessage", operationID)
	open_im_sdk.RevokeMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export DeleteMessage
func DeleteMessage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback("DeleteMessage", operationID)
	open_im_sdk.DeleteMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export DeleteMessageFromLocalStorage
func DeleteMessageFromLocalStorage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback("DeleteMessageFromLocalStorage", operationID)
	open_im_sdk.DeleteMessageFromLocalStorage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export DeleteAllMsgFromLocal
func DeleteAllMsgFromLocal(operationID *C.char) {
	baseCallback := NewBaseCallback("DeleteAllMsgFromLocal", operationID)
	open_im_sdk.DeleteAllMsgFromLocal(baseCallback, C.GoString(operationID))
}

//export DeleteAllMsgFromLocalAndSvr
func DeleteAllMsgFromLocalAndSvr(operationID *C.char) {
	baseCallback := NewBaseCallback("DeleteAllMsgFromLocalAndSvr", operationID)
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(baseCallback, C.GoString(operationID))
}

//export SearchLocalMessages
func SearchLocalMessages(operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback("SearchLocalMessages", operationID)
	open_im_sdk.SearchLocalMessages(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export GetAdvancedHistoryMessageList
func GetAdvancedHistoryMessageList(operationID, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback("GetAdvancedHistoryMessageList", operationID)
	open_im_sdk.GetAdvancedHistoryMessageList(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export GetAdvancedHistoryMessageListReverse
func GetAdvancedHistoryMessageListReverse(operationID *C.char, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback("GetAdvancedHistoryMessageListReverse", operationID)
	open_im_sdk.GetAdvancedHistoryMessageListReverse(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export FindMessageList
func FindMessageList(operationID *C.char, findMessageOptions *C.char) {
	baseCallback := NewBaseCallback("FindMessageList", operationID)
	open_im_sdk.FindMessageList(baseCallback, C.GoString(operationID), C.GoString(findMessageOptions))
}

//export InsertGroupMessageToLocalStorage
func InsertGroupMessageToLocalStorage(operationID *C.char, message *C.char, groupID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback("InsertGroupMessageToLocalStorage", operationID)
	open_im_sdk.InsertGroupMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(groupID), C.GoString(sendID))
}

//export InsertSingleMessageToLocalStorage
func InsertSingleMessageToLocalStorage(operationID *C.char, message *C.char, recvID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback("InsertSingleMessageToLocalStorage", operationID)
	open_im_sdk.InsertSingleMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}

//export SearchConversation
func SearchConversation(operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback("SearchConversation", operationID)
	open_im_sdk.SearchConversation(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export SetMessageLocalEx
func SetMessageLocalEx(operationID *C.char, conversationID *C.char, clientMsgID *C.char, localEx *C.char) {
	baseCallback := NewBaseCallback("SetMessageLocalEx", operationID)
	open_im_sdk.SetMessageLocalEx(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID), C.GoString(localEx))
}

//export GetAtAllTag
func GetAtAllTag(operationID *C.char) *C.char {
	return C.CString(open_im_sdk.GetAtAllTag(C.GoString(operationID)))
}

//export CreateAdvancedTextMessage
func CreateAdvancedTextMessage(operationID, text, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedTextMessage(C.GoString(operationID), C.GoString(text),
		C.GoString(messageEntityList)))
}

//export CreateAdvancedQuoteMessage
func CreateAdvancedQuoteMessage(operationID, text, message, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedQuoteMessage(C.GoString(operationID), C.GoString(text),
		C.GoString(message), C.GoString(messageEntityList)))
}

//export CreateImageMessage
func CreateImageMessage(operationID, imagePath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath)))
}

//export CreateSoundMessage
func CreateSoundMessage(operationID, soundPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

//export CreateSoundMessageByURL
func CreateSoundMessageByURL(operationID, soundBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo)))
}

//export CreateVideoMessage
func CreateVideoMessage(operationID, videoPath *C.char, videoType *C.char, duration C.int64_t,
	snapshotPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath),
		C.GoString(videoType), int64(duration), C.GoString(snapshotPath)))
}

//export CreateVideoMessageByURL
func CreateVideoMessageByURL(operationID, videoBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo)))
}

//export CreateFileMessage
func CreateFileMessage(operationID, filePath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName)))
}

//export CreateMergerMessage
func CreateMergerMessage(operationID, messageList, title, summaryList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateMergerMessage(C.GoString(operationID), C.GoString(messageList),
		C.GoString(title), C.GoString(summaryList)))
}

//export CreateFaceMessage
func CreateFaceMessage(operationID *C.char, index C.int, data *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFaceMessage(C.GoString(operationID), int(index), C.GoString(data)))
}

//export MarkMessagesAsReadByMsgID
func MarkMessagesAsReadByMsgID(operationID *C.char, conversationID, clientMsgIDs *C.char) {
	baseCallback := NewBaseCallback("MarkMessagesAsReadByMsgID", operationID)
	open_im_sdk.MarkMessagesAsReadByMsgID(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgIDs))
}

//export CreateFileMessageByURL
func CreateFileMessageByURL(operationID, fileBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo)))
}

//export CreateFileMessageFromFullPath
func CreateFileMessageFromFullPath(operationID, fileFullPath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath),
		C.GoString(fileName)))
}

//export CreateSoundMessageFromFullPath
func CreateSoundMessageFromFullPath(operationID, soundFullPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundFullPath),
		int64(duration)))
}

//export CreateVideoMessageFromFullPath
func CreateVideoMessageFromFullPath(operationID, videoFullPath, videoType *C.char, duration C.int64_t,
	snapshotFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath),
		C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath)))
}

// =====================================================file===============================================

//export UploadFile
func UploadFile(operationID *C.char, req *C.char, uuid *C.char) {
	// TODO:uuid not use
	baseCallback := NewBaseCallback("UploadFile", operationID)
	uploadFileCallback := NewUploadFileCallback("UploadFile", operationID)
	open_im_sdk.UploadFile(baseCallback, C.GoString(operationID), C.GoString(req), uploadFileCallback)
}

// =====================================================relation===============================================

//export AcceptFriendApplication
func AcceptFriendApplication(operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback("AcceptFriendApplication", operationID)
	open_im_sdk.AcceptFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export AddBlack
func AddBlack(operationID *C.char, blackUserID *C.char, ex *C.char) {
	baseCallback := NewBaseCallback("AddBlack", operationID)
	open_im_sdk.AddBlack(baseCallback, C.GoString(operationID), C.GoString(blackUserID), C.GoString(ex))
}

//export AddFriend
func AddFriend(operationID *C.char, userIDReqMsg *C.char) {
	baseCallback := NewBaseCallback("AddFriend", operationID)
	open_im_sdk.AddFriend(baseCallback, C.GoString(operationID), C.GoString(userIDReqMsg))
}

//export CheckFriend
func CheckFriend(operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback("CheckFriend", operationID)
	open_im_sdk.CheckFriend(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export DeleteFriend
func DeleteFriend(operationID *C.char, friendUserID *C.char) {
	baseCallback := NewBaseCallback("DeleteFriend", operationID)
	open_im_sdk.DeleteFriend(baseCallback, C.GoString(operationID), C.GoString(friendUserID))
}

//export GetBlackList
func GetBlackList(operationID *C.char) {
	baseCallback := NewBaseCallback("GetBlackList", operationID)
	open_im_sdk.GetBlackList(baseCallback, C.GoString(operationID))
}

//export GetFriendApplicationListAsApplicant
func GetFriendApplicationListAsApplicant(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("GetFriendApplicationListAsApplicant", operationID)
	open_im_sdk.GetFriendApplicationListAsApplicant(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export GetFriendApplicationListAsRecipient
func GetFriendApplicationListAsRecipient(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("GetFriendApplicationListAsRecipient", operationID)
	open_im_sdk.GetFriendApplicationListAsRecipient(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export GetFriendApplicationUnhandledCount
func GetFriendApplicationUnhandledCount(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("GetFriendApplicationUnhandledCount", operationID)
	open_im_sdk.GetFriendApplicationUnhandledCount(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export GetFriendList
func GetFriendList(operationID *C.char, filterBlack C._Bool) {
	baseCallback := NewBaseCallback("GetFriendList", operationID)
	open_im_sdk.GetFriendList(baseCallback, C.GoString(operationID), bool(filterBlack))
}

//export GetFriendListPage
func GetFriendListPage(operationID *C.char, offset C.int32_t, count C.int32_t, filterBlack C._Bool) {
	baseCallback := NewBaseCallback("GetFriendListPage", operationID)
	open_im_sdk.GetFriendListPage(baseCallback, C.GoString(operationID), int32(offset), int32(count), bool(filterBlack))
}

//export GetSpecifiedFriendsInfo
func GetSpecifiedFriendsInfo(operationID *C.char, userIDList *C.char, filterBlack C._Bool) {
	baseCallback := NewBaseCallback("GetSpecifiedFriendsInfo", operationID)
	open_im_sdk.GetSpecifiedFriendsInfo(baseCallback, C.GoString(operationID), C.GoString(userIDList), bool(filterBlack))
}

//export RefuseFriendApplication
func RefuseFriendApplication(operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback("RefuseFriendApplication", operationID)
	open_im_sdk.RefuseFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export RemoveBlack
func RemoveBlack(operationID *C.char, removeUserID *C.char) {
	baseCallback := NewBaseCallback("RemoveBlack", operationID)
	open_im_sdk.RemoveBlack(baseCallback, C.GoString(operationID), C.GoString(removeUserID))
}

//export SearchFriends
func SearchFriends(operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback("SearchFriends", operationID)
	open_im_sdk.SearchFriends(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export UpdateFriends
func UpdateFriends(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("UpdateFriends", operationID)
	open_im_sdk.UpdateFriends(baseCallback, C.GoString(operationID), C.GoString(req))
}

// =====================================================group===============================================

//export CreateGroup
func CreateGroup(operationID, groupReqInfo *C.char) {
	baseCallback := NewBaseCallback("CreateGroup", operationID)
	open_im_sdk.CreateGroup(baseCallback, C.GoString(operationID), C.GoString(groupReqInfo))
}

//export JoinGroup
func JoinGroup(operationID, groupID, reqMsg *C.char, joinSource C.int32_t, ex *C.char) {
	baseCallback := NewBaseCallback("JoinGroup", operationID)
	open_im_sdk.JoinGroup(baseCallback, C.GoString(operationID), C.GoString(groupID), C.GoString(reqMsg),
		int32(joinSource), C.GoString(ex))
}

//export InviteUserToGroup
func InviteUserToGroup(operationID, groupID, reason, userIDList *C.char) {
	baseCallback := NewBaseCallback("InviteUserToGroup", operationID)
	open_im_sdk.InviteUserToGroup(baseCallback, C.GoString(operationID), C.GoString(groupID), C.GoString(reason),
		C.GoString(userIDList))
}

//export GetJoinedGroupList
func GetJoinedGroupList(operationID *C.char) {
	baseCallback := NewBaseCallback("GetJoinedGroupList", operationID)
	open_im_sdk.GetJoinedGroupList(baseCallback, C.GoString(operationID))
}

//export getJoinedGroupListPage
func getJoinedGroupListPage(operationID *C.char, offset, count C.int32_t) {
	baseCallback := NewBaseCallback("GetJoinedGroupListPage", operationID)
	open_im_sdk.GetJoinedGroupListPage(baseCallback, C.GoString(operationID), int32(offset), int32(count))
}

//export SearchGroups
func SearchGroups(operationID, searchParam *C.char) {
	baseCallback := NewBaseCallback("SearchGroups", operationID)
	open_im_sdk.SearchGroups(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export GetSpecifiedGroupsInfo
func GetSpecifiedGroupsInfo(operationID, groupIDList *C.char) {
	baseCallback := NewBaseCallback("GetSpecifiedGroupsInfo", operationID)
	open_im_sdk.GetSpecifiedGroupsInfo(baseCallback, C.GoString(operationID), C.GoString(groupIDList))
}

//export SetGroupInfo
func SetGroupInfo(operationID, groupInfo *C.char) {
	baseCallback := NewBaseCallback("SetGroupInfo", operationID)
	open_im_sdk.SetGroupInfo(baseCallback, C.GoString(operationID), C.GoString(groupInfo))
}

//export GetGroupApplicationListAsRecipient
func GetGroupApplicationListAsRecipient(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("GetGroupApplicationListAsRecipient", operationID)
	open_im_sdk.GetGroupApplicationListAsRecipient(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export GetGroupApplicationListAsApplicant
func GetGroupApplicationListAsApplicant(operationID *C.char, req *C.char) {
	baseCallback := NewBaseCallback("GetGroupApplicationListAsApplicant", operationID)
	open_im_sdk.GetGroupApplicationListAsApplicant(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export GetGroupApplicationUnhandledCount
func GetGroupApplicationUnhandledCount(operationID, req *C.char) {
	baseCallback := NewBaseCallback("GetGroupApplicationUnhandledCount", operationID)
	open_im_sdk.GetGroupApplicationUnhandledCount(baseCallback, C.GoString(operationID), C.GoString(req))
}

//export AcceptGroupApplication
func AcceptGroupApplication(operationID, groupID, fromUserID, handleMsg *C.char) {
	baseCallback := NewBaseCallback("AcceptGroupApplication", operationID)
	open_im_sdk.AcceptGroupApplication(baseCallback, C.GoString(operationID), C.GoString(groupID),
		C.GoString(fromUserID), C.GoString(handleMsg))
}

//export RefuseGroupApplication
func RefuseGroupApplication(operationID, groupID, fromUserID, handleMsg *C.char) {
	baseCallback := NewBaseCallback("RefuseGroupApplication", operationID)
	open_im_sdk.RefuseGroupApplication(baseCallback, C.GoString(operationID), C.GoString(groupID),
		C.GoString(fromUserID), C.GoString(handleMsg))
}

//export GetGroupMemberList
func GetGroupMemberList(operationID, groupID *C.char, filter, offset, count C.int32_t) {
	baseCallback := NewBaseCallback("GetGroupMemberList", operationID)
	open_im_sdk.GetGroupMemberList(baseCallback, C.GoString(operationID), C.GoString(groupID), int32(filter),
		int32(offset), int32(count))
}

//export GetSpecifiedGroupMembersInfo
func GetSpecifiedGroupMembersInfo(operationID, groupID, userIDList *C.char) {
	baseCallback := NewBaseCallback("GetSpecifiedGroupMembersInfo", operationID)
	open_im_sdk.GetSpecifiedGroupMembersInfo(baseCallback, C.GoString(operationID), C.GoString(groupID),
		C.GoString(userIDList))
}

//export SearchGroupMembers
func SearchGroupMembers(operationID, searchParam *C.char) {
	baseCallback := NewBaseCallback("SearchGroupMembers", operationID)
	open_im_sdk.SearchGroupMembers(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export SetGroupMemberInfo
func SetGroupMemberInfo(operationID *C.char, groupMemberInfo *C.char) {
	baseCallback := NewBaseCallback("SetGroupMemberInfo", operationID)
	open_im_sdk.SetGroupMemberInfo(baseCallback, C.GoString(operationID), C.GoString(groupMemberInfo))
}

//export GetGroupMemberOwnerAndAdmin
func GetGroupMemberOwnerAndAdmin(operationID, groupID *C.char) {
	baseCallback := NewBaseCallback("GetGroupMemberOwnerAndAdmin", operationID)
	open_im_sdk.GetGroupMemberOwnerAndAdmin(baseCallback, C.GoString(operationID), C.GoString(groupID))
}

//export GetGroupMemberListByJoinTimeFilter
func GetGroupMemberListByJoinTimeFilter(operationID, groupID *C.char, offset,
	count C.int32_t, joinTimeBegin, joinTimeEnd C.int64_t, filterUserIDList *C.char) {
	baseCallback := NewBaseCallback("GetGroupMemberListByJoinTimeFilter", operationID)
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(baseCallback, C.GoString(operationID), C.GoString(groupID),
		int32(offset), int32(count), int64(joinTimeBegin), int64(joinTimeEnd), C.GoString(filterUserIDList))
}

//export KickGroupMember
func KickGroupMember(operationID, groupID, reason, userIDList *C.char) {
	baseCallback := NewBaseCallback("KickGroupMember", operationID)
	open_im_sdk.KickGroupMember(baseCallback, C.GoString(operationID), C.GoString(groupID), C.GoString(reason),
		C.GoString(userIDList))
}

//export ChangeGroupMemberMute
func ChangeGroupMemberMute(operationID, groupID, userID *C.char, mutedSeconds C.int) {
	baseCallback := NewBaseCallback("ChangeGroupMemberMute", operationID)
	open_im_sdk.ChangeGroupMemberMute(baseCallback, C.GoString(operationID), C.GoString(groupID), C.GoString(userID),
		int(mutedSeconds))
}

//export ChangeGroupMute
func ChangeGroupMute(operationID, groupID *C.char, isMute C._Bool) {
	baseCallback := NewBaseCallback("ChangeGroupMute", operationID)
	open_im_sdk.ChangeGroupMute(baseCallback, C.GoString(operationID), C.GoString(groupID), bool(isMute))
}

//export TransferGroupOwner
func TransferGroupOwner(operationID, groupID, newOwnerUserID *C.char) {
	baseCallback := NewBaseCallback("TransferGroupOwner", operationID)
	open_im_sdk.TransferGroupOwner(baseCallback, C.GoString(operationID), C.GoString(groupID),
		C.GoString(newOwnerUserID))
}

//export DismissGroup
func DismissGroup(operationID, groupID *C.char) {
	baseCallback := NewBaseCallback("DismissGroup", operationID)
	open_im_sdk.DismissGroup(baseCallback, C.GoString(operationID), C.GoString(groupID))
}

//export GetUsersInGroup
func GetUsersInGroup(operationID, groupID, userIDList *C.char) {
	baseCallback := NewBaseCallback("GetUsersInGroup", operationID)
	open_im_sdk.GetUsersInGroup(baseCallback, C.GoString(operationID), C.GoString(groupID), C.GoString(userIDList))
}

//export IsJoinGroup
func IsJoinGroup(operationID, groupID *C.char) {
	baseCallback := NewBaseCallback("IsJoinGroup", operationID)
	open_im_sdk.IsJoinGroup(baseCallback, C.GoString(operationID), C.GoString(groupID))
}

//export QuitGroup
func QuitGroup(operationID, groupID *C.char) {
	baseCallback := NewBaseCallback("QuitGroup", operationID)
	open_im_sdk.QuitGroup(baseCallback, C.GoString(operationID), C.GoString(groupID))
}

// =====================================================third===============================================

//export UploadLogs
func UploadLogs(operationID *C.char, line C.int, ex *C.char, uuid *C.char) {
	// TODO: uuid not use
	baseCallback := NewBaseCallback("UploadLogs", operationID)
	uploadLogCallback := NewUploadLogProgressCallback("UploadLogs", operationID)
	open_im_sdk.UploadLogs(baseCallback, C.GoString(operationID), int(line), C.GoString(ex), uploadLogCallback)
}

//export Logs
func Logs(operationID *C.char, logLevel C.int, file *C.char, line C.int, msgs *C.char, err *C.char, keyAndValue *C.char) {
	baseCallback := NewBaseCallback("Logs", operationID)
	open_im_sdk.Logs(baseCallback, C.GoString(operationID), int(logLevel), C.GoString(file), int(line), C.GoString(msgs), C.GoString(err), C.GoString(keyAndValue))
}

//export GetSdkVersion
func GetSdkVersion() *C.char {
	return C.CString(open_im_sdk.GetSdkVersion())
}

// =====================================================init_login===============================================

//export InitSDK
func InitSDK(
	imListener C.Openim_Listener, port C.Dart_Port_DL,
	operationID *C.char, config *C.char) C._Bool {

	initSDKInstance(imListener, port)
	if getSDKInstance() == nil {
		return C._Bool(false)
	}

	SetGroupListener("InitSDK", operationID)
	SetConversationListener("InitSDK", operationID)
	SetAdvancedMsgListener("InitSDK", operationID)
	SetBatchMsgListener("InitSDK", operationID)
	SetUserListener("InitSDK", operationID)
	SetFriendListener("InitSDK", operationID)
	SetCustomBusinessListener("InitSDK", operationID)

	callback := NewConnCallback("InitSDK", operationID)
	return C._Bool(open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config)))
}

//export Login
func Login(operationID, userID, token *C.char) {
	baseCallback := NewBaseCallback("Login", operationID)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(userID), C.GoString(token))
}

//export Logout
func Logout(operationID *C.char) {
	baseCallback := NewBaseCallback("Logout", operationID)
	open_im_sdk.Logout(baseCallback, C.GoString(operationID))
}

//export UnInitSDK
func UnInitSDK(operationID *C.char) {
	open_im_sdk.UnInitSDK(C.GoString(operationID))
}

//export ImLogin
func ImLogin(operationID, uid, token *C.char) {
	baseCallback := NewBaseCallback("ImLogin", operationID)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
}

//export ImLogout
func ImLogout(operationID *C.char) {
	baseCallback := NewBaseCallback("ImLogout", operationID)
	open_im_sdk.Logout(baseCallback, C.GoString(operationID))
}

//export SetAppBackgroundStatus
func SetAppBackgroundStatus(operationID *C.char, isBackground C._Bool) {
	baseCallback := NewBaseCallback("SetAppBackgroundStatus", operationID)
	open_im_sdk.SetAppBackgroundStatus(baseCallback, C.GoString(operationID), bool(isBackground))
}

//export NetworkStatusChanged
func NetworkStatusChanged(operationID *C.char) {
	baseCallback := NewBaseCallback("NetworkStatusChanged", operationID)
	open_im_sdk.NetworkStatusChanged(baseCallback, C.GoString(operationID))
}

//export GetLoginStatus
func GetLoginStatus(operationID *C.char) C.int {
	return C.int(open_im_sdk.GetLoginStatus(C.GoString(operationID)))
}

//export GetLoginUserID
func GetLoginUserID() *C.char {
	return C.CString(open_im_sdk.GetLoginUserID())
}

//export UpdateFcmToken
func UpdateFcmToken(operationID, fcmToken *C.char, expireTime C.int64_t) {
	baseCallback := NewBaseCallback("UpdateFcmToken", operationID)
	open_im_sdk.UpdateFcmToken(baseCallback, C.GoString(operationID), C.GoString(fcmToken), int64(expireTime))
}

//export SetAppBadge
func SetAppBadge(operationID *C.char, appUnreadCount C.int32_t) {
	baseCallback := NewBaseCallback("SetAppBadge", operationID)
	open_im_sdk.SetAppBadge(baseCallback, C.GoString(operationID), int32(appUnreadCount))
}

// =====================================================user===============================================

//export GetUsersInfo
func GetUsersInfo(operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback("GetUsersInfo", operationID)
	open_im_sdk.GetUsersInfo(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export GetUsersInfoFromSrv
func GetUsersInfoFromSrv(operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback("GetUsersInfoFromSrv", operationID)
	open_im_sdk.GetUsersInfo(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export SetSelfInfo
func SetSelfInfo(operationID *C.char, userInfo *C.char) {
	baseCallback := NewBaseCallback("SetSelfInfo", operationID)
	open_im_sdk.SetSelfInfo(baseCallback, C.GoString(operationID), C.GoString(userInfo))
}

//export GetSelfUserInfo
func GetSelfUserInfo(operationID *C.char) {
	baseCallback := NewBaseCallback("GetSelfUserInfo", operationID)
	open_im_sdk.GetSelfUserInfo(baseCallback, C.GoString(operationID))
}

// =====================================================online===============================================

//export SubscribeUsersStatus
func SubscribeUsersStatus(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback("SubscribeUsersStatus", operationID)
	open_im_sdk.SubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export UnsubscribeUsersStatus
func UnsubscribeUsersStatus(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback("UnsubscribeUsersStatus", operationID)
	open_im_sdk.UnsubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export GetSubscribeUsersStatus
func GetSubscribeUsersStatus(operationID *C.char) {
	baseCallback := NewBaseCallback("GetSubscribeUsersStatus", operationID)
	open_im_sdk.GetSubscribeUsersStatus(baseCallback, C.GoString(operationID))
}

//export GetUserStatus
func GetUserStatus(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback("GetUserStatus", operationID)
	open_im_sdk.GetUserStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

func main() {

}
