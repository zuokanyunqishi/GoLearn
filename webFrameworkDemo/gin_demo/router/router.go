package router

import (
	"net/http"
	"speed/app/http/controllers"
	"speed/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsport string = "8086"

var OnlineMap = make(map[string]*websocket.Conn)

func Router(c *gin.Engine) {

	c.Use(middleware.CorsMiddleware(), middleware.Trace)
	c.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello word!")
	})

	{
		c.GET("/index", controllers.HelloC.Index).Use()
	}
	apiGroup := c.Group("api/", func(ctx *gin.Context) {
		ctx.Next()
	})
	{
		apiGroup.POST("login", controllers.LoginC.Login)
		apiGroup.POST("register", controllers.Register.Register)

		// 商品相关接口（公开接口）
		apiGroup.GET("products", controllers.ProductC.List)       // 商品列表
		apiGroup.GET("products/:id", controllers.ProductC.Detail) // 商品详情

		// 商品分类相关接口（公开接口）
		apiGroup.GET("categories", controllers.ProductCategoryC.List)      // 商品分类列表
		apiGroup.GET("categories/tree", controllers.ProductCategoryC.Tree) // 商品分类树
	}

	apiAuth := apiGroup.Use(middleware.JWTAuthMiddleware())
	{
		apiAuth.GET("me", controllers.UserC.Me)
		apiAuth.POST("me/edit", controllers.UserC.Edit)

		// 购物车相关接口（需要认证）
		apiAuth.POST("cart", controllers.CartC.Add)                                // 添加商品到购物车
		apiAuth.GET("cart", controllers.CartC.List)                                // 获取购物车列表
		apiAuth.GET("cart/count", controllers.CartC.Count)                         // 获取购物车商品数量
		apiAuth.POST("cart/update/:id", controllers.CartC.Update)                  // 更新购物车
		apiAuth.POST("cart/del/:id", controllers.CartC.Delete)                     // 删除购物车记录
		apiAuth.POST("cart/batch-selected", controllers.CartC.BatchUpdateSelected) // 批量更新选中状态
		apiAuth.POST("cart/clear", controllers.CartC.Clear)                        // 清空购物车
	}

	//{
	//	c.GET("/first", func(context *gin.Context) {
	//		context.HTML(http.StatusOK, "Firstroom.html", gin.H{"wsport": wsport})
	//
	//	})
	//
	//	c.GET("/diff", func(context *gin.Context) {
	//		context.HTML(http.StatusOK, "Differentroom.html", gin.H{"wsport": wsport})
	//
	//	}).Use(func(context *gin.Context) {
	//
	//	})
	//
	//	c.GET("/some", func(context *gin.Context) {
	//		context.HTML(http.StatusOK, "Sameroom.html", gin.H{"wsport": wsport})
	//
	//	})
	//}
	////webscoket 服务
	//{
	//	c.GET("/ws", func(context *gin.Context) {
	//		//up
	//		var upgrader = websocket.Upgrader{
	//			CheckOrigin: func(r *http.Request) bool {
	//				return true
	//
	//			},
	//		}
	//
	//		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	//		if err != nil {
	//			log.WithCtx(context).Error("握手错误!!!")
	//		}
	//		//read
	//
	//		const (
	//			// Time allowed to write a message to the peer.
	//			writeWait = 10 * time.Second
	//
	//			// Time allowed to read the next pong message from the peer.
	//			pongWait = 60 * time.Second
	//
	//			// Send pings to peer with this period. Must be less than pongWait.
	//			pingPeriod = (pongWait * 9) / 10
	//
	//			// Maximum message size allowed from peer.
	//			maxMessageSize = 512
	//		)
	//
	//		var (
	//			newline = []byte{'\n'}
	//			space   = []byte{' '}
	//		)
	//
	//		OnlineMap[uuid.New().String()] = conn
	//		go func() {
	//			conn.SetReadLimit(maxMessageSize)
	//			conn.SetReadDeadline(time.Now().Add(pongWait))
	//			//conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	//			for {
	//				_, message, err := conn.ReadMessage()
	//
	//				if err != nil {
	//					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	//						fmt.Printf("error: %v", err)
	//					}
	//					break
	//				}
	//				message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
	//				fmt.Printf("收到消息--%s--,%s\n", string(message), conn.RemoteAddr())
	//				conn.WriteMessage(websocket.TextMessage, []byte("我是websocket"))
	//				fmt.Println("返回消息---", "我是websocket")
	//
	//			}
	//		}()
	//
	//	})
	//}

}
