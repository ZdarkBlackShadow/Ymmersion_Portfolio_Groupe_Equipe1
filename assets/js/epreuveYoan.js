document.addEventListener("DOMContentLoaded", function () {
    const formPasswd = document.getElementById("formPasswd");
    if (formPasswd) {
        formPasswd.addEventListener("submit", function (e) {
            e.preventDefault();
            const passwd = document.getElementById("passwd").value;

            fetch("/checkPasswd", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ passwd: passwd})
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Erreur lors de la vérification du mot de passe")
                }
                return response.json();
            })
            .then(data => {
                document.getElementById("result").innerText = data.text;
        })
        .catch(error => {
             console.error("Erreur : ", error);
             document.getElementById("result").innerText = "Une erreur est survenue";
         });
        });
    }
});