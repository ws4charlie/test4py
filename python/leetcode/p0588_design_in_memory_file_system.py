
# http://lixinchengdu.github.io/algorithmbook/leetcode/design-in-memory-file-system.html

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
