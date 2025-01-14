package main

type Application struct {
	Name      string            `json:"name"`
	Metadata  Metadata          `json:"metadata,omitempty"`
	Env       map[string]string `json:"env,omitempty"`
	Services  Services          `json:"services,omitempty"`
	Processes Processes         `json:"processes,omitempty"`
	Replicas  uint              `json:"replicas,omitempty"`
}

type Metadata struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Routes []Route

type Route struct {
	URL      string        `json:"fqdn,omitempty"`
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

// Services represents a slice of Service
type Services []Service

// Service contains the specification for an existing Cloud Foundry service required by the application
type Service struct {
	// Name represents the name of the Cloud Foundry service required by the application
	Name string `json:"name"`
	// Parameters contain the k/v relationship for the aplication to connect to the service
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// Tag
}

// Processes represents a slice of Processes
type Processes []Process

// Process contains the abstraction of the specification of a Cloud Foundry Process
type Process struct {
	// Name represents the name of the process
	Name string `json:"name"`
	// Image represents the pull spec of the container image
	Image string `json:"image"`
	// Amount of memory requested by the process
	Memory string `json:"memory,omitempty"`
	// Amount of persistent disk requested by the process
	Disk string `json:"disk,omitempty"`
	// LivenessProbeInitialDelay represents the numbef of seconds before start checking the first health of the process.
	LivenessProbeInitialDelay uint `json:"livenessProbeInitialDelay,omitempty"`
	// LivenessProbeRetries represents the number of retries to attempt before failing the liveness probe of the process
	LivenessProbeRetries uint `json:"livenessProbeRetries,omitempty"`
	// LivenessProbe represents the endpoint location where to perform the liveness probe check
	LivenessProbe string `json:"livenessProbe,omitempty"`
	// LivenessProbeTimeout represents the number of seconds in which the liveness probe can be considered as timedout
	LivenessProbeTimeout uint `json:"livenessProbeTimeout,omitempty"`
	// Command represents the command used to run the process
	Command []string `json:"command,omitempty"`
	// Replicas represents the number of instances for this process to run
	Replicas uint `json:"replicas"`
	// Env define the list of k/v values to inject to the running container
	Env map[string]interface{} `json:"env,omitempty"`
	// Routes represent the routes that are made available by the process's open port.
	Routes Routes `json:"routes,omitempty"`
}
