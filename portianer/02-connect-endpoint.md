เรามาทำการเชื่อม Endpoints ต่างๆ กัน

เราได้ทำการติดตั้ง การติดตั้ง Docker Engine และทำการเปิดใช้งาน Docker Remote api ไว้เรียบร้อยแล้ว ซึ่งสามารถเชื่อมต่อโดยใช้ private ip เพื่อความปลอดภัยควรให้ docker remote api เป็น private
จากตัวอย่าง ได้ทำการเตรียม เครื่องไว้ทั้งหมด 3 เครื่อง

![alt text](https://miro.medium.com/max/3040/1*F-LTYT1Ladv52vOicz70NQ.png)

เข้ามาที่ Endpoints เลือก Docker environment
![alt text](https://miro.medium.com/max/5512/1*tbSt2yiFiNkX401gjJt-Ow.png)

ทำการใส่ endpoint ด้วย port 4243 อ้างอิงจากวิธีการ เปิดใช้งาน Docker Remote api และ Public IP ด้วย IP ของ เครื่อง จากนั้นกด Add endpoint
https://medium.com/@krutpong/%E0%B9%80%E0%B8%9B%E0%B8%B4%E0%B8%94%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%87%E0%B8%B2%E0%B8%99-docker-remote-api-%E0%B8%9A%E0%B8%99-ubuntu-16-04-af93f03b4e05
![alt text](https://miro.medium.com/max/4376/1*08SkA4JiZsGed1aFsOuB3w.png)

จะพบว่า มีเครื่อง demo-docker-api-1 ที่เรา add remote เข้าไปปรากฎขึ้นมา ทำการ add เครื่องจนครบที่เรามี

![alt text](https://miro.medium.com/max/4476/1*K68Rg3lLjiAR6kRPehM7vw.png)
