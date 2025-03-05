package my_sentinel_err

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}
