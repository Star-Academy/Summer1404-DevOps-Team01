#!/usr/bin/env bash

git fetch --all --quiet

for branch in $(git for-each-ref --format='%(refname:short)' refs/heads/); do
	git checkout --quiet "$branch"
	while IFS= read -r file; do
		echo "Branch: $branch | File: $file"
	done < <(git grep -l "TODO" | cut -d: -f1 | sort -u)
done

git checkout --quiet -
