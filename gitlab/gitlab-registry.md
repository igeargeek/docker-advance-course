### ขั้นตอนการ push Docker Image ขึ้นบน Gitlab Registry
- ให้เราเปิด terminal หรือ command line จากนั้นรันคำสั่ง ด้านล่าง เพื่อทำการ Login เข้าสู่ Gitlab

```
docker login registry.gitlab.com
```

จากนั้น ให้เราเข้าไปที่ Folder “my-project” จากขั้นตอนการสร้าง Docker Image ด้านบน ซึ่งในขั้นตอนนี้เราจะเปลี่ยนชื่อ Docker Image ของเราจาก my-project ไปเป็น registry.gitlab.com/[your-username]/my-project ตามด้านล่าง

```
# แบบเก่าที่ build ในเครื่อง
#docker build -t my-project .
# เปลี่ยนเป็นแบบใหม่ 
docker build -t registry.gitlab.com/[your-username]/my-project .
```

หลังจากนั้นก็ทำการ push Docker Image ของเราขึ้น Gitlab ได้เลย โดยใช้คำสั่ง
```
docker push registry.gitlab.com/[your-username]/my-project
```

หลังจากเสร็จเรียบร้อยแล้ว เราก็จะเห็น Docker Image ของเราขึ้นมาอยู่บน Gitlab แล้ว และ เราสามารถ copy ลิงค์ไปใช้ได้ทันที ซึ่งการใช้งาน Docker Image จะสามารถใช้งานได้แค่ Gitlab User ของเราเท่านั้นนะ (ในกรณีที่ project ของเราเป็น private ) ซึ่งวิธีการโหลด Docker Image ไปใช้ ทำได้โดยใช้คำสั่ง
```
docker pull registry.gitlab.com/[your-username]/my-project
```

แค่นี้เราก็มีที่เก็บ Source Code และ ที่เก็บ Docker Image ของเราแล้ว


### สรุป command ต่างๆ ที่ใช้ในบทนี้

Docker Command
- docker login registry.gitlab.com — ล็อกอินเข้าสู่ Gitlab Registry
- docker build -t [project-name] . — สร้าง Docker Image
- docker pull registry.gitlab.com/[your-username]/my-project — ดึง Docker Image
- docker push registry.gitlab.com/[your-username]/my-project — อัพโหลด Docker Image ไปยัง Gitlab Registry
