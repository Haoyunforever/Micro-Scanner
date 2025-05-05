package http

import (
	"github.com/Haoyunforever/Micro-Scanner/idl/pb"
	"github.com/Haoyunforever/Micro-Scanner/pkg/ctl"
	log "github.com/Haoyunforever/Micro-Scanner/pkg/logger"
	"github.com/Haoyunforever/Micro-Scanner/pkg/utils"
	"github.com/Haoyunforever/Micro-Scanner/types"
	"github.com/Haoyunforever/Micro-Scanner/website/gateway/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserRegister Bind 绑定参数错误"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		log.LogrusObj.Errorf("UserRegister:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC 服务调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctl.RespError(ctx, err, "UserLogin Bind 绑定参数错误"))
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		log.LogrusObj.Errorf("UserLogin:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin RPC调用失败"))
		return
	}
	token, err := utils.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		log.LogrusObj.Errorf("UserLogin:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin 生成Token失败"))
		return
	}
	res := &types.TokenData{
		User:  userResp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
