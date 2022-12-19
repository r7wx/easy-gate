/*
MIT License

Copyright (c) 2022 r7wx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package routine

import (
	"github.com/r7wx/easy-gate/internal/config"
	"github.com/r7wx/easy-gate/internal/service"
)

func (r *Routine) getServices(cfg *config.Config) []service.Service {
	type serviceWrapper struct {
		service service.Service
		index   int
	}

	servicePChan := make(chan serviceWrapper)
	for index, cfgService := range cfg.Services {
		go func(index int, cfgService config.Service) {
			servicePChan <- serviceWrapper{
				service: service.Service{
					Icon:     r.getIconData(cfgService),
					Name:     cfgService.Name,
					URL:      cfgService.URL,
					Category: cfgService.Category,
					Groups:   cfgService.Groups,
				},
				index: index,
			}
		}(index, cfgService)
	}

	processedServices := map[int]service.Service{}
	for i := 1; i <= len(cfg.Services); i++ {
		processedService := <-servicePChan
		processedServices[processedService.index] = processedService.service
	}

	services := []service.Service{}
	for index := range cfg.Services {
		services = append(services,
			processedServices[index])
	}

	return services
}
