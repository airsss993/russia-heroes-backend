package service

type Services struct {
	AdminService AdminService
}

func NewServices(adminService *AdminService) *Services {
	return &Services{
		AdminService: *adminService,
	}
}
