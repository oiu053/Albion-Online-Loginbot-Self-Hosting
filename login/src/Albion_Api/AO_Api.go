package albionAPI

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var baseURL = "https://gameinfo.albiononline.com/api/gameinfo"

func Getplayerdeaths_by_ID(PlayerID string) (data string) {
	URL := baseURL + "/players/" + PlayerID + "/deaths"

	return request(URL)
}

/*
// Pve
func Getplayerkills_by_ID(PlayerID string) (data string) {
	URL := baseURL + "/players/" + PlayerID + "/deaths"

	return request(URL)
}
func GetTop50GuildPlayerstatics_week_Pve(GuildID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=week&limit=50&offset=0&type=PvE&region=Total"

	return request(URL)
}
func GetTop50AlliancePlayerstatics_week_PvE(AllianceID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=week&limit=50&offset=0&type=PvE&region=Total&allianceId=" + AllianceID

	return request(URL)
}
func GetTop50GuildPlayerstatics_month_Pve(GuildID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=month&limit=50&offset=0&type=PvE&region=Total"

	return request(URL)
}
func GetTop50AlliancePlayerstatics_month_PvE(AllianceID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=month&limit=50&offset=0&type=PvE&region=Total&allianceId=" + AllianceID

	return request(URL)
}

///Gathering

func GetTop50GuildPlayerstatics_week_Gathering(GuildID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=week&limit=50&offset=0&type=Gathering"

	return request(URL)
}
func GetTop50AlliancePlayerstatics_week_Gathering(AllianceID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=week&limit=50&offset=0&type=Gathering&allianceId=" + AllianceID

	return request(URL)
}
func GetTop50GuildPlayerstatics_month_Gathering(GuildID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=month&limit=50&offset=0&type=Gathering"

	return request(URL)
}
func GetTop50AlliancePlayerstatics_month_Gathering(AllianceID string) (data string) {
	URL := baseURL + "/players/statistics" + "?range=month&limit=50&offset=0&type=Gathering&allianceId=" + AllianceID

	return request(URL)
}

*/
//Playerinfo

func Getplayerinfo_by_ID(PlayerID string) (data string) {
	URL := baseURL + "/players/" + PlayerID

	return request(URL)
}

//Guildinfo

func GetGuildinfo_by_ID(GuildID string) (data string) {
	URL := baseURL + "/guilds/" + GuildID

	return request(URL)
}

func GetguildMembers_by_ID(GuildID string) (data []string) {
	URL := baseURL + "/guilds/" + GuildID + "/members"

	data1 := request(URL)

	data2 := strings.Split(data1, ",")

	for _, v := range data2 {

		if strings.HasPrefix(v, `"Name":"`) {
			data3 := strings.TrimPrefix(strings.TrimSuffix(v, `"`), `"Name":"`)

			data = append(data, data3)
		}

	}

	return data
}

func GetGuildData_By_ID(GuildID string) (data string) {
	URL := baseURL + "/guilds/" + GuildID + "/DATA"

	return request(URL)
}

// Guildkills
func GetGuildKills_month(GuildID string) (data string) {
	URL := baseURL + "/Guilds" + GuildID + "top?limit=9999&offset=1&range=month&region=Total"

	return request(URL)
}
func GetGuildKills_week(GuildID string) (data string) {
	URL := baseURL + "/Guilds" + GuildID + "top?limit=9999&offset=1&range=week&region=Total"

	return request(URL)
}

// Alliance
func GetAllianceinfo_by_ID(AllianceID string) (data string) {
	URL := baseURL + "/alliances/" + AllianceID

	return request(URL)
}

//BAttles

func GEtBattleData_by_ID(BAttleID string) (data string) {
	URL := baseURL + "/battles/" + BAttleID

	return request(URL)
}

// Items
func GetItemData_by_ID(ItemID string) (data string) {
	URL := baseURL + "" + ItemID + "/data"

	return request(URL)
}

// KillInfo
func KillInfo_BY_event_ID(KillID string) (data string) {
	URL := baseURL + "/events/" + KillID

	return request(URL)
}

/* Search
func Search(wanted string) (data string) {
	URL := baseURL + "/search?q=" + wanted
	a := request(URL)
	return a
}*/

func SearchPlayer(wanted string) (playersdata string) {

	URL := baseURL + "/search?q=" + wanted
	b := request(URL)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[1]

	B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), strings.ToLower(Optvar)) {
				return player
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), strings.ToLower(Optvar)) {
			return Ab
		}
	}
	return "-"
}
func SearchGuild(wanted string) (playersdata string) {

	URL := baseURL + "/search?q=" + wanted
	b := request(URL)

	//b := strings.ToLower(a)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[0]
	C = strings.TrimPrefix(C, "{")
	C = strings.TrimPrefix(C, `"guilds":[`)

	B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), Optvar) {
				return player
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), Optvar) {
			return Ab
		}
	}
	return "-"
}

func SearchGuildID(wanted string) (playersdata string) {

	URL := baseURL + "/search?q=" + wanted
	b := request(URL)

	//b := strings.ToLower(a)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[0]
	B := strings.TrimPrefix(C, `{"guilds":[{`)

	//B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), Optvar) {
				Z := strings.Split(player, ",")

				ID := Z[0]

				ID = strings.TrimPrefix(ID, `"Id":"`)
				ID = strings.TrimSuffix(ID, `"`)

				return ID
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), Optvar) {
			Z := strings.Split(Ab, ",")

			ID := Z[0]

			ID = strings.TrimPrefix(ID, `"Id":"`)
			ID = strings.TrimSuffix(ID, `"`)

			return ID
		}
	}
	return "-"
}

func SearchPlayerID(wanted string) (playersdata string) {

	URL := baseURL + "/search?q=" + wanted
	b := request(URL)

	//b := strings.ToLower(a)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[1]

	B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), Optvar) {
				Z := strings.Split(player, ",")

				ID := Z[0]

				ID = strings.TrimPrefix(ID, `"Id":"`)
				ID = strings.TrimSuffix(ID, `"`)

				return ID
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), Optvar) {
			Z := strings.Split(Ab, ",")

			ID := Z[0]

			ID = strings.TrimPrefix(ID, `"Id":"`)
			ID = strings.TrimSuffix(ID, `"`)

			return ID
		}
	}
	return "-"
}
func SearchPlayerGuildName(wanted string) (playersdata string, errors bool) {

	URL := baseURL + "/search?q=" + wanted

	//fmt.Println(URL)

	b := request(URL)
	//fmt.Println(URL)

	//b := strings.ToLower(a)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[1]

	B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), Optvar) {
				Z := strings.Split(player, ",")

				ID := Z[3]
				//fmt.Print(Z)

				ID = strings.TrimPrefix(ID, `"GuildName":"`)
				ID = strings.TrimSuffix(ID, `"`)

				return ID, false
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), Optvar) {
			Z := strings.Split(Ab, ",")

			ID := Z[3]
			//fmt.Print(Z)

			ID = strings.TrimPrefix(ID, `"GuildName":"`)
			ID = strings.TrimSuffix(ID, `"`)

			return ID, false
		}
	}
	return "-", true
}

// RequestFunction
func request(url string) (data string) {

	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	a := string(body)
	return a
}
func SearchPlayerAllianceName(wanted string) (playersdata string, errors bool) {

	URL := baseURL + "/search?q=" + wanted

	//fmt.Println(URL)

	b := request(URL)
	//fmt.Println(URL)

	//b := strings.ToLower(a)

	split1 := strings.Split(b, `,"players":[`)

	C := split1[1]

	B := strings.TrimPrefix(C, "{")

	Ab := strings.TrimSuffix(B, "}]}")

	Test := strings.Contains(Ab, "},{")

	if Test {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		A := strings.Split(Ab, "},{")
		for _, player := range A {
			if strings.Contains(strings.ToLower(player), Optvar) {
				Z := strings.Split(player, ",")

				AID := Z[5]
				//fmt.Print(Z)

				//fmt.Print(AID)

				AID = strings.TrimPrefix(AID, `"AllianceName":"`)
				AID = strings.TrimSuffix(AID, `"`)

				if AID == "" {
					AID = "-"
				}

				if AID == `"AllianceName":null` {
					AID = "-"
				}
				//fmt.Print(AID)

				return AID, false
			}
		}
	} else {
		Optvar := `","name":"` + strings.ToLower(wanted) + `"`
		if strings.Contains(strings.ToLower(Ab), Optvar) {
			Z := strings.Split(Ab, ",")

			AID := Z[5]

			AID = strings.TrimPrefix(AID, `"AllianceName":"`)
			AID = strings.TrimSuffix(AID, `"`)

			if AID == "" {
				AID = "-"
			}

			if AID == `"AllianceName":null` {
				AID = "-"
			}
			//fmt.Print(AID)

			return AID, false
		}
	}
	return "-", true
}
