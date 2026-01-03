package route

import (
	"GoLearn/chat/client/controller"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/util"
)

var (
	IntoRoomController           = new(controller.IntoRoomRoute)
	FriendsAddController         = new(controller.FriendsAddRoute)
	UserFriendsController        = new(controller.UserFriendsRoute)
	OnlineUsersController        = new(controller.OnlineUsersRoute)
	ChatRoomListController       = new(controller.ChatRoomListRoute)
	LoginResponseController      = new(controller.LoginResponseRoute)
	ChatRoomCreateController     = new(controller.ChatRoomCreateRoute)
	RegisterResponseController   = new(controller.RegisterResponseRoute)
	RoomMessageReceiveController = new(controller.RoomMessageReceiveRoute)
)

func AddRoute(serv *util.Client) {
	serv.AddRouter(messageType.LoginResponse, LoginResponseController)
	serv.AddRouter(messageType.LoginRegisterResponse, RegisterResponseController)
	serv.AddRouter(messageType.UserFriendsListResponse, UserFriendsController)
	serv.AddRouter(messageType.FriendsAddResponse, FriendsAddController)
	serv.AddRouter(messageType.OnlineUserListResponse, OnlineUsersController)
	serv.AddRouter(messageType.ChatRoomListResponse, ChatRoomListController)
	serv.AddRouter(messageType.ChatRoomCreateResponse, ChatRoomCreateController)
	serv.AddRouter(messageType.IntoRoomResponse, IntoRoomController)
	serv.AddRouter(messageType.ChatRoomMessageReceive, RoomMessageReceiveController)
}
