package caller

import "runtime"

func services() string{
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}

	return ""
}

func Controller() string {
	return services()
}
