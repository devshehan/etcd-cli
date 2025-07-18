package internal

import (
	"etcd_cli_pickme/internal/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	projectDir, _      = os.Getwd()
	propertiesFilePath = filepath.Join(projectDir, "properties.yaml")
	lockfilePath       = filepath.Join(projectDir, ".initialized.lock")
)

func EnsureInitialized(groupId string) error {
	if isInitialized() {
		return nil
	}

	data, err := fetchTheDataFromTheGitlab(groupId)
	if err != nil {
		return err
	}

	if err = savePropertiesFile(data); err != nil {
		return err
	}

	if err = createLockFile(); err != nil {
		return nil
	}
	return nil
}

func isInitialized() bool {
	_, err := os.Stat(lockfilePath)
	return !os.IsNotExist(err)
}

func RemoveLockFile() error {
	err := os.Remove(lockfilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}

func fetchTheDataFromTheGitlab(groupId string) ([]byte, error) {
	url := fmt.Sprintf(
		"https://%s/api/v4/projects/%s/repository/files/properties.yaml/raw?ref=main",
		config.AppCfg.SelfHostedDomain,
		groupId,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", config.AppCfg.GitLabToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch properties: %v", err)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func savePropertiesFile(data []byte) error {
	return os.WriteFile(propertiesFilePath, data, 0644)
}

func createLockFile() error {
	return os.WriteFile(lockfilePath, []byte("initialized"), 0644)
}
