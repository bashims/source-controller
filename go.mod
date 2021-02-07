module github.com/fluxcd/source-controller

go 1.15

replace github.com/fluxcd/source-controller/api => ./api

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/blang/semver/v4 v4.0.0
	github.com/cyphar/filepath-securejoin v0.2.2
	github.com/fluxcd/pkg/apis/meta v0.7.0
	github.com/fluxcd/pkg/gittestserver v0.1.0
	github.com/fluxcd/pkg/helmtestserver v0.1.0
	github.com/fluxcd/pkg/lockedfile v0.0.5
	github.com/fluxcd/pkg/runtime v0.8.1
	github.com/fluxcd/pkg/ssh v0.0.5
	github.com/fluxcd/pkg/untar v0.0.5
	github.com/fluxcd/pkg/version v0.0.1
	github.com/fluxcd/source-controller/api v0.7.4
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/go-logr/logr v0.3.0
	github.com/libgit2/git2go/v31 v31.4.7
	github.com/minio/minio-go/v7 v7.0.5
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	helm.sh/helm/v3 v3.5.0
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.0
	sigs.k8s.io/yaml v1.2.0
)
