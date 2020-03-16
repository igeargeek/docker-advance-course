# ประเภทของ Docker Network
โดยปกติแล้ว Docker จะสามารถแบ่ง Network ออกได้เป็น 5 ประเภทประกอบด้วย (ไม่รวม custom และพวก 3rd party)

 - None คือการรัน container ในระบบปิด ที่ไม่สามารถออก network ภายนอกได้ ปกติเรามักสร้าง None network ไว้เชื่อมกับ custom network driver ของเราเอง
 - Bridge คือการรัน container ที่สามารถเชื่อมต่อไปยัง host network หรือถ้าพูดง่ายๆก็คือ าสามารถออก internet ได้ โดยถ้าเราสั่งรัน container โดยไม่ระบุ network driver ก็จะได้ Bridge มาเป็น default network เสมอ ซึ่ง Bridge จะมี Subnet และ Gateway เป็นของตัวเอง
 - Host คือการรัน container บน host network เลย จะไม่มีการแยก network ระหว่าง container และ docker host เหมือนใน Bridge
 - Overlay คือการรัน container ที่มีการคุยกันข้าม docker daemons (host network) ถ้ามีโจทย์แนวๆการทำ distributed network และต้องการให้ container สามารถคุยกันได้ บน host ที่ต่างกัน Overlay ก็จะเป็น network อีกประเภทหนึ่ง ที่เราควรศึกษาเพิ่มดู
 - MacVlan คือการรัน container ที่สามารถทำให้ container วิ่งไปยัง router ที่ต้องการได้ โดยไม่ต้องผ่าน host network


![alt text](https://www.ostechnix.com/wp-content/uploads/2017/01/Portainer-Chromium_019.png)

เมื่อเราแยก บาง Container ออกจาก stack แล้วยังต้องการให้ Container ที่แยกออกยังสามารถติดต่อกับ Container เดิมได้

wordpress.stack
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
     networks:
      - network-public
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
     networks:
       - network-public
networks:
  network-public:
    external: true
volumes:
    db_data:
```


pma.stack
```
version: '3'
services:
   
   pma:
     image: phpmyadmin/phpmyadmin
     ports:
       - 8088:80
     restart: always
     environment:
       PMA_HOST: db
       PMA_PORT: 3306
     networks:
       - network-public
networks:
  network-public:
    external: true
```
