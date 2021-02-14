package main

import sw "./todoclient"

type app struct {
	todoclient sw.APIClient
	userclient userapi.APIClient
}
