package grayservice

import (
	"deploy_app/gateway"
	"deploy_app/destination"
	"deploy_app/virtualservice"
	"deploy_app/deployment"
)



type Service struct {
	Gateway               *gateway.GateWay                        `json:"gateway"`
	Destinationrule       *destination.DestinationRule            `json:"destination"`
	Virtulservice        *virtualservice.VirtualService          `json:"virtrulservice"`
}

type Graypolicy struct {
	Deployment            deployment.Deployment                  `json:"deployment"`
	Destinationrule       destination.DestinationRule            `json:"destination"`
	Virtualservice        virtualservice.VirtualService          `json:"virtrulservice"`

}

func (svc *Service) initService(appid,proc,image string,versions []string,rules map[string]int, envs []deployment.Env, port int) *Service {

	return &Service{
		Gateway:               svc.Gateway.GetGateway(appid,port,proc),
		Destinationrule:       svc.Destinationrule.GetDestinationRule(appid,versions),
		Virtulservice:        svc.Virtulservice.GetVs(appid,rules,port),

	}

}


func (gp *Graypolicy) getgraypolicy(appid,proc,image,version string,versions []string,rules map[string]int, envs []deployment.Env, port int) *Graypolicy {

	return &Graypolicy{
		Deployment:            *gp.Deployment.GetDeploy(appid,image,version,envs,port),
		Destinationrule:       *gp.Destinationrule.GetDestinationRule(appid,versions),
		Virtualservice:        *gp.Virtualservice.GetVs(appid,rules,port),

	}

}
