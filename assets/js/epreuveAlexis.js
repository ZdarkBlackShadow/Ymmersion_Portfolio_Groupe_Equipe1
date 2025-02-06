// Fonction existante pour gérer les onglets
function openTab(tabId) {
    // Masquer tous les contenus d'onglet
    const tabContents = document.querySelectorAll('.tab-content');
    tabContents.forEach(tab => tab.classList.remove('active-content'));

    // Retirer la classe active de tous les onglets
    const tabs = document.querySelectorAll('.tab');
    tabs.forEach(tab => tab.classList.remove('active-tab'));

    // Afficher le contenu de l'onglet sélectionné et activer l'onglet
    document.getElementById(tabId).classList.add('active-content');
    const activeTab = Array.from(tabs).find(tab => tab.classList.contains(tabId));
    activeTab.classList.add('active-tab');

    // Changer la couleur de fond de chaque onglet
    tabs.forEach(tab => {
        tab.style.backgroundColor = "#ffffff"; // Réinitialise la couleur de fond pour tous
    });

    // Appliquer une nouvelle couleur de fond pour l'onglet actif
    activeTab.style.backgroundColor = "#0ed300"; // Changer la couleur de fond de l'onglet actif
} 

// Lorsque le DOM est chargé, on ajoute l'écouteur d'événement pour le formulaire
document.addEventListener("DOMContentLoaded", function() {
    // Sélectionner le formulaire par son ID
    const formPassword = document.getElementById("formPassword");
    if (formPassword) {
        formPassword.addEventListener("submit", function(e) {
            e.preventDefault(); // Empêche le rechargement de la page

            // Récupérer la valeur de la longueur saisie par l'utilisateur
            const length = parseInt(document.getElementById("length").value, 10);

            // Envoyer la requête POST en JSON vers le serveur
            fetch("/generatePasswd", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ length: length })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Erreur lors de la génération du mot de passe");
                }
                return response.json();
            })
            .then(data => {
                // Afficher le mot de passe généré dans le div avec l'ID "result"
                document.getElementById("result").innerText = data.mdp;
            })
            .catch(error => {
                console.error("Erreur:", error);
                document.getElementById("result").innerText = "Une erreur est survenue.";
            });
        });
    }
});
