package messageType

const (
	LoginRequest            = uint32(iota) + uint32(1) // 1  登录消息
	LoginResponse                                      // 2  登录结果
	LoginRegisterRequest                               // 3  注册请求
	LoginRegisterResponse                              // 4  注册响应
	UserFriendsListRequest                             // 5  好友列表请求
	UserFriendsListResponse                            // 6  好友列表响应
	FriendsAddRequest                                  // 7  好友添加请求
	FriendsAddResponse                                 // 8  好友添加响应
	OnlineUserListRequest                              // 9  在线用户列表请求
	OnlineUserListResponse                             // 10 在线用户列表请求
	ChatRoomListRequest                                // 11 聊天室列表请求
	ChatRoomListResponse                               // 12 聊天室列表响应
	ChatRoomCreateRequest                              // 13 聊天室创建请求
	ChatRoomCreateResponse                             // 14 聊天室创建响应
	ChatRoomMessageSend                                // 15 单人对聊天室发送信息
	ChatRoomMessageReceive                             // 16 聊天室群发消息
	IntoRoomRequest                                    // 17 进入聊天室请求
	IntoRoomResponse                                   // 18 进入聊天室响应
)
