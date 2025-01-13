package main

type Application struct {
	Name     string            `json:"name"`
	Docker   DockerSpec        `json:"dockerSpec,omitempty"`
	Env      map[string]string `json:"env,omitempty"`
	Route    RouteSpec         `json:"routeSpec,omitempty"`
	Service  ServiceSpec       `json:"serviceSpec,omitempty"`
	Process  Processes         `json:"process,omitempty"`
	Sidecar  Sidecars          `json:"sidecar,omitempty"`
	Metadata Metadata          `json:"metadata,omitempty"`
}

type DockerSpec struct {
	Image        string `json:"image"`
	AuthRequired bool   `json:"authRequired,omitempty"`
}

type Metadata struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Routes []Route

type Route struct {
	FQDN     string        `json:"fqdn,omitempty"`
	Protocol RouteProtocol `json:"protocol,omitempty"`
}

type RouteSpec struct {
	Type   RouteType `json:"type"`
	Routes Routes    `json:"routes"`
}

type RouteType string

const (
	Default RouteType = "default"
	Random  RouteType = "random"
)

type RouteProtocol string

const (
	HTTP  RouteProtocol = "http"
	HTTPS RouteProtocol = "https"
	TCP   RouteProtocol = "tcp"
)

type ServiceSpec struct {
}

