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
    if (homeButton) {
        let scaleUp = true;

        setInterval(() => {
            homeButton.style.transform = scaleUp ? "scale(1.05)" : "scale(1)";
            scaleUp = !scaleUp;
        }, 1000);

        homeButton.addEventListener("click", function () {
            window.location.href = "templates/accueil.html";
        });
    }
});

document.addEventListener("DOMContentLoaded", function() {
    Prism.highlightAll();
});

function showSection(sectionId) {
    const sections = document.querySelectorAll('.section-content');
    sections.forEach(section => {
        section.style.display = 'none';
    });

    const selectedSection = document.getElementById(sectionId);
    if (selectedSection) {
        selectedSection.style.display = 'block';
    }
}