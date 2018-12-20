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

func (svc *Service) initService(appid,proc,image string,versions []string,rules map[string]int, envs []deployment.Env, gw_port,port int) *Service {

	return &Service{
		Gateway:               svc.Gateway.GetGateway(appid,gw_port,proc),
		Destinationrule:       svc.Destinationrule.GetDestinationRule(appid,versions),
		Virtulservice:         svc.Virtulservice.GetVs(appid,rules,port),

	}

}


//新版本version,各版本及对应值流量权重组成的rules
func (gp *Graypolicy) initgray(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port int) *Graypolicy {
	var b_version []string
	for i,_ := range rules{
		b_version = append(b_version,i)
	}

	return &Graypolicy{
		Deployment:            *gp.Deployment.GetDeploy(appid,image,version,envs,port,0),
		Destinationrule:       *gp.Destinationrule.GetDestinationRule(appid,b_version),
		Virtualservice:        *gp.Virtualservice.GetVs(appid,rules,port),

	}

}


func (gp *Graypolicy) updatepolicy(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port,replices int) *Graypolicy {

	return &Graypolicy{
		Deployment:            *gp.Deployment.GetDeploy(appid,image,version,envs,port,replices),
		Virtualservice:        *gp.Virtualservice.GetVs(appid,rules,port),

	}

}
