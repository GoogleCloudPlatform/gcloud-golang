module cloud.google.com/go/internal/gapicgen

go 1.13

require (
	cloud.google.com/go/internal/gensnippets v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-github/v34 v34.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/shurcooL/githubv4 v0.0.0-20201206200315-234843c633fa
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	golang.org/x/oauth2 v0.0.0-20210413134643-5e61552d6c78
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.4.0
)

replace cloud.google.com/go/internal/gensnippets => ../gensnippets

replace cloud.google.com/go/internal/godocfx => ../godocfx
