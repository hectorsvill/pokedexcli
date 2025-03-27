package main

type StatsResult struct {
	StatsResult []Stats `json:"stats"`
}

type Stats struct {
	Base_Stat int  `json:"base_stat"`
	Stat      Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	Base_Stat int
}

func (sr StatsResult) getStats() []Stat {
	stats := []Stat{}
	for _, stat := range sr.StatsResult {
		stat.Stat.Base_Stat = stat.Base_Stat
		stats = append(stats, stat.Stat)

	}
	return stats
}

