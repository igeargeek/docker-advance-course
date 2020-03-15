ลอง Deploy WordPress ด้วย Stacks แบบง่ายๆ กัน
อันนี้เป็น WordPress Docker Compose เนี้ยแหละครับ โดยกำหนดให้มี services อยู่ 3 ตัว ซึ่งมี database (db) , wordpress และ phpmyadmin (pma) ตามนี้ครับ

```
version: '3'
services:
  db:
     image: mysql:5.7
     volumes:
       - db_data:/var/lib/mysql
     restart: always
     environment:
       MYSQL_ROOT_PASSWORD: ${MYSQL_DATABASE_PASSWORD}
       MYSQL_DATABASE: wpdb
       MYSQL_USER: wpuser
       MYSQL_PASSWORD: wpass
  wordpress:
     image: wordpress:latest
     ports:
       - :80
     restart: always
     environment:
       WORDPRESS_DB_HOST: db:3306
       WORDPRESS_DB_NAME: wpdb
       WORDPRESS_DB_USER: wpuser
       WORDPRESS_DB_PASSWORD: wpass
     depends_on:
       - db
       
   pma:
     image: phpmyadmin/phpmyadmin
     ports:
       - 8088:80
     restart: always
     environment:
       PMA_HOST: db
       PMA_PORT: 3306
       
volumes:
    db_data:
```

จากนั้นเข้าไป Local Docker Engine เลือกเมนู Stacks คลิกปุ่ม +Add stack ดังรูป
![alt text](https://miro.medium.com/max/2032/0*2T5taJLaNJgNqJDt.png)

หน้าต่าง Create stack ให้กำหนดชื่อ (Name) แล้วเลือก Build method เป็น Web editor จากนั้นก็ Copy WordPress Template ข้างต้นมาวางครับ
ใน WordPress Template ผมให้กำหนดค่าของ Environment Variables ด้วย
MYSQL_DATABASE_PASSWORD อันนี้เป็นรหัสผ่านของ root บน MySQL Container นะครับ

![alt text](https://miro.medium.com/max/1186/0*vbxgu3jIWLVSG6_V.png)

เสร็จก็กดปุ่ม Deploy the stack โลดครับ เจ้าตัว Portainer จะทำการสร้าง stack ตามที่เรากำหนดจาก Template เมื่อเสร็จแล้วก็กลับไปที่หน้า Stacks อีกครั้ง
เมื่อเสร็จแล้วก็จะมี stack ที่เราสร้างแสดงขึ้นมา ลองคลิกเข้าไปที่ wordpress1 ดูครับ
![alt text](https://miro.medium.com/max/2048/0*eqzEmvKYtWwoPZc3.png)

เราจะเห็นว่ามี services หรือ container ด้วยกัน 3 ตัวที่มีสถานะ Running อยู่
![alt text](https://miro.medium.com/max/1684/0*3mj3HVOZxyKaCvL5.png)

ลองเปิด Browser เข้าไปที่ WordPress container ตามที่ port ที่ระบบ mapping ไว้ดูครับ
![alt text](https://miro.medium.com/max/2034/0*KllMd9d7PhZOuHFK.png)

จากนั้นเลือกเข้า phpmyadmin แล้ว login
![alt text](https://miro.medium.com/max/2048/0*pSl-KP9oVxHUlIb5.png)
