package service

import (
	"GoLearn/chat/client/service/chatRoom"
	"GoLearn/chat/client/service/user"
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/util"
	"GoLearn/chat/util/color"
	"GoLearn/chat/util/zlog"
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"
)

type Client struct {
	draw        *util.IO
	ClientServe *util.Client
	user        *user.User
	currentRoom baseCommon.ChatRoom
}

func (c *Client) Run() {

	for {
		c.draw.Writeln(
			"-----------  在线聊天系统 -----------",
			"       1.  注册账号            ",
			"       2.  登陆系统            ",
			"       3.  退出系统             ",
			"-----------  在线聊天系统 -----------",
		)
		c.processInput(c.draw.Read())

	}
}

func (c *Client) processInput(input string) {
	switch input {
	case "1":
		c.registerProcess()
	case "2":
		//
		c.loginProcess()
	case "3":
		os.Exit(0)
	default:
		c.draw.Writeln(color.Red("输入错误请重新输入"))

	}

}

func (c *Client) loginProcess() {
	c.draw.Writeln(
		[]string{
			"-----------  在线聊天系统 -----------",
			"",
			"用户名：",
		}...)

	var userName, userPwd string
	userName = c.draw.Read()
	c.draw.Writeln("密码：")
	userPwd = c.draw.Read()
	tmpUser := user.NewUser(c.ClientServe)
	err := tmpUser.Login(userName, userPwd)
	if err != nil {
		zlog.Error(err.Error())
		c.draw.Writeln(color.Red(err.Error()))
		return
	}
	c.user = tmpUser
	c.ChatMenu()
}

func (c *Client) GetServer() *util.Client {
	return c.ClientServe
}

func (c *Client) registerProcess() {
	c.draw.Writeln(color.Blue("-----------注册用户-----------------"))
	c.draw.Writeln("用户名: ")
	userName := c.draw.Read()
	var userPwd, cPwd string
	for {
		c.draw.Writeln("密码: ")
		userPwd = c.draw.Read()
		c.draw.Writeln("确认密码: ")
		cPwd = c.draw.Read()
		if cPwd == userPwd {
			break
		}
		c.draw.Writeln(color.Red("密码和确认密码不一致 !,请重新输入"))
	}
	err := user.NewUser(c.ClientServe).Register(userName, userPwd, cPwd)
	if err != nil {
		c.draw.Writeln(color.Red(err.Error()))
		return
	}
	c.draw.Writeln(color.Green("注册成功,请登录..."))
}

// 聊天目录
func (c *Client) ChatMenu() {
	for {
		c.draw.Writeln([]string{
			color.Green(fmt.Sprintf("--------  欢迎 %s 登录 -------", c.user.UserName)),
			"        1. 消息列表",
			"        2. 好友列表",
			"        3. 添加好友",
			"        4. 聊天室",
			"        5. 退出系统",
		}...)
		c.ChatMenuInput()

	}

}
func (c *Client) ChatMenuInput() {
	switch c.draw.Read() {
	case "1":
		c.draw.Writeln(color.Red("暂不支持"))
	case "2":
		c.processFriendsList()
	case "3":
		c.ProcessAddFriends()
	case "4":
		c.ProcessChatRoom()
	case "5":
		c.draw.Writeln(color.Green("退出系统中。。。"))
		os.Exit(0)
	default:
		c.draw.Writeln(color.Red("输入错误请重新输入"))
	}

}

func (c *Client) processFriendsList() {
	friends, err := c.user.GetFriends()
	if err != nil {
		c.draw.Writeln(color.Red(err.Error()))
		return
	}
	// 渲染好友数据

	tableModel := table.NewWriter()
	tableModel.AppendHeader([]interface{}{"好友id", "好友姓名", "好友昵称", "当前状态"})
	for _, friend := range friends {
		tableModel.AppendRow(table.Row{friend.UserId, friend.UserName, friend.NickName, friend.Status})
	}
	tableModel.SetColumnConfigs([]table.ColumnConfig{
		{Align: text.AlignCenter, Number: 1},
		{Align: text.AlignCenter, Number: 2},
		{Align: text.AlignCenter, Number: 3},
		{Align: text.AlignCenter, Number: 4},
	})

	c.draw.Writeln("")
	c.draw.Writeln(fmt.Sprintf("----%s 的好友------", c.user.UserName))
	c.draw.Writeln(tableModel.Render())
	c.draw.Writeln("")

InputRetry:
	c.draw.Writeln([]string{
		" 1. 返回上级菜单 ",
		" 2. 退出系统   ",
	}...)
	readInput := c.draw.Read()
	if readInput == "2" {
		os.Exit(0)
	}

	if readInput == "1" {
		return
	}
	c.draw.Writeln(color.Red("输入错误请重新输入..."))
	goto InputRetry

}

func (c *Client) ProcessAddFriends() {
	users, err := c.user.GetOnlineUsers()
	if err != nil {
		c.draw.Writeln("暂时无用户" + err.Error())
	}

	tableModel := table.NewWriter()
	tableModel.AppendHeader([]interface{}{"用户id", "用户姓名", "用户昵称", "当前状态"})
	for _, cuser := range users {
		if cuser.UserName == c.user.UserName {
			continue
		}
		tableModel.AppendRow(table.Row{cuser.UserId, cuser.UserName, cuser.NickName, cuser.Status})
	}
	tableModel.SetColumnConfigs([]table.ColumnConfig{
		{Align: text.AlignCenter, Number: 1},
		{Align: text.AlignCenter, Number: 2},
		{Align: text.AlignCenter, Number: 3},
		{Align: text.AlignCenter, Number: 4},
	})

	c.draw.Writeln("")
	c.draw.Writeln("----在线用户列表------")
	c.draw.Writeln(tableModel.Render())
	c.draw.Writeln("")
	c.draw.Writeln("请输入用户id,添加好友")
	userId := uint32(c.draw.ReadInt())
	err = c.user.AddFriend(userId)
	if err != nil {
		c.draw.Writeln(color.Red(err.Error()))
		return
	}
	c.draw.Writeln(color.Green("添加成功..."))

}

func (c *Client) ChatRoomMenu() {
	// 渲染当前聊天室列表 (可用) todo
	list, err := chatRoom.ChatRooms(c.user.GetConn(), c.user.GetToken())
	if err != nil {
		c.draw.Writeln(color.Red(err.Error()))
		return
	}

	if len(list) > 0 {
		c.draw.Writeln(
			c.RenderChatRoomList(list),
		)
	}

	c.draw.Writeln([]string{
		color.Green("----------------------------------------"),
		"1. 创建聊天室",
		"2. 进入聊天室",
		"3. 返回登录后界面",
		"4. 退出系统",
		"",
	}...)
RetryInput:
	input := c.draw.Read()
	switch input {
	case "1":
		c.ProcessCreateChatRoom()
	case "2":
		c.ProcessIntoChatRoom()
	case "3":
		return
	case "4":
		c.draw.Writeln("退出系统中...")
	default:
		c.InputErrorNotice()
		goto RetryInput

	}

}

func (c *Client) InputErrorNotice() {
	c.draw.Writeln(color.Red("输入错误,请重新数据"))
}

func (c *Client) RenderChatRoomList(list []baseCommon.ChatRoom) string {
	tableModel := table.NewWriter()
	tableModel.AppendHeader([]interface{}{"聊天室id", "聊天室名字", "创建者姓名", "创建者用户id"})
	for _, room := range list {
		tableModel.AppendRow(table.Row{room.RoomId, room.RoomName, room.CreateUser.UserName, room.CreateUser.UserId})
	}

	tableModel.SetColumnConfigs([]table.ColumnConfig{
		{Align: text.AlignCenter, Number: 1},
		{Align: text.AlignCenter, Number: 2},
		{Align: text.AlignCenter, Number: 3},
		{Align: text.AlignCenter, Number: 4},
	})
	return tableModel.Render()
}

func (c *Client) ProcessChatRoom() {
	c.ChatRoomMenu()
}

func (c *Client) ProcessCreateChatRoom() {
	c.draw.Writeln([]string{"------创建聊天室-------", ""}...)
	var room baseCommon.ChatRoom
	c.draw.Writeln("请输入聊天室名称: ")
	room.RoomName = c.draw.Read()
	c.draw.Writeln("请输入聊天室描述:")
	room.RoomDescribe = c.draw.Read()
	result := chatRoom.CreateRoom(c.user.GetConn(), c.user.GetToken(), room)
	if result != nil {
		c.draw.Writeln([]string{"聊天室创建失败", color.Red(result.Error())}...)
		return
	}

	c.SuccessNotice("创建聊天室成功")

}

func (c *Client) SuccessNotice(str ...string) {

	for i, value := range str {
		str[i] = color.Green(value)
	}
	c.draw.Writeln(str...)
}

func (c *Client) FailedNotice(strs ...string) {
	for i, value := range strs {
		strs[i] = color.Green(value)
	}
	c.draw.Writeln(strs...)
}

func (c *Client) ProcessIntoChatRoom() {
	// 输入聊天室的id 进入聊天室,连接成功后,渲染当前聊天室的人员,
	c.draw.Writeln("请输入聊天室 id :")
	roomId := uint32(c.draw.ReadInt())

	room, err := c.user.IntoRoom(roomId) // todo
	if err != nil {
		c.FailedNotice(err.Error())
		return
	}
	c.currentRoom = room
	ctx, cancel := context.WithCancel(context.Background())
	// 后台启动一个协程接收聊天室消息,
	go func(ctx context.Context) {

		for {
			select {
			case roomMessage := <-c.user.ReceiveRoomMessages():
				if roomMessage.User.UserId != c.user.UserId {
					c.draw.Writeln(
						color.Blue(roomMessage.User.UserName) + "说: " + roomMessage.Message,
					)
				}
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second * 1)
			}
		}

	}(ctx)
	// 接收输入参数,发消息
RoomMessageLoopBegin:
	c.draw.Writeln("请输入 1:发送消息 ,2:退出聊天室")
	input := c.draw.Read()
	switch input {
	case "1":
		message := c.draw.Read()
		c.user.SendMessage(c.currentRoom.RoomId, message)
	case "2":
		cancel()
		return
	default:
		c.InputErrorNotice()
	}
	goto RoomMessageLoopBegin

}
func NewClient() *Client {

	return &Client{draw: util.NewIO(), ClientServe: util.NewClientServer()}
}
