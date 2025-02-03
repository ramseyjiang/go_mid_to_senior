package filterrooms

import "sort"

func FilterRooms(treasureRooms []string, instructions [][]string) []string {
	// Match the source and dest
	sourceToDest := make(map[string]string)
	for _, inst := range instructions {
		source, dest := inst[0], inst[1]
		sourceToDest[source] = dest
	}

	// record the source room is pointed by which rooms
	inSources := make(map[string]map[string]bool)
	for _, inst := range instructions {
		source, dest := inst[0], inst[1]
		if source == dest {
			continue
		}
		if _, ok := inSources[dest]; !ok {
			inSources[dest] = make(map[string]bool)
		}
		inSources[dest][source] = true
	}

	// Collect all room names
	rooms := make(map[string]bool)
	for _, room := range instructions {
		rooms[room[0]] = true
		rooms[room[1]] = true
	}

	// Collect all rooms have treasure.
	treasureMap := make(map[string]bool)
	for _, t := range treasureRooms {
		treasureMap[t] = true
	}

	result := []string{}
	for room := range rooms {
		// Check the condition 2, this room's own instruction must point to a treasure room.
		dest, ok := sourceToDest[room]
		if !ok {
			continue
		}

		// Check the condition 1, at least two *other* rooms must have instructions pointing to this room.
		if sources, ok := inSources[room]; !ok || len(sources) < 2 {
			continue
		}

		if treasureMap[dest] {
			result = append(result, room)
		}
	}

	sort.Strings(result)
	return result
}

// FilterRooms2 returns all rooms that satisfy:
// 1. At least two *other* rooms have instructions pointing to this room.
// 2. This room's instruction immediately points to a treasure room.
func FilterRooms2(treasureRooms []string, instructions [][]string) []string {
	// Build a set for treasure rooms.
	treasureSet := make(map[string]struct{})
	for _, room := range treasureRooms {
		treasureSet[room] = struct{}{}
	}

	// Build mappings:
	// sourceToDest: each room's own instruction (source -> destination)
	// incoming: maps each room to the set of rooms that point to it.
	sourceToDest := make(map[string]string)
	incoming := make(map[string]map[string]bool)
	for _, pair := range instructions {
		if len(pair) != 2 {
			continue
		}
		src, dest := pair[0], pair[1]
		sourceToDest[src] = dest

		if incoming[dest] == nil {
			incoming[dest] = make(map[string]bool)
		}
		incoming[dest][src] = true
	}

	// Initialize result as an empty (non-nil) slice.
	// var result []string // Don't use this way, it declared the result slice as a nil slice
	result := []string{}

	// For each room that issues an instruction, check the conditions.
	for room, dest := range sourceToDest {
		// Condition 2: This room's own instruction must point to a treasure room.
		if _, ok := treasureSet[dest]; !ok {
			continue
		}

		// Condition 1: At least two *other* rooms must have instructions pointing to this room.
		sources, ok := incoming[room]
		if !ok {
			continue
		}

		count := 0
		for src := range sources {
			// check the source and dest are the same or not. It is used to check point to itself or not.
			if src != room {
				count++
			}
		}
		if count >= 2 {
			result = append(result, room)
		}
	}

	sort.Strings(result)
	return result
}
