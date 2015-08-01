package main

import "testing"

func TestGetUpdatedAt(t *testing.T) {
	var bodyTest string = `{"id":12425945,"name":"raspicam","full_name":"dhowden/raspicam","owner":{"login":"dhowden","id":2822,"avatar_url":"https://avatars.githubusercontent.com/u/2822?v=2","gravatar_id":"1a7c9eee4f52559c408d8e26073a69a1","url":"https://api.github.com/users/dhowden","html_url":"https://github.com/dhowden","followers_url":"https://api.github.com/users/dhowden/followers","following_url":"https://api.github.com/users/dhowden/following{/other_user}","gists_url":"https://api.github.com/users/dhowden/gists{/gist_id}","starred_url":"https://api.github.com/users/dhowden/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/dhowden/subscriptions","organizations_url":"https://api.github.com/users/dhowden/orgs","repos_url":"https://api.github.com/users/dhowden/repos","events_url":"https://api.github.com/users/dhowden/events{/privacy}","received_events_url":"https://api.github.com/users/dhowden/received_events","type":"User","site_admin":false},"private":false,"html_url":"https://github.com/dhowden/raspicam","description":"","fork":false,"url":"https://api.github.com/repos/dhowden/raspicam","forks_url":"https://api.github.com/repos/dhowden/raspicam/forks","keys_url":"https://api.github.com/repos/dhowden/raspicam/keys{/key_id}","collaborators_url":"https://api.github.com/repos/dhowden/raspicam/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/dhowden/raspicam/teams","hooks_url":"https://api.github.com/repos/dhowden/raspicam/hooks","issue_events_url":"https://api.github.com/repos/dhowden/raspicam/issues/events{/number}","events_url":"https://api.github.com/repos/dhowden/raspicam/events","assignees_url":"https://api.github.com/repos/dhowden/raspicam/assignees{/user}","branches_url":"https://api.github.com/repos/dhowden/raspicam/branches{/branch}","tags_url":"https://api.github.com/repos/dhowden/raspicam/tags","blobs_url":"https://api.github.com/repos/dhowden/raspicam/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/dhowden/raspicam/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/dhowden/raspicam/git/refs{/sha}","trees_url":"https://api.github.com/repos/dhowden/raspicam/git/trees{/sha}","statuses_url":"https://api.github.com/repos/dhowden/raspicam/statuses/{sha}","languages_url":"https://api.github.com/repos/dhowden/raspicam/languages","stargazers_url":"https://api.github.com/repos/dhowden/raspicam/stargazers","contributors_url":"https://api.github.com/repos/dhowden/raspicam/contributors","subscribers_url":"https://api.github.com/repos/dhowden/raspicam/subscribers","subscription_url":"https://api.github.com/repos/dhowden/raspicam/subscription","commits_url":"https://api.github.com/repos/dhowden/raspicam/commits{/sha}","git_commits_url":"https://api.github.com/repos/dhowden/raspicam/git/commits{/sha}","comments_url":"https://api.github.com/repos/dhowden/raspicam/comments{/number}","issue_comment_url":"https://api.github.com/repos/dhowden/raspicam/issues/comments/{number}","contents_url":"https://api.github.com/repos/dhowden/raspicam/contents/{+path}","compare_url":"https://api.github.com/repos/dhowden/raspicam/compare/{base}...{head}","merges_url":"https://api.github.com/repos/dhowden/raspicam/merges","archive_url":"https://api.github.com/repos/dhowden/raspicam/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/dhowden/raspicam/downloads","issues_url":"https://api.github.com/repos/dhowden/raspicam/issues{/number}","pulls_url":"https://api.github.com/repos/dhowden/raspicam/pulls{/number}","milestones_url":"https://api.github.com/repos/dhowden/raspicam/milestones{/number}","notifications_url":"https://api.github.com/repos/dhowden/raspicam/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/dhowden/raspicam/labels{/name}","releases_url":"https://api.github.com/repos/dhowden/raspicam/releases{/id}","created_at":"2013-08-28T06:04:26Z","updated_at":"2014-02-25T12:24:03Z","pushed_at":"2014-02-25T12:24:02Z","git_url":"git://github.com/dhowden/raspicam.git","ssh_url":"git@github.com:dhowden/raspicam.git","clone_url":"https://github.com/dhowden/raspicam.git","svn_url":"https://github.com/dhowden/raspicam","homepage":null,"size":168,"stargazers_count":1,"watchers_count":1,"language":"Go","has_issues":true,"has_downloads":true,"has_wiki":true,"forks_count":0,"mirror_url":null,"open_issues_count":0,"forks":0,"open_issues":0,"watchers":1,"default_branch":"master","network_count":0,"subscribers_count":1}`

	_, err := getUpdatedAt([]byte(bodyTest))
	if err != nil {
		t.Errorf("getUpdatedAt returned an error: %v", err)
	}
}