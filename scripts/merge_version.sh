#!/bin/sh

git log -1 --pretty=format:%s | sed "s@Merge branch 'version/@@" | sed "s@'@@"
