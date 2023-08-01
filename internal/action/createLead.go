package action

var (
	// action type
	CreateLeadRequestType = "CreateLeadRequest"
)

// Тип Action с type "CreateLeadRequest"
type CreateLeadAction struct {
	Action
	Data CreateLeadActionData `json:"Data"`
}

// Data у action c type "CreateLeadRequest"
type CreateLeadActionData struct {
	Source struct {
		CourceCode       string `json:"CourceCode"`
		CampaignID       string `json:"CampaignId"`
		CampaignMemberID string `json:"CampaignMemberId"`
	} `json:"Source"`
	Lead struct {
		FirstName     string `json:"FirstName"`
		LastName      string `json:"LastName"`
		MiddleName    string `json:"MiddleName"`
		JobTitle      string `json:"JobTitle"`
		MobilePhone   string `json:"MobilePhone"`
		Telephone1    string `json:"Telephone1"`
		EMailAddress1 string `json:"EMailAddress1"`
		EMailAddress2 string `json:"EMailAddress2"`
		EMailAddress3 string `json:"EMailAddress3"`
		Products      []struct {
			Code string `json:"Code"`
		} `json:"Products"`
	} `json:"Lead"`
	Region struct {
		Name      string `json:"Name"`
		FiasCode  string `json:"FiasCode"`
		OkatoCode string `json:"OkatoCode"`
	} `json:"Region"`
	Company struct {
		Name          string `json:"Name"`
		SparkID       string `json:"SparkId"`
		Inn           string `json:"INN"`
		Kpp           string `json:"KPP"`
		Type          int    `json:"Type"`
		IsNotResident bool   `json:"IsNotResident"`
	} `json:"Company"`
}
