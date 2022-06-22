/*
let myCourses = ["Learn CSS Animations", "UI Design Fundamentals", "Intro to Clean Code"]

function render(myCourses)
{
    for(let i = 0; i < myCourses.length; i++)
    {
        console.log(myCourses[i])
    }
}
render(myCourses)
*/

// Save to Local Storage
/*

localStorage.setItem("Sports", "Football")
localStorage.setItem("Courses", "Website")
let print = localStorage.getItem("Sports")

console.log(print)

localStorage.clear()
*/

// Adding addEventListener and object in array
/*
let data = [
    {
        player: "Jane",
        score: 52
    },
    {
        player: "Mark",
        score: 41
    }
]

const log_btn = document.getElementById("log-btn")

log_btn.addEventListener("click", function(){console.log(data[0].score)})
*/

// Generate a sentence
/*
const description = "best fruits"
const array = ["Oranges", "Apples"]

function generateSentence(desc, arr)
{
    let output = arr[0]
    const lastIndex= arr.length - 1
    for(let i = 1; i < lastIndex; i++)
    {
        output += ", " + arr[i]
    }
    output += " and " + arr[lastIndex]
    const sentence = `The ${arr.length} ${desc} are ${output}`
    console.log(sentence)
}

generateSentence(description, array)
*/

// Render Images

const imgs = [
    "https://www.w3schools.com/howto/img_avatar.png",
    "https://www.w3schools.com/howto/img_avatar.png",
    "https://www.w3schools.com/howto/img_avatar.png"
]

const img_disp = document.getElementById("img-disp")

function renderImages()
{
    let imagesTag = ""
    for(let i = 0; i < imgs.length; i++)
    {
        imagesTag += `<img class="team_img" src="${imgs[i]}">`
    }
    img_disp.innerHTML = imagesTag
}
renderImages()