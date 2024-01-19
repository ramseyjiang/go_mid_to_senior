/*
We have a bin of robot parts in a factory. Each part goes to a robot with a specific, unique name. Each part will be
described by a string, with the name of the robot and the part name separated by an underscore, like "Rosie_arm".

All robots are made of the same types of parts, and we have a string of all of the parts required to form a complete
robot. Given a list of available parts, return the collection of robot names for which we can build at least one
complete robot.

Sample Input:

all_parts = [
"Rosie_claw",
"Rosie_sensors",
"Dustie_case",
"Optimus_sensors",
"Rust_sensors",
"Rosie_case",
"Rust_case",
"Optimus_speaker",
"Rosie_wheels",
"Rosie_speaker",
"Dustie_case",
"Dustie_arms",
"Rust_claw",
"Dustie_case",
"Dustie_speaker",
"Optimus_case",
"Optimus_wheels",
"Rust_legs",
"Optimus_sensors" ]

required_parts_1 = "sensors,case,speaker,wheels"
required_parts_2 = "sensors,case,speaker,wheels,claw"
required_parts_3 = "sensors,case,screws"

Expected Output (output can be in any order):

get_robots(all_parts, required_parts_1) => ["Optimus", "Rosie"]
get_robots(all_parts, required_parts_2) => ["Rosie"]
get_robots(all_parts, required_parts_3) => []

N: Number of elements in `all_parts`
P: Number of elements in `required_parts`
*/

package main
import (
"fmt"
"strings"

)

func main() {
required_parts_1 := "sensors,case,speaker,wheels"
required_parts_2 := "sensors,case,speaker,wheels,claw"
required_parts_3 := "sensors,case,screws"

all_parts := []string{
"Rosie_claw",
"Rosie_sensors",
"Dustie_case",
"Optimus_sensors",
"Rust_sensors",
"Rosie_case",
"Rust_case",
"Optimus_speaker",
"Rosie_wheels",
"Rosie_speaker",
"Dustie_case",
"Dustie_arms",
"Rust_claw",
"Dustie_case",
"Dustie_speaker",
"Optimus_case",
"Optimus_wheels",
"Rust_legs",
"Optimus_sensors",
}