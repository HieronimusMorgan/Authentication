package handler

import (
	"authentication/internal/services"
	"authentication/internal/utils"
	"authentication/package/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleHandler struct {
	RoleService *services.RoleService
}

func NewRoleHandler(db *gorm.DB) *RoleHandler {
	return &RoleHandler{RoleService: services.NewRoleService(db)}
}

func extractAndValidateToken(context *gin.Context) (utils.TokenClaims, error) {
	token, err := utils.ExtractClaims(context.GetHeader("Authorization"))
	if err != nil {
		response.SendResponse(context, 401, "Unauthorized", nil, err.Error())
	}
	return *token, err
}

func (h RoleHandler) AddRole(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendResponse(ctx, 400, "Invalid request", nil, err.Error())
		return
	}

	token, err := extractAndValidateToken(ctx)
	if err != nil {
		return
	}

	role, err := h.RoleService.RegisterRole(&req, token.ClientID)
	if err != nil {
		response.SendResponse(ctx, 500, "Failed to register role", nil, err.Error())
		return
	}
	response.SendResponse(ctx, 201, "Role registered successfully", role, nil)
}

func (h RoleHandler) UpdateRole(ctx *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"optional"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendResponse(ctx, 400, "Invalid request", nil, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		response.SendResponse(ctx, 400, "Role ID is required", nil, nil)
		return
	}

	roleID, err := utils.ConvertToUint(id)
	if err != nil {
		response.SendResponse(ctx, 400, "Role ID must be a number", nil, err.Error())
		return
	}

	if req.Name == "" {
		response.SendResponse(ctx, 400, "Role name is required", nil, nil)
		return
	}

	token, err := extractAndValidateToken(ctx)
	if err != nil {
		return
	}

	role, err := h.RoleService.UpdateRole(roleID, &req, token.ClientID)
	if err != nil {
		response.SendResponse(ctx, 500, "Failed to update role", nil, err.Error())
		return
	}
	response.SendResponse(ctx, 200, "Role updated successfully", role, nil)
}

func (h RoleHandler) GetListRole(ctx *gin.Context) {
	token, err := extractAndValidateToken(ctx)
	if err != nil {
		return
	}

	roles, err := h.RoleService.GetListRole(token.ClientID)
	if err != nil {
		response.SendResponse(ctx, 500, "Failed to get list of roles", nil, err.Error())
		return
	}

	response.SendResponse(ctx, 200, "Roles retrieved successfully", roles, nil)
}

func (h RoleHandler) GetRoleById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.SendResponse(ctx, 400, "Role ID is required", nil, nil)
		return
	}

	roleID, err := utils.ConvertToUint(id)
	if err != nil {
		response.SendResponse(ctx, 400, "Role ID must be a number", nil, err.Error())
		return
	}

	token, err := extractAndValidateToken(ctx)
	if err != nil {
		return
	}

	role, err := h.RoleService.GetRoleById(roleID, token.ClientID)
	if err != nil {
		response.SendResponse(ctx, 500, "Failed to retrieve role", nil, err.Error())
		return
	}

	response.SendResponse(ctx, 200, "Role retrieved successfully", role, nil)
}

func (h RoleHandler) DeleteRoleById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.SendResponse(ctx, 400, "Role ID is required", nil, nil)
		return
	}

	roleID, err := utils.ConvertToUint(id)
	if err != nil {
		response.SendResponse(ctx, 400, "Role ID must be a number", nil, err.Error())
		return
	}

	token, err := extractAndValidateToken(ctx)
	if err != nil {
		return
	}

	err = h.RoleService.DeleteRole(roleID, token.ClientID)
	if err != nil {
		response.SendResponse(ctx, 500, "Failed to delete role", nil, err.Error())
		return
	}

	response.SendResponse(ctx, 200, "Role deleted successfully", nil, nil)
}

func (h RoleHandler) GetListRoleUsers(context *gin.Context) {
	token, err := extractAndValidateToken(context)
	if err != nil {
		return
	}

	users, err := h.RoleService.GetListRoleUsers(token.ClientID)
	if err != nil {
		response.SendResponse(context, 500, "Failed to get list of role users", nil, err.Error())
		return
	}

	response.SendResponse(context, 200, "Role users retrieved successfully", users, nil)
}
