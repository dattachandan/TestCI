// Copyright (c) 2015 Chandan Datta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package TestCI

import "errors"

func IsItTestedOrNot(word string) (bool, error) {
	fmt.Printf("Hello, world.\n")
	switch word {
	case "Tested":
		return true, nil
	case "Not":
		err := errors.New("Oh noooooo, it's bar!")
		return false, err
	}
	return false, nil
}
