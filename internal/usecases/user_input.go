package usecases

import (
	"errors"
	"etcd_cli_pickme/internal/config"
	"fmt"
	"strconv"
)

func GetEnvValueFromUser() (envValue, grpId string, err error) {
	for {
		envValue, grpId, err = fetchEnvFromUser()
		if err == nil {
			return envValue, grpId, err
		}
	}
}

func fetchEnvFromUser() (envValue, grpId string, err error) {
	var envUserInsertedValue string

	fmt.Print("Enter environment (dev-1) (stage-2) (live-3): ")
	_, err = fmt.Scan(&envUserInsertedValue)
	if err != nil {
		return "", "", err
	}

	envUserValueInt64, err := strconv.ParseInt(envUserInsertedValue, 10, 64)
	if err != nil {
		fmt.Println("error occurred parsing the input value.")
		return "", "", err
	}

	if envUserValueInt64 < 1 || envUserValueInt64 > 3 {
		fmt.Println("invalid input")
		return "", "", errors.New("invalid input")
	}

	envKey, grpId := selectEnvKey(envUserValueInt64)
	return envKey, grpId, nil
}

func selectEnvKey(envValue int64) (envKey, grpId string) {
	switch envValue {
	case 1:
		return "dev", config.AppCfg.GitLabCfg.ProjectIds.Dev
	case 2:
		return "stage", config.AppCfg.GitLabCfg.ProjectIds.Dev
	case 3:
		return "live", config.AppCfg.GitLabCfg.ProjectIds.Dev
	}
	return envKey, grpId
}

func GetSearchTermFromTheUser() (searchTerm string, err error) {
	var searchTermValue string
	fmt.Print("Search Prefix (/system/delivery-services/) : ")
	_, err = fmt.Scan(&searchTermValue)
	if err != nil {
		return "", err
	}
	return searchTermValue, nil
}
