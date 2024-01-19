package robotparts

import "strings"

func GetRobotParts(allParts []string, requiredParts string) []string {
	// Splitting the required parts into a slice
	parts_needed := strings.Split(requiredParts, ",")

	// Map to hold the parts each robot has
	robot_parts := make(map[string]map[string]bool)

	// Iterate over all parts and populate the robot_parts map
	for _, part := range allParts {
		split := strings.Split(part, "_")
		robotName, partName := split[0], split[1]

		if robot_parts[robotName] == nil {
			robot_parts[robotName] = make(map[string]bool)
		}
		robot_parts[robotName][partName] = true
	}

	// List to hold the names of robots that can be completely built
	complete_robots := []string{}

	// Check each robot for required parts
	for robot, parts := range robot_parts {
		complete := true
		for _, needed := range parts_needed {
			if !parts[needed] {
				complete = false
				break
			}
		}
		if complete {
			complete_robots = append(complete_robots, robot)
		}
	}

	return complete_robots
}
