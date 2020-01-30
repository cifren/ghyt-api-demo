package config

import "os"

func GetParameters() map[string]interface{} {

	return map[string]interface{}{
        "github": map[string]string{
            "github_account": "",
            "github_secret": "",
        },
        "youtrack": map[string]string{
            "url": os.Getenv("YOUTRACK_URL"),
            "token": os.Getenv("YOUTRACK_TOKEN"),
        },
    }
}
