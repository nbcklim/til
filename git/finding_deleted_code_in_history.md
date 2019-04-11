# Finding Deleted Code in Git History

## `git log -S`

The docs: https://git-scm.com/docs/git-log#Documentation/git-log.txt--Sltstringgt

Basically, you can use this command to find all instances of a string in your repo's history. 

Specifically, you will search all *changes* for this string.

## Searching this repo (not much to choose from right now)
```
$ git log -S Today
commit edb9e7af58a8c84724dc548ffab563b1ec98d6ee 
Author: nbcklim <kevin.lim@nbcuni.com>
Date:   Mon Apr 8 11:33:58 2019 -0700

    Initial commit
```

## the string `Today` was added to the code as of that commit
```
$ git show edb9e7af58a8c84724dc548ffab563b1ec98d6ee
commit edb9e7af58a8c84724dc548ffab563b1ec98d6ee
Author: nbcklim <kevin.lim@nbcuni.com>
Date:   Mon Apr 8 11:33:58 2019 -0700

    Initial commit

diff --git a/README.md b/README.md
new file mode 100644
index 0000000..d233451
--- /dev/null
+++ b/README.md
@@ -0,0 +1,2 @@
+# til
+Today I learned
```

## Commit messages are NOT searched

`$ git log -S "Initial commit"` doesn't yield any results

via https://blog.danlew.net/2019/02/19/git-log-s/