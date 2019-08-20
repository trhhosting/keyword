package gogit

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
)

// Basic example of how to clone a repository using clone options.
func Clone(url, directory string) (err error) {
	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
	return
}
