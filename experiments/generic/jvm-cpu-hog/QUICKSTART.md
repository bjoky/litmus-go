# JVM Fault Injection - Quick Start Guide

## What Was Created

A complete proof-of-concept structure for JVM fault injection in litmus-go with placeholders for actual implementation.

## File Structure

### Experiment Files
- `experiments/generic/jvm-cpu-hog/experiment/jvm-cpu-hog.go` - Main experiment logic
- `experiments/generic/jvm-cpu-hog/rbac.yaml` - Kubernetes RBAC
- `experiments/generic/jvm-cpu-hog/test/test.yml` - Test manifest
- `experiments/generic/jvm-cpu-hog/README.md` - Documentation

### Chaos Library
- `chaoslib/litmus/jvm-chaos/lib/jvm-chaos.go` - Orchestration logic
- `chaoslib/litmus/jvm-chaos/helper/jvm-helper.go` - Helper pod logic

### Supporting Code
- `pkg/generic/jvm-chaos/types/types.go` - Type definitions
- `pkg/generic/jvm-chaos/environment/environment.go` - ENV parsing

## Key Features

✅ Standard Litmus experiment flow
✅ Helper pod pattern (like stress-chaos)
✅ RBAC configuration
✅ JVM-specific parameters (threads, memory, Java version)
✅ Proper error handling structure
✅ Event and result recording

## What's Missing (By Design)

These are placeholders for future implementation:

1. **Java Agent JAR** - Needs to be built separately
2. **Process Discovery** - Finding Java PIDs in containers
3. **Agent Attachment** - Using Java Attach API
4. **Actual Chaos Logic** - CPU/Memory stress implementation
5. **Build Integration** - Makefile and Dockerfile updates

## Next Development Steps

### Immediate (Phase 1)
1. Create Java agent project in `test-tools/custom/jvm-fault-injector/`
2. Implement CPU stress (Fibonacci)
3. Implement memory stress (allocation)
4. Build agent JAR

### Short-term (Phase 2)
1. Add Java build to Makefile
2. Include JAR in litmus-go image
3. Update Dockerfiles

### Medium-term (Phase 3)
1. Complete helper pod logic
2. Implement process discovery
3. Add agent attachment
4. Test end-to-end

### Long-term (Phase 4)
1. Create chaos-charts
2. Add to Chaos Studio
3. Documentation and examples

## How to Extend

### Adding Memory Hog Variant
Copy the structure and modify:
```bash
cp -r experiments/generic/jvm-cpu-hog experiments/generic/jvm-memory-hog
# Update experiment name and parameters
```

### Adding New JVM Parameters
1. Add to `pkg/generic/jvm-chaos/types/types.go`
2. Add to `pkg/generic/jvm-chaos/environment/environment.go`
3. Use in chaos library

## Testing Strategy

1. **Unit Tests**: Test individual functions
2. **Integration Tests**: Test with mock Java apps
3. **E2E Tests**: Test with real Java applications

## Important Notes

- This follows the helper pod pattern (not exec-based)
- Requires Java 17+ in target containers
- Helper pod needs Java runtime to attach agent
- Agent JAR must be included in litmus-go image

## References

- Proposal: `litmus/proposals/jvm-fault-injection.md`
- Similar experiments: `experiments/generic/pod-cpu-hog/`
- Helper pattern: `chaoslib/litmus/stress-chaos/`
- Developer guide: `contribute/developer-guide/README.md`
