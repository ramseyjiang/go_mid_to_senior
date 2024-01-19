package robotparts

import "strings"

func GetRobotParts(allParts []string, requiredParts string) []string {
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
