package gateway



type GateWay struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json: "kind"`
	Metadata   Metadata `json: "metadata"`
	Spec       Spec     `json: "spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	Selector Selector `json: "selector"`
	Servers  []Server `json: "servers"`
}

type Selector struct {
	Istio string `json:"istio"`
}

type Server struct {
	Port  Port     `json: "port"`
	Hosts []string `json: "hosts"`
}

type Port struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

func (vm *GateWay) GetGateway(appid string, number int,pro string) *GateWay{
	g_name := appid + "-gateway"

	return &GateWay{
		ApiVersion :     "networking.istio.io/v1alpha3",
		Kind:            "Gateway",
		Metadata:   	  Metadata{g_name},
		Spec:             *getspec(pro,number),
	}

}

func convertPort(pro string, port int) *Port {
	return &Port{
		Name:       pro,
		Number:     int(port),
		Protocol:   pro,
	}
}

func getspec(pro string, number int) *Spec{

	return &Spec{
		Selector:     Selector{"ingressgateway"},
		Servers:      []Server{*getserver(pro,number)},
	}

}

func getserver(pro string,number int)  *Server{
	g_host := []string{"*"}
	return &Server{
		Port:  *convertPort(pro,number),
		Hosts: g_host,
	}

}


