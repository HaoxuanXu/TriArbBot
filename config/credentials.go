package config

import (
	"strings"

	"github.com/HaoxuanXu/TriArbBot/config/credentials/live"
	"github.com/HaoxuanXu/TriArbBot/config/credentials/paper"
)

type credentials struct {
	API_KEY    string `json:"api_key"`
	API_SECRET string `json:"api_secret"`
	BASE_URL   string `json:"base_url"`
}

func GetCredentials(accountType, serverType string) credentials {
	var creds credentials

	if strings.ToLower(accountType) == LIVE_ACCOUNT {
		if strings.ToLower(serverType) == STAGING_SERVER {
			creds = credentials{
				API_KEY:    live.API_KEY_STAGING,
				API_SECRET: live.API_SECRET_STAGING,
				BASE_URL:   live.BASE_URL,
			}
		} else if strings.ToLower(serverType) == PRODUCTION_SERVER {
			creds = credentials{
				API_KEY:    live.API_KEY_PROD,
				API_SECRET: live.API_SECRET_PROD,
				BASE_URL:   live.BASE_URL,
			}
		}
	} else if strings.ToLower(accountType) == PAPER_ACCOUNT {
		creds = credentials{
			API_KEY:    paper.API_KEY,
			API_SECRET: paper.API_SECRET,
			BASE_URL:   paper.BASE_URL,
		}
	}

	return creds
}
