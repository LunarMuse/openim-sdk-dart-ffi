package main

/*
#include <stdio.h>
#include <stdbool.h>
#include "openim_sdk_ffi.h"
static void callOnMethodChannel(Openim_Listener listener, Dart_Port_DL port, char *message) {
    listener.onMethodChannel(port, message);
}
*/
import "C"
import "github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"

// ================================================begin global callback================================================
// =======================================================Begin Base====================================================
type BaseCallback struct {
	//cCallback   C.CB_S_I_S_S
	//operationID string

	callMethodName string
	operationID    *C.char
}

func NewBaseCallback(callMethodName string, operationID *C.char) *BaseCallback {
	return &BaseCallback{callMethodName, operationID}
}

func (b BaseCallback) OnError(errCode int32, errMsg string) {
	//C.Call_CB_S_I_S_S(b.cCallback, C.CString(b.operationID), C.int(errCode), C.CString(errMsg), NO_DATA)

	m := make(map[string]any)
	m["method"] = C.CString("OnError")
	m["errCode"] = C.int(errCode)
	m["data"] = C.CString(errMsg)
	m["operationId"] = b.operationID
	m["callMethodName"] = C.CString(b.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (b BaseCallback) OnSuccess(data string) {
	//C.Call_CB_S_I_S_S(b.cCallback, C.CString(b.operationID), NO_ERR, NO_ERR_MSG, C.CString(data))

	m := make(map[string]any)
	m["method"] = C.CString("OnSuccess")
	m["errCode"] = C.int(0)
	m["data"] = C.CString(data)
	m["operationId"] = b.operationID
	m["callMethodName"] = C.CString(b.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ========================================================End Base=====================================================

// ==================================================Begin SendMsgCallBack==============================================
type SendMessageCallback struct {
	//cCallback   C.CB_S_I_S_S_I
	//operationID string

	callMethodName string
	operationID    *C.char
}

func NewSendMessageCallback(callMethodName string, operationID *C.char) *SendMessageCallback {
	return &SendMessageCallback{callMethodName, operationID}
}

func (s SendMessageCallback) OnError(errCode int32, errMsg string) {
	//C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), C.int(errCode), C.CString(errMsg), NO_DATA, NO_PROGRESS)

	data := make(map[string]any)
	data["method"] = C.CString("OnError")
	data["errCode"] = C.int(errCode)
	data["data"] = C.CString(errMsg)
	data["operationId"] = s.operationID
	data["callMethodName"] = C.CString(s.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (s SendMessageCallback) OnSuccess(data string) {
	//C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), NO_ERR, NO_ERR_MSG, C.CString(data), NO_PROGRESS)

	m := make(map[string]any)
	m["method"] = C.CString("OnSuccess")
	m["errCode"] = C.int(0)
	m["data"] = C.CString(data)
	m["operationId"] = s.operationID
	m["callMethodName"] = C.CString(s.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (s SendMessageCallback) OnProgress(progress int) {
	//C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), NO_ERR, NO_ERR_MSG, NO_DATA, C.int(progress))

	m := make(map[string]any)
	m["method"] = C.CString("OnProgress")
	m["errCode"] = C.int(0)
	m["data"] = C.int(progress)
	m["operationId"] = s.operationID
	m["callMethodName"] = C.CString(s.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ===================================================End SendMsgCallBack===============================================

// ==================================================Begin OnConnListener===============================================
type ConnCallback struct {
	callMethodName string
	operationID    *C.char
}

func NewConnCallback(callMethodName string, operationID *C.char) *ConnCallback {
	return &ConnCallback{callMethodName: callMethodName, operationID: operationID}
}

func (c ConnCallback) OnConnecting() {
	//C.Call_CB_I_S(c.cCallback, CONNECTING, NO_DATA)

	m := make(map[string]any)
	m["method"] = C.CString("OnConnecting")
	m["data"] = C.CString("")
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)
	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConnCallback) OnConnectSuccess() {
	//C.Call_CB_I_S(c.cCallback, CONNECT_SUCCESS, NO_DATA)
	m := make(map[string]any)
	m["method"] = C.CString("OnConnectSuccess")
	m["data"] = C.CString("")
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	//C.Call_CB_I_S(c.cCallback, CONNECT_FAILED, C.CString(StructToJsonString(Base{ErrCode: errCode, ErrMsg: errMsg})))
	m := make(map[string]any)
	m["method"] = C.CString("OnConnectFailed")
	m["data"] = C.CString("")
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConnCallback) OnKickedOffline() {
	//C.Call_CB_I_S(c.cCallback, KICKED_OFFLINE, NO_DATA)
	m := make(map[string]any)
	m["method"] = C.CString("OnKickedOffline")
	m["data"] = C.CString("")
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConnCallback) OnUserTokenExpired() {
	//C.Call_CB_I_S(c.cCallback, USER_TOKEN_EXPIRED, NO_DATA)

	m := make(map[string]any)
	m["method"] = C.CString("OnUserTokenExpired")
	m["data"] = C.CString("")
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConnCallback) OnUserTokenInvalid(errMsg string) {
	//C.Call_CB_I_S(c.cCallback, USER_TOKEN_INVALID, C.CString(errMsg))
	m := make(map[string]any)
	m["method"] = C.CString("OnUserTokenInvalid")
	m["data"] = C.CString(errMsg)
	m["operationId"] = c.operationID
	m["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ====================================================End OnConnListener===============================================

// =================================================Begin OnGroupListener===============================================
type GroupCallback struct {
	//cCallback C.CB_I_S
	callMethodName string
	operationID    *C.char
}

func NewGroupCallback(callMethodName string, operationID *C.char) *GroupCallback {
	return &GroupCallback{callMethodName: callMethodName, operationID: operationID}
}

func (g GroupCallback) OnJoinedGroupAdded(groupInfo string) {
	//C.Call_CB_I_S(g.cCallback, JOINED_GROUP_ADDED, C.CString(groupInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnJoinedGroupAdded")
	data["data"] = C.CString(groupInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnJoinedGroupDeleted(groupInfo string) {
	//C.Call_CB_I_S(g.cCallback, JOINED_GROUP_DELETED, C.CString(groupInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnJoinedGroupDeleted")
	data["data"] = C.CString(groupInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupMemberAdded(groupMemberInfo string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_ADDED, C.CString(groupMemberInfo))

	data := make(map[string]any)
	data["method"] = C.CString("GROUP_MEMBER_ADDED")
	data["data"] = C.CString(groupMemberInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupMemberDeleted(groupMemberInfo string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_DELETED, C.CString(groupMemberInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupMemberDeleted")
	data["data"] = C.CString(groupMemberInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupApplicationAdded(groupApplication string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_ADDED, C.CString(groupApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupApplicationAdded")
	data["data"] = C.CString(groupApplication)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupApplicationDeleted(groupApplication string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_DELETED, C.CString(groupApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupApplicationDeleted")
	data["data"] = C.CString(groupApplication)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupInfoChanged(groupInfo string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_INFO_CHANGED, C.CString(groupInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupInfoChanged")
	data["data"] = C.CString(groupInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupDismissed(groupInfo string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_DISMISSED, C.CString(groupInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupDismissed")
	data["data"] = C.CString(groupInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupMemberInfoChanged(groupMemberInfo string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_INFO_CHANGED, C.CString(groupMemberInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupMemberInfoChanged")
	data["data"] = C.CString(groupMemberInfo)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupApplicationAccepted(groupApplication string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_ACCEPTED, C.CString(groupApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupApplicationAccepted")
	data["data"] = C.CString(groupApplication)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (g GroupCallback) OnGroupApplicationRejected(groupApplication string) {
	//C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_REJECTED, C.CString(groupApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnGroupApplicationRejected")
	data["data"] = C.CString(groupApplication)
	data["operationId"] = g.operationID
	data["callMethodName"] = C.CString(g.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ===================================================End OnGroupListener===============================================

// ================================================Begin OnFriendshipListener===========================================
type FriendCallback struct {
	callMethodName string
	operationID    *C.char
}

func NewFriendCallback(callMethodName string, operationID *C.char) *FriendCallback {
	return &FriendCallback{callMethodName: callMethodName, operationID: operationID}
}

func (f FriendCallback) OnFriendApplicationAdded(friendApplication string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_ADDED, C.CString(friendApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendApplicationAdded")
	data["data"] = C.CString(friendApplication)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendApplicationDeleted(friendApplication string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_DELETED, C.CString(friendApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendApplicationDeleted")
	data["data"] = C.CString(friendApplication)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendApplicationAccepted(friendApplication string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_ACCEPTED, C.CString(friendApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendApplicationAccepted")
	data["data"] = C.CString(friendApplication)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendApplicationRejected(friendApplication string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_REJECTED, C.CString(friendApplication))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendApplicationRejected")
	data["data"] = C.CString(friendApplication)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendAdded(friendInfo string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_ADDED, C.CString(friendInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendAdded")
	data["data"] = C.CString(friendInfo)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendDeleted(friendInfo string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_DELETED, C.CString(friendInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendDeleted")
	data["data"] = C.CString(friendInfo)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnFriendInfoChanged(friendInfo string) {
	//C.Call_CB_I_S(f.cCallback, FRIEND_INFO_CHANGED, C.CString(friendInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnFriendInfoChanged")
	data["data"] = C.CString(friendInfo)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnBlackAdded(blackInfo string) {
	//C.Call_CB_I_S(f.cCallback, BLACK_ADDED, C.CString(blackInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnBlackAdded")
	data["data"] = C.CString(blackInfo)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (f FriendCallback) OnBlackDeleted(blackInfo string) {
	//C.Call_CB_I_S(f.cCallback, BLACK_DELETED, C.CString(blackInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnBlackDeleted")
	data["data"] = C.CString(blackInfo)
	data["operationId"] = f.operationID
	data["callMethodName"] = C.CString(f.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ==================================================End OnFriendshipListener===========================================

// ==============================================Begin OnConversationListener===========================================
type ConversationCallback struct {
	callMethodName string
	operationID    *C.char
}

func NewConversationCallback(callMethodName string, operationID *C.char) *ConversationCallback {
	return &ConversationCallback{callMethodName: callMethodName, operationID: operationID}
}

func (c ConversationCallback) OnSyncServerStart(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	//C.Call_CB_I_S(c.cCallback, SYNC_SERVER_START, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnSyncServerStart")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnSyncServerFinish(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	//C.Call_CB_I_S(c.cCallback, SYNC_SERVER_FINISH, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnSyncServerFinish")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnSyncServerProgress(progress int) {
	m := make(map[string]any)
	m["progress"] = progress
	//C.Call_CB_I_S(c.cCallback, SYNC_SERVER_PROGRESS, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnSyncServerProgress")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnSyncServerFailed(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	//C.Call_CB_I_S(c.cCallback, SYNC_SERVER_FAILED, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnSyncServerFailed")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnNewConversation(conversationList string) {
	//C.Call_CB_I_S(c.cCallback, NEW_CONVERSATION, C.CString(conversationList))

	data := make(map[string]any)
	data["method"] = C.CString("OnNewConversation")
	data["data"] = C.CString(StructToJsonString(conversationList))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnConversationChanged(conversationList string) {
	//C.Call_CB_I_S(c.cCallback, CONVERSATION_CHANGED, C.CString(conversationList))

	data := make(map[string]any)
	data["method"] = C.CString("OnConversationChanged")
	data["data"] = C.CString(StructToJsonString(conversationList))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	//C.Call_CB_I_S(c.cCallback, TOTAL_UNREAD_MESSAGE_COUNT_CHANGED, C.CString(IntToString(totalUnreadCount)))

	data := make(map[string]any)
	data["method"] = C.CString("OnTotalUnreadMessageCountChanged")
	data["data"] = C.CString(IntToString(totalUnreadCount))
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (c ConversationCallback) OnConversationUserInputStatusChanged(change string) {
	//C.Call_CB_I_S(c.cCallback, CONVERSATION_USER_INPUT_STATUS_CHANGED, C.CString(change))

	data := make(map[string]any)
	data["method"] = C.CString("OnConversationUserInputStatusChanged")
	data["data"] = C.CString(change)
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ================================================End OnConversationListener===========================================

// ===============================================Begin OnAdvancedMsgListener===========================================
type AdvancedMsgCallback struct {
	callMethodName string
	operationID    *C.char
}

func NewAdvancedMsgCallback(callMethodName string, operationID *C.char) *AdvancedMsgCallback {
	return &AdvancedMsgCallback{callMethodName: callMethodName, operationID: operationID}
}

func (a AdvancedMsgCallback) OnRecvNewMessage(message string) {
	//C.Call_CB_I_S(a.cCallback, RECV_NEW_MESSAGE, C.CString(message))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvNewMessage")
	data["data"] = C.CString(message)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message = StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvC2CReadReceipt(msgReceiptList string) {
	//C.Call_CB_I_S(a.cCallback, RECV_C2C_READ_RECEIPT, C.CString(msgReceiptList))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvC2CReadReceipt")
	data["data"] = C.CString(msgReceiptList)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnNewRecvMessageRevoked(messageRevoked string) {
	//C.Call_CB_I_S(a.cCallback, NEW_RECV_MESSAGE_REVOKED, C.CString(messageRevoked))

	data := make(map[string]any)
	data["method"] = C.CString("OnNewRecvMessageRevoked")
	data["data"] = C.CString(messageRevoked)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvOfflineNewMessage(message string) {
	//C.Call_CB_I_S(a.cCallback, RECV_OFFLINE_NEW_MESSAGE, C.CString(message))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvOfflineNewMessage")
	data["data"] = C.CString(message)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message = StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnMsgDeleted(message string) {
	//C.Call_CB_I_S(a.cCallback, MSG_DELETED, C.CString(message))

	data := make(map[string]any)
	data["method"] = C.CString("OnMsgDeleted")
	data["data"] = C.CString(message)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message = StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvOnlineOnlyMessage(message string) {
	//C.Call_CB_I_S(a.cCallback, RECV_ONLINE_ONLY_MESSAGE, C.CString(message))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvOnlineOnlyMessage")
	data["data"] = C.CString(message)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message = StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// TODO: why not
func (a AdvancedMsgCallback) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	//C.Call_CB_I_S(a.cCallback, RECV_GROUP_READ_RECEIPT, C.CString(groupMsgReceiptList))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvGroupReadReceipt")
	data["data"] = C.CString(groupMsgReceiptList)
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionList"] = reactionExtensionList
	//C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_CHANGED, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvMessageExtensionsChanged")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionKeyList"] = reactionExtensionKeyList
	//C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_DELETED, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvMessageExtensionsDeleted")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionList"] = reactionExtensionList
	//C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_ADDED, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvMessageExtensionsAdded")
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = a.operationID
	data["callMethodName"] = C.CString(a.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// =================================================End OnAdvancedMsgListener===========================================

// =================================================Begin OnUserListener================================================
type UserCallback struct {
	//cCallback C.CB_I_S

	callMethodName string
	operationID    *C.char
}

func NewUserCallback(callMethodName string, operationID *C.char) *UserCallback {
	return &UserCallback{callMethodName, operationID}
}

func (u UserCallback) OnSelfInfoUpdated(userInfo string) {
	//C.Call_CB_I_S(u.cCallback, SELF_INFO_UPDATED, C.CString(userInfo))

	data := make(map[string]any)
	data["method"] = C.CString("OnSelfInfoUpdated")
	data["data"] = C.CString(userInfo)
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UserCallback) OnUserStatusChanged(statusMap string) {
	//C.Call_CB_I_S(u.cCallback, USER_STATUS_CHANGED, C.CString(statusMap))

	data := make(map[string]any)
	data["method"] = C.CString("OnUserStatusChanged")
	//data["errCode"] = // TODO:
	data["data"] = C.CString(statusMap)
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// TODO:why not
func (u UserCallback) OnUserCommandAdd(userCommand string)    {}
func (u UserCallback) OnUserCommandDelete(userCommand string) {}
func (u UserCallback) OnUserCommandUpdate(userCommand string) {}

// ===================================================End OnUserListener================================================

// ============================================Begin OnCustomBusinessListener===========================================
type CustomBusinessCallback struct {
	//cCallback C.CB_I_S
	callMethodName string
	operationID    *C.char
}

func NewCustomBusinessCallback(callMethodName string, operationID *C.char) *CustomBusinessCallback {
	return &CustomBusinessCallback{callMethodName, operationID}
}

func (c CustomBusinessCallback) OnRecvCustomBusinessMessage(businessMessage string) {
	//C.Call_CB_I_S(c.cCallback, RECV_CUSTOM_BUSINESS_MESSAGE, C.CString(businessMessage))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvCustomBusinessMessage")
	data["data"] = C.CString(businessMessage)
	data["operationId"] = c.operationID
	data["callMethodName"] = C.CString(c.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// =============================================End OnCustomBusinessListener============================================

// ============================================Begin OnMessageKvInfoListener============================================
// TODO:
// =============================================End OnMessageKvInfoListener=============================================

// ==============================================Begin OnListenerForService=============================================
// TODO:
// ================================================End OnListenerForService=============================================

// ==============================================Begin OnSignalingListener==============================================
// TODO:
// ================================================End OnSignalingListener==============================================

// ==============================================Begin UploadFileCallback===============================================
type UploadFileCallback struct {
	//cCallback C.CB_I_S

	callMethodName string
	operationID    *C.char
}

func NewUploadFileCallback(callMethodName string, operationID *C.char) *UploadFileCallback {
	return &UploadFileCallback{callMethodName, operationID}
}

func (u UploadFileCallback) Open(size int64) {
	//C.Call_CB_I_S(u.cCallback, OPEN, C.CString(IntToString(size)))

	m := make(map[string]any)
	m["method"] = C.CString("Open")
	m["errCode"] = C.int(0)
	m["data"] = C.CString(IntToString(size))
	m["operationId"] = u.operationID
	m["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(m)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) PartSize(partSize int64, num int) {
	m := make(map[string]any)
	m["partSize"] = partSize
	m["num"] = num
	//C.Call_CB_I_S(u.cCallback, PART_SIZE, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("PartSize")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) HashPartProgress(index int, size int64, partHash string) {
	m := make(map[string]any)
	m["index"] = index
	m["size"] = size
	m["partHash"] = partHash
	//C.Call_CB_I_S(u.cCallback, HASH_PART_PROGRESS, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("HashPartProgress")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) HashPartComplete(partsHash string, fileHash string) {
	m := make(map[string]any)
	m["partsHash"] = partsHash
	m["fileHash"] = fileHash
	//C.Call_CB_I_S(u.cCallback, HASH_PART_COMPLETE, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("HashPartComplete")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) UploadID(uploadID string) {
	//C.Call_CB_I_S(u.cCallback, UPLOAD_ID, C.CString(uploadID))

	data := make(map[string]any)
	data["method"] = C.CString("UploadID")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(uploadID)
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) UploadPartComplete(index int, partSize int64, partHash string) {
	m := make(map[string]any)
	m["index"] = index
	m["partSize"] = partSize
	m["partHash"] = partHash
	//C.Call_CB_I_S(u.cCallback, UPLOAD_PART_COMPLETE, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("UploadPartComplete")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) UploadComplete(fileSize int64, streamSize int64, storageSize int64) {
	m := make(map[string]any)
	m["fileSize"] = fileSize
	m["streamSize"] = streamSize
	m["storageSize"] = storageSize
	//C.Call_CB_I_S(u.cCallback, UPLOAD_COMPLETE, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("UploadComplete")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (u UploadFileCallback) Complete(size int64, url string, typ int) {
	m := make(map[string]any)
	m["size"] = size
	m["url"] = url
	m["typ"] = typ
	//C.Call_CB_I_S(u.cCallback, COMPLETE, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("Complete")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = u.operationID
	data["callMethodName"] = C.CString(u.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ================================================End UploadFileCallback===============================================

// ===============================================Begin UploadLogProgress===============================================
type UploadLogProgressCallback struct {
	//cCallback C.CB_I_S

	callMethodName string
	operationID    *C.char
}

func NewUploadLogProgressCallback(callMethodName string, operationID *C.char) *UploadLogProgressCallback {
	return &UploadLogProgressCallback{callMethodName, operationID}
}

func (l UploadLogProgressCallback) OnProgress(current, size int64) {
	m := make(map[string]any)
	m["current"] = current
	m["size"] = size
	//C.Call_CB_I_S(l.cCallback, ON_PROGRESS, C.CString(StructToJsonString(m)))

	data := make(map[string]any)
	data["method"] = C.CString("Complete")
	data["errCode"] = C.int(0)
	data["data"] = C.CString(StructToJsonString(m))
	data["operationId"] = l.operationID
	data["callMethodName"] = C.CString(l.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

// ================================================End UploadLogProgress================================================

func SetGroupListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetGroupListener(NewGroupCallback(callMethodName, operationID))
}

func SetConversationListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetConversationListener(NewConversationCallback(callMethodName, operationID))
}

func SetAdvancedMsgListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetAdvancedMsgListener(NewAdvancedMsgCallback(callMethodName, operationID))
}

func SetBatchMsgListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetBatchMsgListener(NewBatchMessageCallback(callMethodName, operationID))
}

func SetUserListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetUserListener(NewUserCallback(callMethodName, operationID))
}

func SetFriendListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetFriendListener(NewFriendCallback(callMethodName, operationID))
}

func SetCustomBusinessListener(callMethodName string, operationID *C.char) {
	open_im_sdk.SetCustomBusinessListener(NewCustomBusinessCallback(callMethodName, operationID))
}

//	func SetMessageKvInfoListener(callMethodName string, operationID *C.char) {
//		open_im_sdk.SetMessageKvInfoListener(NewMessageKVCallback(callMethodName, operationID))
//	}
//

// ================================================end global callback==================================================
type BatchMessageCallback struct {
	callMethodName string
	operationID    *C.char
}

func NewBatchMessageCallback(callMethodName string, operationID *C.char) *BatchMessageCallback {
	return &BatchMessageCallback{callMethodName: callMethodName, operationID: operationID}
}

func (b BatchMessageCallback) OnRecvNewMessages(messageList string) {
	//C.Call_CB_I_S(b.cCallback, RECV_NEW_MESSAGES, C.CString(messageList))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvNewMessages")
	data["data"] = C.CString(messageList)
	data["operationId"] = b.operationID
	data["callMethodName"] = C.CString(b.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}

func (b BatchMessageCallback) OnRecvOfflineNewMessages(messageList string) {
	//C.Call_CB_I_S(b.cCallback, RECV_OFFLINE_NEW_MESSAGES, C.CString(messageList))

	data := make(map[string]any)
	data["method"] = C.CString("OnRecvOfflineNewMessages")
	data["data"] = C.CString(messageList)
	data["operationId"] = b.operationID
	data["callMethodName"] = C.CString(b.callMethodName)
	message := StructToJsonString(data)

	C.callOnMethodChannel(instance.listener, instance.port, C.CString(message))
}
