package git

import (
	"fmt"
	"os"

	"github.com/Originate/git-town/src/command"
)

// GetRootDirectory returns the path of the rood directory of the current repository,
// i.e. the directory that contains the ".git" folder.
func GetRootDirectory() string {
	return command.New("git", "rev-parse", "--show-toplevel").Output()
}

// HasConflicts returns whether the local repository currently has unresolved merge conflicts.
func HasConflicts() bool {
	return command.New("git", "status").OutputContainsText("Unmerged paths")
}

// HasOpenChanges returns whether the local repository contains uncommitted changes.
func HasOpenChanges() bool {
	return command.New("git", "status", "--porcelain").Output() != ""
}

// HasShippableChanges returns whether the supplied branch has an changes
// not currently on the main branchName
func HasShippableChanges(branchName string) bool {
	return command.New("git", "diff", GetMainBranch()+".."+branchName).Output() != ""
}

// IsMergeInProgress returns whether the local repository is in the middle of
// an unfinished merge process.
func IsMergeInProgress() bool {
	_, err := os.Stat(fmt.Sprintf("%s/.git/MERGE_HEAD", GetRootDirectory()))
	return err == nil
}

// IsRebaseInProgress returns whether the local repository is in the middle of
// an unfinished rebase process.
func IsRebaseInProgress() bool {
	return command.New("git", "status").OutputContainsText("rebase in progress")
}
