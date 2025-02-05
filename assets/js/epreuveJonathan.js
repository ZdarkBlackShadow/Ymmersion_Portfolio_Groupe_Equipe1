document.addEventListener("DOMContentLoaded", function () {
    const elements = document.querySelectorAll(".fade-in");

    function showElements() {
        elements.forEach((el) => {
            if (el.getBoundingClientRect().top < window.innerHeight - 50) {
                el.classList.add("show");
            }
        });
    }
    window.addEventListener("scroll", showElements);
    showElements();
    const homeButton = document.getElementById("homeButton");
    let scaleUp = true;

    setInterval(() => {
        if (scaleUp) {
            homeButton.style.transform = "scale(1.05)";
        } else {
            homeButton.style.transform = "scale(1)";
        }
        scaleUp = !scaleUp;
    }, 1000);
    homeButton.addEventListener("click", function () {
        window.location.href = "templates/accueil.html";
    });
});
