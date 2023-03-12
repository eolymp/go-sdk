package ern

// InvalidErr is triggered when ERN has correct format but points to an unknown resource
type InvalidErr struct{}

func (InvalidErr) Error() string {
	return "invalid ERN"
}

// MalformedErr is triggered when string does not follow ERN format
type MalformedErr struct{}

func (MalformedErr) Error() string {
	return "invalid ERN format"
}
