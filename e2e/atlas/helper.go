package atlas

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	atlasEntity    = "atlas"
	clustersEntity = "clusters"
)

func GetHostnameAndPort(cliPath string) (string, error) {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"processes",
		"list")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	var processes []*mongodbatlas.Process
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return "", err
	}

	if len(processes) == 0 {
		return "", fmt.Errorf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil
}

// ExistCluster returns true if there is at least a cluster deployed, false otherwise
func ExistCluster(cliPath string) bool {
	cmd := exec.Command(cliPath,
		atlasEntity,
		clustersEntity,
		"list")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return false
	}

	var clusters []atlas.Cluster
	err = json.Unmarshal(resp, &clusters)

	if err != nil {
		return false
	}

	if len(clusters) > 0 {
		return true
	}

	return false

}

func DeployCluster(cliPath, clusterName string) error {
	cmd := exec.Command(cliPath,
		atlasEntity,
		clustersEntity,
		"create",
		clusterName,
		"--region=US_EAST_1",
		"--members=3",
		"--tier=M10",
		"--provider=AWS",
		"--mdbVersion=4.0",
		"--diskSizeGB=10")
	cmd.Env = os.Environ()
	err := cmd.Run()

	if err != nil {
		return err
	}

	cmd = exec.Command(cliPath,
		"atlas",
		clustersEntity,
		"watch",
		clusterName)
	cmd.Env = os.Environ()
	return cmd.Run()
}

func DeleteCluster(cliPath, clusterName string) error {
	cmd := exec.Command(cliPath, atlasEntity, "clusters", "delete", clusterName, "--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

func GetHostname(cliPath string) (string, error) {
	hostnamePort, err := GetHostnameAndPort(cliPath)
	if err != nil {
		return "", err
	}

	parts := strings.Split(hostnamePort, ":")
	return parts[0], nil
}
