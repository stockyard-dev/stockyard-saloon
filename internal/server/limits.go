package server

type Tier string

const (
	TierFree Tier = "free"
	TierPro  Tier = "pro"
)

type Limits struct {
	Tier        Tier
	Description string
}

func LimitsFor(tier string) Limits {
	if tier == "pro" {
		return Limits{Tier: TierPro, Description: "Unlimited categories and threads"}
	}
	return Limits{Tier: TierFree, Description: "3 categories, 50 threads"}
}

func (l Limits) IsPro() bool {
	return l.Tier == TierPro
}
