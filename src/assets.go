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
