/*
Copyright © 2022 ANTOINE ODDOZ <antoine.oddoz@protonmail.com>
*/

package main

import (
	"github.com/lafusew/gokeep/cmd"
	"github.com/lafusew/gokeep/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
