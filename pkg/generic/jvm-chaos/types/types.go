package types

import (
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// ExperimentDetails holds the experiment-specific details
type ExperimentDetails struct {
	ExperimentName     string
	EngineName         string
	ChaosDuration      int
	ChaosInterval      int
	RampTime           int
	ChaosUID           clientTypes.UID
	InstanceID         string
	ChaosNamespace     string
	ChaosPodName       string
	Timeout            int
	Delay              int
	TargetContainer    string
	TargetPods         string
	PodsAffectedPerc   string
	Sequence           string
	NodeLabel          string
	
	// JVM-specific parameters
	JavaThreads        string // Number of threads for CPU stress
	MemoryPercentage   string // Memory percentage to consume
	MemoryMB           string // Memory in MB to consume
	JavaVersion        string // Target Java version (17+)
	
	// Helper pod configuration
	LIBImage           string
	LIBImagePullPolicy string
	ChaosServiceAccount string
	TerminationGracePeriodSeconds int
	
	// Runtime details
	ContainerRuntime   string
	SocketPath         string
}
