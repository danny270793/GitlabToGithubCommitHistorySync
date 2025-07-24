# GitlabToGithubCommitHistorySync

[![GitlabToGithubCommitHistorySync](https://github.com/danny270793/GitlabToGithubCommitHistorySync/actions/workflows/releaser.yaml/badge.svg)](https://github.com/danny270793/GitlabToGithubCommitHistorySync/actions/workflows/release.yaml)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/danny270793/GitlabToGithubCommitHistorySync)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/danny270793/GitlabToGithubCommitHistorySync/total)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/danny270793/GitlabToGithubCommitHistorySync)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/danny270793/GitlabToGithubCommitHistorySync)

Make commits on a github private repo to sync the commit histpry from all gitlab ussing the events api

### From terminal

Get the last version available on github

```bash
LAST_VERSION=$(curl https://api.github.com/repos/danny270793/GitlabToGithubCommitHistorySync/releases/latest | grep tag_name | cut -d '"' -f 4)
```

Download the last version directly to the binaries folder

For Linux (linux):

```bash
curl -L https://github.com/danny270793/GoHotReloader/releases/download/${LAST_VERSION}/GitlabToGithubCommitHistorySync_${LAST_VERSION}_linux_amd64.tar.gz -o ./GitlabToGithubCommitHistorySync.tar.gz
```

Untar the downloaded file

```bash
tar -xvf ./GitlabToGithubCommitHistorySync.tar.gz
```

Then copy the binary to the binaries folder

```bash
sudo cp ./GitlabToGithubCommitHistorySync /usr/local/bin/GitlabToGithubCommitHistorySync
```

Make it executable the binary

```bash
sudo chmod +x /usr/local/bin/GitlabToGithubCommitHistorySync
```

```bash
GitlabToGithubCommitHistorySync --version
```

## Ussage

Run the binary in the same folder than a `.env` file which must contains the following

```conf
GITLAB_USERID=
GITLAB_ACCESS_TOKEN=

GITHUB_USERNAME=
GITHUB_REPOSITORY=
GITHUB_ACCESS_TOKEN=

SYNC_START_DATE=
SYNC_END_DATE=
```

```bash
GitlabToGithubCommitHistorySync
```

## Follow me

[![YouTube](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white)](https://www.youtube.com/channel/UC5MAQWU2s2VESTXaUo-ysgg)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://www.github.com/danny270793/)
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2aWV3Qm94PSIwLDAsMjU2LDI1NiIgd2lkdGg9IjUwcHgiIGhlaWdodD0iNTBweCIgZmlsbC1ydWxlPSJub256ZXJvIj48ZyBmaWxsPSIjZmZmZmZmIiBmaWxsLXJ1bGU9Im5vbnplcm8iIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBzdHJva2UtbGluZWNhcD0iYnV0dCIgc3Ryb2tlLWxpbmVqb2luPSJtaXRlciIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBzdHJva2UtZGFzaGFycmF5PSIiIHN0cm9rZS1kYXNob2Zmc2V0PSIwIiBmb250LWZhbWlseT0ibm9uZSIgZm9udC13ZWlnaHQ9Im5vbmUiIGZvbnQtc2l6ZT0ibm9uZSIgdGV4dC1hbmNob3I9Im5vbmUiIHN0eWxlPSJtaXgtYmxlbmQtbW9kZTogbm9ybWFsIj48ZyB0cmFuc2Zvcm09InNjYWxlKDUuMTIsNS4xMikiPjxwYXRoIGQ9Ik00MSw0aC0zMmMtMi43NiwwIC01LDIuMjQgLTUsNXYzMmMwLDIuNzYgMi4yNCw1IDUsNWgzMmMyLjc2LDAgNSwtMi4yNCA1LC01di0zMmMwLC0yLjc2IC0yLjI0LC01IC01LC01ek0xNywyMHYxOWgtNnYtMTl6TTExLDE0LjQ3YzAsLTEuNCAxLjIsLTIuNDcgMywtMi40N2MxLjgsMCAyLjkzLDEuMDcgMywyLjQ3YzAsMS40IC0xLjEyLDIuNTMgLTMsMi41M2MtMS44LDAgLTMsLTEuMTMgLTMsLTIuNTN6TTM5LDM5aC02YzAsMCAwLC05LjI2IDAsLTEwYzAsLTIgLTEsLTQgLTMuNSwtNC4wNGgtMC4wOGMtMi40MiwwIC0zLjQyLDIuMDYgLTMuNDIsNC4wNGMwLDAuOTEgMCwxMCAwLDEwaC02di0xOWg2djIuNTZjMCwwIDEuOTMsLTIuNTYgNS44MSwtMi41NmMzLjk3LDAgNy4xOSwyLjczIDcuMTksOC4yNnoiPjwvcGF0aD48L2c+PC9nPjwvc3ZnPg==&logoColor=white&style=for-the-badge)](https://www.linkedin.com/in/danny270793)

## LICENSE

[![GitHub License](https://img.shields.io/github/license/danny270793/GitlabToGithubCommitHistorySync)
](license.md)

## Version

![GitHub Tag](https://img.shields.io/github/v/tag/danny270793/GitlabToGithubCommitHistorySync)

Last update 29/07/2024
