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

    const formTextAnalyse = document.getElementById("formTextAnalyse");
    if (formTextAnalyse) {
        formTextAnalyse.addEventListener("submit", function(e) {
            e.preventDefault();
            const inputText = document.getElementById("inputText").value;
            fetch("/analyzeText", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ inputText: inputText })
            })
            .then(response => { 
                if (!response.ok) {
                    throw new Error("Erreur lors de l'analyse du texte");
                }
                return response.json();
            })
            .then(data => {
                const resultText =
                " - Résultats de l'analyse :\n" +
                " - Nombre de mots : " + data.wordCount + "\n" +
                " - Nombre de caractères (sans espaces) : " + data.charCount + "\n" +
                " - Mot le plus long : \"" + data.longestWord + "\"";
                document.getElementById("result").innerText = resultText; 
            })
            .catch(error => {
                console.error("Erreur:", error);
                document.getElementById("result").innerText = "Une erreur est survenue.";
            });
        });
    }
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




/*

document.addEventListener("DOMContentLoaded", function () {

    const formTextAnalyse = document.getElementById("formTextAnalyse");
    if (formTextAnalyse) {
        formTextAnalyse.addEventListener("submit", function(e) {
            e.preventDefault();
            const inputText = document.getElementById("inputText").value;
            fetch("/analyzeText", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ inputText: inputText })
            })
            .then(response => { 
                if (!response.ok) {
                    throw new Error("Erreur lors de l'analyse du texte");
                }
                return response.json();
            })
            .then(data => {
                const resultText =
                " - Résultats de l'analyse :\n" +
                " - Nombre de mots : " + data.wordCount + "\n" +
                " - Nombre de caractères (sans espaces) : " + data.charCount + "\n" +
                " - Mot le plus long : \"" + data.longestWord + "\"";
                document.getElementById("result").innerText = resultText; 
            })
            .catch(error => {
                console.error("Erreur:", error);
                document.getElementById("result").innerText = "Une erreur est survenue.";
            });
        });
    }
});
*/