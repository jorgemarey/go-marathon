/*
Copyright 2014 Rohith All rights reserved.

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

package marathon

import (
	"fmt"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	AppID             string               `json:"appId"`
	Host              string               `json:"host"`
	ID                string               `json:"id"`
	HealthCheckResult []*HealthCheckResult `json:"healthCheckResults"`
	Ports             []int                `json:"ports"`
	ServicePorts      []int                `json:"servicePorts"`
	StagedAt          string               `json:"stagedAt"`
	StartedAt         string               `json:"startedAt"`
	Version           string               `json:"version"`
}

func (client *Client) AllTasks() (*Tasks, error) {
	tasks := new(Tasks)
	if err := client.ApiGet(MARATHON_API_TASKS, "", tasks); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}

func (client *Client) Tasks(application_id string) (*Tasks, error) {
	tasks := new(Tasks)
	if err := client.ApiGet(fmt.Sprintf("%s%s/tasks", MARATHON_API_APPS, application_id), "", tasks); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}
