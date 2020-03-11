package mh4u

type Monster struct {
	Monster struct {
		ID             int    `json:"id"`
		BaseHp         int    `json:"base_hp"`
		Size           string `json:"size"`
		HpMultLow      string `json:"hp_mult_low"`
		HpMultHigh     string `json:"hp_mult_high"`
		HpMultG        string `json:"hp_mult_g"`
		CrownMiniature string `json:"crown_miniature"`
		CrownLarge     string `json:"crown_large"`
		CrownKing      string `json:"crown_king"`
		RageDuration   int    `json:"rage_duration"`
		RageModAttack  string `json:"rage_mod_attack"`
		RageModDefense string `json:"rage_mod_defense"`
		RageModSpeed   string `json:"rage_mod_speed"`
		LimpLow        int    `json:"limp_low"`
		CapLow         int    `json:"cap_low"`
		LimpHigh       int    `json:"limp_high"`
		CapHigh        int    `json:"cap_high"`
		LimpHighApex   int    `json:"limp_high_apex"`
		CapHighApex    int    `json:"cap_high_apex"`
		LimpG          int    `json:"limp_g"`
		CapG           int    `json:"cap_g"`
		LimpGApex      int    `json:"limp_g_apex"`
		CapGApex       int    `json:"cap_g_apex"`
		LocalName      string `json:"local_name"`
		Link           string `json:"link"`
		Monstersounds  []struct {
			ID        int    `json:"id"`
			MonsterID int    `json:"monster_id"`
			Filename  string `json:"filename"`
		} `json:"monstersounds"`
		Monsterbodyparts []struct {
			ID        int    `json:"id"`
			LocalName string `json:"local_name"`
			Pivot     struct {
				MonsterID         int         `json:"monster_id"`
				MonsterbodypartID int         `json:"monsterbodypart_id"`
				Type              string      `json:"type"`
				ResCut            int         `json:"res_cut"`
				ResImpact         int         `json:"res_impact"`
				ResShot           int         `json:"res_shot"`
				ResFire           int         `json:"res_fire"`
				ResWater          int         `json:"res_water"`
				ResIce            int         `json:"res_ice"`
				ResThunder        int         `json:"res_thunder"`
				ResDragon         int         `json:"res_dragon"`
				ResDizzy          interface{} `json:"res_dizzy"`
			} `json:"pivot"`
		} `json:"monsterbodyparts"`
		Monsterstaggerlimits []struct {
			ID           int         `json:"id"`
			MonsterID    int         `json:"monster_id"`
			Region       string      `json:"region"`
			Value        int         `json:"value"`
			ValueCut     interface{} `json:"value_cut"`
			ExtractColor string      `json:"extract_color"`
		} `json:"monsterstaggerlimits"`
		Weaponspecialattacks []struct {
			ID        int    `json:"id"`
			LocalName string `json:"local_name"`
			Pivot     struct {
				MonsterID             int `json:"monster_id"`
				WeaponspecialattackID int `json:"weaponspecialattack_id"`
				Initial               int `json:"initial"`
				Increase              int `json:"increase"`
				Max                   int `json:"max"`
				Duration              int `json:"duration"`
				Damage                int `json:"damage"`
				ReductionTime         int `json:"reduction_time"`
				ReductionAmount       int `json:"reduction_amount"`
			} `json:"pivot"`
		} `json:"weaponspecialattacks"`
		Itemeffects []struct {
			ID        int    `json:"id"`
			Rarity    int    `json:"rarity"`
			Carry     int    `json:"carry"`
			Sell      int    `json:"sell"`
			Buy       int    `json:"buy"`
			LocalName string `json:"local_name"`
			Link      string `json:"link"`
			Pivot     struct {
				MonsterID int `json:"monster_id"`
				ItemID    int `json:"item_id"`
				Normal    int `json:"normal"`
				Enraged   int `json:"enraged"`
				Fatigued  int `json:"fatigued"`
			} `json:"pivot"`
		} `json:"itemeffects"`
		Canopytrap struct {
			ID        int `json:"id"`
			MonsterID int `json:"monster_id"`
			Normal    int `json:"normal"`
			Enraged   int `json:"enraged"`
			Fatigued  int `json:"fatigued"`
		} `json:"canopytrap"`
		Monsterguildquestitems []interface{} `json:"monsterguildquestitems"`
		Items                  []struct {
			ID        int    `json:"id"`
			Rarity    int    `json:"rarity"`
			Carry     int    `json:"carry"`
			Sell      int    `json:"sell"`
			Buy       int    `json:"buy"`
			LocalName string `json:"local_name"`
			Link      string `json:"link"`
			Pivot     struct {
				MonsterID           int `json:"monster_id"`
				ItemID              int `json:"item_id"`
				Quantity            int `json:"quantity"`
				Percentage          int `json:"percentage"`
				RankID              int `json:"rank_id"`
				MonsteritemmethodID int `json:"monsteritemmethod_id"`
				Monsteritemmethod   struct {
					ID        int    `json:"id"`
					LocalName string `json:"local_name"`
				} `json:"monsteritemmethod"`
				Rank struct {
					ID            int    `json:"id"`
					RarityLowest  int    `json:"rarity_lowest"`
					RarityHighest int    `json:"rarity_highest"`
					LocalName     string `json:"local_name"`
				} `json:"rank"`
			} `json:"pivot"`
		} `json:"items"`
		Quests []struct {
			ID              int         `json:"id"`
			RankID          int         `json:"rank_id"`
			QuestguildID    int         `json:"questguild_id"`
			Star            int         `json:"star"`
			QuesttypeID     int         `json:"questtype_id"`
			Reward          int         `json:"reward"`
			QuestpriorityID int         `json:"questpriority_id"`
			QuestgoalID     int         `json:"questgoal_id"`
			QuestsubgoalID  int         `json:"questsubgoal_id"`
			HrpSub          interface{} `json:"hrp_sub"`
			RewardSub       interface{} `json:"reward_sub"`
			Fee             int         `json:"fee"`
			Time            int         `json:"time"`
			Hrp             int         `json:"hrp"`
			MapID           int         `json:"map_id"`
			MapTime         interface{} `json:"map_time"`
			HrpRequired     interface{} `json:"hrp_required"`
			LocalName       string      `json:"local_name"`
			Link            string      `json:"link"`
			Pivot           struct {
				MonsterID int `json:"monster_id"`
				QuestID   int `json:"quest_id"`
			} `json:"pivot"`
			Questguild struct {
				ID        int    `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"questguild"`
			Questpriority struct {
				ID        int    `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"questpriority"`
			Questgoal struct {
				ID        int    `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"questgoal"`
			Questsubgoal struct {
				ID        int    `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"questsubgoal"`
			Questtype struct {
				ID        int    `json:"id"`
				LocalName string `json:"local_name"`
			} `json:"questtype"`
		} `json:"quests"`
		Maps []struct {
			ID        int    `json:"id"`
			LocalName string `json:"local_name"`
			Link      string `json:"link"`
			Pivot     struct {
				MonsterID int    `json:"monster_id"`
				MapID     int    `json:"map_id"`
				Start     int    `json:"start"`
				Areas     string `json:"areas"`
				Rest      int    `json:"rest"`
			} `json:"pivot"`
		} `json:"maps"`
		Weapons []struct {
			ID             int         `json:"id"`
			WeapontypeID   int         `json:"weapontype_id"`
			WeaponParentID int         `json:"weapon_parent_id"`
			Final          int         `json:"final"`
			Attack         int         `json:"attack"`
			Defense        int         `json:"defense"`
			Affinity       interface{} `json:"affinity"`
			AffinityVirus  interface{} `json:"affinity_virus"`
			Rarity         int         `json:"rarity"`
			Slot           int         `json:"slot"`
			PriceCreate    int         `json:"price_create"`
			PriceUpgrade   int         `json:"price_upgrade"`
			LocalName      string      `json:"local_name"`
			Link           string      `json:"link"`
			Pivot          struct {
				MonsterID int `json:"monster_id"`
				WeaponID  int `json:"weapon_id"`
			} `json:"pivot"`
			Weapontype struct {
				ID          int    `json:"id"`
				Modifier    string `json:"modifier"`
				ArmortypeID int    `json:"armortype_id"`
				Code        string `json:"code"`
				ModelPrefix string `json:"model_prefix"`
				LocalName   string `json:"local_name"`
			} `json:"weapontype"`
		} `json:"weapons"`
	} `json:"monster"`
}
