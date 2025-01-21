package main

// Application represents an interpretation of a runtime Cloud Foundry application. This structure differs in that
// the information it contains has been processed to simplify its transformation to a Kubernetes manifest using MTA
type Application struct {
	// Metadata captures the name, labels and annotations in the application.
	Metadata Metadata `json:",inline"`
	// Env captures the `env` field values in the CF application manifest.
	Env map[string]string `json:"env,omitempty"`
	// Routes represent the routes that are made available by the application.
	Routes Routes `json:"routes,omitempty"`
	// Services captures the `services` field values in the CF application manifest.
	Services Services `json:"services,omitempty"`
	// Processes captures the `processes` field values in the CF application manifest.
	Processes Processes `json:"processes,omitempty"`
	// Sidecars captures the `sidecars` field values in the CF application manifest.
	Sidecars Sidecars `json:"sidecars,omitempty"`
	// Stack represents the `stack` field in the application manifest. The value is captured for information
	// purposes because it has no relevance in Kubernetes.
	Stack string `json:"stack,omitempty"`
	// StartupTimeout captures the maximum elapsed time in which an application that is starting is considered to have failed to respond to checks.
	// An application has to respond to a readiness or health check before the timeout time elapses or else the platform will
	// fail the deployment of the application. By default its 60 seconds.
	// https://github.com/cloudfoundry/docs-dev-guide/blob/96f19d9d67f52ac7418c147d5ddaa79c957eec34/deploy-apps/large-app-deploy.html.md.erb#L35
	StartupTimeout uint `json:"startupTimeout,omitempty"`
	// Replicas configures the number of Cloud Foundry application instances.
	Replicas uint `json:"replicas"`
}

// Metadata captures the name, labels and annotations in the application
type Metadata struct {
	// Name capture the `name` field int CF application manifest
	Name string `json:"name"`
	// Space captures the `space` where the CF application is deployed at runtime. The field is empty if the
	// application is discovered directly from the CF manifest. It is equivalent to a Namespace in Kubernetes.
	Space string `json:"space,omitempty"`
	// Labels capture the labels as defined in the `annotations` field in the CF application manifest
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations capture the annotations as defined in the `labels` field in the CF application manifest
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Routes represents a slice of Routes
type Routes []Route

// Route captures the key elements that define a Route in a string that maps to a URL structure. These values
// are captured as runtime routes, meaning that if the CF Application manifest is configured to disable all routes
// with the `no-route` value, it will translate into an empty slice.
// By default CloudFoundry will always attempt to create a route for each application, unless specified by the field `no-route` when true
// For further details check: https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#no-route
// and https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#random-route
// Example
// ---
//
//	...
//	routes:
//	- route: example.com
//	  protocol: http2
//	- route: www.example.com/foo
//	- route: tcp-example.com:1234
type Route struct {
	// URL captures the Fully Qualified Domain Name of the hostname field in the route. If the hostname contained a port
	// its value it captured in the `Port` field in the Route structure.
	URL string `json:"url"`
	// Protocol captures the protocol type: http, http2 or tcp. Note that the CF `protocol` field is only available
	// for CF deployments that use HTTP/2 routing.
	Protocol RouteProtocol `json:"protocol"`
}

type RouteProtocol string

const (
	HTTP  RouteProtocol = "http"
	HTTPS RouteProtocol = "https"
	TCP   RouteProtocol = "tcp"
)

// Services represents a slice of Service
type Services []Service

// Service contains the specification for an existing Cloud Foundry service required by the application.
// Examples:
// ---
//
//	...
//	services:
//	  - service-1
//	  - name: service-2
//	  - name: service-3
//	    parameters:
//	      key-1: value-1
//	      key-2: [value-2, value-3]
//	      key-3: ... any other kind of value ...
//	  - name: service-4
//	    binding_name: binding-1
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
	// Image represents the pull spec of the container image.
	Image string `json:"image"`
	// Command represents the command used to run the process.
	Command []string `json:"command,omitempty"`
	// DiskQuota represents the amount of persistent disk requested by the process.
	DiskQuota string `json:"disk,omitempty"`
	// Memory represents the amount of memory requested by the process.
	Memory string `json:"memory,omitempty"`
	// HealthCheck captures the health check information
	HealthCheck Probe `json:"healthCheck"`
	// ReadinessCheck captures the readiness check information.
	ReadinessCheck Probe `json:"readinessCheck"`
	// Replicas represents the number of instances for this process to run.
	Replicas uint `json:"replicas"`
	// LogRateLimit represents the maximum amount of logs to be captured per second.
	LogRateLimit string `json:"logRateLimit,omitempty"`
}

type Sidecars []Sidecar

// Sidecar captures the information of a Sidecar process
// https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#sidecars
type Sidecar struct {
	// Name represents the name of the Sidecar
	Name string `json:"name"`
	// ProcessTypes captures the different process types defined for the sidecar.
	// Compared to a Process, which has only one type, sidecar processes can accumulate more than one type.
	ProcessTypes ProcessTypes `json:"processTypes"`
	// Command captures the command to use to run the sidecar
	Command []string `json:"command"`
	// Memory represents the amount of memory to allocate to the sidecar. It's an optional field.
	Memory string `json:"memory,omitempty"`
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

type ProcessTypes []ProcessType

type ProcessType string

const (
	// Web represents a `web` application type
	Web ProcessType = "web"
	// Worker represents a `worker` application type
	Worker ProcessType = "worker"
)
