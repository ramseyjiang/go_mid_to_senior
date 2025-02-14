package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. Initialize graph (adjacency list) and in-degree array
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 2. Build the graph and update in-degree
	for _, prereq := range prerequisites {
		course := prereq[0]
		preCourse := prereq[1]
		// "preCourse" must come before "course"
		graph[preCourse] = append(graph[preCourse], course)
		// Increase inDegree of 'course' because it has a prerequisite
		inDegree[course]++
	}

	// 3. Find all courses with no prerequisites (inDegree == 0)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 4. Process courses in topological order using a queue (Kahn's Algorithm)
	visited := 0
	for len(queue) > 0 {
		// Take a course that has no remaining prerequisites
		current := queue[0]
		queue = queue[1:]
		visited++

		// 5. Decrease inDegree for all courses dependent on 'current'
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			// If a course no longer has prerequisites, add it to the queue
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return visited == numCourses
}
