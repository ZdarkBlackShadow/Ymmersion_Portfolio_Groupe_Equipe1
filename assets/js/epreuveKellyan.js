document.addEventListener("DOMContentLoaded", function () {
    Prism.highlightAll();
  });
  document.addEventListener("DOMContentLoaded", () => {
    const columns = document.querySelectorAll(".column");

    columns.forEach((column, index) => {
      if (index % 2 === 0) {
        column.classList.add("slide-left");
      } else {
        column.classList.add("slide-right");
      }
    });
  });

  document.addEventListener("DOMContentLoaded", function() {
    const formPalindrome = document.getElementById("formPalindrome");
    if (formPalindrome) {
        formPalindrome.addEventListener("submit", function(e) {
            e.preventDefault();
            const text = document.getElementById("text").value;

            fetch("/checkPalindrome", {
                method: "POST",
                headers: { "Content-type": "application/json" },
                body: JSON.stringify({ text: text })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Erreur lors de la vérification du Palindrome");
                }
                return response.json();
            })
            .then(data => {
                document.getElementById("result").innerText = data.text;
            })
            .catch(error => {
                console.error("Erreur:", error);
                document.getElementById("result").innerText = "Une erreur est survenue.";
            });
        });
    }
  });