# Issues-commenter

This app will send a post request to a git issues endpoint with the contents of a specified file as the body.

The original aim for this guy was to post the results of `go test` that would be saved to some file to a git issues api endpoint.

The following env vars must be present:

- TOKEN: the generated api token for your bot/user to be able to post request
- REPORT_PATH: path to the go test report file
- ISSUE_URL: url to the specific issues (pr) you try to post the comment to (check your api docs, could be something like "https://git.example.com/api/v1/repos/example-owner/example-project/issues/69/comments"
