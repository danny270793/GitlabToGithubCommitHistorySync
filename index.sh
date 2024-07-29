set -e

DATE=$1
TIME=$(date +"%T")

now=$(date +%s)
then=$(date -d "${DATE}T${TIME}Z" +%s)
SECONDS_DIFFERENCE=$((now - then))
DAYS_DIFFERENCE=$((SECONDS_DIFFERENCE / 86400))

date > .temp
git add .
GIT_COMMITTER_DATE="${DATE}T${TIME}Z" git commit --date "${DAYS_DIFFERENCE} day ago" -m "sync commit"
git push origin master
