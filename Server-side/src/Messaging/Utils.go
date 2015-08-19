package Messaging

var DebugMode bool = true

func log(message string) {
	if DebugMode {
		println(message)
	}
}

// -- Error Handling

// Returns true if no errors
// Each handler function that is passed in will be run in order
// If one handler fails, then the subsequent handlers are not run
var verbose bool = true

func checkForError(err error, handlers ...func(error) bool) bool {
	if err != nil {
		for _, handler := range handlers {
			if !handler(err) {
				if verbose {
					println(err.Error())
				}

				return false
			}
		}
	}

	return true
}

func warningGenerator(msg string) func(error) bool {
	return func(err error) bool {
		log(msg)
		return true
	}
}

func errorGenerator(msg string) func(error) bool {
	return func(err error) bool {
		log(msg)
		return false
	}
}

func assert(err error) bool {
	println("\nERROR: Remocon has experienced an unexpected error. Shutting down...\n")
	panic(err)
	return false
}

func warning(err error) bool {
	println("\nWARNING: Remocon has experienced a hiccup. Resuming...\n")
	return false
}
