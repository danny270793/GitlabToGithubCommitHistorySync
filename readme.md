# GitlabToGithubCommitHistorySync

[![GitlabToGithubCommitHistorySync](https://github.com/danny270793/GitlabToGithubCommitHistorySync/actions/workflows/releaser.yaml/badge.svg)](https://github.com/danny270793/GitlabToGithubCommitHistorySync/actions/workflows/release.yaml)

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
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/danny270793)

## LICENSE

Licensed under the [MIT](license.md) License

## Version

GitlabToGithubCommitHistorySync version 1.0.0

Last update 29/07/2024
