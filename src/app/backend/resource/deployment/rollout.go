// Copyright 2017 The Kubernetes Authors.
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

package deployment

import (
	"errors"

	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client "k8s.io/client-go/kubernetes"
)

// RollbackSpec is a specification for deployment rollback from an specific revision number
type RollbackSpec struct {
	// revision is the revision number of the replicateSet which we want to rollback
	Revision string `json:"revision"`
}

// RollbackDeployment rollback to a specific replicaSet version
func RollbackDeployment(client client.Interface, rollbackSpec *RollbackSpec, namespace, name string) error {
	deployment, err := client.AppsV1().Deployments(namespace).Get(name, metaV1.GetOptions{})
	if err != nil {
		return err
	}
	currRevision := deployment.Annotations["deployment.kubernetes.io/revision"]
	if currRevision == "1" {
		return errors.New("No revision for rolling back ")
	}
	matchRS, err := GetReplicaSetFromDeployment(client, namespace, name)
	if err != nil {
		return err
	}
	for _, rs := range matchRS {
		if rs.Annotations["deployment.kubernetes.io/revision"] == rollbackSpec.Revision {
			updateDeployment := deployment.DeepCopy()
			updateDeployment.Spec.Template.Spec = rs.Spec.Template.Spec
			_, err = client.AppsV1().Deployments(namespace).Update(updateDeployment)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("No match revisionNumber replicaSet for deployment ")
}

// PauseDeployment is used to pause a deployment
func PauseDeployment(client client.Interface, namespace, deploymentName string) (*v1.Deployment, error) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if deployment.Spec.Paused {
		deployment.Spec.Paused = false
		_, err = client.AppsV1().Deployments(namespace).Update(deployment)
		if err != nil {
			return nil, err
		}
		return deployment, nil
	}
	return nil, errors.New("the deployment is already paused")
}

// ResumeDeployment is used to resume a deployment
func ResumeDeployment(client client.Interface, namespace, deploymentName string) (*v1.Deployment, error) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if !deployment.Spec.Paused {
		deployment.Spec.Paused = true
		_, err = client.AppsV1().Deployments(namespace).Update(deployment)
		if err != nil {
			return nil, err
		}
		return deployment, nil
	}
	return nil, errors.New("the deployment is already resumed")
}

// GetReplicaSetFromDeployment return all replicaSet which is belong to the deployment
func GetReplicaSetFromDeployment(client client.Interface, namespace, deploymentName string) ([]v1.ReplicaSet, error) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	selector, err := metaV1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		return nil, err
	}
	options := metaV1.ListOptions{LabelSelector: selector.String()}
	allRS, err := client.AppsV1().ReplicaSets(namespace).List(options)
	if err != nil {
		return nil, err
	}
	var result []v1.ReplicaSet
	for _, rs := range allRS.Items {
		if metaV1.IsControlledBy(&rs, deployment) {
			result = append(result, rs)
		}
	}
	return result, nil
}
