// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by main. DO NOT EDIT.

package k8s

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Backend) DeepEqual(other *Backend) bool {
	if other == nil {
		return false
	}

	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	if in.NodeName != other.NodeName {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointSlices) DeepEqual(other *EndpointSlices) bool {
	if other == nil {
		return false
	}

	if ((in.epSlices != nil) && (other.epSlices != nil)) || ((in.epSlices == nil) != (other.epSlices == nil)) {
		in, other := &in.epSlices, &other.epSlices
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Endpoints) deepEqual(other *Endpoints) bool {
	if other == nil {
		return false
	}

	if ((in.Backends != nil) && (other.Backends != nil)) || ((in.Backends == nil) != (other.Backends == nil)) {
		in, other := &in.Backends, &other.Backends
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodePortToFrontend) DeepEqual(other *NodePortToFrontend) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(otherValue) {
					return false
				}
			}
		}
	}

	return true
}

// deepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Service) deepEqual(other *Service) bool {
	if other == nil {
		return false
	}

	if in.IsHeadless != other.IsHeadless {
		return false
	}
	if in.TrafficPolicy != other.TrafficPolicy {
		return false
	}
	if in.HealthCheckNodePort != other.HealthCheckNodePort {
		return false
	}
	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.NodePorts != nil) && (other.NodePorts != nil)) || ((in.NodePorts == nil) != (other.NodePorts == nil)) {
		in, other := &in.NodePorts, &other.NodePorts
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.LoadBalancerSourceRanges != nil) && (other.LoadBalancerSourceRanges != nil)) || ((in.LoadBalancerSourceRanges == nil) != (other.LoadBalancerSourceRanges == nil)) {
		in, other := &in.LoadBalancerSourceRanges, &other.LoadBalancerSourceRanges
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.Labels != nil) && (other.Labels != nil)) || ((in.Labels == nil) != (other.Labels == nil)) {
		in, other := &in.Labels, &other.Labels
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if ((in.Selector != nil) && (other.Selector != nil)) || ((in.Selector == nil) != (other.Selector == nil)) {
		in, other := &in.Selector, &other.Selector
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if in.SessionAffinity != other.SessionAffinity {
		return false
	}
	if in.SessionAffinityTimeoutSec != other.SessionAffinityTimeoutSec {
		return false
	}

	return true
}
