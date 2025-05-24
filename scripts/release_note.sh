#!/bin/sh

DEST=ReleaseNotes.md
TEMP=temp_release_notes.md
FORMAT="- %s ([%h](%H))"
VERSION=$(git branch --show-current | sed 's@qa/@@')
SINCE="$(git log -1 --pretty=format:%h -- $DEST)"

echo "VERSION: $VERSION"
echo "SINCE: $SINCE"

mv $DEST $TEMP

echo -e "# Ke\n
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
