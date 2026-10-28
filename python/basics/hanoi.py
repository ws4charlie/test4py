#：CSDN 
#：https://blog.csdn.net/hikobe8/article/details/50479669

def move(n, a, buff, c):
    if n == 1:
        print(a,"->",c)
        return

    # move n-1 plates from a to buff using c as buffer
    move(n-1, a, c, buff)

    # move 1 plate from a to c using buff as buffer
    move(1, a, buff, c)

    # move n-1 plates from buff to c using a as buffer
    move(n-1, buff, a, c)

move(3, "a", "b", "c")
exit()