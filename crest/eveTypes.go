package crest

import ()

type AllianceListFormat struct {
	Pageable
	Items []ItemTypeDescription `json:"items"`
}

type AllianceDetails struct {
	StartDate            string        `json:"startDate"`
	CorporationsCount    int           `json:"corporationsCount"`
	Description          string        `json:"description"`
	ExecutorCorporation  CorpDetails   `json:"executorCorporation"`
	CorporationCount_str string        `json:"corporationCount_str"`
	Deleted              bool          `json:"deleted"`
	CreatorCorporation   CorpDetails   `json:"creatorCorporation"`
	Url                  string        `json:"url"`
	ID_str               string        `json:"id_str"`
	CreatorCharacter     CharDetails   `json:"creatorCharacter"`
	Corporations         []CorpDetails `json:"corporations"`
	ShortName            string        `json:"shortName"`
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
}

// type DogmaAttributesFormat struct {
// 	Pageable
// 	Items []ItemTypeDescription `json:"items"`
// }

type DogmaAttributeFormat struct {
	Description  string  `json:"description"`
	UnitID_str   string  `json:"unitID_str"`
	DisplayName  string  `json:"displayName"`
	Name         string  `json:"name"`
	IconID       int     `json:"iconID"`
	UnitID       int     `json:"unitID"`
	IconID_str   string  `json:"iconID_str"`
	ID           int     `json:"id"`
	ID_str       string  `json:"id_str"`
	HighIsGood   bool    `json:"highIsGood"`
	Stackable    bool    `json:"stackable"`
	DefaultValue float64 `json:"defaultValue"`
	Published    bool    `json:"published"`
}

type DogmaEffectFormat struct {
	EffectCategory_str string `json:"effectCategory_str"`
	PostExpression_str string `json:"postExpression_str"`
	IsAssistance       bool   `json:"isAssistance"`
	Description        string `json:"description"`
	IsOffensive        bool   `json:"isOffensive"`
	DisallowAutoRepeat bool   `json:"disallowAutoRepeat"`
	IsWarpSafe         bool   `json:"isWarpSafe"`
	PreExpression_str  string `json:"preExpression_str"`
	ElectronicChance   bool   `json:"electronicChance"`
	RangeChance        bool   `json:"rangeChance"`
	EffectCategory     int    `json:"effectCategory"`
	ID_str             string `json:"id_str"`
	PostExpression     int    `json:"postExpression"`
	Published          bool   `json:"published"`
	PreExpression      int    `json:"preExpression"`
	DisplayName        string `json:"displayName"`
	ID                 int    `json:"id"`
	Name               string `json:"name"`
}

type IncursionsFormat struct {
	Pageable
	Items []IncursionItem `json:"items"`
}

type IncursionItem struct {
	AggressorFactionID   ItemTypeDescription   `json:"aggressorFactionID"`
	StagingSolarSystem   ItemTypeDescription   `json:"stagingSolarSystem"`
	Influence            float64               `json:"influence"`
	IncursionType        string                `json:"incursionType"`
	State                string                `json:"state"`
	HasBoss              bool                  `json:"hasBoss"`
	InfestedSolarSystems []ItemTypeDescription `json:"infestedSolarSystems"`
	Constellation        ItemTypeDescription   `json:"constellation"`
}

type InsuranceInfoFormat struct {
	Pageable
	Items []InsuranceInfo `json:"items"`
}

type InsuranceInfo struct {
	Type      ItemTypeDescription `json:"type"`
	Insurance []InsurancePayout   `json:"insurance"`
}

type InsurancePayout struct {
	Payout float64 `json:"payout"`
	Cost   float64 `json:"cost"`
	Level  string  `json:"level"`
}

type OpportunitiesTaskFormat struct {
	Pageable
	Items []OpportunitiesTaskItem `json:"items"`
}

type OpportunitiesTaskItem struct {
	NotificationText string `json:"notificationText"`
	ID_str           string `json:"id_str"`
	Description      string `json:"description"`
	Name             string `json:"name"`
	ID               int    `json:"id"`
}

type OpportunityGroupFormat struct {
	Pageable
	Items []OpportunityGroup `json:"items"`
}

type OpportunityGroup struct {
	//imperfect. None of the ItemTypeDescriptions used has "name"
	GroupConnections []ItemIdDetails `json:"groupConnections"`
	AchievementTasks []ItemIdDetails `json:"achievementTasks"`
	OpportunitiesTaskItem
}

type SovStructuresFormat struct {
	Pageable
	Items []SovStructureItem `json:"items"`
}

type SovStructureItem struct {
	Alliance                    ItemTypeDescription `json:"alliance"`
	VulnerabilityOccupancyLevel float32             `json:"vulnerabilityOccupancyLevel"`
	StructureID_str             string              `json:"structureID_str"`
	StructureID                 int                 `json:"structureID"`
	VulnerableStartTime         string              `json:"vulnerableStartTime"`
	SolarSystem                 ItemTypeDescription `json:"solarSystem"`
	VulnerableEndTime           string              `json:"vulnerableEndTime"`
	Type                        ItemTypeDescription `json:"type"`
}

type SovCampaignsFormat struct {
	Pageable
	Items []SovCampaignItem `json:"items"`
}

type SovCampaignItem struct {
	EventType_str     string              `json:"eventType_str"`
	CampaignID        int                 `json:"campaignID"`
	EventType         int                 `json:"eventType"`
	SourceSolarsystem ItemTypeDescription `json:"sourceSolarsystem"`
	Attackers         AttackerInfo        `json:"attackers"`
	CampaignID_str    string              `json:"campaignID_str"`
	SourceItemID      int                 `json:"sourceItemID"`
	StartTime         string              `json:"startTime"`
	SourceItemID_str  string              `json:"sourceItemID_str"`
	Defender          DefenderInfo        `json:""`
	Constellation     ItemTypeDescription `json:""`
}

type AttackerInfo struct {
	Score float32 `json:"score"`
}

type DefenderInfo struct {
	Score    float32             `json:"score"`
	Defender ItemTypeDescription `json:"defender"`
}

// types is a TypeListContainer

type InvTypeDetails struct {
	GraphicID       GraphicIdDetails `json:"graphicID"`
	Capacity        float64          `json:"capacity"`
	Description     string           `json:"description"`
	PortionSize_str string           `json:"portionSize_str"`
	IconID          int              `json:"iconID"`
	PortionSize     int              `json:"portionSize"`
	IconID_str      string           `json:"iconID_str"`
	Volume          float64          `json:"volume"`
	Dogma           InvTypeDogma     `json:"dogma"`
	Radius          float64          `json:"radius"`
	ID_str          string           `json:"id_str"`
	Published       bool             `json:"published"`
	Mass            float64          `json:"mass"`
	ID              int              `json:"id"`
	Name            string           `json:"name"`
}

type GraphicIdDetails struct {
	ID_str string `json:"id_str"`
	ID     int    `json:"id"`
	Href   string `json:"href"`
	SofDNA string `json:"sofDNA"`
}

type InvTypeDogma struct {
	Attributes []InvTypeDogmaAttributes `json:"attributes"`
	Effects    []InvTypeDogmaEffects    `json:"effects"`
}

type InvTypeDogmaAttributes struct {
	Attribute ItemTypeDescription `json:"attribute"`
	Value     float64             `json:"value"`
}

type InvTypeDogmaEffects struct {
	Attribute ItemTypeDescription `json:"attribute"`
	IsDefault bool                `json:"isDefault"`
}

type WarListFormat struct {
	Pageable
	Items []ItemIdDetails `json:"items"`
}

type WarDetails struct {
	TimeFinished  string         `json:"timeFinished"`
	OpenForAllies bool           `json:"openForAllies"`
	TimeStarted   string         `json:"timeStarted"`
	AllyCount     int            `json:"allyCount"`
	TimeDeclared  string         `json:"timeDeclared"`
	Aggressor     WarParticipant `json:"aggressor"`
	Mutual        bool           `json:"mutual"`
	AllyCount_str string         `json:"allyCount_str"`
	Killmails     string         `json:"killmails"`
	ID_str        string         `json:"id_str"`
	Defender      WarParticipant `json:"defender"`
	ID            int            `json:"id"`
}

type WarParticipant struct {
	ShipsKilled     int     `json:"shipsKilled"`
	ShipsKilled_str string  `json:"shipsKilled_str"`
	Name            string  `json:"name"`
	Href            string  `json:"href"`
	ID_str          string  `json:"id_str"`
	Icon            SubHref `json:"icon"`
	ID              int     `json:"id"`
	IskKilled       float64 `json:"iskKilled"`
}
