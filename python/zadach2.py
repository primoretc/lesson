"""
Напишите функцию, которая объединяет два списка, чередуя элементы. Например
с учетом двух списков [1,2,3] и ['a','b','c'] функция должна вернуть [1,a,2,b,3,c]

Усложнил. Если списки разной длины, то в более короткий добавляются пустые элементы

"""
a=[1,2,3,4,5]
b=['a','b','c','d','e']

def ListChered(a,b):
    c=[]
    i=0
    if len(a)==len(b):
        li = len(a)
    elif len(a)>len(b):
        n=len(a)-len(b)
        for i in range(n):
            b.append('')    
        li=len(a)
    elif len(a)<len(b):
        n=len(b)-len(a)
        for i in range(n):
            a.append('')    
        li=len(a)
    for i in range(li):
        c.append(a[i])
        c.append(b[i])
        i+=1
    return c


print(ListChered(a,b))