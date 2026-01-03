package main

import (
	_ "gfDeomo/internal/packed"

	_ "gfDeomo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gfDeomo/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
