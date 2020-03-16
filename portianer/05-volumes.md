![alt text](https://www.ostechnix.com/wp-content/uploads/2017/01/Portainer-Chromium_020.png)

![alt text](https://sysadmin.psu.ac.th/wp-content/uploads/2017/09/docker-types-of-mounts.png)

Docker ให้เราสามารถเลือกใช้วิธีการ mount data เข้าไปให้กับ container อยู่ 3 อย่างคือ
1. Volumes
2. Bind mounts
3. tmpfs mounts

Volumes จะถูกเก็บอยู่ในส่วนของ Host filesystem ที่จัดการโดย Docker เอง (อยู่ที่ /var/lib/docker/volumes)
และนี่เป็นวิธีที่ดีที่สุดในการจัดเก็บข้อมูลที่เป็น persistent data (ตามคำบอกในเว็บเพจ docs.docker.com)

Bind mounts จะถูกเก็บอยู่ในที่ไหนก็ได้ของ Host filesystem เป็นวิธีการที่มีมาตั้งแต่ Docker รุ่นแรก ๆ จึงมีข้อจำกัดเมื่อเทียบกับ Volumes

tmpfs mounts จะถูกเก็บอยู่ในหน่วยความจำของ Host เท่านั้น

```
version: '2'
services:
 openldap:
 image: openldap
 container_name: openldap
 volumes:
  - ldapdatavol:/var/lib/ldap # Volumes
  - ldapconfigvol:/etc/ldap/slapd.d # Volumes
  - /root/prometheus.yml:/etc/prometheus/prometheus.yml # Bind mounts
 ports:
  - "389:389"
  - "636:636"

volumes:
 ldapdatavol:
  external: false
 ldapconfigvol:
  external: false
```

อธิบายได้ดังนี้ ในไฟล์ นี้ เราจะรัน services ชื่อ openldap จาก image ที่สร้างไว้แล้วชื่อว่า openldap โดยรันเป็น container ที่ผมตั้งชื่อว่า openldap โดยจะเก็บข้อมูลไว้ถาวรที่ volume ชื่อ ldapdatavol ซึ่งจะ mapped กับ /var/lib/ldap ใน container และอีกบรรทัดคือ ldapconfigvol จะ mapped กับ /etc/ldap/slapd.d ใน container

ถัดมาด้านล่างของไฟล์ เราจะต้องประกาศ volumes ไว้ด้วยว่า ldapdatavol ไม่ได้เป็น volume ที่สร้างไว้อยู่แล้วก่อนการัน docker-stack ด้วยการประกาศค่าว่า external: false เช่นเดียวกับ volume ชื่อ ldapconfigvol

แต่ถ้าใช้ external: true จะหมายถึง docker-stack จะไม่สร้าง volume ให้ นั่นคือ เราได้สร้างไว้ก่อนแล้ว


```
volumes:
      - /root/prometheus.yml:/etc/prometheus/prometheus.yml
```


```
volumes:
      - dev-storage-data:/app/storage
```
