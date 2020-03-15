#Intro

ตอนนี้ ให้ทดลองเข้า url ของ portainer ด้วย
```
{{$IP_เครื่องนั้นๆ}}:9000
```

จะพบกับหน้าจอที่ให้ตั้งรหัสผ่าน admin สำหรับจัดการ Docker Container ทั้งหมด
![alt text](https://miro.medium.com/max/3304/1*_pnKMxrGYjlLT__l18dtPQ.png)


หลังจากตั้งรหัสผ่านเสร็จจะพบกับหน้าจอนี้ คือการ Connect กับ Docker API เบื้องต้นเราทำการเลือก local ไปก่อนเพื่อเข้าไปจัดการระบบในเครื่องก่อน
![alt text](https://miro.medium.com/max/4640/1*inWlhTKZFKqXMsT-Q1Uq4Q.png)

เราจะได้พบกับหน้าจอจัดการหลักของ portainer
![alt text](https://miro.medium.com/max/5536/1*Ls4ibetf1GhQsJR9lIc6XQ.png)

เมนู Extensions สำรับจัดการ External Authentication กับ Registry Manager
![alt text](https://miro.medium.com/max/5520/1*m_HUG5T0TP3iRmpvoBoQJA.png)

หน้าจอ User สำหรับจัดการ User ที่สามารถเข้ามาใช้งาน Portainer ได้
![alt text](https://miro.medium.com/max/5520/1*m_HUG5T0TP3iRmpvoBoQJA.png)

Endpoints ส่วนนี้คือส่วนที่จะ Add Docker Remote Api ที่เราติดตั้งไว้ที่ Server ต่างๆ ไม่ว่าจะเป็น บน Digital Ocean , AWS EC2 , Google Compute , Etc ก็จะสามารถจัดการได้ที่เดียว
![alt text](https://miro.medium.com/max/5584/1*pS6lhnyDyDTAo9k3uyVYaw.png)

Registries คือส่วนที่ไว้ใช้สำหรับ เก็บ Docker Images ไว้ที่ Server ของเราเอง ซึ่งสามารถที่จะ Build แล้ว Push มาไว้ที่นี่เพื่อความสะดวกในการ Deploy ภายใน Server
![alt text](https://miro.medium.com/max/5540/1*6zdL9F13x1-HCmC77TxvyQ.png)

ส่วนของ Setting จะเป็นสำหรับ จัดการ Config ต่างๆ ไม่ว่าจะเป็น host name การทำ Snapshot ให้กับ Containner เป็นต้น
![alt text](https://miro.medium.com/max/5512/1*CexU6IpyF9V820HU69dvvw.png)
