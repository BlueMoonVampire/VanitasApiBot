package types

type CheckResponse struct {
	User        int
	Reason      string `json:"reason"`
	Message     string
	Enforcer    string `json:"enforcer"`
	Blacklisted bool
}
