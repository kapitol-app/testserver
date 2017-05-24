package dummydata

var member map[string]interface{}

func Member() map[string]interface{} {
	m := map[string]interface{} {
		"id": "tt_86547",
		"api_uri": "https://someapiuri/withsomeotherstuff",
		"first_name": "karl",
		"last_name": "marx",
		"party": "I",
		"leadership_role": "king of the world",
		"twitter_account": "https://twitter.com/karl-marx",
		"facebook_account": "https://facebook.com/karl-marx",
		"url": nil,
		"govtrack_id": "gov_64644848",
		"in_office": false,
		"next_election": nil,
		"total_votes": 42,
		"missed_votes": 9,
		"office": "2278 Prolotaritet Way",
		"phone": "+40 45 87 90 98 76",
		"state": "ge",
	}

	return m
}
