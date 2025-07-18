package main

import (
	"etcd_cli_pickme/internal"
	"etcd_cli_pickme/internal/config"
	"etcd_cli_pickme/internal/usecases"
	"fmt"
	"os"
)

func main() {
	// reset operation handle special case
	if len(os.Args) == 2 && os.Args[1] == "--reset" {
		err := internal.RemoveLockFile()
		if err != nil {
			return
		}
	}

	// get the user env selection + get selected env gitlab group id
	envValue, grpId, err := usecases.GetEnvValueFromUser()
	if err != nil {
		fmt.Println("error occurred while getting env")
		return
	}

	// verify previously fetched credentials from the gitlab
	err = internal.EnsureInitialized(grpId)
	if err != nil {
		return
	}

	// load etcd access credentials from the properties.yaml
	if err = config.LoadETCDConfig(envValue); err != nil {
		fmt.Println("error : ", err)
		return
	}

	// init etcd client
	err = usecases.InitEtcdClient()
	if err != nil {
		fmt.Println("error : ", err)
		return
	}

	// get search term from the user
	etcdKeyValue, err := usecases.GetSearchTermFromTheUser()
	if err != nil {
		fmt.Println("error : ", err)
	}

	// fetch the values from the etcd
	res, err := usecases.GetReadResult(etcdKeyValue)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// view the result
	usecases.ViewResponse(res)

}
