package robotparts

import (
	"slices"
	"sort"
	"strings"
)

func GetRobotParts(allParts []string, requiredParts string) (result []string) {
	partNames := strings.Split(requiredParts, ",")
	robotParts := make(map[string]map[string]bool)

	// Populate robotParts map
	for _, parts := range allParts {
		split := strings.Split(parts, "_")
		robot, part := split[0], split[1]

		if robotParts[robot] == nil {
			robotParts[robot] = make(map[string]bool)
		}
		robotParts[robot][part] = true
	}

	// Find robots with all required parts
	for robot, parts := range robotParts {
		selected := true
		for _, name := range partNames {
			if !parts[name] {
				selected = false
				break
			}
		}
		if selected {
			result = append(result, robot)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})

	return result
}

func GetRobotParts1(allParts []string, requiredParts string) (result []string) {
	// Splitting the required parts into a slice
	partsNeeded := strings.Split(requiredParts, ",")

	// Filter requiredParts from allParts, and assign all possible robots into tmpResult 2D string slice
	tmpResults := make([][]string, len(partsNeeded))
	for k, partNeed := range partsNeeded {
		var tmpParts []string
		for _, part := range allParts {
			// Make sure all elements in tmpParts unique
			if strings.Contains(part, partNeed) && !slices.Contains(tmpParts, part) {
				tmpParts = append(tmpParts, part)
			}
			tmpResults[k] = tmpParts
		}
	}

	if len(tmpResults) == 0 {
		return
	}

	// Get all selected robots unique names
	robots := make([][]string, len(tmpResults))
	for k, tmpResult := range tmpResults {
		var tmpRobot []string
		for _, v := range tmpResult {
			tmp := strings.Split(v, "_")
			tmpRobot = append(tmpRobot, tmp[0])
		}
		robots[k] = tmpRobot
	}

	// sorted the robots
	sort.SliceStable(robots, func(i, j int) bool {
		return len(robots[i]) < len(robots[j])
	})

	interaction := make(map[string]int)
	for i := 0; i < len(robots); i++ {
		// The shortest slice in robots is the result.
		if i == 0 && len(robots[i]) < len(robots[i+1]) {
			result = robots[i]
			break
		} else {
			// If two or more length ties, we should get the interaction of all tie results.
			for _, item := range robots[i] {
				interaction[item]++
			}
		}
	}

	// Iterate through the map to find the max value and keys
	maxCount := 0
	for key, count := range interaction {
		if count > maxCount {
			maxCount = count
			result = []string{key} // Start a new slice with the current key
		} else if count == maxCount {
			result = append(result, key) // Add the key to the result
		}
	}

	return result
}

func GetRobotParts2(allParts []string, requiredParts string) []string {
	// Splitting the required parts into a slice
	partsNeeded := strings.Split(requiredParts, ",")

	// Map to hold the parts each robot has
	robotParts := make(map[string]map[string]bool)

	// Iterate over all parts and populate the robot_parts map
	for _, part := range allParts {
		split := strings.Split(part, "_")
		robotName, partName := split[0], split[1]

		if robotParts[robotName] == nil {
			robotParts[robotName] = make(map[string]bool)
		}
		robotParts[robotName][partName] = true
	}

	// List to hold the names of robots that can be completely built
	completeRobots := []string{}

	// Check each robot for required parts
	for robot, parts := range robotParts {
		complete := true
		for _, needed := range partsNeeded {
			if !parts[needed] {
				complete = false
				break
			}
		}
		if complete {
			completeRobots = append(completeRobots, robot)
		}
	}

	return completeRobots
}
