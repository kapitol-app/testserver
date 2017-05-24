package models

type Senator struct {
	SenateClass string `json:"senate_class"`
	MissedVotesPct float64 `json:"missed_votes_pct"`
	VotesWithPartyPct float64 `json:"votes_with_party_pct"`
	*Member
}

func (s *Senator)ToMap() map[string]interface{} {
	m := s.Member.ToMap()
	m["senate_class"] = s.SenateClass
	m["missed_votes_pct"] = s.MissedVotesPct
	m["votes_with_party_pct"] = s.VotesWithPartyPct

	return m
}
