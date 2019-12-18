package pkg

func ElectionPostChallengeCount(sectors uint64) uint64 {
	if sectors == 0 {
		return 0
	}
	// ceil(sectors / SectorChallengeRatioDiv)
	return (sectors-1)/SectorChallengeRatioDiv + 1
}
