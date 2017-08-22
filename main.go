/*******************************************************************************
*
* Copyright 2017 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import (
	"fmt"
	"os"
)

func main() {
	err := _main()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func _main() error {
	cfg, err := readConfig()
	if err != nil {
		return err
	}

	for _, src := range cfg.Sources {
		err := src.discoverPackages()
		if err != nil {
			return err
		}

		fmt.Printf("source: %s\n", src.Path)
		for _, pkg := range src.Packages {
			fmt.Printf("%#v\n", pkg)
		}
	}

	_ = cfg
	return nil
}
