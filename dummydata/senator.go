package dummydata

func Senator() map[string]interface{} {
	s := Member()
	s["senate_class"] = "57"
	s["missed_votes_pct"] = 5.889
	s["votes_with_party_pct"] = 67.9087

	return s
}
