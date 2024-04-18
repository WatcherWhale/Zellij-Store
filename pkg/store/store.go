package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type StoreMap map[string]string

func getCacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, "zellij-store"), nil
}

func getStorePath(session string) (string, error) {
	storeDir, err := getCacheDir()
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(storeDir); os.IsNotExist(err) {
		err = os.Mkdir(storeDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return path.Join(storeDir, session+".json"), nil
}

func LoadStore(session string) (StoreMap, error) {
	var store StoreMap
	storePath, err := getStorePath(session)

	if err != nil {
		return nil, err
	}

	jsonBytes, err := os.ReadFile(storePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func SaveStore(session string, store StoreMap) error {
	jsonByes, err := json.Marshal(store)
	if err != nil {
		return nil
	}

	storePath, err := getStorePath(session)
	if err != nil {
		return nil
	}

	err = os.WriteFile(storePath, jsonByes, os.ModePerm)
	if err != nil {
		return nil
	}

	return nil
}

func CleanCache() error {
	storePath, err := getCacheDir()

	if err != nil {
		return nil
	}

	paths, err := os.ReadDir(storePath)
	if err != nil {
		return nil
	}

	for _, file := range paths {
		if file.IsDir() {
			continue
		}
		err = os.Remove(path.Join(storePath, file.Name()))
		if err != nil {
			return err
		}

		fmt.Printf("Removed '%s' from the store\n", file.Name())
	}

	return nil
}
