---
marp: true
_class: lead
paginate: true
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---

![bg border:50% left:40% 80%](https://img.fruugo.com/product/8/18/140180188_max.jpg)

# Coding Dojo

---

![bg right:30%](https://images.unsplash.com/photo-1496113329550-ce8886d06d54?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Z3ltJTIwamFwYW58ZW58MHx8MHx8&auto=format&fit=crop&w=1080&q=80)

## Dojos

> A dojo (japanese for `place of the Do`)  is a hall or place for immersive learning or meditation
### 🧩 A Coding Dojo is a meeting where a bunch of coders get together to work on a programming challenge. 
### 🧩 They are there to have fun, learn coding and bring development teams together

---

![bg right:30%](https://images.unsplash.com/photo-1555597408-26bc8e548a46?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&w=1080&q=80)

## ⁉️ Katas

> Japanese word describing a detailed pattern of movements practiced either solo or in pairs.

### 🧩 Code Katas are small coding exercises to be solved by programmes to learn and improve their skills

---

![bg right:30%](https://images.unsplash.com/photo-1629721671030-a83edbb11211?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8dGFyZ2V0fGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=800&q=60)

## 🥅 Goals

- Learn new stuff (techniques, languages, etc.)
- Share your thought process
- **NOT** necessarily solve the problem
- Focus on what matters most
- Use [Test-Driven-Development](https://codingdojo.org/TestDrivenDevelopment/) and do [Baby Steps](https://codingdojo.org/BabySteps/)
- Follow the principles [YAGNI](https://en.wikipedia.org/wiki/You_aren%27t_gonna_need_it), [KISS](https://en.wikipedia.org/wiki/KISS_principle) and [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

---

![bg right:50% 95%](https://www.xeridia.co.uk/sites/default/files/contenidos/blog/test-driven-development.png)

## 🧪 Test Driven Develoment

[Test Driven Development](https://www.youtube.com/watch?v=Jv2uxzhPFl4)

[Smart Coffee Break: Test Driven Development](https://tube.video.bosch.com/media/Smart+Coffee+Break+-+Test+Driven+Development+%28TDD%29+%28S03E04%29/0_s1dwxf4e)

---

![bg right:40% 100%](https://images.unsplash.com/photo-1612611036764-f4d70058f83b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8cnVsZXN8ZW58MHx8MHx8&auto=format&fit=crop&w=800&q=60)

## 🚨 Rules

- Active member
    - has 5 minutes until we switch seats
    - is the only one speaking
    - has to share his screen and should try to express his thoughts to the others
    - Use TDD and do Baby Steps
    - If stuck, should specifically ask the inactive members for help
- Inactive members 
    - should pay close attention to the coding and do **not** do something else
    - only answer questions if explicitly asked
    
---

![bg right:50%](https://images.unsplash.com/photo-1584907797015-7554cd315667?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1080&q=80)

## 🚧 Preparation

For the sake of time, today I selected a programming language and a kata. Next time we will let the team decide what we want to code and which language to use.

We are going to use the already prepared VSCode Development Container, so everybody can just use their local machine to work on the code.

---
## 💻 Why Go? 

![bg right:50% 80%](https://i.ytimg.com/vi/446E-r0rXHI/maxresdefault.jpg)
[🔥 Fireship Video](https://www.youtube.com/watch?v=446E-r0rXHI)
[📖 Go Cheat Sheat](https://devhints.io/go)

---

## Mob Programming

![bg right 50%](https://mob.sh/logo.svg)
To make working a little bit more fun, we are going to use [mob programming](https://mob.sh/). To decide who is next you could e.g. use the [wheel of names](https://wheelofnames.com/). Mob is already installed in the Development Container, so we can just use it:

---
## Usage

Make sure that the exports like in [env.sample](./env.sample) are set correctly. Then you can start your timer with:

> `mob start 5`

Do your work for 5 minutes. Commit your changes with:

> `mob next`

Watch the next person in the loop and have some fun. If the Coding Dojo is done merge the changes with:

> `mob done`

---

## Timer

You can access the timer on this website:

<https://timer.mob.sh/itk-cloud-coding-dojo>

(the URL must match the `MOB_TIMER_ROOM` value)
