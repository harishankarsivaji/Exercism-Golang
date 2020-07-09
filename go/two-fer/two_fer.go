package twofer

// ShareWith returns with message and name.
func ShareWith(name string) string {

	switch {
	case name == "":
		return "One for you, one for me."
	default:
		return "One for " + name + ", one for me."

	}

}
