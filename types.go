package main

type RequestBody struct {
	ID             int     `json:"id"`
	ActivityTime   float64 `json:"activityTime"`
	ActivityType   string  `json:"activityType"`
	StatusCode     string  `json:"statusCode"`
	Status         string  `json:"status"`
	ActivityResult string  `json:"activityResult"`
	UserID         int     `json:"userId"`
	Message        string  `json:"message"`
	Type           string  `json:"type"`
	Data           `json:"data"`
}

type Data struct {
	Message struct {
		Code   string `json:"code"`
		Params struct {
			ClientID     string `json:"clientId"`
			ClientName   string `json:"clientName"`
			AppUserName  string `json:"appUserName"`
			AppUserID    string `json:"appUserId"`
			AppUserEmail string `json:"appUserEmail"`
		} `json:"params"`
	} `json:"message"`
}
