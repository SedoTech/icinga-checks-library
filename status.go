package icinga

// Status defines the host state
type Status int

const (
	// HostStatusUp 0
	HostStatusUp Status = 0
	// HostStatusDown 2
	HostStatusDown Status = 2
)

const (
	// ServiceStatusOk 0
	ServiceStatusOk Status = iota
	// ServiceStatusWarning 1
	ServiceStatusWarning
	// ServiceStatusCritical 2
	ServiceStatusCritical
	// ServiceStatusUnknown 3
	ServiceStatusUnknown
)

// Ordinal returns the int value
func (s Status) Ordinal() int {
	return int(s)
}

func (s Status) String() string {
	names := [...]string{
		"OK",
		"WARNING",
		"CRITICAL",
		"UNKNOWN",
	}
	if s < ServiceStatusOk || s > ServiceStatusUnknown {
		panic("invalid icinga.Status")
	}

	return names[s]
}

// ServiceStatusForEscalationLevel returns a Status for a given EscalationLevel
func ServiceStatusForEscalationLevel(level EscalationLevel) Status {
	switch level {
	case None:
		return ServiceStatusOk
	case Warning:
		return ServiceStatusWarning
	case Critical:
		return ServiceStatusCritical
	default:
		return ServiceStatusUnknown
	}
}

// HostStatusForEscalationLevel returns a Status for a given EscalationLevel
func HostStatusForEscalationLevel(level EscalationLevel) Status {
	switch level {
	case None:
		return HostStatusUp
	case Warning:
		return HostStatusUp
	case Critical:
		return HostStatusDown
	default:
		return HostStatusDown
	}
}