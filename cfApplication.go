package main

// Application represents an interpretation of a runtime Cloud Foundry application. This structure differs in that
// the information it contains has been processed to simplify its transformation to a Kubernetes manifest using MTA
type Application struct {
	Metadata Metadata `json:",inline"`
	// Env captures the `env` field values in the CF application manifest.
	Env map[string]string `json:"env,omitempty"`
	// Services captures the `services` field values in the CF application manifest.
	Services Services `json:"services,omitempty"`
	// Processes captures the `processes` field values in the CF application manifest.
	Processes Processes `json:"processes,omitempty"`
	// Sidecars captures the `sidecars` field values in the CF application manifest.
	Sidecars Processes `json:"sidecars,omitempty"`
	// Instances configures the number of Cloud Foundry application instances.
	Instances uint `json:"instances"`
	// Stack represents the `stack` field in the application manifest. The value is captured for information
	// purposes because it has no relevance in Kubernetes.
	Stack string `json:"stack,omitempty"`
}

// Metadata captures the name, labels and annotations in the application
type Metadata struct {
	// Name capture the `name` field int CF application manifest
	Name string `json:"name"`
	// Labels capture the labels as defined in the `annotations` field in the CF application manifest
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations capture the annotations as defined in the `labels` field in the CF application manifest
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Routes represents a slice of Routes
type Routes []Route

// Route captures the key elements that define a Route: hostname, protocol and port. These values
// are captured as runtime routes, meaning that if the CF Application manifest is configured to disable all routes
// with the `no-route` value, it will translate into an empty slice.
// Unless specified by the process or setting the application field `no-route` to true,
// by default CloudFoundry will always attempt to create a route for each application.
// For further details check: https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#no-route
// and https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#random-route
type Route struct {
	// Hostname contains the hostname that will be used for the route.
	Hostname string `json:"hostname"`
	// Protocol captures the protocol type: http, http2 or tcp.
	Protocol RouteProtocol `json:"protocol"`
	// Port captures the port to use for the route. For RouteProtocol `http`` it is 80; for `http2` it's 443,
	// and for `tcp` it is as defined in the CF application manifest.
	Port uint `json:"port"`
}

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
	// Name represents the name of the Cloud Foundry service required by the application. This field
	// represents the runtime name of the service, captured from the 3 different cases where
	// the service name can be listed.
	// For more information check https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#services-block
	Name string `json:"name"`
	// Parameters contain the k/v relationship for the aplication to bind to the service
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// Processes represents a slice of Processes.
type Processes []Process

// Process represents the abstraction of the specification of a Cloud Foundry Process.
// For more information check https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#processes
type Process struct {
	// Type captures the `type` field in the Process specification. Accepted values are `web` or `worker`
	Type ProcessType `json:"type,omitempty"`
	// Name represents the name of the process.
	Name string `json:"name"`
	// Image represents the pull spec of the container image.
	Image string `json:"image"`
	// Memory represents the amount of memory requested by the process.
	Memory string `json:"memory,omitempty"`
	// DiskQuota represents the amount of persistent disk requested by the process.
	DiskQuota string `json:"disk,omitempty"`
	// HealthCheck captures the health check information
	HealthCheck Probe `json:"healthCheck"`
	// ReadinessCheck captures the readiness check information.
	ReadinessCheck Probe `json:"readinessCheck"`
	// Command represents the command used to run the process.
	Command []string `json:"command,omitempty"`
	// Replicas represents the number of instances for this process to run.
	Replicas uint `json:"replicas"`
	// Env define the list of k/v values to inject to the running container.
	Env map[string]interface{} `json:"env,omitempty"`
	// Routes represent the routes that are made available by the process's open port.
	Routes Routes `json:"routes,omitempty"`
	// LogRateLimit represents the maximum amount of logs to be captured per second.
	LogRateLimit string `json:"logRateLimit,omitempty"`
}

// Probe captures the fields for managing health checks. For more information check https://docs.cloudfoundry.org/devguide/deploy-apps/healthchecks.html
type Probe struct {
	// Endpoint represents the URL location where to perform the probe check.
	Endpoint string `json:"endpoint"`
	// Timeout represents the number of seconds in which the probe check can be considered as timedout.
	// https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#timeout
	Timeout uint `json:"timeout"`
	// Interval represents the number of seconds between probe checks.
	Interval uint `json:"interval"`
}

// ProcessType captures the CF process types as defined by the CF v3 Application manifest
type ProcessType string

const (
	// Web represents a `web` application type
	Web ProcessType = "web"
	// Worker represents a `worker` application type
	Worker ProcessType = "worker"
)
