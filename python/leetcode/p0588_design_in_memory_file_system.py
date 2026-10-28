"""
588. Design In-Memory File System  (Hard, LeetCode Premium)
https://leetcode.com/problems/design-in-memory-file-system/

Design an in-memory file system with this API:

    FileSystem()                    start with an empty file system.
    ls(path)                        if path is a file, return [its name];
                                    if it's a directory, return the names
                                    inside it, sorted lexicographically.
    mkdir(path)                     create the directory path, plus any
                                    missing parent directories.
    addContentToFile(path, content) create the file if it doesn't exist,
                                    otherwise append content to it.
    readContentFromFile(path)       return the file's content.

Example:
    Input:
        ["FileSystem","ls","mkdir","addContentToFile","ls","readContentFromFile"]
        [[],["/"],["/a/b/c"],["/a/b/c/d","hello"],["/"],["/a/b/c/d"]]
    Output:
        [null,[],null,null,["a"],"hello"]

    After the calls the tree looks like this:

        /
        `-- a/
            `-- b/
                `-- c/
                    `-- d        file, content "hello"

Constraints:
    1 <= path.length, filePath.length <= 100
    Paths are absolute: they start with "/" and don't end with "/",
    except the root "/" itself.
    Directory and file names are lowercase letters; names in the same
    directory are unique.
    readContentFromFile and ls are only called on paths that exist.
    1 <= content.length <= 50
    At most 300 calls in total.
"""

import os
from os import path

class FileSystem:
    def ls(self, dir: str):
        if path.isfile(dir):
            return [path.basename(dir)]
        elif path.isdir(dir):
            items = os.listdir(dir)
            items.sort()
            return items
        else:
            print("invalid parameter: ", dir)
            return []

    def mkdir(self, dir: str):
        if path.exists(dir):
            print("dir already exists: ", dir)
        else:
            os.makedirs(dir, exist_ok=True)

    def addContentToFile(self, file, content):
        if path.isfile(file):
            with open(file, "a") as f:
                f.write(content + "\n")
        else:
            with open(file, "w") as f:
                f.write(content + "\n")

    def readContentFromFile(self, file):
        if path.isfile(file):
            return open(file).read()
        else:
            return ""

if __name__ == "__main__":
    fs = FileSystem()
    res = fs.ls("c:/Charlie")
    print(res)

    res = fs.ls("c:/Charlie/test.txt")
    print(res)

    dir = "c:/Charlie/test/case1"
    fs.mkdir(dir)

    file = dir + "/sample1.txt"
    fs.addContentToFile(file, "point1: 0, 100")

    fs.addContentToFile(file, "point1: 1, 123")

    res = fs.readContentFromFile(file)
    print(res)
