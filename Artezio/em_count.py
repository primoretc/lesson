# -*- coding: utf-8 -*-
# посчитаем число е-маил в рассылке, а следовательно и число зареганых пользователей
# выдерну адреса заключенные в <> и сохраню в новый файл
email = 0
ef = open('email_list.txt', 'r') # открываю файл исходник, полученый копипастой из заголовка письма To:
mailFile = open("out_email.txt", "w") # открываю файл в который буду собирать адреса

# бегу по строчкам файла 
for line in ef:
    
    s1 = line.find("<")
    s2 = line.rfind(">")
    adress = line[s1+1:s2] 
    
    mailFile.write(adress + "\n")
    # бегу по символам строки...
    for simvol in line:
        if simvol == ",":
            email += 1
            
# всё закрываю
ef.close()
mailFile.close()

print(f"Число зарегавшихся: {email}")

