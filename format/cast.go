package v4l2_format

type Caster interface {
	CastFormat() Type
}
