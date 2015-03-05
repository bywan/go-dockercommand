package dockercommand

import "fmt"

func (dock *Docker) pullImageIfNotExist(image string) error {
	_, err := dock.client.InspectImage(image)
	if err != nil && err.Error() == "no such image" {
		err = dock.Pull(&PullOptions{image})
	}
	return err
}

func convertEnvMapToSlice(envMap map[string]string) []string {
	envSlice := []string{}
	for key, value := range envMap {
		envSlice = append(envSlice, fmt.Sprintf("%s=%s", key, value))
	}
	return envSlice
}
