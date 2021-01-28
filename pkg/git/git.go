/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package git

import (
	"fmt"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta1"

	"github.com/fluxcd/source-controller/pkg/git/common"
	"github.com/fluxcd/source-controller/pkg/git/git2go"
	"github.com/fluxcd/source-controller/pkg/git/go-git"
)

func CheckoutStrategyForRef(ref *sourcev1.GitRepositoryRef, gitImplementation string) (common.CheckoutStrategy, error) {
	switch gitImplementation {
	case sourcev1.GoGitImplementation:
		return go_git.CheckoutStrategyForRef(ref), nil
	case sourcev1.LibGit2Implementation:
		return git2go.CheckoutStrategyForRef(ref), nil
	default:
		return nil, fmt.Errorf("invalid git implementation %s", gitImplementation)
	}
}

func AuthSecretStrategyForURL(url string, gitImplementation string) (common.AuthSecretStrategy, error) {
	switch gitImplementation {
	case sourcev1.GoGitImplementation:
		return go_git.AuthSecretStrategyForURL(url)
	case sourcev1.LibGit2Implementation:
		return git2go.AuthSecretStrategyForURL(url)
	default:
		return nil, fmt.Errorf("invalid git implementation %s", gitImplementation)
	}
}
