# JVM Fault Injection - Proof of Concept

## Overview

This is a minimal proof-of-concept implementation for JVM fault injection in litmus-go. The structure follows the standard Litmus experiment pattern but contains placeholders for the actual chaos injection logic.

## Created Structure

```
litmus-go/
├── experiments/generic/jvm-cpu-hog/
│   ├── experiment/
│   │   └── jvm-cpu-hog.go          # Main experiment logic
│   ├── test/
│   │   └── test.yml                # Test deployment manifest
│   ├── rbac.yaml                   # RBAC configuration
│   └── README.md                   # Experiment documentation
│
├── chaoslib/litmus/jvm-chaos/
│   ├── lib/
│   │   └── jvm-chaos.go            # Chaos library (orchestration)
│   └── helper/
│       └── jvm-helper.go           # Helper pod logic (runs in helper pod)
│
└── pkg/generic/jvm-chaos/
    ├── types/
    │   └── types.go                # Experiment-specific types
    └── environment/
        └── environment.go          # Environment variable parsing

```

## What's Implemented

### 1. Experiment Structure (`experiments/generic/jvm-cpu-hog/`)
- **jvm-cpu-hog.go**: Main experiment entry point with standard Litmus flow:
  - Pre-chaos checks
  - Chaos injection (placeholder)
  - Post-chaos checks
  - Result recording

### 2. RBAC Configuration
- ServiceAccount, Role, and RoleBinding
- Permissions for pods, jobs, events, and chaos resources

### 3. Test Manifest
- Basic pod specification for testing the experiment
- Environment variables for chaos configuration

### 4. Chaos Library (`chaoslib/litmus/jvm-chaos/`)
- **lib/jvm-chaos.go**: Orchestration logic (placeholder)
- **helper/jvm-helper.go**: Helper pod logic (placeholder)

### 5. Supporting Packages (`pkg/generic/jvm-chaos/`)
- **types.go**: Experiment-specific data structures with JVM parameters
- **environment.go**: Environment variable parsing

## What's NOT Implemented (Placeholders)

1. **Java Agent**: The actual Java agent JAR that will inject chaos
2. **Helper Pod Logic**: Code to attach Java agent to target JVM
3. **Process Discovery**: Finding Java processes in target containers
4. **Chaos Injection**: CPU/Memory stress implementation
5. **Build Integration**: Adding to Makefile and build system
6. **Chaos Charts**: Chart manifests for Chaos Studio

## Next Steps (As per Proposal Phases)

### Phase 1: Java Agent Development
- Create Java agent project in test-tools repository
- Implement CPU stress (Fibonacci calculations)
- Implement memory stress (allocation with/without references)
- Build agent JAR

### Phase 2: Build Integration
- Add Java build to litmus-go Makefile
- Include agent JAR in litmus-go image
- Update Dockerfile

### Phase 3: Complete Chaos Library
- Implement helper pod creation logic
- Add process discovery (find Java PIDs)
- Implement agent attachment using Java Attach API
- Add cleanup and revert logic

### Phase 4: Charts and Studio Integration
- Create chaos-charts manifests
- Add experiment CR
- Add ChartServiceVersion
- Enable in Chaos Studio

## Testing the Proof of Concept

Currently, this is a structural proof-of-concept. To test:

1. The experiment structure follows Litmus patterns
2. RBAC is properly configured
3. Types and environment handling are in place
4. The flow compiles (once integrated with build system)

## Key Design Decisions

1. **Helper Pod Pattern**: Uses helper pods (like stress-chaos) rather than exec-based injection
2. **Separate Chaos Library**: Follows the pattern of having lib/ and helper/ separation
3. **JVM-Specific Parameters**: Added Java-specific tunables (threads, memory, version)
4. **Minimal Dependencies**: Keeps dependencies minimal for proof-of-concept

## Integration Points

To integrate this into the build:

1. Add to `bin/experiment/Dockerfile` - experiment binary
2. Add to `bin/helper/Dockerfile` - helper binary  
3. Update `Makefile` - build targets
4. Add experiment registration in main experiment router

## References

- Proposal: `litmus/proposals/jvm-fault-injection.md`
- Similar pattern: `experiments/generic/pod-cpu-hog/`
- Helper pattern: `chaoslib/litmus/stress-chaos/`
