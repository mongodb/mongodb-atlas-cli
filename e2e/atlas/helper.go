package atlas

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func GetHostnameAndPort(cliPath, atlasEntity string) (string, error) {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"processes",
		"list")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	var processes []*mongodbatlas.Process
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return "", err
	}

	if len(processes) == 0 {
		return "", fmt.Errorf("got=%#v\nwant=%#v\n", 0, "len(processes) > 0")
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil
}

func DeployCluster(cliPath, atlasEntity, clusterName string) error {
	cmd := exec.Command(cliPath,
		atlasEntity,
		"clusters",
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
		"clusters",
		"watch",
		clusterName)
	cmd.Env = os.Environ()
	return cmd.Run()
}

func DeleteCluster(cliPath, atlasEntity, clusterName string) error {
	cmd := exec.Command(cliPath, atlasEntity, "clusters", "delete", clusterName, "--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

func GetHostname(cliPath, atlasEntity string) (string, error) {
	hostnamePort, err := GetHostnameAndPort(cliPath, atlasEntity)
	if err != nil {
		return "", err
	}

	parts := strings.Split(hostnamePort, ":")
	return parts[0], nil
}
