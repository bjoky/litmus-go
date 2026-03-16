# JVM CPU Hog Experiment

## Summary

This experiment injects CPU stress into Java Virtual Machine (JVM) applications to test their resilience under high CPU load conditions.

## Status

**PROOF OF CONCEPT** - This is a placeholder implementation to validate the experiment flow.

## What

The JVM CPU hog experiment will:
- Target Java applications running in containers
- Inject CPU-intensive operations into the JVM using a Java agent
- Use helper pods to inject the Java agent into target JVM processes

## Why

Java applications behave differently under stress compared to container-level stress. JVM-specific stress testing allows:
- Testing garbage collection behavior under load
- Validating JVM tuning parameters
- Ensuring application resilience to CPU contention

## How

The experiment will:
1. Identify target pods running Java applications
2. Deploy helper pods with the Java agent
3. Attach the agent to the target JVM process
4. Execute CPU-intensive operations (e.g., Fibonacci calculations)
5. Monitor and cleanup after the chaos duration

## Implementation Status

- [x] Basic experiment structure
- [x] RBAC configuration
- [x] Test manifest
- [ ] Environment variable handling
- [ ] Helper pod implementation
- [ ] Java agent integration
- [ ] Chaos library implementation

## Next Steps

1. Implement environment variable parsing
2. Create helper pod specification
3. Develop Java agent for CPU stress
4. Integrate with litmus-go build system
5. Add to chaos-charts repository
