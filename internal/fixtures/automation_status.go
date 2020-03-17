package fixtures

import om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"

func AutomationStatus() *om.AutomationStatus {
	return &om.AutomationStatus{
		GoalVersion: 2,
		Processes: []om.ProcessStatus{
			{
				Name:                    "shardedCluster_myShard_0_0",
				Hostname:                "testDeploy-0",
				Plan:                    []string{},
				LastGoalVersionAchieved: 2,
			},
			{
				Name:                    "shardedCluster_myShard_0_1",
				Hostname:                "testDeploy-1",
				Plan:                    []string{},
				LastGoalVersionAchieved: 2,
			},
			{
				Name:                    "shardedCluster_myShard_0_2",
				Plan:                    []string{"Download", "Start", "WaitRsInit"},
				Hostname:                "testDeploy-2",
				LastGoalVersionAchieved: 2,
			},
		},
	}
}
