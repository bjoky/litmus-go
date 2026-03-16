package environment

import (
	"strconv"

	experimentTypes "github.com/litmuschaos/litmus-go/pkg/generic/jvm-chaos/types"
	"github.com/litmuschaos/litmus-go/pkg/types"
	clientTypes "k8s.io/apimachinery/pkg/types"
)

// GetENV fetches all the environment variables for the JVM chaos experiment
func GetENV(experimentDetails *experimentTypes.ExperimentDetails) {
	experimentDetails.ExperimentName = types.Getenv("EXPERIMENT_NAME", "jvm-cpu-hog")
	experimentDetails.ChaosNamespace = types.Getenv("CHAOS_NAMESPACE", "litmus")
	experimentDetails.EngineName = types.Getenv("CHAOSENGINE", "")
	experimentDetails.ChaosDuration, _ = strconv.Atoi(types.Getenv("TOTAL_CHAOS_DURATION", "60"))
	experimentDetails.ChaosInterval, _ = strconv.Atoi(types.Getenv("CHAOS_INTERVAL", "10"))
	experimentDetails.RampTime, _ = strconv.Atoi(types.Getenv("RAMP_TIME", "0"))
	experimentDetails.ChaosUID = clientTypes.UID(types.Getenv("CHAOS_UID", ""))
	experimentDetails.InstanceID = types.Getenv("INSTANCE_ID", "")
	experimentDetails.ChaosPodName = types.Getenv("POD_NAME", "")
	experimentDetails.Delay, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_DELAY", "2"))
	experimentDetails.Timeout, _ = strconv.Atoi(types.Getenv("STATUS_CHECK_TIMEOUT", "180"))
	experimentDetails.TargetContainer = types.Getenv("TARGET_CONTAINER", "")
	experimentDetails.TargetPods = types.Getenv("TARGET_PODS", "")
	experimentDetails.PodsAffectedPerc = types.Getenv("PODS_AFFECTED_PERC", "0")
	experimentDetails.Sequence = types.Getenv("SEQUENCE", "parallel")
	experimentDetails.NodeLabel = types.Getenv("NODE_LABEL", "")
	experimentDetails.ContainerRuntime = types.Getenv("CONTAINER_RUNTIME", "containerd")
	experimentDetails.SocketPath = types.Getenv("SOCKET_PATH", "/run/containerd/containerd.sock")
	experimentDetails.LIBImage = types.Getenv("LIB_IMAGE", "litmuschaos/litmus-go:latest")
	experimentDetails.LIBImagePullPolicy = types.Getenv("LIB_IMAGE_PULL_POLICY", "Always")
	experimentDetails.TerminationGracePeriodSeconds, _ = strconv.Atoi(types.Getenv("TERMINATION_GRACE_PERIOD_SECONDS", "30"))
	
	// JVM-specific environment variables
	experimentDetails.JavaThreads = types.Getenv("JAVA_THREADS", "1")
	experimentDetails.MemoryPercentage = types.Getenv("MEMORY_PERCENTAGE", "50")
	experimentDetails.MemoryMB = types.Getenv("MEMORY_MB", "")
	experimentDetails.JavaVersion = types.Getenv("JAVA_VERSION", "17")
}
