package version

type Upgrade func(v Version)

func EasyUpgrade(v Version) {
	switch v.(type) {
	case *Simple:
		v.(*Simple).Release()
	case *Complex:
		v.(*Complex).ReleasePatch()
	}
}
