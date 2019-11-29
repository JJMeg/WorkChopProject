package runmode

type RunMode string

const (
	Production  = RunMode("production")
	Development = RunMode("development")
	Test        = RunMode("test")
)

func (mode RunMode) IsValid() bool {
	switch mode {
	case Production, Development, Test:
		return true
	default:
		return false
	}
}
