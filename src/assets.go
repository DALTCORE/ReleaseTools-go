package main

/***********
 * Merge request template
 */
func assetTemplateMergeRequestMd() string {
	return "* [ ] Checkout the branch for this merge request\n" +
		"* [ ] Download `composer global require daltcore/release-tools`\n" +
		"* [ ] Navigate in git bash to your project root directory\n" +
		"* [ ] Run `release-tool init\n" +
		"* [ ] Run `release-tool changelog\n" +
		"* [ ] Write awesome code\n"
}

/***********
 * Merge request template
 */
func assetTemplateBugMd() string {
	return "Steps to reproduce\n" +
		"---\n" +
		"\n" +
		"1.\n" +
		"2.\n" +
		"\n" +
		"Results\n" +
		"---\n" +
		"\n" +
		"Describe the bug.\n" +
		"" +
		"Expected Results\n" +
		"---\n" +
		"\n" +
		"Describe the expected behavior.\n" +
		"\n" +
		"Workaround\n" +
		"---\n" +
		"\n" +
		"Describe a possible workaround.\n" +
		"\n" +
		"Attachments\n" +
		"---\n" +
		"\n" +
		"Add a screen shot or logfile that shows the bug.\n"
}

func assetPrepareStub() string {
	return "**Release `:repo` version `:version`**  " +
		"- [X] Create issue  " +
		"- [ ] Notify in Mattermost `rt p notify-upcoming-release :version`  " +
		"- [ ] Merge request *develop > releases/v:version* `rt p develop-to-release :version`  " +
		"- [ ] Checkout releases/v:version `git fetch --all; git checkout releases/v:version`  " +
		"- [ ] Generate changelog `rt mc :version`  " +
		"- [ ] Add changelog to git `git add .; git commit; git push`  " +
		"- [ ] Create merge request *releases/v:version > staging* `rt p release-to-staging :version`  " +
		"- [ ] Wait for merge request *releases/v:version > staging* to be merged  " +
		"- [ ] Notify in Mattermost `rt p notify-accept :version`  " +
		// "- [ ] Merge request *staging > develop* `rt p staging-to-develop :version`  " +
		// "- [ ] Wait for merge request *staging > develop* to be merged  " +
		"- [ ] Merge request *staging > master* `rt p staging-to-master :version`  " +
		"- [ ] Wait for merge request *staging > master* to be merged  " +
		"- [ ] Create tag v:version `rt p create-tag :version`  " +
		"- [ ] Merge request *master > develop* `rt p master-backport`  " +
		"- [ ] Wait for merge request *master > develop* to be merged  " +
		"- [ ] Notify in Mattermost `rt p release-done`  " +
		"- [ ] Close issue  "
}
