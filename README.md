![Proximu](https://capsule-render.vercel.app/api?type=cylinder&height=300&color=gradient&text=Proximu&textBg=false)

A lightweight UDP proxy for **Minecraft: Bedrock Edition**, that protects your server from direct attacks (e.g. DDoS or IP sniffing).

---

## 📦 Requirements

- [Go (1.22+)](https://go.dev/dl/)
- [BDS](https://www.minecraft.net/en-us/download/server/bedrock)

---

## 🧰 what needs to be added

| what needs to be done | made |
|------|:------:|
| protection against DDoS attacks | ❌ |
| filtering by IP | ❌ |
| config for Proximu | ❌ |
| active sessions log | ❌ |
| port forwarding | ✅ |

---

## 🚀 Usage 

the proxy itself is still raw, there are many bugs and problems, for now it works on local ones, for global ones there are none yet.

open `server.properties`

and change the value of these parameters
`server-port`,
`server-portv6`  to 19133
save

then start the proxy and the server, and the server itself should be on `127.0.0.1:19133`, and after starting the proxy it will open on port ``:19132``

and you can go through the proxy like this `127.0.0.1:19132`
or another way
