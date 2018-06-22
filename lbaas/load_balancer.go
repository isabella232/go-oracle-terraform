package lbaas

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-oracle-terraform/client"
)

const waitForLoadBalancerReadyPollInterval = 30 * time.Second  // 30 seconds
const waitForLoadBalancerReadyTimeout = 30 * time.Minute       // 30 minutes
const waitForLoadBalancerDeletePollInterval = 30 * time.Second // 30 seconds
const waitForLoadBalancerDeleteTimeout = 30 * time.Minute      // 30 minutes

// LoadBalancerScheme Scheme types
type LoadBalancerScheme string

const (
	LoadBalancerSchemeInternetFacing LoadBalancerScheme = "INTERNET_FACING"
	LoadBalancerSchemeInternal       LoadBalancerScheme = "INTERNAL"
)

// LoadBalancerDisabled
type LoadBalancerDisabled string

const (
	LoadBalancerDisabledTrue        LoadBalancerDisabled = "TRUE"
	LoadBalancerDisabledFalse       LoadBalancerDisabled = "FALSE"
	LoadBalancerDisabledMaintenance LoadBalancerDisabled = "MAINTENANCE_MODE"
)

// LoadBalancerEffectiveState
type LoadBalancerEffectiveState string

const (
	LoadBalancerEffectiveStateTrue        LoadBalancerEffectiveState = "TRUE"
	LoadBalancerEffectiveStateFalse       LoadBalancerEffectiveState = "FALSE"
	LoadBalancerEffectiveStateMaintenance LoadBalancerEffectiveState = "MAINTENANCE_MODE"
)

// LoadBalancerState
type LoadBalancerState string

const (
	LoadBalancerStateCreationInProgress     LoadBalancerState = "CREATION_IN_PROGRESS"
	LoadBalancerStateCreated                LoadBalancerState = "CREATED"
	LoadBalancerStateHealthy                LoadBalancerState = "HEALTHY"
	LoadBalancerStateInterventionNeeded     LoadBalancerState = "ADMINISTRATOR_INTERVENTION_NEEDED"
	LoadBalancerStateDeletionInProgress     LoadBalancerState = "DELETION_IN_PROGRESS"
	LoadBalancerStateDeleted                LoadBalancerState = "DELETED"
	LoadBalancerStateModificationInProgress LoadBalancerState = "MODIFICATION_IN_PROGRESS"
	LoadBalancerStateCreationFailed         LoadBalancerState = "CREATION_FAILED"
	LoadBalancerStateModificationFailed     LoadBalancerState = "MODIFICATION_FAILED"
	LoadBalancerStateDeletionFailed         LoadBalancerState = "DELETION_FAILED"
	LoadBalancerStateAccessDenied           LoadBalancerState = "ACCESS_DENIED"
	LoadBalancerStateAbandon                LoadBalancerState = "ABANDON"
	LoadBalancerStatePause                  LoadBalancerState = "PAUSE"
	LoadBalancerStateForcePaused            LoadBalancerState = "FORCE_PAUSED"
	LoadBalancerStateResume                 LoadBalancerState = "RESUME"
)

// HttpMethods
type HttpMethod string

const (
	HttpCOPY      HttpMethod = "COPY"
	HttpDELETE    HttpMethod = "DELETE"
	HttpGET       HttpMethod = "GET"
	HttpHEAD      HttpMethod = "HEAD"
	HttpLOCK      HttpMethod = "LOCK"
	HttpMKCOL     HttpMethod = "MKCOL"
	HttpMOVE      HttpMethod = "MOVE"
	HttpOPTIONS   HttpMethod = "OPTIONS"
	HttpPATCH     HttpMethod = "PATCH"
	HttpPOST      HttpMethod = "POST"
	HttpPROPFIND  HttpMethod = "PROPFIND"
	HttpPROPPATCH HttpMethod = "PROPPATCH"
	HttpPUT       HttpMethod = "PUT"
	HttpUNLOCK    HttpMethod = "UNLOCK"
)

// LoadBalancerInfo specifies the Load Balancer obtained from a GET request
type LoadBalancerInfo struct {
	BalancerVIPs             []string                       `json:"balancer_vips"`
	CanonicalHostName        string                         `json:"canonical_host_name"`
	CloudgateCapable         string                         `json:"cloudgate_capable"`
	ComputeSecurityArtifacts []ComputeSecurityArtifactsInfo `json:"compute_security_artifacts"`
	ComputeSite              string                         `json:"compute_site"`
	CreatedOn                string                         `json:"created_on"`
	Description              string                         `json:"description"`
	Disabled                 LoadBalancerDisabled           `json:"disabled"`
	DisplayName              string                         `json:"display_name"`
	HealthCheck              HealthCheckInfo                `json:"health_check"`
	IPNetworkName            string                         `json:"ip_network_name"`
	IsDisabledEffectively    string                         `json:"is_disabled_effectively"`
	Listeners                []ListenerInfo                 `json:"listeners"`
	ModifiedOn               string                         `json:"modified_on"`
	Name                     string                         `json:"name"`
	Owner                    string                         `json:"owner"`
	PermittedMethods         []string                       `json:"permitted_methods"`
	Region                   string                         `json:"region"`
	RestURIs                 []RestURIInfo                  `json:"rest_uri"`
	Scheme                   LoadBalancerScheme             `json:"scheme"`
	ServerPool               string                         `json:"server_pool"`
	State                    LoadBalancerState              `json:"state"`
	Tags                     []string                       `json:"tags"`
	URI                      string                         `json:"uri"`
}

type ComputeSecurityArtifactsInfo struct {
	AddressType  string `json:"address_type"`
	ArtifactType string `json:"artifact_type"`
	URI          string `json:"uri"`
}

type HealthCheckInfo struct {
	AcceptedReturnCodes []string `json:accepted_return_codes`
	Enabled             string   `json:"enabled"`
	HealthyThreshold    int      `json:"healthy_threshold"`
	Interval            int      `json:"interval"`
	Path                string   `json:"path"`
	Timeout             int      `json:"timeout"`
	Type                string   `json:"type"`
	UnhealthyThreshold  int      `json:"unhealthy_threshold"`
}

type RestURIInfo struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

// CreateLoadBalancerInput specifies the create request for a load balancer service instance
type CreateLoadBalancerInput struct {
	Description        string               `json:"description,omitempty"`
	Disabled           LoadBalancerDisabled `json:"disabled"`
	IPNetworkName      string               `json:"ip_network_name,omitempty"`
	Name               string               `json:"name"`
	ParentLoadBalancer string               `json:"parent_vlbr,omitempty"`
	PermittedClients   []string             `json:"permitted_clients,omitempty"`
	PermittedMethods   []string             `json:"permitted_methods,omitempty"`
	Policies           []string             `json:"policies,omitempty"`
	Region             string               `json:"region"`
	Scheme             LoadBalancerScheme   `json:"scheme"`
	ServerPool         string               `json:"server_pool,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
}

// UpdateLoadBalancerInput specifies the create request for a load balancer service instance
type UpdateLoadBalancerInput struct {
	Description        string               `json:"description,omitempty"`
	Disabled           LoadBalancerDisabled `json:"disabled,omitempty"`
	IPNetworkName      string               `json:"ip_network_name,omitempty"`
	Name               string               `json:"name,omitempty"`
	ParentLoadBalancer string               `json:"parent_vlbr,omitempty"`
	PermittedClients   []string             `json:"permitted_clients,omitempty"`
	PermittedMethods   []string             `json:"permitted_methods,omitempty"`
	Policies           []string             `json:"policies,omitempty"`
	ServerPool         string               `json:"server_pool,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
}

// LoadBalancerContext represents a specific loadbalancer instnace by region/name context
type LoadBalancerContext struct {
	Region string
	Name   string
}

// CreateLoadBalancer creates a new Load Balancer instance
func (c *LoadBalancerClient) CreateLoadBalancer(input *CreateLoadBalancerInput) (*LoadBalancerInfo, error) {

	if c.PollInterval == 0 {
		c.PollInterval = waitForLoadBalancerReadyPollInterval
	}
	if c.Timeout == 0 {
		c.Timeout = waitForLoadBalancerReadyTimeout
	}

	var info LoadBalancerInfo
	if err := c.createResource(&input, &info); err != nil {
		return nil, err
	}

	createdStates := []LBaaSState{LBaaSStateCreationInProgress, LBaaSStateCreated, LBaaSStateHealthy}
	erroredStates := []LBaaSState{LBaaSStateCreationFailed, LBaaSStateDeletionInProgress, LBaaSStateDeleted, LBaaSStateDeletionFailed, LBaaSStateAccessDenied, LBaaSStateAdministratorInterventionNeeded}

	// check the initial response
	ready, err := c.checkLoadBalancerState(&info, createdStates, erroredStates)
	if err != nil {
		return nil, err
	}
	if ready {
		return &info, nil
	}
	// else poll till ready
	lb := LoadBalancerContext{
		Region: input.Region,
		Name:   input.Name,
	}
	err = c.WaitForLoadBalancerState(lb, createdStates, erroredStates, c.PollInterval, c.Timeout, &info)
	return &info, err
}

// DeleteLoadBalancer deletes the service instance with the specified input
func (c *LoadBalancerClient) DeleteLoadBalancer(lb LoadBalancerContext) (*LoadBalancerInfo, error) {

	if c.PollInterval == 0 {
		c.PollInterval = waitForLoadBalancerDeletePollInterval
	}
	if c.Timeout == 0 {
		c.Timeout = waitForLoadBalancerDeleteTimeout
	}

	var info LoadBalancerInfo
	if err := c.deleteResource(lb.Region, lb.Name, &info); err != nil {
		return nil, err
	}

	// deletedStates := []LBaaSState{LBaaSStateDeletionInProgress, LBaaSStateDeleted}
	deletedStates := []LBaaSState{LBaaSStateDeleted}
	erroredStates := []LBaaSState{LBaaSStateDeletionFailed, LBaaSStateAccessDenied, LBaaSStateAdministratorInterventionNeeded}

	// check the initial response
	deleted, err := c.checkLoadBalancerState(&info, deletedStates, erroredStates)
	if err != nil {
		return nil, err
	}
	if deleted {
		return &info, nil
	}
	// else poll till deleted
	err = c.WaitForLoadBalancerState(lb, deletedStates, erroredStates, c.PollInterval, c.Timeout, &info)
	if err != nil && client.WasNotFoundError(err) {
		// resource could not be found, thus deleted
		return nil, nil
	}
	return &info, err
}

// GetLoadBalancer fetchs the instance details of the Load Balancer
func (c *LoadBalancerClient) GetLoadBalancer(lb LoadBalancerContext) (*LoadBalancerInfo, error) {
	var info LoadBalancerInfo
	if err := c.getResource(lb.Region, lb.Name, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// UpdateLoadBalancer fetchs the instance details of the Load Balancer
func (c *LoadBalancerClient) UpdateLoadBalancer(lb LoadBalancerContext, input *UpdateLoadBalancerInput) (*LoadBalancerInfo, error) {

	if c.PollInterval == 0 {
		c.PollInterval = waitForLoadBalancerReadyPollInterval
	}
	if c.Timeout == 0 {
		c.Timeout = waitForLoadBalancerReadyTimeout
	}

	var info LoadBalancerInfo
	if err := c.updateResource(lb.Region, lb.Name, &input, &info); err != nil {
		return nil, err
	}

	updatedStates := []LBaaSState{LBaaSStateModificationInProgress, LBaaSStateHealthy}
	erroredStates := []LBaaSState{LBaaSStateModificaitonFailed, LBaaSStateAccessDenied, LBaaSStateAdministratorInterventionNeeded}

	// check the initial response
	ready, err := c.checkLoadBalancerState(&info, updatedStates, erroredStates)
	if err != nil {
		return nil, err
	}
	if ready {
		return &info, nil
	}
	// else poll till ready

	err = c.WaitForLoadBalancerState(lb, updatedStates, erroredStates, c.PollInterval, c.Timeout, &info)
	return &info, err
}

// WaitForLoadBalancerState waits for the resource to be in one of a set of desired states
func (c *LoadBalancerClient) WaitForLoadBalancerState(lb LoadBalancerContext, desiredStates, errorStates []LBaaSState, pollInterval, timeoutSeconds time.Duration, info *LoadBalancerInfo) error {

	var getErr error
	err := c.client.WaitFor("Load Balancer to be ready", pollInterval, timeoutSeconds, func() (bool, error) {
		info, getErr = c.GetLoadBalancer(lb)
		if getErr != nil {
			return false, getErr
		}

		return c.checkLoadBalancerState(info, desiredStates, errorStates)
	})
	return err
}

// check the State, returns in desired state (true), not ready yet (false) or errored state (error)
func (c *LoadBalancerClient) checkLoadBalancerState(info *LoadBalancerInfo, desiredStates, errorStates []LBaaSState) (bool, error) {

	c.client.DebugLogString(fmt.Sprintf("Load Balancer %v state is %v", info.Name, info.State))

	state := LBaaSState(info.State)

	if isStateInLBaaSStates(state, desiredStates) {
		// we're good, return okay
		return true, nil
	}
	if isStateInLBaaSStates(state, errorStates) {
		// not good, return error
		return false, fmt.Errorf("Load Balancer %v in errored state %v", info.Name, info.State)
	}
	// not ready lifecycleTimeout
	return false, nil
}
