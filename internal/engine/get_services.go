package engine

import (
	"github.com/r7wx/easy-gate/internal/group"
	"github.com/r7wx/easy-gate/internal/routine"
	"github.com/r7wx/easy-gate/internal/service"
)

func getServices(status *routine.Status, addr string) []service.Service {
	services := []service.Service{}
	for _, statusService := range status.Services {
		if group.IsAllowed(status.Groups, statusService.Groups, addr) {
			service := service.Service{
				Icon:     statusService.Icon,
				Name:     statusService.Name,
				URL:      statusService.URL,
				Category: statusService.Category,
			}
			services = append(services, service)
		}
	}

	return services
}
