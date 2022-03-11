package constants

// ClusterOperation type
type ClusterOperation string

const (
	// ClusterOperationCreate - OpenShift/k8s cluster create operation
	ClusterOperationCreate ClusterOperation = "create"

	// ClusterOperationDelete - OpenShift/k8s cluster delete operation
	ClusterOperationDelete ClusterOperation = "delete"

	// DefaultIngressDnsNamePrefix the default name prefix of the default ingress dns
	DefaultIngressDnsNamePrefix = "apps"
)

// TODO change these constants to match your own
const (
	// DinosaurOperatorAddonNamespace is the namespace where the Operator is deployed on via the AddOn flow in Production OSD clusters
	DinosaurOperatorAddonNamespace = "redhat-managed-dinosaur-operator"
	// DinosaurOperatorQEAddonNamespace is the namespace where the Operator is deployed on via the AddOn flow in Stage OSD clusters
	DinosaurOperatorQEAddonNamespace = "redhat-managed-dinosaur-operator-qe"
	// FleetshardAddonNamespace is the namespace where the fleetshard Operator is deployed on via the AddOn flow in Production OSD clusters
	FleetshardAddonNamespace = "redhat-fleetshard-operator"
	// FleetshardQEAddonNamespace is the namespace where the fleetshard Operator is deployed on via the AddOn flow in Stage OSD cluster
	FleetshardQEAddonNamespace = "redhat-fleetshard-operator-qe"
	// OpenIDIdentityProviderName is the name of the Identity Provider configured for each OSD cluster
	OpenIDIdentityProviderName = "Dinosaur_SRE"
	// ReadOnlyGroupName is the K8s group name to contain users intended to have the read only access
	ReadOnlyGroupName = "dinosaur-readonly-access"
	// SREGroupName is the K8s group name to contain users intended to have the cluster admin role
	SREGroupName = "dinosaur-sre"
	// ReadOnlyRoleBindingName is the read only binding name
	ReadOnlyRoleBindingName = "dinosaur-dedicated-readers"
	// SRERoleBindingName is the role binding name for sre cluster admin
	SRERoleBindingName = "dinosaur-sre-cluster-admin"
	// The DNS prefixes used for traffic ingress
	ManagedDinosaurIngressDnsNamePrefix = "kas"
)

func (c ClusterOperation) String() string {
	return string(c)
}
