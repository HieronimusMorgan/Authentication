package services

import (
	"authentication/internal/dto/out"
	"authentication/internal/models"
	"authentication/internal/repository"
	"gorm.io/gorm"
)

type RoleService struct {
	RoleRepository *repository.RoleRepository
	UserRepository *repository.UserRepository
}

func NewRoleService(db *gorm.DB) *RoleService {
	roleRepo := repository.NewRoleRepository(db)
	userRepo := repository.NewUserRepository(db)
	return &RoleService{RoleRepository: roleRepo, UserRepository: userRepo}
}

func (s RoleService) RegisterRole(req *struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}, clientID string) (interface{}, error) {
	user, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}

	var role = &models.Role{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   user.FullName,
		UpdatedBy:   user.FullName,
	}
	err = s.RoleRepository.RegisterRole(&role)
	if err != nil {
		return nil, err
	}

	return out.RoleResponse{
		RoleID:      role.RoleID,
		Name:        role.Name,
		Description: role.Description,
	}, nil
}

func (s RoleService) UpdateRole(roleID uint, req *struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}, clientID string) (interface{}, error) {
	user, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}

	role, err := s.RoleRepository.GetRoleByID(roleID)
	if err != nil {
		return nil, err
	}

	role.Name = req.Name
	role.Description = req.Description
	role.UpdatedBy = user.FullName
	err = s.RoleRepository.UpdateRole(&role)

	if err != nil {
		return nil, err
	}

	return out.RoleResponse{
		RoleID:      role.RoleID,
		Name:        role.Name,
		Description: role.Description,
	}, nil
}

func (s RoleService) GetListRole(clientID string) (interface{}, error) {
	_, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}

	roles, err := s.RoleRepository.GetAllRoles()
	if err != nil {
		return nil, err
	}

	var roleResponses []out.RoleResponse
	for _, role := range *roles {
		roleResponses = append(roleResponses, out.RoleResponse{
			RoleID:      role.RoleID,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return roleResponses, nil
}

func (s RoleService) GetRoleById(roleID uint, clientID string) (interface{}, error) {
	_, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return nil, err
	}

	role, err := s.RoleRepository.GetRoleByID(roleID)
	if err != nil {
		return nil, err
	}

	return out.RoleResponse{
		RoleID:      role.RoleID,
		Name:        role.Name,
		Description: role.Description,
	}, nil
}

func (s RoleService) DeleteRole(roleID uint, clientID string) error {
	user, err := s.UserRepository.GetUserByClientID(clientID)
	if err != nil {
		return err
	}

	role, err := s.RoleRepository.GetRoleByID(roleID)
	if err != nil {
		return err
	}

	role.DeletedBy = user.FullName
	err = s.RoleRepository.DeleteRole(&role)
	if err != nil {
		return err
	}

	return nil
}