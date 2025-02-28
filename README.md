<p align="center">
  <img src="static/img/TEXT%20LOGO.PNG" alt="banner"/>
</p>

---

## 📖 Table of contents


1. [**📚 About the project**](#-about-the-project)
2. [**🚀 How to run the project**](#-how-to-run-the-project)
3. [**💻 Technologies**](#-technologies)
4. [**👥 Credits**](#-credits)

---

## 📚 About the project

This project was developed during the Groupie-Tracker project, by Ynov Campus. The goal was to create a Website using Go templates to learn how to handle and use API.

We decided to add some features to the site, such as:
- a homePage to show ten artist at random
- a search bar to find your favorite artist
- a about to know about the project and the team

For the theme of the site we decided to take a bit of inspiration in a well known site called spotify but in thruth there are the one whom copied from us 

The final project repository can be founded [here](https://github.com/zoom26042604/Groupie_Tracker)

And our presentation for the project can be found [here](https://www.canva.com/design/DAGgBIMDB50/dKuQRZQub2fcFYP97AXzcg/edit?utm_content=DAGgBIMDB50&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)

And lastly our trello can be found [here](https://trello.com/invite/b/67602859ddbab9490a52b7ec/ATTIc68593987b3cf4f8af880611d2be9febA4F2F9EE/groupie-tracker)


---

## 🚀 How to run the project

To run the project, you will need to have Go installed on your computer. If you don't have it, you can download it [here](https://golang.org/dl/).

1. Clone the repository:
```bash
git clone https://github.com/zoom26042604/Groupie_Tracker.git
cd Groupie_Tracker
go run main.go
```

2. Open your browser and go to `http://localhost:8080/` to visit the site and find your favorite Band and Artist.

![alt text](/static/img/presentation_image.png)
---

## Project structure

```bash
Groupie_Tracker/
├── GetAPI/
│   ├── GetAPI.go
│   ├── params.go
── handler/
│  ├── Handler.go
├── static/
│   ├── css/
│   │   └── about.css
│   │   └── artistPage.css
│   │   └── searchbar.css
│   │   └── searchPage.css
│   │   └── style.css
│   ├── img/
│   │   ├── TEXT LOGO.PNG
│   │   └── POSTIFY.png
│   └── js/
│       └── filtrevisibility.js
│       └── map.js
│       └── slide.js
│   └── json/
│       └── info.json
├── templates/
│   ├── index.gohtml
│   └── artistPage.gohtml
│   └── about.gohtml
│   └── search.gohtml
│   └── map.gohtml
├── .gitignore
├── go.mod
├── main.go
└── README.md
```

---
## 💻 Technologies

The project was developed using the following technologies:
- [Go](https://golang.org/)
- [GoHTML](https://pkg.go.dev/html/template)
- [CSS]()
- [JavaScript](https://www.javascript.com/)
- [Mapbox](https://www.mapbox.com/)

---

## 👥 Credits

This project was developed by:
<br>
<a href="https://github.com/zoom26042604"><img src="https://avatars.githubusercontent.com/u/186803356?v=4" alt="Nathan FERRE" width="69" height="69"/></a>
<a href="https://github.com/LeRaphouu"><img src="https://avatars.githubusercontent.com/u/188911609?v=4" alt="Raphael BONNET" width="69" height="69"/></a>
<a href="https://github.com/tompass8"><img src="https://avatars.githubusercontent.com/u/183885775?v=4" alt="Tom " width="69" height="69"/></a>
