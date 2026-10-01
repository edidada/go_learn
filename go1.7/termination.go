package main

// returnBeforeTrailingEmptyStatement is valid from Go 1.7.
//
// The second explicit semicolon after the if statement is a trailing empty
// statement; the first terminates the if statement itself.
// Go 1.7 determines termination from the final non-empty statement, so this
// function is known to return even though the statement list ends with `;`.
func returnBeforeTrailingEmptyStatement() int {
	if true {
		return 7
	} else {
		return 0
	}; ;
}
