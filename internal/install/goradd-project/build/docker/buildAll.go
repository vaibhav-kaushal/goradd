package docker

//go:generate gofile mkdir goradd-project/../deploy
//go:generate gofile remove goradd-project/../deploy/docker
//go:generate gofile mkdir goradd-project/../deploy/docker
//go:generate go generate ./buildLinux.go
//go:generate go generate ./makeAssets.go
//go:generate ./zipAssets.sh
//go:generate ./buildContainer.sh
