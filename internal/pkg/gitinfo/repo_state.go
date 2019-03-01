package gitinfo

// These names mostly match gitaction in vcs_info:
// https://github.com/zsh-users/zsh/blob/master/Functions/VCS_Info/Backends/VCS_INFO_get_data_git.
func RepoState() {
	// switch (git_repository_state(repo)) {
	//   case GIT_REPOSITORY_STATE_NONE:
	//     return "";
	//   case GIT_REPOSITORY_STATE_MERGE:
	//     return "merge";
	//   case GIT_REPOSITORY_STATE_REVERT:
	//     return "revert";
	//   case GIT_REPOSITORY_STATE_REVERT_SEQUENCE:
	//     return "revert-seq";
	//   case GIT_REPOSITORY_STATE_CHERRYPICK:
	//     return "cherry";
	//   case GIT_REPOSITORY_STATE_CHERRYPICK_SEQUENCE:
	//     return "cherry-seq";
	//   case GIT_REPOSITORY_STATE_BISECT:
	//     return "bisect";
	//   case GIT_REPOSITORY_STATE_REBASE:
	//     return "rebase";
	//   case GIT_REPOSITORY_STATE_REBASE_INTERACTIVE:
	//     return "rebase-i";
	//   case GIT_REPOSITORY_STATE_REBASE_MERGE:
	//     return "rebase-m";
	//   case GIT_REPOSITORY_STATE_APPLY_MAILBOX:
	//     return "am";
	//   case GIT_REPOSITORY_STATE_APPLY_MAILBOX_OR_REBASE:
	//     return "am/rebase";
	// }
	// return "action";
}
