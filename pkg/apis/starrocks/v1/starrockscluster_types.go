/*
Copyright 2021-present, StarRocks Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StarRocksClusterSpec defines the desired state of StarRocksCluster
type StarRocksClusterSpec struct {
	// Specify a Service Account for starRocksCluster use k8s cluster.
	// +optional
	// Deprecated: component use serviceAccount in own's field.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// StarRocksFeSpec define fe configuration for start fe service.
	StarRocksFeSpec *StarRocksFeSpec `json:"starRocksFeSpec,omitempty"`

	// StarRocksBeSpec define be configuration for start be service.
	StarRocksBeSpec *StarRocksBeSpec `json:"starRocksBeSpec,omitempty"`

	// StarRocksCnSpec define cn configuration for start cn service.
	StarRocksCnSpec *StarRocksCnSpec `json:"starRocksCnSpec,omitempty"`
}

// StarRocksClusterStatus defines the observed state of StarRocksCluster.
type StarRocksClusterStatus struct {
	// Represents the state of cluster. the possible value are: running, failed, pending
	Phase ClusterPhase `json:"phase"`

	// Represents the status of fe. the status have running, failed and creating pods.
	StarRocksFeStatus *StarRocksFeStatus `json:"starRocksFeStatus,omitempty"`

	// Represents the status of be. the status have running, failed and creating pods.
	StarRocksBeStatus *StarRocksBeStatus `json:"starRocksBeStatus,omitempty"`

	// Represents the status of cn. the status have running, failed and creating pods.
	StarRocksCnStatus *StarRocksCnStatus `json:"starRocksCnStatus,omitempty"`
}

// ClusterPhase represent the cluster phase. the possible value for cluster phase are: running, failed, pending.
type ClusterPhase string

// MemberPhase represent the component phase about be, cn, be. the possible value for component phase are: reconciliing, failed, running, waitting.
type MemberPhase string

const (
	// ClusterRunning represents starrocks cluster is running.
	ClusterRunning ClusterPhase = "running"

	// ClusterFailed represents starrocks cluster failed.
	ClusterFailed ClusterPhase = "failed"

	// ClusterPending represents the starrocks cluster is creating
	ClusterPending ClusterPhase = "pending"

	// ClusterDeleting waiting all resource deleted
	ClusterDeleting ClusterPhase = "deleting"
)

const (
	// ComponentReconciling the starrocks have component in starting.
	ComponentReconciling MemberPhase = "reconciling"
	// ComponentFailed have at least one service failed.
	ComponentFailed MemberPhase = "failed"
	// ComponentRunning all components runs available.
	ComponentRunning MemberPhase = "running"
)

// AnnotationOperationValue present the operation for fe, cn, be.
type AnnotationOperationValue string

const (
	// AnnotationRestart represent the user want to restart all fe pods.
	AnnotationRestart AnnotationOperationValue = "restart"
	// AnnotationRestartFinished represent all fe pods have restarted.
	AnnotationRestartFinished AnnotationOperationValue = "finished"
	// AnnotationRestarting represent at least one pod on restarting
	AnnotationRestarting AnnotationOperationValue = "restarting"
)

// Operation response key in annnotation, the annotation key be associated with annotation value represent the process status of sr operation.
type AnnotationOperationKey string

const (
	// AnnotationFERestartKey the fe annotation key for restart
	AnnotationFERestartKey AnnotationOperationKey = "app.starrocks.fe.io/restart"

	// AnnotationBERestartKey the be annotation key for restart be
	AnnotationBERestartKey AnnotationOperationKey = "app.starrocks.be.io/restart"

	// AnnotationCNRestartKey the cn annotation key for restart cn
	AnnotationCNRestartKey AnnotationOperationKey = "app.starrocks.cn.io/restart"
)

type HorizontalScaler struct {
	// the horizontal scaler name
	Name string `json:"name,omitempty"`

	// the horizontal version.
	Version AutoScalerVersion `json:"version,omitempty"`
}

// StarRocksCluster defines a starrocks cluster deployment.
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=src
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FeStatus",type=string,JSONPath=`.status.starRocksFeStatus.phase`
// +kubebuilder:printcolumn:name="CnStatus",type=string,JSONPath=`.status.starRocksCnStatus.phase`
// +kubebuilder:printcolumn:name="BeStatus",type=string,JSONPath=`.status.starRocksBeStatus.phase`
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// +genclient
type StarRocksCluster struct {
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired state of the starrocks cluster.
	Spec StarRocksClusterSpec `json:"spec,omitempty"`

	// Most recent observed status of the starrocks cluster
	Status StarRocksClusterStatus `json:"status,omitempty"`
}

// StorageVolume defines additional PVC template for StatefulSets and volumeMount for pods that mount this PVC
type StorageVolume struct {
	// name of a storage volume.
	// +kubebuilder:validation:Pattern=[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
	Name string `json:"name"`

	// storageClassName is the name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty"`

	// StorageSize is a valid memory size type based on powers-of-2, so 1Mi is 1024Ki.
	// Supported units:Mi, Gi, GiB, Ti, Ti, Pi, Ei, Ex: `512Mi`.
	// +kubebuilder:validation:Pattern:="(^0|([0-9]*l[.])?[0-9]+((M|G|T|E|P)i))$"
	StorageSize string `json:"storageSize"`

	// MountPath specify the path of volume mount.
	MountPath string `json:"mountPath,omitempty"`
}

// StarRocksClusterList contains a list of StarRocksCluster
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type StarRocksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StarRocksCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StarRocksCluster{}, &StarRocksClusterList{})
}
