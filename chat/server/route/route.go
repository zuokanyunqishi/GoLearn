package route

import (
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/server/controller"
	"GoLearn/chat/util"
)

var LoginController = new(controller.Login)
var RegisterController = new(controller.Register)
var UserFriendsController = new(controller.UserFriendsList)
var OnlineUserController = new(controller.OnlineUserRoute)
var FriendsAddController = new(controller.FriendsAdd)
var ChatRoomListController = new(controller.ChatRoomListRoute)
var ChatRoomCreateController = new(controller.ChatRoomCreate)
var RoomMessageRevController = new(controller.RoomMessageReceive)
var IntoRoomController = new(controller.IntoRoomRouter)

func AddRoute(serv *util.Server) {
	serv.AddRouter(messageType.LoginRequest, LoginController)
	serv.AddRouter(messageType.LoginRegisterRequest, RegisterController)
	serv.AddRouter(messageType.UserFriendsListRequest, UserFriendsController)
	serv.AddRouter(messageType.OnlineUserListRequest, OnlineUserController)
	serv.AddRouter(messageType.FriendsAddRequest, FriendsAddController)
	serv.AddRouter(messageType.ChatRoomListRequest, ChatRoomListController)
	serv.AddRouter(messageType.ChatRoomCreateRequest, ChatRoomCreateController)
	serv.AddRouter(messageType.ChatRoomMessageSend, RoomMessageRevController)
	serv.AddRouter(messageType.IntoRoomRequest, IntoRoomController)
}
