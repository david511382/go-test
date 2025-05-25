#!/bin/sh

REPO_NAME=go-test
REPO_URL="https://github.com/david511382/$REPO_NAME/commit"
DEST=ReleaseNotes.md
TEMP=temp_release_notes.md
FORMAT="- %s ([%h]($REPO_URL/%H))"
VERSION=$(git branch --show-current | sed 's@version/@@')
SINCE="$(git log -1 --pretty=format:%h -- $DEST)"

echo "VERSION: $VERSION"
echo "SINCE: $SINCE"

mv $DEST $TEMP

echo -e "# $REPO_NAME\n
## $VERSION\n
### Issues\n" >> $DEST
git log --pretty=format:"$FORMAT" --grep="fix" $SINCE..HEAD >> $DEST
echo -e "\n
### Features\n" >> $DEST
git log --pretty=format:"$FORMAT" --grep="\(feat\|cut\)" -i $SINCE..HEAD >> $DEST
echo "" >> $DEST

tail -n +2 $TEMP | while read LINE
do
    echo "$LINE" >> $DEST
done
rm $TEMP

git log --pretty=format:"$FORMAT" --grep="Revert" -i $SINCE..HEAD